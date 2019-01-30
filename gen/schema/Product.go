package schema

type ProductTrait struct {

	// The overall rating, based on a collection of reviews or ratings, of the item.
	//
	// RangeIncludes:
	// - http://schema.org/AggregateRating
	//
	AggregateRating *AggregateRating `json:"aggregateRating,omitempty" jsonld:"http://schema.org/aggregateRating"`

	// An intended audience, i.e. a group for whom something was created.
	//
	// RangeIncludes:
	// - http://schema.org/Audience
	//
	Audience *Audience `json:"audience,omitempty" jsonld:"http://schema.org/audience"`

	// An award won by or for this item.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Award string `json:"award,omitempty" jsonld:"http://schema.org/award"`

	// Awards won by or for this item.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Awards string `json:"awards,omitempty" jsonld:"http://schema.org/awards"`

	// The height of the item.
	//
	// RangeIncludes:
	// - http://schema.org/Distance
	// - http://schema.org/QuantitativeValue
	//
	Height interface{} `json:"height,omitempty" jsonld:"http://schema.org/height"`

	// A material that something is made from, e.g. leather, wool, cotton, paper.
	//
	// RangeIncludes:
	// - http://schema.org/Product
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	Material interface{} `json:"material,omitempty" jsonld:"http://schema.org/material"`

	// An offer to provide this item&#x2014;for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
	//
	// RangeIncludes:
	// - http://schema.org/Offer
	//
	Offers *Offer `json:"offers,omitempty" jsonld:"http://schema.org/offers"`

	// A review of the item.
	//
	// RangeIncludes:
	// - http://schema.org/Review
	//
	Review *Review `json:"review,omitempty" jsonld:"http://schema.org/review"`

	// Review of the item.
	//
	// RangeIncludes:
	// - http://schema.org/Review
	//
	Reviews *Review `json:"reviews,omitempty" jsonld:"http://schema.org/reviews"`

	// The width of the item.
	//
	// RangeIncludes:
	// - http://schema.org/Distance
	// - http://schema.org/QuantitativeValue
	//
	Width interface{} `json:"width,omitempty" jsonld:"http://schema.org/width"`

	// The GTIN-12 code of the product, or the product to which the offer refers. The GTIN-12 is the 12-digit GS1 Identification Key composed of a U.P.C. Company Prefix, Item Reference, and Check Digit used to identify trade items. See <a href=\"http://www.gs1.org/barcodes/technical/idkeys/gtin\">GS1 GTIN Summary</a> for more details.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Gtin12 string `json:"gtin12,omitempty" jsonld:"http://schema.org/gtin12"`

	// The GTIN-13 code of the product, or the product to which the offer refers. This is equivalent to 13-digit ISBN codes and EAN UCC-13. Former 12-digit UPC codes can be converted into a GTIN-13 code by simply adding a preceeding zero. See <a href=\"http://www.gs1.org/barcodes/technical/idkeys/gtin\">GS1 GTIN Summary</a> for more details.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Gtin13 string `json:"gtin13,omitempty" jsonld:"http://schema.org/gtin13"`

	// The GTIN-14 code of the product, or the product to which the offer refers. See <a href=\"http://www.gs1.org/barcodes/technical/idkeys/gtin\">GS1 GTIN Summary</a> for more details.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Gtin14 string `json:"gtin14,omitempty" jsonld:"http://schema.org/gtin14"`

	// The <a href=\"http://apps.gs1.org/GDD/glossary/Pages/GTIN-8.aspx\">GTIN-8</a> code of the product, or the product to which the offer refers. This code is also known as EAN/UCC-8 or 8-digit EAN. See <a href=\"http://www.gs1.org/barcodes/technical/idkeys/gtin\">GS1 GTIN Summary</a> for more details.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Gtin8 string `json:"gtin8,omitempty" jsonld:"http://schema.org/gtin8"`

	// A predefined value from OfferItemCondition or a textual description of the condition of the product or service, or the products or services included in the offer.
	//
	// RangeIncludes:
	// - http://schema.org/OfferItemCondition
	//
	ItemCondition *OfferItemCondition `json:"itemCondition,omitempty" jsonld:"http://schema.org/itemCondition"`

	// The Manufacturer Part Number (MPN) of the product, or the product to which the offer refers.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Mpn string `json:"mpn,omitempty" jsonld:"http://schema.org/mpn"`

	// The Stock Keeping Unit (SKU), i.e. a merchant-specific identifier for a product or service, or the product to which the offer refers.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Sku string `json:"sku,omitempty" jsonld:"http://schema.org/sku"`

	// A category for the item. Greater signs or slashes can be used to informally indicate a category hierarchy.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Thing
	// - http://schema.org/PhysicalActivityCategory
	//
	Category interface{} `json:"category,omitempty" jsonld:"http://schema.org/category"`

	// A property-value pair representing an additional characteristics of the entitity, e.g. a product feature or another characteristic for which there is no matching property in schema.org.<br/><br/>\n\nNote: Publishers should be aware that applications designed to use specific schema.org properties (e.g. http://schema.org/width, http://schema.org/color, http://schema.org/gtin13, ...) will typically expect such data to be provided using those properties, rather than using the generic property/value mechanism.
	//
	// RangeIncludes:
	// - http://schema.org/PropertyValue
	//
	AdditionalProperty *PropertyValue `json:"additionalProperty,omitempty" jsonld:"http://schema.org/additionalProperty"`

	// An associated logo.
	//
	// RangeIncludes:
	// - http://schema.org/ImageObject
	// - http://schema.org/URL
	//
	Logo interface{} `json:"logo,omitempty" jsonld:"http://schema.org/logo"`

	// The brand(s) associated with a product or service, or the brand(s) maintained by an organization or business person.
	//
	// RangeIncludes:
	// - http://schema.org/Brand
	// - http://schema.org/Organization
	//
	Brand interface{} `json:"brand,omitempty" jsonld:"http://schema.org/brand"`

	// The weight of the product or person.
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	//
	Weight *QuantitativeValue `json:"weight,omitempty" jsonld:"http://schema.org/weight"`

	// A pointer to another, somehow related product (or multiple products).
	//
	// RangeIncludes:
	// - http://schema.org/Product
	// - http://schema.org/Service
	//
	IsRelatedTo interface{} `json:"isRelatedTo,omitempty" jsonld:"http://schema.org/isRelatedTo"`

	// A pointer to another, functionally similar product (or multiple products).
	//
	// RangeIncludes:
	// - http://schema.org/Product
	// - http://schema.org/Service
	//
	IsSimilarTo interface{} `json:"isSimilarTo,omitempty" jsonld:"http://schema.org/isSimilarTo"`

	// The color of the product.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Color string `json:"color,omitempty" jsonld:"http://schema.org/color"`

	// The depth of the item.
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	// - http://schema.org/Distance
	//
	Depth interface{} `json:"depth,omitempty" jsonld:"http://schema.org/depth"`

	// A pointer to another product (or multiple products) for which this product is an accessory or spare part.
	//
	// RangeIncludes:
	// - http://schema.org/Product
	//
	IsAccessoryOrSparePartFor *Product `json:"isAccessoryOrSparePartFor,omitempty" jsonld:"http://schema.org/isAccessoryOrSparePartFor"`

	// A pointer to another product (or multiple products) for which this product is a consumable.
	//
	// RangeIncludes:
	// - http://schema.org/Product
	//
	IsConsumableFor *Product `json:"isConsumableFor,omitempty" jsonld:"http://schema.org/isConsumableFor"`

	// The manufacturer of the product.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	//
	Manufacturer *Organization `json:"manufacturer,omitempty" jsonld:"http://schema.org/manufacturer"`

	// The model of the product. Use with the URL of a ProductModel or a textual representation of the model identifier. The URL of the ProductModel can be from an external source. It is recommended to additionally provide strong product identifiers via the gtin8/gtin13/gtin14 and mpn properties.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/ProductModel
	//
	Model interface{} `json:"model,omitempty" jsonld:"http://schema.org/model"`

	// The product identifier, such as ISBN. For example: <code>meta itemprop=\"productID\" content=\"isbn:123-456-789\"</code>.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	ProductID string `json:"productID,omitempty" jsonld:"http://schema.org/productID"`

	// The date of production of the item, e.g. vehicle.
	//
	// RangeIncludes:
	// - http://schema.org/Date
	//
	ProductionDate *Date `json:"productionDate,omitempty" jsonld:"http://schema.org/productionDate"`

	// The date the item e.g. vehicle was purchased by the current owner.
	//
	// RangeIncludes:
	// - http://schema.org/Date
	//
	PurchaseDate *Date `json:"purchaseDate,omitempty" jsonld:"http://schema.org/purchaseDate"`

	// The release date of a product or product model. This can be used to distinguish the exact variant of a product.
	//
	// RangeIncludes:
	// - http://schema.org/Date
	//
	ReleaseDate *Date `json:"releaseDate,omitempty" jsonld:"http://schema.org/releaseDate"`

}

type Product struct {
	MetaTrait
	ProductTrait
	ThingTrait
	AdditionalTrait
}

func NewProduct() (x Product) {
	x.Type = "http://schema.org/Product"

	return
}
