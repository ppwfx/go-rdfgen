package schema

type TechArticleTrait struct {

	// Prerequisites needed to fulfill steps in article.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Dependencies string `json:"dependencies,omitempty" jsonld:"http://schema.org/dependencies"`

	// Proficiency needed for this content; expected values: 'Beginner', 'Expert'.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	ProficiencyLevel string `json:"proficiencyLevel,omitempty" jsonld:"http://schema.org/proficiencyLevel"`

}

type TechArticle struct {
	MetaTrait
	TechArticleTrait
	ArticleTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewTechArticle() (x TechArticle) {
	x.Type = "http://schema.org/TechArticle"

	return
}
