package schema

type URLTrait struct {

}

type URL struct {
	MetaTrait
	URLTrait
	TextTrait
	AdditionalTrait
}

func NewURL() (x URL) {
	x.Type = "http://schema.org/URL"

	return
}
