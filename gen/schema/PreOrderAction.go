package schema

type PreOrderActionTrait struct {

}

type PreOrderAction struct {
	MetaTrait
	PreOrderActionTrait
	TradeActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewPreOrderAction() (x PreOrderAction) {
	x.Type = "http://schema.org/PreOrderAction"

	return
}
