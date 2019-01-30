package schema

type DataDownloadTrait struct {

	// A technique or technology used in a <a class=\"localLink\" href=\"http://schema.org/Dataset\">Dataset</a> (or <a class=\"localLink\" href=\"http://schema.org/DataDownload\">DataDownload</a>, <a class=\"localLink\" href=\"http://schema.org/DataCatalog\">DataCatalog</a>),\ncorresponding to the method used for measuring the corresponding variable(s) (described using <a class=\"localLink\" href=\"http://schema.org/variableMeasured\">variableMeasured</a>). This is oriented towards scientific and scholarly dataset publication but may have broader applicability; it is not intended as a full representation of measurement, but rather as a high level summary for dataset discovery.<br/><br/>\n\nFor example, if <a class=\"localLink\" href=\"http://schema.org/variableMeasured\">variableMeasured</a> is: molecule concentration, <a class=\"localLink\" href=\"http://schema.org/measurementTechnique\">measurementTechnique</a> could be: \"mass spectrometry\" or \"nmr spectroscopy\" or \"colorimetry\" or \"immunofluorescence\".<br/><br/>\n\nIf the <a class=\"localLink\" href=\"http://schema.org/variableMeasured\">variableMeasured</a> is \"depression rating\", the <a class=\"localLink\" href=\"http://schema.org/measurementTechnique\">measurementTechnique</a> could be \"Zung Scale\" or \"HAM-D\" or \"Beck Depression Inventory\".<br/><br/>\n\nIf there are several <a class=\"localLink\" href=\"http://schema.org/variableMeasured\">variableMeasured</a> properties recorded for some given data object, use a <a class=\"localLink\" href=\"http://schema.org/PropertyValue\">PropertyValue</a> for each <a class=\"localLink\" href=\"http://schema.org/variableMeasured\">variableMeasured</a> and attach the corresponding <a class=\"localLink\" href=\"http://schema.org/measurementTechnique\">measurementTechnique</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	MeasurementTechnique interface{} `json:"measurementTechnique,omitempty" jsonld:"http://schema.org/measurementTechnique"`

}

type DataDownload struct {
	MetaTrait
	DataDownloadTrait
	MediaObjectTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewDataDownload() (x DataDownload) {
	x.Type = "http://schema.org/DataDownload"

	return
}
