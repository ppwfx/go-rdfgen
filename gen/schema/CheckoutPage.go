package schema

type CheckoutPageTrait struct {

}

type CheckoutPage struct {
	MetaTrait
	CheckoutPageTrait
	WebPageTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewCheckoutPage() (x CheckoutPage) {
	x.Type = "http://schema.org/CheckoutPage"

	return
}
