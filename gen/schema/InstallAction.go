package schema

type InstallActionTrait struct {

}

type InstallAction struct {
	MetaTrait
	InstallActionTrait
	ConsumeActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewInstallAction() (x InstallAction) {
	x.Type = "http://schema.org/InstallAction"

	return
}
