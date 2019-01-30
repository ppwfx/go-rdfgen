package schema

type BookTrait struct {

	// Indicates whether the book is an abridged edition.
	//
	// RangeIncludes:
	// - http://schema.org/Boolean
	//
	Abridged bool `json:"abridged,omitempty" jsonld:"http://schema.org/abridged"`

	// The edition of the book.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	BookEdition string `json:"bookEdition,omitempty" jsonld:"http://schema.org/bookEdition"`

	// The format of the book.
	//
	// RangeIncludes:
	// - http://schema.org/BookFormatType
	//
	BookFormat *BookFormatType `json:"bookFormat,omitempty" jsonld:"http://schema.org/bookFormat"`

	// The illustrator of the book.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Illustrator *Person `json:"illustrator,omitempty" jsonld:"http://schema.org/illustrator"`

	// The ISBN of the book.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Isbn string `json:"isbn,omitempty" jsonld:"http://schema.org/isbn"`

	// The number of pages in the book.
	//
	// RangeIncludes:
	// - http://schema.org/Integer
	//
	NumberOfPages *Integer `json:"numberOfPages,omitempty" jsonld:"http://schema.org/numberOfPages"`

}

type Book struct {
	MetaTrait
	BookTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewBook() (x Book) {
	x.Type = "http://schema.org/Book"

	return
}
