package schema

type BorrowActionTrait struct {

	// A sub property of participant. The person that lends the object being borrowed.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	// - http://schema.org/Organization
	//
	Lender interface{} `json:"lender,omitempty" jsonld:"http://schema.org/lender"`

}

type BorrowAction struct {
	MetaTrait
	BorrowActionTrait
	TransferActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewBorrowAction() (x BorrowAction) {
	x.Type = "http://schema.org/BorrowAction"

	return
}
