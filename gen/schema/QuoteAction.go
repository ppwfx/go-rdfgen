package schema

type QuoteActionTrait struct {

}

type QuoteAction struct {
	MetaTrait
	QuoteActionTrait
	TradeActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewQuoteAction() (x QuoteAction) {
	x.Type = "http://schema.org/QuoteAction"

	return
}
