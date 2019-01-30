package schema

type ComputerLanguageTrait struct {

}

type ComputerLanguage struct {
	MetaTrait
	ComputerLanguageTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewComputerLanguage() (x ComputerLanguage) {
	x.Type = "http://schema.org/ComputerLanguage"

	return
}
