package schema

type DownloadActionTrait struct {

}

type DownloadAction struct {
	MetaTrait
	DownloadActionTrait
	TransferActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewDownloadAction() (x DownloadAction) {
	x.Type = "http://schema.org/DownloadAction"

	return
}
