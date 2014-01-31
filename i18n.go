package i18n

import (
	// standard library
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	// third party
	"launchpad.net/goyaml"
)

// TranslatorFactory is a struct which contains the info necessary for creating
// Translator "instances". It also "caches" previously created Translators.
// Because of this caching, you can request a Translator for a specific locale
// multiple times and always get a pointer to the same Translator instance.
type TranslatorFactory struct {
	messagesPaths []string
	rulesPaths    []string
	translators   map[string]*Translator
	fallback      *Translator
}

// Translator is a struct which contains all the rules and messages necessary
// to do internationalization for a specific locale. Most functionality in this
// package is accessed through a Translator instance.
type Translator struct {
	messages map[string]string
	locale   string
	rules    *translatorRules
	fallback *Translator
}

// translatorError implements the error interface for use in this package. it
// keeps an optional reference to a Translator instance, which it uses to
// include which locale the error occurs with in the error message returned by
// the Error() method
type translatorError struct {
	translator *Translator
	message    string
}

// Error satisfies the error interface requirements
func (e translatorError) Error() string {
	if e.translator != nil {
		return "translator error (locale: " + e.translator.locale + ") - " + e.message
	}
	return "translator error - " + e.message
}

// NewTranslatorFactory returns a TranslatorFactory instance with the specified
// paths and fallback locale.  If a fallback locale is specified, it
// automatically creates the fallback Translator instance. Several errors can
// occur during this process, and those are all returned in the errors slice.
// Even if errors are returned, this function should still return a working
// Translator if the fallback works.
//
// If multiple rulesPaths or messagesPaths are provided, they loaded in the
// order they appear in the slice, with values added later overriding any rules
// or messages loaded earlier.
//
// One lat thing about the messagesPaths. You can organize your locale messages
// files in this messagesPaths directory in 2 different ways.
//
//  1) Place *.yaml files in that directory directly, named after locale codes -
//
//     messages/
//       en.yaml
//       fr.yaml
//
//  2) Place subdirectores in that directory, named after locale codes and
//     containing *.yaml files
//
//     messages/
//       en/
//         front-end.yaml
//         email.yaml
//       fr/
//         front-end.yaml
//         email.yaml
//
//  Using the second way allows you to organize your messages into multiple
//  files.
func NewTranslatorFactory(rulesPaths []string, messagesPaths []string, fallbackLocale string) (f *TranslatorFactory, errors []error) {
	f = new(TranslatorFactory)

	if len(rulesPaths) == 0 {
		errors = append(errors, translatorError{message: "rules paths empty"})
	}

	if len(messagesPaths) == 0 {
		errors = append(errors, translatorError{message: "messages paths empty"})
	}

	foundRules := fallbackLocale == ""
	foundMessages := fallbackLocale == ""

	for _, path := range rulesPaths {
		path = strings.TrimRight(path, "/")
		_, err := os.Stat(path)
		if err != nil {
			errors = append(errors, translatorError{message: "can't read rules path " + path + ": " + err.Error()})
		}

		if !foundRules {
			_, err = os.Stat(path + "/" + fallbackLocale + ".yaml")
			if err == nil {
				foundRules = true
			}
		}
	}

	for _, path := range messagesPaths {
		path = strings.TrimRight(path, "/")
		_, err := os.Stat(path)
		if err != nil {
			errors = append(errors, translatorError{message: "can't read messages path " + path + ": " + err.Error()})
		}

		if !foundMessages {
			_, err = os.Stat(path + "/" + fallbackLocale + ".yaml")
			if err == nil {
				foundMessages = true
			} else {
				info, err := os.Stat(path + "/" + fallbackLocale)
				if err == nil && info.IsDir() {
					files, _ := filepath.Glob(path + "/" + fallbackLocale + "/*.yaml")
					for _, file := range files {
						_, err = os.Stat(file)
						if err == nil {
							foundMessages = true
							break
						}
					}
				}
			}
		}
	}

	if !foundRules {
		errors = append(errors, translatorError{message: "found no rules for fallback locale"})
	}

	if !foundMessages {
		errors = append(errors, translatorError{message: "found no messages for fallback locale"})
	}

	f.rulesPaths = rulesPaths
	f.messagesPaths = messagesPaths
	f.translators = map[string]*Translator{}

	// load and check the fallback locale
	if fallbackLocale != "" {
		var errs []error
		f.fallback, errs = f.GetTranslator(fallbackLocale)
		for _, err := range errs {
			errors = append(errors, err)
		}
	}

	return
}

// GetTranslator returns an Translator instance for the requested locale. If you
// request the same locale multiple times, a pointed to the same Translator will
// be returned each time.
func (f *TranslatorFactory) GetTranslator(localeCode string) (t *Translator, errors []error) {

	fallback := f.getFallback(localeCode)

	if t, ok := f.translators[localeCode]; ok {
		return t, nil
	}

	if !f.LocaleExists(localeCode) {
		errors = append(errors, translatorError{message: "could not find rules and messages for locale " + localeCode})
	}

	rules := new(translatorRules)
	files := []string{}

	// TODO: the rules loading logic is fairly complex, and there are some
	// specific cases we are not testing for yet. We need to test that the
	// fallback locale rules do not influence the rules loaded, and that the
	// base rules do.

	// the load the base (default) rule values
	// the step above
	for _, path := range f.rulesPaths {
		path = strings.TrimRight(path, "/")
		files = append(files, path+"/base.yaml")
	}

	// load less specific fallback locale rules
	parts := strings.Split(localeCode, "-")
	if len(parts) > 1 {
		for i, _ := range parts {
			fb := strings.Join(parts[0:i+1], "-")
			for _, path := range f.rulesPaths {
				path = strings.TrimRight(path, "/")
				files = append(files, path+"/"+fb+".yaml")
			}
		}
	}

	// finally load files for this specific locale
	for _, path := range f.rulesPaths {
		path = strings.TrimRight(path, "/")
		files = append(files, path+"/"+localeCode+".yaml")
	}

	errs := rules.load(files)
	for _, err := range errs {
		errors = append(errors, err)
	}

	messages, errs := loadMessages(localeCode, f.messagesPaths)
	for _, err := range errs {
		errors = append(errors, err)
	}

	t = new(Translator)
	t.locale = localeCode
	t.messages = messages
	t.fallback = fallback
	t.rules = rules

	f.translators[localeCode] = t

	return
}

// getFallback returns the best fallback for this locale. It first checks for
// less specific versions of the locale before falling back to the global
// fallback if it exists.
func (f *TranslatorFactory) getFallback(localeCode string) *Translator {

	if f.fallback != nil && localeCode == f.fallback.locale {
		return nil
	}

	separator := "-"

	// for a "multipart" locale code, find the most appropriate fallback
	// start by taking off the last "part"
	// if you run out of parts, use the factory's fallback

	fallback := f.fallback
	parts := strings.Split(localeCode, separator)
	for len(parts) > 1 {
		parts = parts[0 : len(parts)-1]
		fb := strings.Join(parts, separator)

		if f.LocaleExists(fb) {
			fallback, _ = f.GetTranslator(fb)
			break
		}
	}

	return fallback
}

// LocaleExists checks to see if any messages files exist for the requested
// locale string.
func (f *TranslatorFactory) LocaleExists(localeCode string) bool {
	for _, path := range f.messagesPaths {
		path = strings.TrimRight(path, "/")
		_, err := os.Stat(path + "/" + localeCode + ".yaml")
		if err == nil {
			return true
		}

		info, err := os.Stat(path + "/" + localeCode)
		if err == nil && info.IsDir() {
			files, _ := filepath.Glob(path + "/" + localeCode + "/*.yaml")
			for _, file := range files {
				_, err := os.Stat(file)
				if err == nil {
					return true
				}
			}
		}
	}

	return false
}

// Translate returns the translated message, performang any substitutions
// requested in the substitutions map. If neither this translator nor its
// fallback translator (or the fallback's fallback and so on) have a translation
// for the requested key, and empty string and an error will be returned.
func (t *Translator) Translate(key string, substitutions map[string]string) (translation string, errors []error) {
	if _, ok := t.messages[key]; !ok {
		if t.fallback != nil && t.fallback != t {
			return t.fallback.Translate(key, substitutions)
		}

		errors = append(errors, translatorError{translator: t, message: "key not found: " + key})
		return
	}

	translation, errors = t.subsitute(t.messages[key], substitutions)
	return
}

// Pluralize returns the translation for a message containing a plural. The
// plural form used is based on the number float64 and the number displayed in
// the translated string is the numberStr string. If neither this translator nor
// its fallback translator (or the fallback's fallback and so on) have a
// translation for the requested key, and empty string and an error will be
// returned.
func (t *Translator) Pluralize(key string, number float64, numberStr string) (translation string, errors []error) {

	if _, ok := t.messages[key]; !ok {
		if t.fallback != nil && t.fallback != t {
			return t.fallback.Pluralize(key, number, numberStr)
		}

		errors = append(errors, translatorError{translator: t, message: "key not found: " + key})
		return
	}

	form := (t.rules.PluralRuleFunc)(number)

	parts := strings.Split(t.messages[key], "|")

	if form > len(parts)-1 {
		errors = append(errors, translatorError{translator: t, message: "too few plural variations: " + key})
		form = len(parts) - 1
	}

	var errs []error
	translation, errs = t.subsitute(parts[form], map[string]string{"n": numberStr})
	for _, err := range errs {
		errors = append(errors, err)
	}
	return
}

// subsitute returns a string copy of the input str string will all keys in the
// substitutions map replaced with their value.
func (t *Translator) subsitute(str string, substitutions map[string]string) (substituted string, errors []error) {

	substituted = str

	for find, replace := range substitutions {
		if strings.Index(str, "{"+find+"}") == -1 {
			errors = append(errors, translatorError{translator: t, message: "substitution not found: " + str + ", " + replace})
		}
		substituted = strings.Replace(substituted, "{"+find+"}", replace, -1)
	}

	return
}

// loadMessages loads all messages from the properly named locale message yaml
// files in the requested messagesPaths.  if multiple paths are provided, paths
// further down the list take precedence over earlier paths.
func loadMessages(locale string, messagesPaths []string) (messages map[string]string, errors []error) {

	messages = make(map[string]string)

	found := false
	for _, path := range messagesPaths {
		path = strings.TrimRight(path, "/")
		file := path + "/" + locale + ".yaml"

		_, statErr := os.Stat(file)
		if statErr == nil {
			contents, readErr := ioutil.ReadFile(file)
			if readErr != nil {
				errors = append(errors, translatorError{message: "can't open messages file: " + readErr.Error()})
			} else {
				newmap := map[string]string{}
				yamlErr := goyaml.Unmarshal(contents, &newmap)
				if yamlErr != nil {
					errors = append(errors, translatorError{message: "can't load messages YAML: " + yamlErr.Error()})
				} else {
					found = true
					for key, value := range newmap {
						messages[key] = value
					}
				}
			}
		}

		// now look for a directory named after this locale and get an *.yaml children
		dir := path + "/" + locale
		info, statErr := os.Stat(dir)
		if statErr == nil && info.IsDir() {
			// found the directory - now look for *.yaml files
			files, globErr := filepath.Glob(dir + "/*.yaml")
			if globErr != nil {
				errors = append(errors, translatorError{message: "can't glob messages files: " + globErr.Error()})
			}
			for _, file := range files {
				contents, readErr := ioutil.ReadFile(file)
				if readErr != nil {
					errors = append(errors, translatorError{message: "can't open messages file: " + readErr.Error()})
				} else {
					newmap := map[string]string{}
					yamlErr := goyaml.Unmarshal(contents, &newmap)
					if yamlErr != nil {
						errors = append(errors, translatorError{message: "can't load messages YAML: " + yamlErr.Error()})
					} else {
						found = true
						for key, value := range newmap {
							messages[key] = value
						}
					}
				}
			}
		}
	}

	if !found {
		errors = append(errors, translatorError{message: "no messages files found: " + locale})
	}

	return
}
