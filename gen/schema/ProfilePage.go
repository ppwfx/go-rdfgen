package schema

type ProfilePageTrait struct {

}

type ProfilePage struct {
	MetaTrait
	ProfilePageTrait
	WebPageTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewProfilePage() (x ProfilePage) {
	x.Type = "http://schema.org/ProfilePage"

	return
}
