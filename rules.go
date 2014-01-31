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

			yamlErr := goyaml.Unmarshal(contents, t)

			if yamlErr != nil {
				errors = append(errors, translatorError{message: "can't load rules YAML: " + yamlErr.Error()})
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
