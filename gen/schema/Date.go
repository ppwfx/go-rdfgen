package schema

type DateTrait struct {

}

type Date struct {
	MetaTrait
	DateTrait
	AdditionalTrait
}

func NewDate() (x Date) {
	x.Type = "http://schema.org/Date"

	return
}
