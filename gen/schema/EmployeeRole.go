package schema

type EmployeeRoleTrait struct {

	// The base salary of the job or of an employee in an EmployeeRole.
	//
	// RangeIncludes:
	// - http://schema.org/MonetaryAmount
	// - http://schema.org/Number
	// - http://schema.org/PriceSpecification
	//
	BaseSalary interface{} `json:"baseSalary,omitempty" jsonld:"http://schema.org/baseSalary"`

	// The currency (coded using <a href=\"http://en.wikipedia.org/wiki/ISO_4217\">ISO 4217</a> ) used for the main salary information in this job posting or for this employee.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	SalaryCurrency string `json:"salaryCurrency,omitempty" jsonld:"http://schema.org/salaryCurrency"`

}

type EmployeeRole struct {
	MetaTrait
	EmployeeRoleTrait
	OrganizationRoleTrait
	RoleTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewEmployeeRole() (x EmployeeRole) {
	x.Type = "http://schema.org/EmployeeRole"

	return
}
