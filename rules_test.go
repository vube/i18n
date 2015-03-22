package i18n

import (
	"reflect"

	. "gopkg.in/check.v1"
)

func (s *MySuite) TestLoad(c *C) {
	t := new(TranslatorRules)

	errs := t.load([]string{"does/not/exist/xx.yaml"})
	c.Check(errs, Not(HasLen), 0)

	errs = t.load([]string{"data/rules/root.yaml", s.rulesDir + "/en.yaml"})
	c.Check(errs, HasLen, 0)
	c.Check(t.Currencies, Not(HasLen), 0)
	c.Check(t.Currencies["USD"].Symbol, Equals, "US$")
	c.Check(t.Direction, Equals, "LTR")
	c.Check(t.Numbers.Symbols.Decimal, Equals, "≥")
	c.Check(t.Numbers.Symbols.Group, Equals, "≤")
	c.Check(t.Numbers.Symbols.Negative, Equals, "-")
	c.Check(t.Numbers.Symbols.Percent, Equals, "ﬁ")
	c.Check(t.Numbers.Symbols.Permille, Equals, "‰")
	c.Check(t.Numbers.Formats.Currency, Equals, "¤\u00a0#,##0.00")
	c.Check(t.Numbers.Formats.Decimal, Equals, "#,##0.###")
	c.Check(t.Numbers.Formats.Percent, Equals, "#,##0%")
	c.Check(t.Plural, Equals, "2A")
	c.Check(funcEquals(t.PluralRuleFunc, pluralRule2A), Equals, true)

	// basic check for all complete locales
	locales := []string{
		"aa", "af", "agq", "ak", "am", "ar", "as", "asa", "ast", "az", "bas",
		"be", "bem", "bez", "bg", "bm", "bn", "bo", "br", "brx", "bs", "byn",
		"ca", "cgg", "chr", "cs", "cy", "da", "dav", "de", "dje", "dua", "dyo",
		"dz", "ebu", "ee", "el", "en", "eo", "es", "et", "eu", "ewo", "fa",
		"ff", "fi", "fil", "fo", "fr", "fur", "ga", "gd", "gl", "gsw", "gu",
		"guz", "gv", "ha", "haw", "he", "hi", "hr", "ht", "hu", "hy", "ia",
		"id", "ig", "ii", "is", "it", "ja", "jgo", "jmc", "ka", "kab", "kam",
		"kde", "kea", "khq", "ki", "kk", "kkj", "kl", "kln", "km", "kn", "ko",
		"kok", "ks", "ksb", "ksf", "ksh", "kw", "ky", "lag", "lg", "ln", "lo",
		"lt", "lu", "luo", "luy", "lv", "mas", "mer", "mfe", "mg", "mgh", "mgo",
		"mk", "ml", "mn", "mr", "ms", "mt", "mua", "my", "naq", "nb", "nd",
		"ne", "nl", "nmg", "nn", "nnh", "no", "nr", "nso", "nus", "nyn", "om",
		"or", "os", "pa", "pl", "ps", "pt", "rm", "rn", "ro", "rof", "ru", "rw",
		"rwk", "sah", "saq", "sbp", "se", "seh", "ses", "sg", "shi", "si", "sk",
		"sl", "sn", "so", "sq", "sr", "ss", "ssy", "st", "sv", "sw", "swc",
		"ta", "te", "teo", "tg", "th", "ti", "tig", "tn", "to", "tr", "ts",
		"twq", "tzm", "uk", "ur", "uz", "vai", "ve", "vi", "vo", "vun", "wae",
		"wal", "xh", "xog", "yav", "yo", "zh-hans", "zh", "zu",
	}

	for _, l := range locales {
		t := new(TranslatorRules)
		errs = t.load([]string{"data/rules/" + l + ".yaml"})
		c.Check(errs, HasLen, 0)
	}

	// these are missing plural rule - on purpose
	locales = []string{
		"en-au",
		"en-gb",
		"fr-ca",
		"nl-be",
		"pt-br",
	}

	for _, l := range locales {
		t := new(TranslatorRules)
		errs = t.load([]string{"data/rules/" + l + ".yaml"})
		c.Log(l)
		c.Check(errs, HasLen, 1)
	}
}

func funcEquals(f1 func(float64) int, f2 func(float64) int) bool {
	return reflect.ValueOf(f1) == reflect.ValueOf(f2)
}
