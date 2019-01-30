package schema

type SoftwareSourceCodeTrait struct {

	// What type of code sample: full (compile ready) solution, code snippet, inline code, scripts, template.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	SampleType string `json:"sampleType,omitempty" jsonld:"http://schema.org/sampleType"`

	// Runtime platform or script interpreter dependencies (Example - Java v1, Python2.3, .Net Framework 3.0).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Runtime string `json:"runtime,omitempty" jsonld:"http://schema.org/runtime"`

	// Runtime platform or script interpreter dependencies (Example - Java v1, Python2.3, .Net Framework 3.0).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	RuntimePlatform string `json:"runtimePlatform,omitempty" jsonld:"http://schema.org/runtimePlatform"`

	// The computer programming language.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/ComputerLanguage
	//
	ProgrammingLanguage interface{} `json:"programmingLanguage,omitempty" jsonld:"http://schema.org/programmingLanguage"`

	// What type of code sample: full (compile ready) solution, code snippet, inline code, scripts, template.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	CodeSampleType string `json:"codeSampleType,omitempty" jsonld:"http://schema.org/codeSampleType"`

	// Link to the repository where the un-compiled, human readable code and related code is located (SVN, github, CodePlex).
	//
	// RangeIncludes:
	// - http://schema.org/URL
	//
	CodeRepository string `json:"codeRepository,omitempty" jsonld:"http://schema.org/codeRepository"`

	// Target Operating System / Product to which the code applies.  If applies to several versions, just the product name can be used.
	//
	// RangeIncludes:
	// - http://schema.org/SoftwareApplication
	//
	TargetProduct *SoftwareApplication `json:"targetProduct,omitempty" jsonld:"http://schema.org/targetProduct"`

}

type SoftwareSourceCode struct {
	MetaTrait
	SoftwareSourceCodeTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewSoftwareSourceCode() (x SoftwareSourceCode) {
	x.Type = "http://schema.org/SoftwareSourceCode"

	return
}
