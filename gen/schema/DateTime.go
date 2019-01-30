package schema

type DateTimeTrait struct {

}

type DateTime struct {
	MetaTrait
	DateTimeTrait
	AdditionalTrait
}

func NewDateTime() (x DateTime) {
	x.Type = "http://schema.org/DateTime"

	return
}
