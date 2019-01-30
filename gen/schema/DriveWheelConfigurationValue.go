package schema

type DriveWheelConfigurationValueTrait struct {

}

type DriveWheelConfigurationValue struct {
	MetaTrait
	DriveWheelConfigurationValueTrait
	QualitativeValueTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewDriveWheelConfigurationValue() (x DriveWheelConfigurationValue) {
	x.Type = "http://schema.org/DriveWheelConfigurationValue"

	return
}
