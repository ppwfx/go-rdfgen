package schema

type TextTrait struct {

}

type Text struct {
	MetaTrait
	TextTrait
	AdditionalTrait
}

func NewText() (x Text) {
	x.Type = "http://schema.org/Text"

	return
}
