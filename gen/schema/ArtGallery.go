package schema

type ArtGalleryTrait struct {

}

type ArtGallery struct {
	MetaTrait
	ArtGalleryTrait
	EntertainmentBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewArtGallery() (x ArtGallery) {
	x.Type = "http://schema.org/ArtGallery"

	return
}
