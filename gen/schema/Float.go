package schema

type FloatTrait struct {

}

type Float struct {
	MetaTrait
	FloatTrait
	NumberTrait
	AdditionalTrait
}

func NewFloat() (x Float) {
	x.Type = "http://schema.org/Float"

	return
}
