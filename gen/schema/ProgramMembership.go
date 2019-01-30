package schema

type ProgramMembershipTrait struct {

	// A member of an Organization or a ProgramMembership. Organizations can be members of organizations; ProgramMembership is typically for individuals.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	// - http://schema.org/Organization
	//
	Member interface{} `json:"member,omitempty" jsonld:"http://schema.org/member"`

	// A member of this organization.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	// - http://schema.org/Organization
	//
	Members interface{} `json:"members,omitempty" jsonld:"http://schema.org/members"`

	// The organization (airline, travelers' club, etc.) the membership is made with.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	//
	HostingOrganization *Organization `json:"hostingOrganization,omitempty" jsonld:"http://schema.org/hostingOrganization"`

	// A unique identifier for the membership.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	MembershipNumber string `json:"membershipNumber,omitempty" jsonld:"http://schema.org/membershipNumber"`

	// The program providing the membership.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	ProgramName string `json:"programName,omitempty" jsonld:"http://schema.org/programName"`

}

type ProgramMembership struct {
	MetaTrait
	ProgramMembershipTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewProgramMembership() (x ProgramMembership) {
	x.Type = "http://schema.org/ProgramMembership"

	return
}
