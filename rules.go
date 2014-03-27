package i18n

import (
	// standard library
	"io/ioutil"
	"os"

	// third party
	"launchpad.net/goyaml"
)

// constants for text directionality
// using strings rather than iotas for easy yaml unmarshalling
const (
	direction_ltr = "LTR"
	direction_rtl = "RTL"
)

// translatorRules is a struct containing all of the information unmarshalled
// from a locale rules file.
type translatorRules struct {
	Plural         string `yaml:"plural,omitempty"`
	PluralRuleFunc pluralRule
	Direction      string `yaml:"direction,omitempty"`
	Numbers        struct {
		Symbols struct {
			Decimal  string `yaml:"decimal,omitempty"`
			Group    string `yaml:"group,omitempty"`
			Negative string `yaml:"negative,omitempty"`
			Percent  string `yaml:"percent,omitempty"`
			Permille string `yaml:"permille,omitempty"`
		} `yaml:"symbols,omitempty"`
		Formats struct {
			Decimal  string `yaml:"decimal,omitempty"`
			Currency string `yaml:"currency,omitempty"`
			Percent  string `yaml:"percent,omitempty"`
		} `yaml:"formats,omitempty"`
	} `yaml:"numbers,omitempty"`
	Currencies map[string]currency `yaml:"currencies,omitempty"`
}

// currency is a struct that's used in the above translatorRules struct for
// capturing the rule info for a single currency
type currency struct {
	Symbol      string `yaml:"symbol,omitempty"`
	Translation string `yaml:"display,omitempty"`
}

// load unmarshalls rule data from yaml files into the translator's rules
func (t *translatorRules) load(files []string) (errors []error) {

	for _, file := range files {
		_, statErr := os.Stat(file)
		if statErr == nil {
			contents, readErr := ioutil.ReadFile(file)

			if readErr != nil {
				errors = append(errors, translatorError{message: "can't open rules file: " + readErr.Error()})
			}

			tNew := new(translatorRules)
			yamlErr := goyaml.Unmarshal(contents, tNew)

			if yamlErr != nil {
				errors = append(errors, translatorError{message: "can't load rules YAML: " + yamlErr.Error()})
			} else {
				t.merge(tNew)
			}
		}
	}

	// set the plural rule func
	pRule, ok := pluralRules[t.Plural]
	if ok {
		t.PluralRuleFunc = pRule
	} else {
		if t.Plural == "" {
			errors = append(errors, translatorError{message: "missing plural rule: " + t.Plural})

		} else {
			errors = append(errors, translatorError{message: "invalid plural rule: " + t.Plural})
		}
		t.PluralRuleFunc = pluralRules["1"]
	}

	if t.Direction == "" {
		errors = append(errors, translatorError{message: "missing direction rule"})
		t.Direction = direction_ltr
	} else if t.Direction != direction_ltr && t.Direction != direction_rtl {
		errors = append(errors, translatorError{message: "invalid direction rule: " + t.Direction})
		t.Direction = direction_ltr
	}

	return
}

// merge takes another translatorRules instance and safely merges its metadata
// into this instance. this replaces yaml marshalling directly into the same
// instance - as that doesn't do what we want for deep merging.
func (t *translatorRules) merge(tNew *translatorRules) {
	if tNew.Plural != "" {
		t.Plural = tNew.Plural
	}

	if tNew.PluralRuleFunc != nil {
		t.PluralRuleFunc = tNew.PluralRuleFunc
	}

	if tNew.Direction != "" {
		t.Direction = tNew.Direction
	}

	if tNew.Numbers.Symbols.Decimal != "" {
		t.Numbers.Symbols.Decimal = tNew.Numbers.Symbols.Decimal
	}
	if tNew.Numbers.Symbols.Group != "" {
		t.Numbers.Symbols.Group = tNew.Numbers.Symbols.Group
	}
	if tNew.Numbers.Symbols.Negative != "" {
		t.Numbers.Symbols.Negative = tNew.Numbers.Symbols.Negative
	}
	if tNew.Numbers.Symbols.Percent != "" {
		t.Numbers.Symbols.Percent = tNew.Numbers.Symbols.Percent
	}
	if tNew.Numbers.Symbols.Permille != "" {
		t.Numbers.Symbols.Permille = tNew.Numbers.Symbols.Permille
	}

	if tNew.Numbers.Formats.Decimal != "" {
		t.Numbers.Formats.Decimal = tNew.Numbers.Formats.Decimal
	}
	if tNew.Numbers.Formats.Currency != "" {
		t.Numbers.Formats.Currency = tNew.Numbers.Formats.Currency
	}
	if tNew.Numbers.Formats.Percent != "" {
		t.Numbers.Formats.Percent = tNew.Numbers.Formats.Percent
	}

	if tNew.Currencies != nil {
		if t.Currencies == nil {
			t.Currencies = tNew.Currencies
		} else {
			for i, c := range tNew.Currencies {
				if _, ok := t.Currencies[i]; !ok {
					t.Currencies[i] = c
				} else {

					curr := t.Currencies[i]

					if c.Symbol != "" {
						curr.Symbol = c.Symbol
					}
					if c.Translation != "" {
						curr.Translation = c.Translation
					}
					t.Currencies[i] = curr
				}
			}
		}
	}
}
