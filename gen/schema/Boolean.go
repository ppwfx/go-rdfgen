package schema

type BooleanTrait struct {

}

type Boolean struct {
	MetaTrait
	BooleanTrait
	AdditionalTrait
}

func NewBoolean() (x Boolean) {
	x.Type = "http://schema.org/Boolean"

	return
}
