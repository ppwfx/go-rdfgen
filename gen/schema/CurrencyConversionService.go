package schema

type CurrencyConversionServiceTrait struct {

}

type CurrencyConversionService struct {
	MetaTrait
	CurrencyConversionServiceTrait
	FinancialProductTrait
	ServiceTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewCurrencyConversionService() (x CurrencyConversionService) {
	x.Type = "http://schema.org/CurrencyConversionService"

	return
}
