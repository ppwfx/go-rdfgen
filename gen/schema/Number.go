package schema

type NumberTrait struct {

}

type Number struct {
	MetaTrait
	NumberTrait
	AdditionalTrait
}

func NewNumber() (x Number) {
	x.Type = "http://schema.org/Number"

	return
}
