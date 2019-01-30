package schema

type BreadcrumbListTrait struct {

}

type BreadcrumbList struct {
	MetaTrait
	BreadcrumbListTrait
	ItemListTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewBreadcrumbList() (x BreadcrumbList) {
	x.Type = "http://schema.org/BreadcrumbList"

	return
}
