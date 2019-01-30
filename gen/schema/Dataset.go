package schema

type DatasetTrait struct {

	// A technique or technology used in a <a class=\"localLink\" href=\"http://schema.org/Dataset\">Dataset</a> (or <a class=\"localLink\" href=\"http://schema.org/DataDownload\">DataDownload</a>, <a class=\"localLink\" href=\"http://schema.org/DataCatalog\">DataCatalog</a>),\ncorresponding to the method used for measuring the corresponding variable(s) (described using <a class=\"localLink\" href=\"http://schema.org/variableMeasured\">variableMeasured</a>). This is oriented towards scientific and scholarly dataset publication but may have broader applicability; it is not intended as a full representation of measurement, but rather as a high level summary for dataset discovery.<br/><br/>\n\nFor example, if <a class=\"localLink\" href=\"http://schema.org/variableMeasured\">variableMeasured</a> is: molecule concentration, <a class=\"localLink\" href=\"http://schema.org/measurementTechnique\">measurementTechnique</a> could be: \"mass spectrometry\" or \"nmr spectroscopy\" or \"colorimetry\" or \"immunofluorescence\".<br/><br/>\n\nIf the <a class=\"localLink\" href=\"http://schema.org/variableMeasured\">variableMeasured</a> is \"depression rating\", the <a class=\"localLink\" href=\"http://schema.org/measurementTechnique\">measurementTechnique</a> could be \"Zung Scale\" or \"HAM-D\" or \"Beck Depression Inventory\".<br/><br/>\n\nIf there are several <a class=\"localLink\" href=\"http://schema.org/variableMeasured\">variableMeasured</a> properties recorded for some given data object, use a <a class=\"localLink\" href=\"http://schema.org/PropertyValue\">PropertyValue</a> for each <a class=\"localLink\" href=\"http://schema.org/variableMeasured\">variableMeasured</a> and attach the corresponding <a class=\"localLink\" href=\"http://schema.org/measurementTechnique\">measurementTechnique</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	MeasurementTechnique interface{} `json:"measurementTechnique,omitempty" jsonld:"http://schema.org/measurementTechnique"`

	// The International Standard Serial Number (ISSN) that identifies this serial publication. You can repeat this property to identify different formats of, or the linking ISSN (ISSN-L) for, this serial publication.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Issn string `json:"issn,omitempty" jsonld:"http://schema.org/issn"`

	// The range of temporal applicability of a dataset, e.g. for a 2011 census dataset, the year 2011 (in ISO 8601 time interval format).
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	DatasetTimeInterval *DateTime `json:"datasetTimeInterval,omitempty" jsonld:"http://schema.org/datasetTimeInterval"`

	// The range of temporal applicability of a dataset, e.g. for a 2011 census dataset, the year 2011 (in ISO 8601 time interval format).
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	Temporal *DateTime `json:"temporal,omitempty" jsonld:"http://schema.org/temporal"`

	// The variableMeasured property can indicate (repeated as necessary) the  variables that are measured in some dataset, either described as text or as pairs of identifier and description using PropertyValue.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/PropertyValue
	//
	VariableMeasured interface{} `json:"variableMeasured,omitempty" jsonld:"http://schema.org/variableMeasured"`

	// A data catalog which contains this dataset.
	//
	// RangeIncludes:
	// - http://schema.org/DataCatalog
	//
	Catalog *DataCatalog `json:"catalog,omitempty" jsonld:"http://schema.org/catalog"`

	// A downloadable form of this dataset, at a specific location, in a specific format.
	//
	// RangeIncludes:
	// - http://schema.org/DataDownload
	//
	Distribution *DataDownload `json:"distribution,omitempty" jsonld:"http://schema.org/distribution"`

	// A data catalog which contains this dataset (this property was previously 'catalog', preferred name is now 'includedInDataCatalog').
	//
	// RangeIncludes:
	// - http://schema.org/DataCatalog
	//
	IncludedDataCatalog *DataCatalog `json:"includedDataCatalog,omitempty" jsonld:"http://schema.org/includedDataCatalog"`

	// A data catalog which contains this dataset.
	//
	// RangeIncludes:
	// - http://schema.org/DataCatalog
	//
	IncludedInDataCatalog *DataCatalog `json:"includedInDataCatalog,omitempty" jsonld:"http://schema.org/includedInDataCatalog"`

	// The range of spatial applicability of a dataset, e.g. for a dataset of New York weather, the state of New York.
	//
	// RangeIncludes:
	// - http://schema.org/Place
	//
	Spatial *Place `json:"spatial,omitempty" jsonld:"http://schema.org/spatial"`

}

type Dataset struct {
	MetaTrait
	DatasetTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewDataset() (x Dataset) {
	x.Type = "http://schema.org/Dataset"

	return
}
