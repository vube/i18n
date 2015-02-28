package i18n

import (
	. "gopkg.in/check.v1"
)

func (s *MySuite) TestIsInt(c *C) {
	// positive test cases
	c.Check(isInt(float64(0)), Equals, true)
	c.Check(isInt(float64(1)), Equals, true)
	c.Check(isInt(float64(0.0)), Equals, true)
	c.Check(isInt(float64(1.0000)), Equals, true)
	c.Check(isInt(float64(-50)), Equals, true)

	// negative
	c.Check(isInt(float64(0.1)), Equals, false)
	c.Check(isInt(float64(-0.1)), Equals, false)
	c.Check(isInt(float64(0.00000000000001)), Equals, false)
}

func (s *MySuite) TestPluralRule1(c *C) {
	c.Check(pluralRule1(float64(0)), Equals, 0)
	c.Check(pluralRule1(float64(0.5)), Equals, 0)
	c.Check(pluralRule1(float64(100)), Equals, 0)
}

func (s *MySuite) TestPluralRule2A(c *C) {
	// first form
	c.Check(pluralRule2A(float64(-1)), Equals, 0)
	c.Check(pluralRule2A(float64(1)), Equals, 0)

	// second form
	c.Check(pluralRule2A(float64(0)), Equals, 1)
	c.Check(pluralRule2A(float64(0.5)), Equals, 1)
	c.Check(pluralRule2A(float64(2)), Equals, 1)
}

func (s *MySuite) TestPluralRule2B(c *C) {
	// first form
	c.Check(pluralRule2B(float64(-1)), Equals, 0)
	c.Check(pluralRule2B(float64(0)), Equals, 0)
	c.Check(pluralRule2B(float64(1)), Equals, 0)

	// second form
	c.Check(pluralRule2B(float64(0.5)), Equals, 1)
	c.Check(pluralRule2B(float64(2)), Equals, 1)
}

func (s *MySuite) TestPluralRule2C(c *C) {
	// first form
	c.Check(pluralRule2C(float64(-1)), Equals, 0)
	c.Check(pluralRule2C(float64(0)), Equals, 0)
	c.Check(pluralRule2C(float64(0.5)), Equals, 0)
	c.Check(pluralRule2C(float64(1)), Equals, 0)
	c.Check(pluralRule2C(float64(1.5)), Equals, 0)

	// second form
	c.Check(pluralRule2C(float64(2)), Equals, 1)
	c.Check(pluralRule2C(float64(2.5)), Equals, 1)
	c.Check(pluralRule2C(float64(100)), Equals, 1)
}

func (s *MySuite) TestPluralRule2D(c *C) {
	// first form
	c.Check(pluralRule2D(float64(-1)), Equals, 0)
	c.Check(pluralRule2D(float64(1)), Equals, 0)
	c.Check(pluralRule2D(float64(21)), Equals, 0)

	// second form
	c.Check(pluralRule2D(float64(0)), Equals, 1)
	c.Check(pluralRule2D(float64(0.5)), Equals, 1)
	c.Check(pluralRule2D(float64(2)), Equals, 1)
	c.Check(pluralRule2D(float64(11)), Equals, 1)
}

func (s *MySuite) TestPluralRule2E(c *C) {
	// first form
	c.Check(pluralRule2E(float64(-1)), Equals, 0)
	c.Check(pluralRule2E(float64(0)), Equals, 0)
	c.Check(pluralRule2E(float64(1)), Equals, 0)
	c.Check(pluralRule2E(float64(11)), Equals, 0)
	c.Check(pluralRule2E(float64(12)), Equals, 0)
	c.Check(pluralRule2E(float64(98)), Equals, 0)
	c.Check(pluralRule2E(float64(99)), Equals, 0)

	// second form
	c.Check(pluralRule2E(float64(0.5)), Equals, 1)
	c.Check(pluralRule2E(float64(2)), Equals, 1)
	c.Check(pluralRule2E(float64(10)), Equals, 1)
	c.Check(pluralRule2E(float64(100)), Equals, 1)
}

func (s *MySuite) TestPluralRule2F(c *C) {
	// first form
	c.Check(pluralRule2F(float64(-1)), Equals, 0)
	c.Check(pluralRule2F(float64(0)), Equals, 0)
	c.Check(pluralRule2F(float64(1)), Equals, 0)
	c.Check(pluralRule2F(float64(2)), Equals, 0)
	c.Check(pluralRule2F(float64(11)), Equals, 0)
	c.Check(pluralRule2F(float64(12)), Equals, 0)
	c.Check(pluralRule2F(float64(20)), Equals, 0)
	c.Check(pluralRule2F(float64(40)), Equals, 0)

	// second form
	c.Check(pluralRule2F(float64(0.5)), Equals, 1)
	c.Check(pluralRule2F(float64(3)), Equals, 1)
	c.Check(pluralRule2F(float64(10)), Equals, 1)
}

func (s *MySuite) TestPluralRule3A(c *C) {
	// first form
	c.Check(pluralRule3A(float64(0)), Equals, 0)

	// second form
	c.Check(pluralRule3A(float64(-1)), Equals, 1)
	c.Check(pluralRule3A(float64(1)), Equals, 1)
	c.Check(pluralRule3A(float64(21)), Equals, 1)

	// third form
	c.Check(pluralRule3A(float64(0.5)), Equals, 2)
	c.Check(pluralRule3A(float64(2)), Equals, 2)
	c.Check(pluralRule3A(float64(10)), Equals, 2)
	c.Check(pluralRule3A(float64(11)), Equals, 2)
}

func (s *MySuite) TestPluralRule3B(c *C) {
	// first form
	c.Check(pluralRule3B(float64(-1)), Equals, 0)
	c.Check(pluralRule3B(float64(1)), Equals, 0)

	// second form
	c.Check(pluralRule3B(float64(-2)), Equals, 1)
	c.Check(pluralRule3B(float64(2)), Equals, 1)

	// third form
	c.Check(pluralRule3B(float64(0)), Equals, 2)
	c.Check(pluralRule3B(float64(0.5)), Equals, 2)
	c.Check(pluralRule3B(float64(3)), Equals, 2)
	c.Check(pluralRule3B(float64(11)), Equals, 2)
}

func (s *MySuite) TestPluralRule3C(c *C) {
	// first form
	c.Check(pluralRule3C(float64(-1)), Equals, 0)
	c.Check(pluralRule3C(float64(1)), Equals, 0)

	// second form
	c.Check(pluralRule3C(float64(0)), Equals, 1)
	c.Check(pluralRule3C(float64(-11)), Equals, 1)
	c.Check(pluralRule3C(float64(11)), Equals, 1)
	c.Check(pluralRule3C(float64(19)), Equals, 1)
	c.Check(pluralRule3C(float64(111)), Equals, 1)
	c.Check(pluralRule3C(float64(119)), Equals, 1)

	// third form
	c.Check(pluralRule3C(float64(0.5)), Equals, 2)
	c.Check(pluralRule3C(float64(20)), Equals, 2)
	c.Check(pluralRule3C(float64(21)), Equals, 2)
}

func (s *MySuite) TestPluralRule3D(c *C) {
	// first form
	c.Check(pluralRule3D(float64(-1)), Equals, 0)
	c.Check(pluralRule3D(float64(1)), Equals, 0)
	c.Check(pluralRule3D(float64(21)), Equals, 0)

	// second form
	c.Check(pluralRule3D(float64(-2)), Equals, 1)
	c.Check(pluralRule3D(float64(2)), Equals, 1)
	c.Check(pluralRule3D(float64(9)), Equals, 1)
	c.Check(pluralRule3D(float64(22)), Equals, 1)
	c.Check(pluralRule3D(float64(29)), Equals, 1)

	// third form
	c.Check(pluralRule3D(float64(0)), Equals, 2)
	c.Check(pluralRule3D(float64(0.5)), Equals, 2)
	c.Check(pluralRule3D(float64(11)), Equals, 2)
	c.Check(pluralRule3D(float64(19)), Equals, 2)
}

func (s *MySuite) TestPluralRule3E(c *C) {
	// first form
	c.Check(pluralRule3E(float64(-1)), Equals, 0)
	c.Check(pluralRule3E(float64(1)), Equals, 0)

	// second form
	c.Check(pluralRule3E(float64(-2)), Equals, 1)
	c.Check(pluralRule3E(float64(2)), Equals, 1)
	c.Check(pluralRule3E(float64(3)), Equals, 1)
	c.Check(pluralRule3E(float64(4)), Equals, 1)

	// third form
	c.Check(pluralRule3E(float64(0)), Equals, 2)
	c.Check(pluralRule3E(float64(0.5)), Equals, 2)
	c.Check(pluralRule3E(float64(5)), Equals, 2)
	c.Check(pluralRule3E(float64(9)), Equals, 2)
	c.Check(pluralRule3E(float64(11)), Equals, 2)
	c.Check(pluralRule3E(float64(12)), Equals, 2)
	c.Check(pluralRule3E(float64(14)), Equals, 2)
}

func (s *MySuite) TestPluralRule3F(c *C) {
	// first form
	c.Check(pluralRule3F(float64(0)), Equals, 0)

	// second form
	c.Check(pluralRule3F(float64(-0.5)), Equals, 1)
	c.Check(pluralRule3F(float64(0.5)), Equals, 1)
	c.Check(pluralRule3F(float64(1)), Equals, 1)
	c.Check(pluralRule3F(float64(1.5)), Equals, 1)

	// third form
	c.Check(pluralRule3F(float64(-2)), Equals, 2)
	c.Check(pluralRule3F(float64(2)), Equals, 2)
	c.Check(pluralRule3F(float64(3)), Equals, 2)
}

func (s *MySuite) TestPluralRule3G(c *C) {
	// first form
	c.Check(pluralRule3G(float64(-0.5)), Equals, 0)
	c.Check(pluralRule3G(float64(0)), Equals, 0)
	c.Check(pluralRule3G(float64(0.5)), Equals, 0)
	c.Check(pluralRule3G(float64(1)), Equals, 0)

	// second form
	c.Check(pluralRule3G(float64(-2)), Equals, 1)
	c.Check(pluralRule3G(float64(2)), Equals, 1)
	c.Check(pluralRule3G(float64(3)), Equals, 1)
	c.Check(pluralRule3G(float64(9)), Equals, 1)
	c.Check(pluralRule3G(float64(10)), Equals, 1)

	// third
	c.Check(pluralRule3G(float64(1.5)), Equals, 2)
	c.Check(pluralRule3G(float64(11)), Equals, 2)
	c.Check(pluralRule3G(float64(12)), Equals, 2)
}

func (s *MySuite) TestPluralRule3H(c *C) {
	// first form
	c.Check(pluralRule3H(float64(0)), Equals, 0)

	// second form
	c.Check(pluralRule3H(float64(-1)), Equals, 1)
	c.Check(pluralRule3H(float64(1)), Equals, 1)

	// third form
	c.Check(pluralRule3H(float64(0.5)), Equals, 2)
	c.Check(pluralRule3H(float64(1.5)), Equals, 2)
	c.Check(pluralRule3H(float64(2)), Equals, 2)
	c.Check(pluralRule3H(float64(10)), Equals, 2)
	c.Check(pluralRule3H(float64(11)), Equals, 2)
}

func (s *MySuite) TestPluralRule3I(c *C) {
	// first form
	c.Check(pluralRule3I(float64(-1)), Equals, 0)
	c.Check(pluralRule3I(float64(1)), Equals, 0)

	// second form
	c.Check(pluralRule3I(float64(-2)), Equals, 1)
	c.Check(pluralRule3I(float64(2)), Equals, 1)
	c.Check(pluralRule3I(float64(3)), Equals, 1)
	c.Check(pluralRule3I(float64(4)), Equals, 1)
	c.Check(pluralRule3I(float64(22)), Equals, 1)
	c.Check(pluralRule3I(float64(23)), Equals, 1)
	c.Check(pluralRule3I(float64(24)), Equals, 1)

	// third form
	c.Check(pluralRule3I(float64(0)), Equals, 2)
	c.Check(pluralRule3I(float64(0.5)), Equals, 2)
	c.Check(pluralRule3I(float64(5)), Equals, 2)
	c.Check(pluralRule3I(float64(9)), Equals, 2)
	c.Check(pluralRule3I(float64(12)), Equals, 2)
	c.Check(pluralRule3I(float64(13)), Equals, 2)
	c.Check(pluralRule3I(float64(14)), Equals, 2)
	c.Check(pluralRule3I(float64(15)), Equals, 2)
}

func (s *MySuite) TestPluralRule4A(c *C) {
	// first form
	c.Check(pluralRule4A(float64(-1)), Equals, 0)
	c.Check(pluralRule4A(float64(1)), Equals, 0)

	// second form
	c.Check(pluralRule4A(float64(-2)), Equals, 1)
	c.Check(pluralRule4A(float64(2)), Equals, 1)

	// third form
	c.Check(pluralRule4A(float64(-10)), Equals, 2)
	c.Check(pluralRule4A(float64(10)), Equals, 2)
	c.Check(pluralRule4A(float64(20)), Equals, 2)
	c.Check(pluralRule4A(float64(100)), Equals, 2)

	// fourth form
	c.Check(pluralRule4A(float64(0)), Equals, 3)
	c.Check(pluralRule4A(float64(0.5)), Equals, 3)
	c.Check(pluralRule4A(float64(3)), Equals, 3)
	c.Check(pluralRule4A(float64(9)), Equals, 3)
	c.Check(pluralRule4A(float64(11)), Equals, 3)
}

func (s *MySuite) TestPluralRule4B(c *C) {
	// first form
	c.Check(pluralRule4B(float64(-1)), Equals, 0)
	c.Check(pluralRule4B(float64(1)), Equals, 0)
	c.Check(pluralRule4B(float64(21)), Equals, 0)

	// second form
	c.Check(pluralRule4B(float64(-2)), Equals, 1)
	c.Check(pluralRule4B(float64(2)), Equals, 1)
	c.Check(pluralRule4B(float64(3)), Equals, 1)
	c.Check(pluralRule4B(float64(4)), Equals, 1)
	c.Check(pluralRule4B(float64(22)), Equals, 1)
	c.Check(pluralRule4B(float64(23)), Equals, 1)
	c.Check(pluralRule4B(float64(24)), Equals, 1)

	// third form
	c.Check(pluralRule4B(float64(-5)), Equals, 2)
	c.Check(pluralRule4B(float64(0)), Equals, 2)
	c.Check(pluralRule4B(float64(5)), Equals, 2)
	c.Check(pluralRule4B(float64(6)), Equals, 2)
	c.Check(pluralRule4B(float64(8)), Equals, 2)
	c.Check(pluralRule4B(float64(9)), Equals, 2)
	c.Check(pluralRule4B(float64(11)), Equals, 2)
	c.Check(pluralRule4B(float64(12)), Equals, 2)
	c.Check(pluralRule4B(float64(13)), Equals, 2)
	c.Check(pluralRule4B(float64(14)), Equals, 2)

	// fourth form
	c.Check(pluralRule4B(float64(0.5)), Equals, 3)
	c.Check(pluralRule4B(float64(1.5)), Equals, 3)
}

func (s *MySuite) TestPluralRule4C(c *C) {
	// first form
	c.Check(pluralRule4C(float64(-1)), Equals, 0)
	c.Check(pluralRule4C(float64(1)), Equals, 0)

	// second form
	c.Check(pluralRule4C(float64(-2)), Equals, 1)
	c.Check(pluralRule4C(float64(2)), Equals, 1)
	c.Check(pluralRule4C(float64(3)), Equals, 1)
	c.Check(pluralRule4C(float64(4)), Equals, 1)
	c.Check(pluralRule4C(float64(22)), Equals, 1)
	c.Check(pluralRule4C(float64(23)), Equals, 1)
	c.Check(pluralRule4C(float64(24)), Equals, 1)

	// third form
	c.Check(pluralRule4C(float64(-10)), Equals, 2)
	c.Check(pluralRule4C(float64(10)), Equals, 2)
	c.Check(pluralRule4C(float64(11)), Equals, 2)
	c.Check(pluralRule4C(float64(12)), Equals, 2)
	c.Check(pluralRule4C(float64(13)), Equals, 2)
	c.Check(pluralRule4C(float64(14)), Equals, 2)
	c.Check(pluralRule4C(float64(15)), Equals, 2)
	c.Check(pluralRule4C(float64(16)), Equals, 2)
	c.Check(pluralRule4C(float64(18)), Equals, 2)
	c.Check(pluralRule4C(float64(19)), Equals, 2)
	c.Check(pluralRule4C(float64(20)), Equals, 2)
	c.Check(pluralRule4C(float64(21)), Equals, 2)
	c.Check(pluralRule4C(float64(25)), Equals, 2)
	c.Check(pluralRule4C(float64(26)), Equals, 2)
	c.Check(pluralRule4C(float64(28)), Equals, 2)
	c.Check(pluralRule4C(float64(29)), Equals, 2)

	// fourth form
	c.Check(pluralRule4C(float64(0.5)), Equals, 3)
	c.Check(pluralRule4C(float64(1.5)), Equals, 3)
}

func (s *MySuite) TestPluralRule4D(c *C) {
	// first form
	c.Check(pluralRule4D(float64(-1)), Equals, 0)
	c.Check(pluralRule4D(float64(1)), Equals, 0)
	c.Check(pluralRule4D(float64(101)), Equals, 0)

	// second form
	c.Check(pluralRule4D(float64(-2)), Equals, 1)
	c.Check(pluralRule4D(float64(2)), Equals, 1)
	c.Check(pluralRule4D(float64(102)), Equals, 1)

	// third form
	c.Check(pluralRule4D(float64(-3)), Equals, 2)
	c.Check(pluralRule4D(float64(3)), Equals, 2)
	c.Check(pluralRule4D(float64(4)), Equals, 2)
	c.Check(pluralRule4D(float64(103)), Equals, 2)
	c.Check(pluralRule4D(float64(104)), Equals, 2)

	// fourth form
	c.Check(pluralRule4D(float64(0)), Equals, 3)
	c.Check(pluralRule4D(float64(0.5)), Equals, 3)
	c.Check(pluralRule4D(float64(5)), Equals, 3)
	c.Check(pluralRule4D(float64(10)), Equals, 3)
	c.Check(pluralRule4D(float64(11)), Equals, 3)
	c.Check(pluralRule4D(float64(12)), Equals, 3)
	c.Check(pluralRule4D(float64(13)), Equals, 3)
	c.Check(pluralRule4D(float64(14)), Equals, 3)
}

func (s *MySuite) TestPluralRule4E(c *C) {
	// first form
	c.Check(pluralRule4E(float64(-1)), Equals, 0)
	c.Check(pluralRule4E(float64(1)), Equals, 0)

	// second form
	c.Check(pluralRule4E(float64(-2)), Equals, 1)
	c.Check(pluralRule4E(float64(0)), Equals, 1)
	c.Check(pluralRule4E(float64(2)), Equals, 1)
	c.Check(pluralRule4E(float64(10)), Equals, 1)
	c.Check(pluralRule4E(float64(102)), Equals, 1)
	c.Check(pluralRule4E(float64(110)), Equals, 1)

	// third form
	c.Check(pluralRule4E(float64(-11)), Equals, 2)
	c.Check(pluralRule4E(float64(11)), Equals, 2)
	c.Check(pluralRule4E(float64(19)), Equals, 2)
	c.Check(pluralRule4E(float64(111)), Equals, 2)
	c.Check(pluralRule4E(float64(119)), Equals, 2)

	// fourth form
	c.Check(pluralRule4E(float64(0.5)), Equals, 3)
	c.Check(pluralRule4E(float64(20)), Equals, 3)
	c.Check(pluralRule4E(float64(21)), Equals, 3)
	c.Check(pluralRule4E(float64(22)), Equals, 3)
	c.Check(pluralRule4E(float64(29)), Equals, 3)
}

func (s *MySuite) TestPluralRule4F(c *C) {
	// first form
	c.Check(pluralRule4F(float64(-1)), Equals, 0)
	c.Check(pluralRule4F(float64(1)), Equals, 0)
	c.Check(pluralRule4F(float64(11)), Equals, 0)

	// second form
	c.Check(pluralRule4F(float64(-2)), Equals, 1)
	c.Check(pluralRule4F(float64(2)), Equals, 1)
	c.Check(pluralRule4F(float64(12)), Equals, 1)

	// third form
	c.Check(pluralRule4F(float64(-3)), Equals, 2)
	c.Check(pluralRule4F(float64(3)), Equals, 2)
	c.Check(pluralRule4F(float64(10)), Equals, 2)
	c.Check(pluralRule4F(float64(13)), Equals, 2)
	c.Check(pluralRule4F(float64(19)), Equals, 2)

	// fourth form
	c.Check(pluralRule4F(float64(0)), Equals, 3)
	c.Check(pluralRule4F(float64(0.5)), Equals, 3)
	c.Check(pluralRule4F(float64(20)), Equals, 3)
	c.Check(pluralRule4F(float64(21)), Equals, 3)
	c.Check(pluralRule4F(float64(22)), Equals, 3)
	c.Check(pluralRule4F(float64(23)), Equals, 3)
	c.Check(pluralRule4F(float64(29)), Equals, 3)
	c.Check(pluralRule4F(float64(101)), Equals, 3)
	c.Check(pluralRule4F(float64(101)), Equals, 3)
	c.Check(pluralRule4F(float64(102)), Equals, 3)
	c.Check(pluralRule4F(float64(103)), Equals, 3)
	c.Check(pluralRule4F(float64(109)), Equals, 3)
}

func (s *MySuite) TestPluralRule5A(c *C) {
	// first form
	c.Check(pluralRule5A(float64(-1)), Equals, 0)
	c.Check(pluralRule5A(float64(1)), Equals, 0)

	// second form
	c.Check(pluralRule5A(float64(-2)), Equals, 1)
	c.Check(pluralRule5A(float64(2)), Equals, 1)

	// third form
	c.Check(pluralRule5A(float64(-3)), Equals, 2)
	c.Check(pluralRule5A(float64(3)), Equals, 2)
	c.Check(pluralRule5A(float64(4)), Equals, 2)
	c.Check(pluralRule5A(float64(5)), Equals, 2)
	c.Check(pluralRule5A(float64(6)), Equals, 2)

	// fourth form
	c.Check(pluralRule5A(float64(-7)), Equals, 3)
	c.Check(pluralRule5A(float64(7)), Equals, 3)
	c.Check(pluralRule5A(float64(8)), Equals, 3)
	c.Check(pluralRule5A(float64(9)), Equals, 3)
	c.Check(pluralRule5A(float64(10)), Equals, 3)

	// fifth form
	c.Check(pluralRule5A(float64(0)), Equals, 4)
	c.Check(pluralRule5A(float64(0.5)), Equals, 4)
	c.Check(pluralRule5A(float64(11)), Equals, 4)
	c.Check(pluralRule5A(float64(12)), Equals, 4)
	c.Check(pluralRule5A(float64(13)), Equals, 4)
	c.Check(pluralRule5A(float64(14)), Equals, 4)
	c.Check(pluralRule5A(float64(15)), Equals, 4)
	c.Check(pluralRule5A(float64(16)), Equals, 4)
	c.Check(pluralRule5A(float64(17)), Equals, 4)
	c.Check(pluralRule5A(float64(18)), Equals, 4)
	c.Check(pluralRule5A(float64(19)), Equals, 4)
	c.Check(pluralRule5A(float64(20)), Equals, 4)
}

func (s *MySuite) TestPluralRule5B(c *C) {
	// first form
	c.Check(pluralRule5B(float64(-1)), Equals, 0)
	c.Check(pluralRule5B(float64(1)), Equals, 0)
	c.Check(pluralRule5B(float64(21)), Equals, 0)
	c.Check(pluralRule5B(float64(61)), Equals, 0)
	c.Check(pluralRule5B(float64(81)), Equals, 0)
	c.Check(pluralRule5B(float64(101)), Equals, 0)

	// second form
	c.Check(pluralRule5B(float64(-2)), Equals, 1)
	c.Check(pluralRule5B(float64(2)), Equals, 1)
	c.Check(pluralRule5B(float64(22)), Equals, 1)
	c.Check(pluralRule5B(float64(62)), Equals, 1)
	c.Check(pluralRule5B(float64(82)), Equals, 1)
	c.Check(pluralRule5B(float64(102)), Equals, 1)

	// third form
	c.Check(pluralRule5B(float64(-3)), Equals, 2)
	c.Check(pluralRule5B(float64(3)), Equals, 2)
	c.Check(pluralRule5B(float64(4)), Equals, 2)
	c.Check(pluralRule5B(float64(9)), Equals, 2)
	c.Check(pluralRule5B(float64(23)), Equals, 2)
	c.Check(pluralRule5B(float64(24)), Equals, 2)
	c.Check(pluralRule5B(float64(29)), Equals, 2)
	c.Check(pluralRule5B(float64(63)), Equals, 2)
	c.Check(pluralRule5B(float64(64)), Equals, 2)
	c.Check(pluralRule5B(float64(69)), Equals, 2)
	c.Check(pluralRule5B(float64(83)), Equals, 2)
	c.Check(pluralRule5B(float64(84)), Equals, 2)
	c.Check(pluralRule5B(float64(89)), Equals, 2)
	c.Check(pluralRule5B(float64(103)), Equals, 2)
	c.Check(pluralRule5B(float64(104)), Equals, 2)
	c.Check(pluralRule5B(float64(109)), Equals, 2)

	// fourth form
	c.Check(pluralRule5B(float64(-1000000)), Equals, 3)
	c.Check(pluralRule5B(float64(1000000)), Equals, 3)
	c.Check(pluralRule5B(float64(2000000)), Equals, 3)
	c.Check(pluralRule5B(float64(10000000)), Equals, 3)

	// fourth form
	c.Check(pluralRule5B(float64(0)), Equals, 4)
	c.Check(pluralRule5B(float64(0.5)), Equals, 4)
	c.Check(pluralRule5B(float64(10)), Equals, 4)
	c.Check(pluralRule5B(float64(11)), Equals, 4)
	c.Check(pluralRule5B(float64(12)), Equals, 4)
	c.Check(pluralRule5B(float64(13)), Equals, 4)
	c.Check(pluralRule5B(float64(14)), Equals, 4)
	c.Check(pluralRule5B(float64(19)), Equals, 4)
	c.Check(pluralRule5B(float64(20)), Equals, 4)
	c.Check(pluralRule5B(float64(71)), Equals, 4)
	c.Check(pluralRule5B(float64(72)), Equals, 4)
	c.Check(pluralRule5B(float64(73)), Equals, 4)
	c.Check(pluralRule5B(float64(74)), Equals, 4)
	c.Check(pluralRule5B(float64(79)), Equals, 4)
	c.Check(pluralRule5B(float64(91)), Equals, 4)
	c.Check(pluralRule5B(float64(92)), Equals, 4)
	c.Check(pluralRule5B(float64(93)), Equals, 4)
	c.Check(pluralRule5B(float64(94)), Equals, 4)
	c.Check(pluralRule5B(float64(99)), Equals, 4)
	c.Check(pluralRule5B(float64(100)), Equals, 4)
	c.Check(pluralRule5B(float64(1000)), Equals, 4)
	c.Check(pluralRule5B(float64(10000)), Equals, 4)
	c.Check(pluralRule5B(float64(100000)), Equals, 4)
}

func (s *MySuite) TestPluralRule6A(c *C) {
	// first form
	c.Check(pluralRule6A(float64(0)), Equals, 0)

	// second form
	c.Check(pluralRule6A(float64(-1)), Equals, 1)
	c.Check(pluralRule6A(float64(1)), Equals, 1)

	// third form
	c.Check(pluralRule6A(float64(-2)), Equals, 2)
	c.Check(pluralRule6A(float64(2)), Equals, 2)

	// fourth form
	c.Check(pluralRule6A(float64(-3)), Equals, 3)
	c.Check(pluralRule6A(float64(3)), Equals, 3)
	c.Check(pluralRule6A(float64(4)), Equals, 3)
	c.Check(pluralRule6A(float64(9)), Equals, 3)
	c.Check(pluralRule6A(float64(10)), Equals, 3)
	c.Check(pluralRule6A(float64(103)), Equals, 3)
	c.Check(pluralRule6A(float64(104)), Equals, 3)
	c.Check(pluralRule6A(float64(109)), Equals, 3)
	c.Check(pluralRule6A(float64(110)), Equals, 3)

	// fifth form
	c.Check(pluralRule6A(float64(-11)), Equals, 4)
	c.Check(pluralRule6A(float64(11)), Equals, 4)
	c.Check(pluralRule6A(float64(12)), Equals, 4)
	c.Check(pluralRule6A(float64(98)), Equals, 4)
	c.Check(pluralRule6A(float64(99)), Equals, 4)
	c.Check(pluralRule6A(float64(111)), Equals, 4)
	c.Check(pluralRule6A(float64(112)), Equals, 4)
	c.Check(pluralRule6A(float64(198)), Equals, 4)
	c.Check(pluralRule6A(float64(199)), Equals, 4)

	// sixth form
	c.Check(pluralRule6A(float64(0.5)), Equals, 5)
	c.Check(pluralRule6A(float64(100)), Equals, 5)
	c.Check(pluralRule6A(float64(102)), Equals, 5)
	c.Check(pluralRule6A(float64(200)), Equals, 5)
	c.Check(pluralRule6A(float64(202)), Equals, 5)
}

func (s *MySuite) TestPluralRule6B(c *C) {
	// first form
	c.Check(pluralRule6B(float64(0)), Equals, 0)

	// second form
	c.Check(pluralRule6B(float64(-1)), Equals, 1)
	c.Check(pluralRule6B(float64(1)), Equals, 1)

	// third form
	c.Check(pluralRule6B(float64(-2)), Equals, 2)
	c.Check(pluralRule6B(float64(2)), Equals, 2)

	// fourth form
	c.Check(pluralRule6B(float64(-3)), Equals, 3)
	c.Check(pluralRule6B(float64(3)), Equals, 3)

	// fifth form
	c.Check(pluralRule6B(float64(-6)), Equals, 4)
	c.Check(pluralRule6B(float64(6)), Equals, 4)

	// sixth form
	c.Check(pluralRule6B(float64(0.5)), Equals, 5)
	c.Check(pluralRule6B(float64(4)), Equals, 5)
	c.Check(pluralRule6B(float64(5)), Equals, 5)
	c.Check(pluralRule6B(float64(7)), Equals, 5)
	c.Check(pluralRule6B(float64(8)), Equals, 5)
	c.Check(pluralRule6B(float64(9)), Equals, 5)
	c.Check(pluralRule6B(float64(10)), Equals, 5)
	c.Check(pluralRule6B(float64(11)), Equals, 5)
	c.Check(pluralRule6B(float64(12)), Equals, 5)
	c.Check(pluralRule6B(float64(13)), Equals, 5)
	c.Check(pluralRule6B(float64(16)), Equals, 5)
}
