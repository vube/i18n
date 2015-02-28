package i18n

import (
	"io/ioutil"
	"os"
	"testing"

	. "gopkg.in/check.v1"
)

// passes control of tests off to go-check
func Test(t *testing.T) { TestingT(t) }

type MySuite struct {
	messagesDir       string
	rulesDir          string
	messagesDirStyle1 string
	messagesDirStyle2 string
}

var _ = Suite(&MySuite{})

var (
	enMessages = `
WELCOME      : "Howdy!"
WELCOME_USER : "Howdy, {user}!"
`
	enRules = `
direction: LTR
plural: 2A
numbers:
  symbols:
    decimal:  "≥"
    group:    "≤"
    percent:  "ﬁ"
`
)

// fixtures
func (s *MySuite) SetUpSuite(c *C) {
	tmpDir := c.MkDir()

	s.messagesDir = tmpDir + "/messages"
	s.messagesDirStyle1 = tmpDir + "messages-1"
	s.messagesDirStyle2 = tmpDir + "messages-2"
	s.rulesDir = tmpDir + "/rules"

	err := os.Mkdir(s.messagesDir, os.FileMode(0777))
	c.Assert(err, IsNil)

	err = os.Mkdir(s.messagesDirStyle1, os.FileMode(0777))
	c.Assert(err, IsNil)

	err = os.Mkdir(s.messagesDirStyle2, os.FileMode(0777))
	c.Assert(err, IsNil)

	err = os.Mkdir(s.rulesDir, os.FileMode(0777))
	c.Assert(err, IsNil)

	err = os.Mkdir(s.messagesDir+"/en", os.FileMode(0777))
	c.Assert(err, IsNil)

	err = os.Mkdir(s.messagesDirStyle2+"/en", os.FileMode(0777))
	c.Assert(err, IsNil)

	// write the messages yaml files - mixed style
	err = ioutil.WriteFile(s.messagesDir+"/en.yaml", []byte(enMessages), os.FileMode(0777))
	c.Assert(err, IsNil)

	contents, readErr := ioutil.ReadFile(s.messagesDir + "/en.yaml")
	c.Assert(readErr, IsNil)
	c.Assert(string(contents), Equals, enMessages)

	err = ioutil.WriteFile(s.messagesDir+"/en/default.yaml", []byte(`GOODBYE : "So long!"`), os.FileMode(0777))
	c.Assert(err, IsNil)

	contents, readErr = ioutil.ReadFile(s.messagesDir + "/en/default.yaml")
	c.Assert(readErr, IsNil)
	c.Assert(string(contents), Equals, `GOODBYE : "So long!"`)

	// write the messages yaml files - style 1
	err = ioutil.WriteFile(s.messagesDirStyle1+"/en.yaml", []byte(enMessages), os.FileMode(0777))
	c.Assert(err, IsNil)

	contents, readErr = ioutil.ReadFile(s.messagesDir + "/en/default.yaml")
	c.Assert(readErr, IsNil)
	c.Assert(string(contents), Equals, `GOODBYE : "So long!"`)

	// write the messages yaml files - style 2
	err = ioutil.WriteFile(s.messagesDirStyle2+"/en/default.yaml", []byte(`GOODBYE : "So long!"`), os.FileMode(0777))
	c.Assert(err, IsNil)

	contents, readErr = ioutil.ReadFile(s.messagesDirStyle2 + "/en/default.yaml")
	c.Assert(readErr, IsNil)
	c.Assert(string(contents), Equals, `GOODBYE : "So long!"`)

	// write the rules yaml files
	err = ioutil.WriteFile(s.rulesDir+"/en.yaml", []byte(enRules), os.FileMode(0777))
	c.Assert(err, IsNil)

	contents, readErr = ioutil.ReadFile(s.rulesDir + "/en.yaml")
	c.Assert(readErr, IsNil)
	c.Assert(string(contents), Equals, enRules)
}

func (s *MySuite) SetUpTest(c *C) {
	// nothing to do here
}

func (s *MySuite) TearDownSuite(c *C) {
	// nothing to do here
}

func (s *MySuite) TearDownTest(c *C) {
	// clear out any loaded data
	numberFormats = map[string]*numberFormat{}
}

func (s *MySuite) TestNewTranslatorFactory(c *C) {
	// test success cases
	f, errors := NewTranslatorFactory(
		[]string{"data/rules"},
		[]string{"data/messages"},
		"",
	)
	c.Check(errors, HasLen, 0)
	c.Assert(f, NotNil)
	c.Assert(f.rulesPaths, HasLen, 1)
	c.Check(f.rulesPaths[0], Equals, "data/rules")
	c.Assert(f.messagesPaths, HasLen, 1)
	c.Check(f.messagesPaths[0], Equals, "data/messages")
	c.Check(f.translators, HasLen, 0)
	c.Check(f.fallback, IsNil)

	f, errors = NewTranslatorFactory(
		[]string{"data/rules", s.rulesDir},
		[]string{"data/messages", s.messagesDir},
		"en",
	)
	c.Check(errors, HasLen, 0)
	c.Assert(f, NotNil)
	c.Assert(f.rulesPaths, HasLen, 2)
	c.Check(f.rulesPaths[0], Equals, "data/rules")
	c.Check(f.rulesPaths[1], Equals, s.rulesDir)
	c.Assert(f.messagesPaths, HasLen, 2)
	c.Check(f.messagesPaths[0], Equals, "data/messages")
	c.Check(f.messagesPaths[1], Equals, s.messagesDir)
	c.Check(f.translators, HasLen, 1)
	c.Check(f.fallback, NotNil)
	c.Check(f.fallback.locale, Equals, "en")

	f, errors = NewTranslatorFactory(
		[]string{"data/rules"},
		[]string{s.messagesDirStyle1},
		"en",
	)
	c.Check(errors, HasLen, 0)
	c.Assert(f, NotNil)
	c.Assert(f.messagesPaths, HasLen, 1)
	c.Check(f.messagesPaths[0], Equals, s.messagesDirStyle1)

	f, errors = NewTranslatorFactory(
		[]string{"data/rules"},
		[]string{s.messagesDirStyle2},
		"en",
	)
	c.Check(errors, HasLen, 0)
	c.Assert(f, NotNil)
	c.Assert(f.messagesPaths, HasLen, 1)
	c.Check(f.messagesPaths[0], Equals, s.messagesDirStyle2)

	// test with no paths
	f, errors = NewTranslatorFactory(
		[]string{},
		[]string{},
		"",
	)
	c.Check(errors, HasLen, 2)

	// test with invalid paths
	f, errors = NewTranslatorFactory(
		[]string{"does-not-exist"},
		[]string{"does-not-exist"},
		"",
	)
	c.Check(errors, HasLen, 2)
}

func (s *MySuite) TestGetTranslator(c *C) {

	f, errors := NewTranslatorFactory(
		[]string{"data/rules"},
		[]string{"data/messages"},
		"en",
	)

	c.Assert(f, NotNil)
	c.Check(errors, HasLen, 0)

	tEn, errors := f.GetTranslator("en")

	c.Assert(tEn, NotNil)
	c.Check(errors, HasLen, 0)
	c.Check(tEn.messages, Not(HasLen), 0)
	c.Check(tEn.locale, Equals, "en")
	c.Check(tEn.rules, NotNil)
	c.Check(tEn.fallback, IsNil)

	tFr, errors := f.GetTranslator("fr")

	c.Assert(tFr, NotNil)
	c.Check(errors, HasLen, 0)
	c.Check(tFr.messages, Not(HasLen), 0)
	c.Check(tFr.locale, Equals, "fr")
	c.Check(tFr.rules, NotNil)
	c.Assert(tFr.fallback, NotNil)
	c.Check(tFr.fallback, Equals, tEn)

	tFrCa, errors := f.GetTranslator("fr-ca")

	c.Assert(tFrCa, NotNil)
	c.Assert(tFrCa.fallback, NotNil)
	c.Check(tFrCa.fallback, Equals, tFr)

	_, errors = f.GetTranslator("does-not-exist")

	c.Check(errors, Not(HasLen), 0)

	f, _ = NewTranslatorFactory(
		[]string{"data/rules"},
		[]string{s.messagesDir},
		"en",
	)

	c.Assert(f, NotNil)

	tEn, errors = f.GetTranslator("en")

	c.Assert(tEn, NotNil)
	c.Check(errors, HasLen, 0)
	c.Check(tEn.messages, HasLen, 3)

	f, _ = NewTranslatorFactory(
		[]string{"data/rules"},
		[]string{s.messagesDirStyle1},
		"en",
	)

	c.Assert(f, NotNil)

	tEn, errors = f.GetTranslator("en")

	c.Assert(tEn, NotNil)
	c.Check(errors, HasLen, 0)
	c.Check(tEn.messages, HasLen, 2)

	f, _ = NewTranslatorFactory(
		[]string{"data/rules"},
		[]string{s.messagesDirStyle2},
		"en",
	)

	c.Assert(f, NotNil)

	tEn, errors = f.GetTranslator("en")

	c.Assert(tEn, NotNil)
	c.Check(errors, HasLen, 0)
	c.Check(tEn.messages, HasLen, 1)
}

func (s *MySuite) TestGetFallback(c *C) {
	f, _ := NewTranslatorFactory(
		[]string{"data/rules"},
		[]string{"data/messages"},
		"en",
	)

	c.Assert(f, NotNil)

	c.Check(f.getFallback("en"), IsNil)
	c.Assert(f.getFallback("fr"), NotNil)
	c.Check(f.getFallback("fr").locale, Equals, "en")
	c.Assert(f.getFallback("fr-ca"), NotNil)
	c.Check(f.getFallback("fr-ca").locale, Equals, "fr")
}

func (s *MySuite) TestLocaleExists(c *C) {
	f, _ := NewTranslatorFactory(
		[]string{"data/rules"},
		[]string{"data/messages"},
		"en",
	)

	c.Assert(f, NotNil)

	exists, errs := f.LocaleExists("en")
	c.Check(exists, Equals, true)
	c.Check(errs, HasLen, 0)

	exists, errs = f.LocaleExists("does-not-exit")
	c.Check(exists, Equals, false)
	c.Check(errs, HasLen, 0)

	f, _ = NewTranslatorFactory(
		[]string{"data/rules"},
		[]string{s.messagesDir},
		"en",
	)

	c.Assert(f, NotNil)

	exists, errs = f.LocaleExists("en")
	c.Check(exists, Equals, true)
	c.Check(errs, HasLen, 0)

	f, _ = NewTranslatorFactory(
		[]string{"data/rules"},
		[]string{s.messagesDirStyle1},
		"en",
	)

	c.Assert(f, NotNil)

	exists, errs = f.LocaleExists("en")
	c.Check(exists, Equals, true)
	c.Check(errs, HasLen, 0)

	f, _ = NewTranslatorFactory(
		[]string{"data/rules"},
		[]string{s.messagesDirStyle2},
		"en",
	)

	c.Assert(f, NotNil)

	exists, errs = f.LocaleExists("en")
	c.Check(exists, Equals, true)
	c.Check(errs, HasLen, 0)
}

func (s *MySuite) TestTranslate(c *C) {

	f, errors := NewTranslatorFactory(
		[]string{"data/rules"},
		[]string{"data/messages"},
		"en",
	)

	c.Assert(f, NotNil)
	c.Check(errors, HasLen, 0)

	tEn, _ := f.GetTranslator("en")

	c.Assert(tEn, NotNil)

	// test basic translation
	welcome, errors := tEn.Translate("WELCOME", map[string]string{})
	c.Check(errors, HasLen, 0)
	c.Check(welcome, Equals, "Welcome!")

	// test translation with placeholder
	welcomeUser, errors := tEn.Translate("WELCOME_USER", map[string]string{"user": "Mother Goose"})
	c.Check(errors, HasLen, 0)
	c.Check(welcomeUser, Equals, "Welcome, Mother Goose!")

	// test for error when message doesn't exist
	notFound, errors := tEn.Translate("THIS_KEY_DOES_NOT_EXIST", map[string]string{})
	c.Check(errors, HasLen, 1)
	c.Check(notFound, Equals, "")

	// test for error when substitution does not exist
	noSubstitution, errors := tEn.Translate("WELCOME", map[string]string{"no": "substitution"})
	c.Check(errors, HasLen, 1)
	c.Check(noSubstitution, Equals, "Welcome!")

	// test fallback
	tFr, _ := f.GetTranslator("fr")

	welcome, errors = tFr.Translate("WELCOME", map[string]string{})

	c.Check(errors, HasLen, 0)
	c.Check(welcome, Equals, "Welcome!")

	// test double fallback
	tfrCa, _ := f.GetTranslator("fr-ca")

	welcome, errors = tfrCa.Translate("WELCOME", map[string]string{})

	c.Check(errors, HasLen, 0)
	c.Check(welcome, Equals, "Welcome!")
}

func (s *MySuite) TestPluralize(c *C) {

	f, errors := NewTranslatorFactory(
		[]string{"data/rules"},
		[]string{"data/messages"},
		"en",
	)

	c.Assert(f, NotNil)
	c.Check(errors, HasLen, 0)

	tEn, _ := f.GetTranslator("en")

	c.Assert(tEn, NotNil)

	p, errors := tEn.Pluralize("TIME_UNIT_DAY_FUTURE", 0, "0")
	c.Check(errors, HasLen, 0)
	c.Check(p, Equals, "In 0 days")

	p, errors = tEn.Pluralize("TIME_UNIT_DAY_FUTURE", 0.5, "0,5")
	c.Check(errors, HasLen, 0)
	c.Check(p, Equals, "In 0,5 days")

	p, errors = tEn.Pluralize("TIME_UNIT_DAY_FUTURE", 1, "one")
	c.Check(errors, HasLen, 0)
	c.Check(p, Equals, "In one day")

	p, errors = tEn.Pluralize("TIME_UNIT_DAY_FUTURE", -1, "-1")
	c.Check(errors, HasLen, 0)
	c.Check(p, Equals, "In -1 day")

	p, errors = tEn.Pluralize("TIME_UNIT_DAY_FUTURE", 1.5, "one and a half")
	c.Check(errors, HasLen, 0)
	c.Check(p, Equals, "In one and a half days")

	p, errors = tEn.Pluralize("TIME_UNIT_DAY_FUTURE", 2, "2.0")
	c.Check(errors, HasLen, 0)
	c.Check(p, Equals, "In 2.0 days")

	p, errors = tEn.Pluralize("WELCOME", 2, "2.0")
	c.Check(errors, HasLen, 2)
	c.Check(p, Equals, "Welcome!")

}

func (s *MySuite) TestDirection(c *C) {
	f, errors := NewTranslatorFactory(
		[]string{"data/rules"},
		[]string{"data/messages"},
		"en",
	)

	c.Assert(f, NotNil)
	c.Check(errors, HasLen, 0)

	tEn, _ := f.GetTranslator("en")
	tAr, _ := f.GetTranslator("ar")

	c.Assert(tEn, NotNil)
	c.Assert(tAr, NotNil)

	c.Check(tEn.Direction(), Equals, "LTR")
	c.Check(tAr.Direction(), Equals, "RTL")
}

func (s *MySuite) TestSubstitute(c *C) {

	f, errors := NewTranslatorFactory(
		[]string{"data/rules"},
		[]string{"data/messages"},
		"en",
	)

	c.Assert(f, NotNil)
	c.Check(errors, HasLen, 0)

	tEn, _ := f.GetTranslator("en")

	sub, errors := tEn.substitute("A {n} noise annoys a {n} {o}.", map[string]string{
		"o":         "oyster",
		"n":         "noisy",
		"not there": "not there",
	})

	c.Check(errors, HasLen, 1)
	c.Check(sub, Equals, "A noisy noise annoys a noisy oyster.")
}

func (s *MySuite) TestLoadMessages(c *C) {
	messages, errors := loadMessages("en", []string{"data/messages", s.messagesDir})
	c.Check(errors, HasLen, 0)
	c.Check(messages["TIME_UNIT_DAY"], Equals, "{n} day|{n} days")
	c.Check(messages["WELCOME"], Equals, "Howdy!")
	c.Check(messages["GOODBYE"], Equals, "So long!")

	messages, errors = loadMessages("xx", []string{"does/not/exist"})
	c.Check(errors, Not(HasLen), 0)
	c.Check(messages, HasLen, 0)
}
