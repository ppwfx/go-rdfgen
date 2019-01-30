package schema

type NewsMediaOrganizationTrait struct {

	// For a <a class=\"localLink\" href=\"http://schema.org/NewsMediaOrganization\">NewsMediaOrganization</a> or other news-related <a class=\"localLink\" href=\"http://schema.org/Organization\">Organization</a>, a statement about public engagement activities (for news media, the newsroom’s), including involving the public - digitally or otherwise -- in coverage decisions, reporting and activities after publication.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	// - http://schema.org/URL
	//
	ActionableFeedbackPolicy interface{} `json:"actionableFeedbackPolicy,omitempty" jsonld:"http://schema.org/actionableFeedbackPolicy"`

	// For an <a class=\"localLink\" href=\"http://schema.org/Organization\">Organization</a> (e.g. <a class=\"localLink\" href=\"http://schema.org/NewsMediaOrganization\">NewsMediaOrganization</a>), a statement describing (in news media, the newsroom’s) disclosure and correction policy for errors.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	// - http://schema.org/URL
	//
	CorrectionsPolicy interface{} `json:"correctionsPolicy,omitempty" jsonld:"http://schema.org/correctionsPolicy"`

	// Statement on diversity policy by an <a class=\"localLink\" href=\"http://schema.org/Organization\">Organization</a> e.g. a <a class=\"localLink\" href=\"http://schema.org/NewsMediaOrganization\">NewsMediaOrganization</a>. For a <a class=\"localLink\" href=\"http://schema.org/NewsMediaOrganization\">NewsMediaOrganization</a>, a statement describing the newsroom’s diversity policy on both staffing and sources, typically providing staffing data.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	// - http://schema.org/URL
	//
	DiversityPolicy interface{} `json:"diversityPolicy,omitempty" jsonld:"http://schema.org/diversityPolicy"`

	// For an <a class=\"localLink\" href=\"http://schema.org/Organization\">Organization</a> (often but not necessarily a <a class=\"localLink\" href=\"http://schema.org/NewsMediaOrganization\">NewsMediaOrganization</a>), a report on staffing diversity issues. In a news context this might be for example ASNE or RTDNA (US) reports, or self-reported.
	//
	// RangeIncludes:
	// - http://schema.org/URL
	// - http://schema.org/Article
	//
	DiversityStaffingReport interface{} `json:"diversityStaffingReport,omitempty" jsonld:"http://schema.org/diversityStaffingReport"`

	// Statement about ethics policy, e.g. of a <a class=\"localLink\" href=\"http://schema.org/NewsMediaOrganization\">NewsMediaOrganization</a> regarding journalistic and publishing practices, or of a <a class=\"localLink\" href=\"http://schema.org/Restaurant\">Restaurant</a>, a page describing food source policies. In the case of a <a class=\"localLink\" href=\"http://schema.org/NewsMediaOrganization\">NewsMediaOrganization</a>, an ethicsPolicy is typically a statement describing the personal, organizational, and corporate standards of behavior expected by the organization.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	// - http://schema.org/URL
	//
	EthicsPolicy interface{} `json:"ethicsPolicy,omitempty" jsonld:"http://schema.org/ethicsPolicy"`

	// For an <a class=\"localLink\" href=\"http://schema.org/Organization\">Organization</a> (often but not necessarily a <a class=\"localLink\" href=\"http://schema.org/NewsMediaOrganization\">NewsMediaOrganization</a>), a description of organizational ownership structure; funding and grants. In a news/media setting, this is with particular reference to editorial independence.   Note that the <a class=\"localLink\" href=\"http://schema.org/funder\">funder</a> is also available and can be used to make basic funder information machine-readable.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/CreativeWork
	// - http://schema.org/URL
	// - http://schema.org/AboutPage
	//
	OwnershipFundingInfo interface{} `json:"ownershipFundingInfo,omitempty" jsonld:"http://schema.org/ownershipFundingInfo"`

	// For an <a class=\"localLink\" href=\"http://schema.org/Organization\">Organization</a> (typically a <a class=\"localLink\" href=\"http://schema.org/NewsMediaOrganization\">NewsMediaOrganization</a>), a statement about policy on use of unnamed sources and the decision process required.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	// - http://schema.org/URL
	//
	UnnamedSourcesPolicy interface{} `json:"unnamedSourcesPolicy,omitempty" jsonld:"http://schema.org/unnamedSourcesPolicy"`

	// For a <a class=\"localLink\" href=\"http://schema.org/NewsMediaOrganization\">NewsMediaOrganization</a>, a statement on coverage priorities, including any public agenda or stance on issues.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	// - http://schema.org/URL
	//
	MissionCoveragePrioritiesPolicy interface{} `json:"missionCoveragePrioritiesPolicy,omitempty" jsonld:"http://schema.org/missionCoveragePrioritiesPolicy"`

	// Disclosure about verification and fact-checking processes for a <a class=\"localLink\" href=\"http://schema.org/NewsMediaOrganization\">NewsMediaOrganization</a> or other fact-checking <a class=\"localLink\" href=\"http://schema.org/Organization\">Organization</a>.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	// - http://schema.org/URL
	//
	VerificationFactCheckingPolicy interface{} `json:"verificationFactCheckingPolicy,omitempty" jsonld:"http://schema.org/verificationFactCheckingPolicy"`

	// For a <a class=\"localLink\" href=\"http://schema.org/NewsMediaOrganization\">NewsMediaOrganization</a> or other news-related <a class=\"localLink\" href=\"http://schema.org/Organization\">Organization</a>, a statement explaining when authors of articles are not named in bylines.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	// - http://schema.org/URL
	//
	NoBylinesPolicy interface{} `json:"noBylinesPolicy,omitempty" jsonld:"http://schema.org/noBylinesPolicy"`

	// For a <a class=\"localLink\" href=\"http://schema.org/NewsMediaOrganization\">NewsMediaOrganization</a>, a link to the masthead page or a page listing top editorial management.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	// - http://schema.org/URL
	//
	Masthead interface{} `json:"masthead,omitempty" jsonld:"http://schema.org/masthead"`

}

type NewsMediaOrganization struct {
	MetaTrait
	NewsMediaOrganizationTrait
	OrganizationTrait
	ThingTrait
	AdditionalTrait
}

func NewNewsMediaOrganization() (x NewsMediaOrganization) {
	x.Type = "http://schema.org/NewsMediaOrganization"

	return
}
