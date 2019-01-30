package schema

type WatchActionTrait struct {

}

type WatchAction struct {
	MetaTrait
	WatchActionTrait
	ConsumeActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewWatchAction() (x WatchAction) {
	x.Type = "http://schema.org/WatchAction"

	return
}
