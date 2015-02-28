package i18n

import (
	. "gopkg.in/check.v1"
)

func (s *MySuite) TestFormatCurrency(c *C) {

	f, errors := NewTranslatorFactory(
		[]string{"data/rules"},
		[]string{"data/messages"},
		"en",
	)

	c.Assert(f, NotNil)
	c.Check(errors, HasLen, 0)

	tEn, _ := f.GetTranslator("en")

	cur, err := tEn.FormatCurrency(12345.6789, "USD")
	c.Check(err, IsNil)
	c.Check(cur, Equals, "$12,345.68")

	cur, err = tEn.FormatCurrency(-12345.6789, "USD")
	c.Check(err, IsNil)
	c.Check(cur, Equals, "($12,345.68)")

	cur, err = tEn.FormatCurrency(12345.6789, "WHAT???")
	c.Check(err, NotNil)
	c.Check(cur, Equals, "WHAT???12,345.68")

	// try some really big numbers to make sure weird floaty stuff doesn't
	// happen
	cur, err = tEn.FormatCurrency(12345000000000.6789, "USD")
	c.Check(err, IsNil)
	c.Check(cur, Equals, "$12,345,000,000,000.68")

	cur, err = tEn.FormatCurrency(-12345000000000.6789, "USD")
	c.Check(err, IsNil)
	c.Check(cur, Equals, "($12,345,000,000,000.68)")

	// Try something that needs a partial fallback
	tSaq, _ := f.GetTranslator("saq")

	cur, err = tSaq.FormatCurrency(12345.6789, "USD")
	c.Check(err, IsNil)
	c.Check(cur, Equals, "US$12,345.68")

	// And one more for with some unusual symbols for good measure
	tAr, _ := f.GetTranslator("ar")

	cur, err = tAr.FormatCurrency(-12345.6789, "USD")
	c.Check(err, IsNil)
	c.Check(cur, Equals, "US$ 12,345.68-")
}

func (s *MySuite) TestFormatCurrencyWhole(c *C) {

	f, errors := NewTranslatorFactory(
		[]string{"data/rules"},
		[]string{"data/messages"},
		"en",
	)

	c.Assert(f, NotNil)
	c.Check(errors, HasLen, 0)

	tEn, _ := f.GetTranslator("en")

	cur, err := tEn.FormatCurrencyWhole(12345.6789, "USD")
	c.Check(err, IsNil)
	c.Check(cur, Equals, "$12,346")

	cur, err = tEn.FormatCurrencyWhole(-12345.6789, "USD")
	c.Check(err, IsNil)
	c.Check(cur, Equals, "($12,346)")

	cur, err = tEn.FormatCurrencyWhole(12345.6789, "WHAT???")
	c.Check(err, NotNil)
	c.Check(cur, Equals, "WHAT???12,346")

	// try some really big numbers to make sure weird floaty stuff doesn't
	// happen
	cur, err = tEn.FormatCurrencyWhole(12345000000000.6789, "USD")
	c.Check(err, IsNil)
	c.Check(cur, Equals, "$12,345,000,000,001")

	cur, err = tEn.FormatCurrencyWhole(-12345000000000.6789, "USD")
	c.Check(err, IsNil)
	c.Check(cur, Equals, "($12,345,000,000,001)")
}

func (s *MySuite) TestFormatNumber(c *C) {

	f, errors := NewTranslatorFactory(
		[]string{"data/rules"},
		[]string{"data/messages"},
		"en",
	)

	c.Assert(f, NotNil)
	c.Check(errors, HasLen, 0)

	// check basic english
	tEn, _ := f.GetTranslator("en")

	num := tEn.FormatNumber(12345.6789)
	c.Check(num, Equals, "12,345.679")

	num = tEn.FormatNumber(-12345.6789)
	c.Check(num, Equals, "-12,345.679")

	num = tEn.FormatNumber(123456789)
	c.Check(num, Equals, "123,456,789")

	// check Hindi - different group sizes
	tHi, _ := f.GetTranslator("hi")

	num = tHi.FormatNumber(12345.6789)
	c.Check(num, Equals, "12,345.679")

	num = tHi.FormatNumber(-12345.6789)
	c.Check(num, Equals, "-12,345.679")

	num = tHi.FormatNumber(123456789)
	c.Check(num, Equals, "12,34,56,789")

	// check Uzbek - something with a partial fallback
	tUz, _ := f.GetTranslator("uz")

	num = tUz.FormatNumber(12345.6789)
	c.Check(num, Equals, "12 345,679")

	num = tUz.FormatNumber(-12345.6789)
	c.Check(num, Equals, "-12 345,679")

	num = tUz.FormatNumber(123456789)
	c.Check(num, Equals, "123 456 789")

	format := &(numberFormat{
		positivePrefix:   "p",
		positiveSuffix:   "P%",
		negativePrefix:   "n",
		negativeSuffix:   "N‰",
		multiplier:       2,
		minDecimalDigits: 2,
		maxDecimalDigits: 5,
		minIntegerDigits: 10,
		groupSizeFinal:   6,
		groupSizeMain:    3,
	})

	tEn.rules.Numbers.Symbols.Decimal = ".."
	tEn.rules.Numbers.Symbols.Group = ",,"
	tEn.rules.Numbers.Symbols.Negative = "--"
	tEn.rules.Numbers.Symbols.Percent = "%%"
	tEn.rules.Numbers.Symbols.Permille = "‰‰"

	// check numbers with too few integer digits & too many decimal digits
	num = tEn.formatNumber(format, 1.12341234)
	c.Check(num, Equals, "p0,,000,,000002..24682P%%")
	num = tEn.formatNumber(format, -1.12341234)
	c.Check(num, Equals, "n0,,000,,000002..24682N‰‰")

	// check numbers with more than enough integer digits  & too few decimal digits
	num = tEn.formatNumber(format, 1234123412341234)
	c.Check(num, Equals, "p2,,468,,246,,824,,682468..00P%%")
	num = tEn.formatNumber(format, -1234123412341234)
	c.Check(num, Equals, "n2,,468,,246,,824,,682468..00N‰‰")
}

func (s *MySuite) TestFormatNumberWhole(c *C) {

	f, errors := NewTranslatorFactory(
		[]string{"data/rules"},
		[]string{"data/messages"},
		"en",
	)

	c.Assert(f, NotNil)
	c.Check(errors, HasLen, 0)

	tEn, _ := f.GetTranslator("en")

	num := tEn.FormatNumberWhole(12345.6789)
	c.Check(num, Equals, "12,346")

	num = tEn.FormatNumberWhole(-12345.6789)
	c.Check(num, Equals, "-12,346")
}

func (s *MySuite) TestFormatPercent(c *C) {

	f, errors := NewTranslatorFactory(
		[]string{"data/rules"},
		[]string{"data/messages"},
		"en",
	)

	c.Assert(f, NotNil)
	c.Check(errors, HasLen, 0)

	tEn, _ := f.GetTranslator("en")

	cur := tEn.FormatPercent(0.01234)
	c.Check(cur, Equals, "1%")

	cur = tEn.FormatPercent(0.1234)
	c.Check(cur, Equals, "12%")

	cur = tEn.FormatPercent(1.234)
	c.Check(cur, Equals, "123%")

	cur = tEn.FormatPercent(12.34)
	c.Check(cur, Equals, "1,234%")
}

func (s *MySuite) TestParseFormat(c *C) {
	f, errors := NewTranslatorFactory(
		[]string{"data/rules"},
		[]string{"data/messages"},
		"en",
	)

	c.Assert(f, NotNil)
	c.Check(errors, HasLen, 0)

	tEn, _ := f.GetTranslator("en")

	format := tEn.parseFormat("#0", true)
	c.Assert(format, NotNil)
	c.Check(format.groupSizeFinal, Equals, 0)
	c.Check(format.groupSizeMain, Equals, 0)
	c.Check(format.maxDecimalDigits, Equals, 0)
	c.Check(format.minDecimalDigits, Equals, 0)
	c.Check(format.minIntegerDigits, Equals, 1)
	c.Check(format.multiplier, Equals, 1)
	c.Check(format.negativePrefix, Equals, tEn.rules.Numbers.Symbols.Negative)
	c.Check(format.negativeSuffix, Equals, "")
	c.Check(format.positivePrefix, Equals, "")
	c.Check(format.positiveSuffix, Equals, "")

	format = tEn.parseFormat("#0%", true)
	c.Assert(format, NotNil)
	c.Check(format.groupSizeFinal, Equals, 0)
	c.Check(format.groupSizeMain, Equals, 0)
	c.Check(format.maxDecimalDigits, Equals, 0)
	c.Check(format.minDecimalDigits, Equals, 0)
	c.Check(format.minIntegerDigits, Equals, 1)
	c.Check(format.multiplier, Equals, 100)
	c.Check(format.negativePrefix, Equals, tEn.rules.Numbers.Symbols.Negative)
	c.Check(format.negativeSuffix, Equals, "%")
	c.Check(format.positivePrefix, Equals, "")
	c.Check(format.positiveSuffix, Equals, "%")

	format = tEn.parseFormat("#0‰", true)
	c.Assert(format, NotNil)
	c.Check(format.groupSizeFinal, Equals, 0)
	c.Check(format.groupSizeMain, Equals, 0)
	c.Check(format.maxDecimalDigits, Equals, 0)
	c.Check(format.minDecimalDigits, Equals, 0)
	c.Check(format.minIntegerDigits, Equals, 1)
	c.Check(format.multiplier, Equals, 1000)
	c.Check(format.negativePrefix, Equals, tEn.rules.Numbers.Symbols.Negative)
	c.Check(format.negativeSuffix, Equals, "‰")
	c.Check(format.positivePrefix, Equals, "")
	c.Check(format.positiveSuffix, Equals, "‰")

	format = tEn.parseFormat("#0P;#0N", true)
	c.Assert(format, NotNil)
	c.Check(format.groupSizeFinal, Equals, 0)
	c.Check(format.groupSizeMain, Equals, 0)
	c.Check(format.maxDecimalDigits, Equals, 0)
	c.Check(format.minDecimalDigits, Equals, 0)
	c.Check(format.minIntegerDigits, Equals, 1)
	c.Check(format.multiplier, Equals, 1)
	c.Check(format.negativePrefix, Equals, "")
	c.Check(format.negativeSuffix, Equals, "N")
	c.Check(format.positivePrefix, Equals, "")
	c.Check(format.positiveSuffix, Equals, "P")

	format = tEn.parseFormat("P#0;N#0", true)
	c.Assert(format, NotNil)
	c.Check(format.groupSizeFinal, Equals, 0)
	c.Check(format.groupSizeMain, Equals, 0)
	c.Check(format.maxDecimalDigits, Equals, 0)
	c.Check(format.minDecimalDigits, Equals, 0)
	c.Check(format.minIntegerDigits, Equals, 1)
	c.Check(format.multiplier, Equals, 1)
	c.Check(format.negativePrefix, Equals, "N")
	c.Check(format.negativeSuffix, Equals, "")
	c.Check(format.positivePrefix, Equals, "P")
	c.Check(format.positiveSuffix, Equals, "")

	format = tEn.parseFormat("#00000.00000", true)
	c.Assert(format, NotNil)
	c.Check(format.groupSizeFinal, Equals, 0)
	c.Check(format.groupSizeMain, Equals, 0)
	c.Check(format.maxDecimalDigits, Equals, 5)
	c.Check(format.minDecimalDigits, Equals, 5)
	c.Check(format.minIntegerDigits, Equals, 5)
	c.Check(format.multiplier, Equals, 1)
	c.Check(format.negativePrefix, Equals, tEn.rules.Numbers.Symbols.Negative)
	c.Check(format.negativeSuffix, Equals, "")
	c.Check(format.positivePrefix, Equals, "")
	c.Check(format.positiveSuffix, Equals, "")

	format = tEn.parseFormat("#0.#####", true)
	c.Assert(format, NotNil)
	c.Check(format.groupSizeFinal, Equals, 0)
	c.Check(format.groupSizeMain, Equals, 0)
	c.Check(format.maxDecimalDigits, Equals, 5)
	c.Check(format.minDecimalDigits, Equals, 0)
	c.Check(format.minIntegerDigits, Equals, 1)
	c.Check(format.multiplier, Equals, 1)
	c.Check(format.negativePrefix, Equals, tEn.rules.Numbers.Symbols.Negative)
	c.Check(format.negativeSuffix, Equals, "")
	c.Check(format.positivePrefix, Equals, "")
	c.Check(format.positiveSuffix, Equals, "")

	format = tEn.parseFormat("##,#0", true)
	c.Assert(format, NotNil)
	c.Check(format.groupSizeFinal, Equals, 2)
	c.Check(format.groupSizeMain, Equals, 2)
	c.Check(format.maxDecimalDigits, Equals, 0)
	c.Check(format.minDecimalDigits, Equals, 0)
	c.Check(format.minIntegerDigits, Equals, 1)
	c.Check(format.multiplier, Equals, 1)
	c.Check(format.negativePrefix, Equals, tEn.rules.Numbers.Symbols.Negative)
	c.Check(format.negativeSuffix, Equals, "")
	c.Check(format.positivePrefix, Equals, "")
	c.Check(format.positiveSuffix, Equals, "")

	format = tEn.parseFormat("##,###,#0", true)
	c.Assert(format, NotNil)
	c.Check(format.groupSizeFinal, Equals, 2)
	c.Check(format.groupSizeMain, Equals, 3)
	c.Check(format.maxDecimalDigits, Equals, 0)
	c.Check(format.minDecimalDigits, Equals, 0)
	c.Check(format.minIntegerDigits, Equals, 1)
	c.Check(format.multiplier, Equals, 1)
	c.Check(format.negativePrefix, Equals, tEn.rules.Numbers.Symbols.Negative)
	c.Check(format.negativeSuffix, Equals, "")
	c.Check(format.positivePrefix, Equals, "")
	c.Check(format.positiveSuffix, Equals, "")

	// test includeDecimalDigits true vs false
	enDecimal := "#,##0.###"
	formatWith := tEn.parseFormat(enDecimal, true)
	formatWithout := tEn.parseFormat(enDecimal, false)
	c.Assert(formatWith, NotNil)
	c.Check(formatWith.maxDecimalDigits, Equals, 3)
	c.Check(formatWith.minDecimalDigits, Equals, 0)
	c.Assert(formatWithout, NotNil)
	c.Check(formatWithout.maxDecimalDigits, Equals, 0)
	c.Check(formatWithout.minDecimalDigits, Equals, 0)

	enCurrency := "¤#,##0.00;(¤#,##0.00)"
	formatWith = tEn.parseFormat(enCurrency, true)
	formatWithout = tEn.parseFormat(enCurrency, false)
	c.Assert(formatWith, NotNil)
	c.Check(formatWith.maxDecimalDigits, Equals, 2)
	c.Check(formatWith.minDecimalDigits, Equals, 2)
	c.Assert(formatWithout, NotNil)
	c.Check(formatWithout.maxDecimalDigits, Equals, 0)
	c.Check(formatWithout.minDecimalDigits, Equals, 0)
}

func (s *MySuite) TestChunkString(c *C) {
	str := ""
	size := 0
	chunks := chunkString(str, size)
	c.Check(chunks, HasLen, 0)

	str = ""
	size = 1
	chunks = chunkString(str, size)
	c.Check(chunks, HasLen, 0)

	str = "What noise annoys a noisy oyster?"
	size = 0
	chunks = chunkString(str, size)
	c.Assert(chunks, HasLen, 1)
	c.Check(chunks[0], Equals, str)

	str = "What noise annoys a noisy oyster?"
	size = 1
	chunks = chunkString(str, size)
	c.Assert(chunks, HasLen, 33)
	c.Check(chunks[0], Equals, "W")
	c.Check(chunks[1], Equals, "h")
	c.Check(chunks[2], Equals, "a")

	str = "What noise annoys a noisy oyster?"
	size = 2
	chunks = chunkString(str, size)
	c.Assert(chunks, HasLen, 17)
	c.Check(chunks[0], Equals, "W")
	c.Check(chunks[1], Equals, "ha")
	c.Check(chunks[2], Equals, "t ")

	str = "What noise annoys a noisy oyster?"
	size = 3
	chunks = chunkString(str, size)
	c.Assert(chunks, HasLen, 11)
	c.Check(chunks[0], Equals, "Wha")
	c.Check(chunks[1], Equals, "t n")
	c.Check(chunks[2], Equals, "ois")

	str = "What noise annoys a noisy oyster?"
	size = 33
	chunks = chunkString(str, size)
	c.Assert(chunks, HasLen, 1)
	c.Check(chunks[0], Equals, str)

	str = "What noise annoys a noisy oyster?"
	size = 133
	chunks = chunkString(str, size)
	c.Assert(chunks, HasLen, 1)
	c.Check(chunks[0], Equals, str)
}

func (s *MySuite) TestNumberRound(c *C) {
	num := float64(0)
	dec := 0
	rounded := numberRound(num, dec)
	c.Check(rounded, Equals, "0")

	num = float64(1)
	dec = 1
	rounded = numberRound(num, dec)
	c.Check(rounded, Equals, "1")

	// test round down
	num = 1.2
	dec = 0
	rounded = numberRound(num, dec)
	c.Check(rounded, Equals, "1")

	// test round up
	num = 1.6
	dec = 0
	rounded = numberRound(num, dec)
	c.Check(rounded, Equals, "2")

	num = 1.51
	dec = 0
	rounded = numberRound(num, dec)
	c.Check(rounded, Equals, "2")

	// test round to even
	num = 1.5
	dec = 0
	rounded = numberRound(num, dec)
	c.Check(rounded, Equals, "2")

	num = 2.5
	dec = 0
	rounded = numberRound(num, dec)
	c.Check(rounded, Equals, "2")

	// a few more
	num = 1.99
	dec = 1
	rounded = numberRound(num, dec)
	c.Check(rounded, Equals, "2")

	num = 1.23456789
	dec = 2
	rounded = numberRound(num, dec)
	c.Check(rounded, Equals, "1.23")

	num = 1.23456789
	dec = 3
	rounded = numberRound(num, dec)
	c.Check(rounded, Equals, "1.235")
}
