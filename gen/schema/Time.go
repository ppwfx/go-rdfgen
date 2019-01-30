package schema

type TimeTrait struct {

}

type Time struct {
	MetaTrait
	TimeTrait
	AdditionalTrait
}

func NewTime() (x Time) {
	x.Type = "http://schema.org/Time"

	return
}
