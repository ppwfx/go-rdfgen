package schema

type DataTypeTrait struct {

}

type DataType struct {
	MetaTrait
	DataTypeTrait
	AdditionalTrait
}

func NewDataType() (x DataType) {
	x.Type = "http://schema.org/DataType"

	return
}
