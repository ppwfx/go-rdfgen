package schema

type APIReferenceTrait struct {

	// Library file name e.g., mscorlib.dll, system.web.dll.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Assembly string `json:"assembly,omitempty" jsonld:"http://schema.org/assembly"`

	// Associated product/technology version. e.g., .NET Framework 4.5.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	AssemblyVersion string `json:"assemblyVersion,omitempty" jsonld:"http://schema.org/assemblyVersion"`

	// Indicates whether API is managed or unmanaged.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	ProgrammingModel string `json:"programmingModel,omitempty" jsonld:"http://schema.org/programmingModel"`

	// Type of app development: phone, Metro style, desktop, XBox, etc.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	TargetPlatform string `json:"targetPlatform,omitempty" jsonld:"http://schema.org/targetPlatform"`

	// Library file name e.g., mscorlib.dll, system.web.dll.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	ExecutableLibraryName string `json:"executableLibraryName,omitempty" jsonld:"http://schema.org/executableLibraryName"`

}

type APIReference struct {
	MetaTrait
	APIReferenceTrait
	TechArticleTrait
	ArticleTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewAPIReference() (x APIReference) {
	x.Type = "http://schema.org/APIReference"

	return
}
