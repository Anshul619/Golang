package main

import (
	"fmt"
	"strconv"
)

type numberBoxContaining struct {
	number int
}

func (n numberBoxContaining) Number() int {
	return n.number
}

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %0.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %0.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

type AnotherFancyNumber struct {
	n string
}

func (a AnotherFancyNumber) Value() string {
	return a.n
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	switch fnb.(type) {
	case FancyNumber:
		i, _ := strconv.Atoi(fnb.Value())
		return i
	}
	return 0
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	return fmt.Sprintf("This is a fancy box containing the number %0.1f", float64(ExtractFancyNumber(fnb)))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch i.(type) {
	case int:
		return DescribeNumber(float64(i.(int)))
	case float64:
		return DescribeNumber(i.(float64))
	case NumberBox:
		return DescribeNumberBox(i.(NumberBox))
	case FancyNumberBox:
		return DescribeFancyNumberBox(i.(FancyNumberBox))
	}

	return "Return to sender"
}

func main() {
	fmt.Println(DescribeNumber(-12.345))
	fmt.Println(DescribeNumberBox(numberBoxContaining{12}))
	fmt.Println(ExtractFancyNumber(FancyNumber{"10"}))
	// Output: 10
	fmt.Println(ExtractFancyNumber(AnotherFancyNumber{"4"}))
	// Output: 0
	fmt.Println(DescribeFancyNumberBox(FancyNumber{"10"}))
	// Output: This is a fancy box containing the number 10.0
	fmt.Println(DescribeFancyNumberBox(AnotherFancyNumber{"4"}))
	// Output: This is a fancy box containing the number 0.0

	fmt.Println(DescribeAnything(numberBoxContaining{12}))
	// Output: This is a box containing the number 12.3
	fmt.Println(DescribeAnything("some string"))
	// Output: Return to sender

	fmt.Println(DescribeAnything(42))
}
