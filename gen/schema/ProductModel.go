package schema

type ProductModelTrait struct {

	// A pointer to a base product from which this product is a variant. It is safe to infer that the variant inherits all product features from the base model, unless defined locally. This is not transitive.
	//
	// RangeIncludes:
	// - http://schema.org/ProductModel
	//
	IsVariantOf *ProductModel `json:"isVariantOf,omitempty" jsonld:"http://schema.org/isVariantOf"`

	// A pointer from a previous, often discontinued variant of the product to its newer variant.
	//
	// RangeIncludes:
	// - http://schema.org/ProductModel
	//
	PredecessorOf *ProductModel `json:"predecessorOf,omitempty" jsonld:"http://schema.org/predecessorOf"`

	// A pointer from a newer variant of a product  to its previous, often discontinued predecessor.
	//
	// RangeIncludes:
	// - http://schema.org/ProductModel
	//
	SuccessorOf *ProductModel `json:"successorOf,omitempty" jsonld:"http://schema.org/successorOf"`

}

type ProductModel struct {
	MetaTrait
	ProductModelTrait
	ProductTrait
	ThingTrait
	AdditionalTrait
}

func NewProductModel() (x ProductModel) {
	x.Type = "http://schema.org/ProductModel"

	return
}
