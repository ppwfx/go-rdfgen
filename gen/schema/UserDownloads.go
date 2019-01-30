package schema

type UserDownloadsTrait struct {

}

type UserDownloads struct {
	MetaTrait
	UserDownloadsTrait
	UserInteractionTrait
	EventTrait
	ThingTrait
	AdditionalTrait
}

func NewUserDownloads() (x UserDownloads) {
	x.Type = "http://schema.org/UserDownloads"

	return
}
