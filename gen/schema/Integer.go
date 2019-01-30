package schema

type IntegerTrait struct {

}

type Integer struct {
	MetaTrait
	IntegerTrait
	NumberTrait
	AdditionalTrait
}

func NewInteger() (x Integer) {
	x.Type = "http://schema.org/Integer"

	return
}
