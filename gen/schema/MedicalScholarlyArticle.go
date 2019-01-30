package schema

type MedicalScholarlyArticleTrait struct {

	// The type of the medical article, taken from the US NLM MeSH publication type catalog. See also <a href=\"http://www.nlm.nih.gov/mesh/pubtypes.html\">MeSH documentation</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	PublicationType string `json:"publicationType,omitempty" jsonld:"http://schema.org/publicationType"`

}

type MedicalScholarlyArticle struct {
	MetaTrait
	MedicalScholarlyArticleTrait
	ScholarlyArticleTrait
	ArticleTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewMedicalScholarlyArticle() (x MedicalScholarlyArticle) {
	x.Type = "http://schema.org/MedicalScholarlyArticle"

	return
}
