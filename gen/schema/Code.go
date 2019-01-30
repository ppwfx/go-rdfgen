package schema

type CodeTrait struct {

}

type Code struct {
	MetaTrait
	CodeTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewCode() (x Code) {
	x.Type = "http://schema.org/Code"

	return
}
