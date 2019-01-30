package schema

type AlignmentObjectTrait struct {

	// A category of alignment between the learning resource and the framework node. Recommended values include: 'assesses', 'teaches', 'requires', 'textComplexity', 'readingLevel', 'educationalSubject', and 'educationalLevel'.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	AlignmentType string `json:"alignmentType,omitempty" jsonld:"http://schema.org/alignmentType"`

	// The framework to which the resource being described is aligned.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	EducationalFramework string `json:"educationalFramework,omitempty" jsonld:"http://schema.org/educationalFramework"`

	// The description of a node in an established educational framework.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	TargetDescription string `json:"targetDescription,omitempty" jsonld:"http://schema.org/targetDescription"`

	// The name of a node in an established educational framework.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	TargetName string `json:"targetName,omitempty" jsonld:"http://schema.org/targetName"`

	// The URL of a node in an established educational framework.
	//
	// RangeIncludes:
	// - http://schema.org/URL
	//
	TargetUrl string `json:"targetUrl,omitempty" jsonld:"http://schema.org/targetUrl"`

}

type AlignmentObject struct {
	MetaTrait
	AlignmentObjectTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewAlignmentObject() (x AlignmentObject) {
	x.Type = "http://schema.org/AlignmentObject"

	return
}
