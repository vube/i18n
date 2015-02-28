package i18n

import (
	. "gopkg.in/check.v1"
)

func (s *MySuite) TestLen(c *C) {
	i := new(i18nSorter)
	i.toBeSorted = []interface{}{"a", "b", "c"}

	c.Check(i.Len(), Equals, 3)
}

func (s *MySuite) TestSwap(c *C) {
	i := new(i18nSorter)
	i.toBeSorted = []interface{}{"a", "b", "c"}
	i.Swap(0, 2)

	c.Check(i.toBeSorted[0], Equals, "c")
	c.Check(i.toBeSorted[1], Equals, "b")
	c.Check(i.toBeSorted[2], Equals, "a")
}

func (s *MySuite) TestLess(c *C) {
	i := new(i18nSorter)
	i.toBeSorted = []interface{}{"a", "b", "c"}
	i.getComparisonValueFunc = func(n interface{}) string {
		if s, ok := n.(string); ok {
			return s
		}
		return ""
	}

	c.Check(i.Less(0, 0), Equals, false)
	c.Check(i.Less(0, 1), Equals, true)
	c.Check(i.Less(0, 2), Equals, true)
	c.Check(i.Less(1, 0), Equals, false)
	c.Check(i.Less(1, 1), Equals, false)
	c.Check(i.Less(1, 2), Equals, true)
	c.Check(i.Less(2, 0), Equals, false)
	c.Check(i.Less(2, 1), Equals, false)
	c.Check(i.Less(2, 2), Equals, false)
}

func (s *MySuite) TestSortUniversal(c *C) {

	toBeSorted := []interface{}{"apple", "beet", "carrot", "ȧpricot", "ḃanana", "ċlementine"}
	getComparisonValueFunc := func(i interface{}) string {
		if s, ok := i.(string); ok {
			return s
		}
		return ""
	}

	SortUniversal(toBeSorted, getComparisonValueFunc)

	c.Assert(toBeSorted, HasLen, 6)
	c.Check(toBeSorted[0], Equals, "apple")
	c.Check(toBeSorted[1], Equals, "ȧpricot")
	c.Check(toBeSorted[2], Equals, "beet") // unicode normalization puts "b" before "ḃ"
	c.Check(toBeSorted[3], Equals, "ḃanana")
	c.Check(toBeSorted[4], Equals, "carrot")
	c.Check(toBeSorted[5], Equals, "ċlementine")
}

func (s *MySuite) TestSortLocal(c *C) {
	toBeSorted := []interface{}{"apple", "beet", "carrot", "ȧpricot", "ḃanana", "ċlementine"}
	getComparisonValueFunc := func(i interface{}) string {
		if s, ok := i.(string); ok {
			return s
		}
		return ""
	}

	SortLocal("en", toBeSorted, getComparisonValueFunc)

	c.Assert(toBeSorted, HasLen, 6)
	c.Check(toBeSorted[0], Equals, "apple")
	c.Check(toBeSorted[1], Equals, "ȧpricot")
	c.Check(toBeSorted[2], Equals, "ḃanana")
	c.Check(toBeSorted[3], Equals, "beet")
	c.Check(toBeSorted[4], Equals, "carrot")
	c.Check(toBeSorted[5], Equals, "ċlementine")
}

func (s *MySuite) TestSort(c *C) {

	f, errors := NewTranslatorFactory(
		[]string{"data/rules"},
		[]string{"data/messages"},
		"en",
	)

	c.Assert(f, NotNil)
	c.Check(errors, HasLen, 0)

	// there will be a collator for "en", so SortLocal will be used
	tEn, _ := f.GetTranslator("en")
	toBeSorted := []interface{}{"apple", "beet", "carrot", "ȧpricot", "ḃanana", "ċlementine"}
	getComparisonValueFunc := func(i interface{}) string {
		if s, ok := i.(string); ok {
			return s
		}
		return ""
	}

	tEn.Sort(toBeSorted, getComparisonValueFunc)
	c.Assert(toBeSorted, HasLen, 6)
	c.Check(toBeSorted[0], Equals, "apple")
	c.Check(toBeSorted[1], Equals, "ȧpricot")
	c.Check(toBeSorted[2], Equals, "ḃanana")
	c.Check(toBeSorted[3], Equals, "beet")
	c.Check(toBeSorted[4], Equals, "carrot")
	c.Check(toBeSorted[5], Equals, "ċlementine")

	// there will not be a collator for "does-not-exist", so SortUniversal will be used
	tWhat, _ := f.GetTranslator("does-not-exist")

	tWhat.Sort(toBeSorted, getComparisonValueFunc)
	c.Assert(toBeSorted, HasLen, 6)
	c.Check(toBeSorted[0], Equals, "apple")
	c.Check(toBeSorted[1], Equals, "ȧpricot")
	c.Check(toBeSorted[2], Equals, "beet") // unicode normalization puts "b" before "ḃ"
	c.Check(toBeSorted[3], Equals, "ḃanana")
	c.Check(toBeSorted[4], Equals, "carrot")
	c.Check(toBeSorted[5], Equals, "ċlementine")
}
