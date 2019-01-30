package schema

type SoftwareApplicationTrait struct {

	// A link to a screenshot image of the app.
	//
	// RangeIncludes:
	// - http://schema.org/ImageObject
	// - http://schema.org/URL
	//
	Screenshot interface{} `json:"screenshot,omitempty" jsonld:"http://schema.org/screenshot"`

	// Type of software application, e.g. 'Game, Multimedia'.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	ApplicationCategory interface{} `json:"applicationCategory,omitempty" jsonld:"http://schema.org/applicationCategory"`

	// Subcategory of the application, e.g. 'Arcade Game'.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	ApplicationSubCategory interface{} `json:"applicationSubCategory,omitempty" jsonld:"http://schema.org/applicationSubCategory"`

	// The name of the application suite to which the application belongs (e.g. Excel belongs to Office).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	ApplicationSuite string `json:"applicationSuite,omitempty" jsonld:"http://schema.org/applicationSuite"`

	// Device required to run the application. Used in cases where a specific make/model is required to run the application.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	AvailableOnDevice string `json:"availableOnDevice,omitempty" jsonld:"http://schema.org/availableOnDevice"`

	// Countries for which the application is not supported. You can also provide the two-letter ISO 3166-1 alpha-2 country code.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	CountriesNotSupported string `json:"countriesNotSupported,omitempty" jsonld:"http://schema.org/countriesNotSupported"`

	// Countries for which the application is supported. You can also provide the two-letter ISO 3166-1 alpha-2 country code.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	CountriesSupported string `json:"countriesSupported,omitempty" jsonld:"http://schema.org/countriesSupported"`

	// Device required to run the application. Used in cases where a specific make/model is required to run the application.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Device string `json:"device,omitempty" jsonld:"http://schema.org/device"`

	// If the file can be downloaded, URL to download the binary.
	//
	// RangeIncludes:
	// - http://schema.org/URL
	//
	DownloadUrl string `json:"downloadUrl,omitempty" jsonld:"http://schema.org/downloadUrl"`

	// Features or modules provided by this application (and possibly required by other applications).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	FeatureList interface{} `json:"featureList,omitempty" jsonld:"http://schema.org/featureList"`

	// Size of the application / package (e.g. 18MB). In the absence of a unit (MB, KB etc.), KB will be assumed.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	FileSize string `json:"fileSize,omitempty" jsonld:"http://schema.org/fileSize"`

	// URL at which the app may be installed, if different from the URL of the item.
	//
	// RangeIncludes:
	// - http://schema.org/URL
	//
	InstallUrl string `json:"installUrl,omitempty" jsonld:"http://schema.org/installUrl"`

	// Minimum memory requirements.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	MemoryRequirements interface{} `json:"memoryRequirements,omitempty" jsonld:"http://schema.org/memoryRequirements"`

	// Operating systems supported (Windows 7, OSX 10.6, Android 1.6).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	OperatingSystem string `json:"operatingSystem,omitempty" jsonld:"http://schema.org/operatingSystem"`

	// Permission(s) required to run the app (for example, a mobile app may require full internet access or may run only on wifi).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Permissions string `json:"permissions,omitempty" jsonld:"http://schema.org/permissions"`

	// Processor architecture required to run the application (e.g. IA64).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	ProcessorRequirements string `json:"processorRequirements,omitempty" jsonld:"http://schema.org/processorRequirements"`

	// Description of what changed in this version.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	ReleaseNotes interface{} `json:"releaseNotes,omitempty" jsonld:"http://schema.org/releaseNotes"`

	// Component dependency requirements for application. This includes runtime environments and shared libraries that are not included in the application distribution package, but required to run the application (Examples: DirectX, Java or .NET runtime).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	Requirements interface{} `json:"requirements,omitempty" jsonld:"http://schema.org/requirements"`

	// Additional content for a software application.
	//
	// RangeIncludes:
	// - http://schema.org/SoftwareApplication
	//
	SoftwareAddOn *SoftwareApplication `json:"softwareAddOn,omitempty" jsonld:"http://schema.org/softwareAddOn"`

	// Software application help.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	//
	SoftwareHelp *CreativeWork `json:"softwareHelp,omitempty" jsonld:"http://schema.org/softwareHelp"`

	// Component dependency requirements for application. This includes runtime environments and shared libraries that are not included in the application distribution package, but required to run the application (Examples: DirectX, Java or .NET runtime).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	SoftwareRequirements interface{} `json:"softwareRequirements,omitempty" jsonld:"http://schema.org/softwareRequirements"`

	// Version of the software instance.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	SoftwareVersion string `json:"softwareVersion,omitempty" jsonld:"http://schema.org/softwareVersion"`

	// Storage requirements (free space required).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	StorageRequirements interface{} `json:"storageRequirements,omitempty" jsonld:"http://schema.org/storageRequirements"`

	// Supporting data for a SoftwareApplication.
	//
	// RangeIncludes:
	// - http://schema.org/DataFeed
	//
	SupportingData *DataFeed `json:"supportingData,omitempty" jsonld:"http://schema.org/supportingData"`

}

type SoftwareApplication struct {
	MetaTrait
	SoftwareApplicationTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewSoftwareApplication() (x SoftwareApplication) {
	x.Type = "http://schema.org/SoftwareApplication"

	return
}
