package schema

type LanguageTrait struct {

}

type Language struct {
	MetaTrait
	LanguageTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewLanguage() (x Language) {
	x.Type = "http://schema.org/Language"

	return
}
