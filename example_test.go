package i18n_test

import (
	"fmt"
	"github.com/vube/i18n"
)

func ExampleNewTranslatorFactory() {
	// creates a new Factory using "en" as the global fallback locale
	f, _ := i18n.NewTranslatorFactory(
		[]string{"data/rules"},
		[]string{"data/messages"},
		"en",
	)

	_ = f
}

func ExampleTranslatorFactory_GetTranslator() {
	f, _ := i18n.NewTranslatorFactory(
		[]string{"data/rules"},
		[]string{"data/messages"},
		"en",
	)

	// gets a translator for the Canadian French locale
	tFrCa, _ := f.GetTranslator("fr-ca")

	_ = tFrCa
}

func ExampleTranslator_Translate() {
	f, _ := i18n.NewTranslatorFactory(
		[]string{"data/rules"},
		[]string{"data/messages"},
		"en",
	)

	tEn, _ := f.GetTranslator("en")

	// performs 2 translations, one with a substitution
	t1, _ := tEn.Translate("WELCOME", map[string]string{})
	t2, _ := tEn.Translate("WELCOME_USER", map[string]string{"user": "Mother Goose"})

	fmt.Printf("Basic Translation : %s\n", t1)
	fmt.Printf("Substitution      : %s\n", t2)

	// Output:
	// Basic Translation : Welcome!
	// Substitution      : Welcome, Mother Goose!
}

func ExampleTranslator_Pluralize() {
	f, _ := i18n.NewTranslatorFactory(
		[]string{"data/rules"},
		[]string{"data/messages"},
		"en",
	)

	tFr, _ := f.GetTranslator("fr")

	// performs 4 plural translations
	p1, _ := tFr.Pluralize("UNIT_DAY", 0, "0")
	p2, _ := tFr.Pluralize("UNIT_DAY", 0.5, "0,5")
	p3, _ := tFr.Pluralize("UNIT_DAY", 1, "one")
	p4, _ := tFr.Pluralize("UNIT_DAY", 2000, "2K")

	fmt.Printf("Plurilization : %s\n", p1)
	fmt.Printf("Plurilization : %s\n", p2)
	fmt.Printf("Plurilization : %s\n", p3)
	fmt.Printf("Plurilization : %s\n", p4)

	// Output:
	// Plurilization : 0 jour
	// Plurilization : 0,5 jour
	// Plurilization : one jour
	// Plurilization : 2K jours
}

func ExampleTranslator_FormatCurrency() {
	f, _ := i18n.NewTranslatorFactory(
		[]string{"data/rules"},
		[]string{"data/messages"},
		"en",
	)

	tEn, _ := f.GetTranslator("en")

	// performs 2 currency formats - one positive, one negative
	c1, _ := tEn.FormatCurrency(12345000000000.6789, "USD")
	c2, _ := tEn.FormatCurrency(-12345000000000.6789, "USD")

	fmt.Printf("Currency : %s\n", c1)
	fmt.Printf("Currency : %s\n", c2)

	// Output:
	// Currency : $12,345,000,000,000.68
	// Currency : ($12,345,000,000,000.68)
}

func ExampleTranslator_FormatNumber() {
	f, _ := i18n.NewTranslatorFactory(
		[]string{"data/rules"},
		[]string{"data/messages"},
		"en",
	)

	tEn, _ := f.GetTranslator("en")

	// performs 2 number formats - one positive, one negative
	n1 := tEn.FormatNumber(12345000000000.6789)
	n2 := tEn.FormatNumber(-12345000000000.6789)

	fmt.Printf("Number : %s\n", n1)
	fmt.Printf("Number : %s\n", n2)

	// Output:
	// Number : 12,345,000,000,000.679
	// Number : -12,345,000,000,000.679
}

func ExampleTranslator_FormatPercent() {
	f, _ := i18n.NewTranslatorFactory(
		[]string{"data/rules"},
		[]string{"data/messages"},
		"en",
	)

	tEn, _ := f.GetTranslator("en")

	// performs 3 percent formats
	p1 := tEn.FormatPercent(0.01234)
	p2 := tEn.FormatPercent(0.5678)
	p3 := tEn.FormatPercent(12.34)

	fmt.Printf("Percent : %s\n", p1)
	fmt.Printf("Percent : %s\n", p2)
	fmt.Printf("Percent : %s\n", p3)

	// Output:
	// Percent : 1%
	// Percent : 57%
	// Percent : 1,234%
}

type Food struct {
	Name string
}

func ExampleTranslator_Sort() {
	f, _ := i18n.NewTranslatorFactory(
		[]string{"data/rules"},
		[]string{"data/messages"},
		"en",
	)

	tEn, _ := f.GetTranslator("en")

	toSort := []interface{}{
		Food{Name: "apple"},
		Food{Name: "beet"},
		Food{Name: "carrot"},
		Food{Name: "ȧpricot"},
		Food{Name: "ḃanana"},
		Food{Name: "ċlementine"},
	}

	fmt.Printf("Before Sort : %v\n", toSort)

	// sorts the food list
	tEn.Sort(toSort, func(i interface{}) string {
		if food, ok := i.(Food); ok {
			return food.Name
		}
		return ""
	})
	fmt.Printf("After Sort  : %v\n", toSort)

	// Output:
	// Before Sort : [{apple} {beet} {carrot} {ȧpricot} {ḃanana} {ċlementine}]
	// After Sort  : [{apple} {ȧpricot} {ḃanana} {beet} {carrot} {ċlementine}]
}

func ExampleSortUniversal() {

	toSort := []interface{}{
		Food{Name: "apple"},
		Food{Name: "beet"},
		Food{Name: "carrot"},
		Food{Name: "ȧpricot"},
		Food{Name: "ḃanana"},
		Food{Name: "ċlementine"},
	}

	fmt.Printf("Before Sort : %v\n", toSort)

	// sorts the list
	i18n.SortUniversal(toSort, func(i interface{}) string {
		if food, ok := i.(Food); ok {
			return food.Name
		}
		return ""
	})
	fmt.Printf("After Sort  : %v\n", toSort)

	// Output:
	// Before Sort : [{apple} {beet} {carrot} {ȧpricot} {ḃanana} {ċlementine}]
	// After Sort  : [{apple} {ȧpricot} {beet} {ḃanana} {carrot} {ċlementine}]
}

func ExampleSortLocal() {

	toSort := []interface{}{
		Food{Name: "apple"},
		Food{Name: "beet"},
		Food{Name: "carrot"},
		Food{Name: "ȧpricot"},
		Food{Name: "ḃanana"},
		Food{Name: "ċlementine"},
	}

	fmt.Printf("Before Sort : %v\n", toSort)

	// sorts the list
	i18n.SortLocal("en", toSort, func(i interface{}) string {
		if food, ok := i.(Food); ok {
			return food.Name
		}
		return ""
	})
	fmt.Printf("After Sort  : %v\n", toSort)

	// Output:
	// Before Sort : [{apple} {beet} {carrot} {ȧpricot} {ḃanana} {ċlementine}]
	// After Sort  : [{apple} {ȧpricot} {ḃanana} {beet} {carrot} {ċlementine}]
}
