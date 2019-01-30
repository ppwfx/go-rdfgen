package schema

type TableTrait struct {

}

type Table struct {
	MetaTrait
	TableTrait
	WebPageElementTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewTable() (x Table) {
	x.Type = "http://schema.org/Table"

	return
}
