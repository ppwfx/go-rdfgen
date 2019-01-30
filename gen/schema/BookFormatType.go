package schema

type BookFormatTypeTrait struct {

}

type BookFormatType struct {
	MetaTrait
	BookFormatTypeTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewBookFormatType() (x BookFormatType) {
	x.Type = "http://schema.org/BookFormatType"

	return
}
