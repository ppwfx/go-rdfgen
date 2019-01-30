package schema

type LendActionTrait struct {

	// A sub property of participant. The person that borrows the object being lent.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Borrower *Person `json:"borrower,omitempty" jsonld:"http://schema.org/borrower"`

}

type LendAction struct {
	MetaTrait
	LendActionTrait
	TransferActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewLendAction() (x LendAction) {
	x.Type = "http://schema.org/LendAction"

	return
}
