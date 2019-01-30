package schema

type OrganizationTrait struct {

	// The overall rating, based on a collection of reviews or ratings, of the item.
	//
	// RangeIncludes:
	// - http://schema.org/AggregateRating
	//
	AggregateRating *AggregateRating `json:"aggregateRating,omitempty" jsonld:"http://schema.org/aggregateRating"`

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

	// A person or organization that supports (sponsors) something through some kind of financial contribution.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/Person
	//
	Funder interface{} `json:"funder,omitempty" jsonld:"http://schema.org/funder"`

	// The publishingPrinciples property indicates (typically via <a class=\"localLink\" href=\"http://schema.org/URL\">URL</a>) a document describing the editorial principles of an <a class=\"localLink\" href=\"http://schema.org/Organization\">Organization</a> (or individual e.g. a <a class=\"localLink\" href=\"http://schema.org/Person\">Person</a> writing a blog) that relate to their activities as a publisher, e.g. ethics or diversity policies. When applied to a <a class=\"localLink\" href=\"http://schema.org/CreativeWork\">CreativeWork</a> (e.g. <a class=\"localLink\" href=\"http://schema.org/NewsArticle\">NewsArticle</a>) the principles are those of the party primarily responsible for the creation of the <a class=\"localLink\" href=\"http://schema.org/CreativeWork\">CreativeWork</a>.<br/><br/>\n\nWhile such policies are most typically expressed in natural language, sometimes related information (e.g. indicating a <a class=\"localLink\" href=\"http://schema.org/funder\">funder</a>) can be expressed using schema.org terminology.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	// - http://schema.org/URL
	//
	PublishingPrinciples interface{} `json:"publishingPrinciples,omitempty" jsonld:"http://schema.org/publishingPrinciples"`

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

	// A person or organization that supports a thing through a pledge, promise, or financial contribution. e.g. a sponsor of a Medical Study or a corporate sponsor of an event.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/Person
	//
	Sponsor interface{} `json:"sponsor,omitempty" jsonld:"http://schema.org/sponsor"`

	// The location of for example where the event is happening, an organization is located, or where an action takes place.
	//
	// RangeIncludes:
	// - http://schema.org/Place
	// - http://schema.org/PostalAddress
	// - http://schema.org/Text
	//
	Location interface{} `json:"location,omitempty" jsonld:"http://schema.org/location"`

	// The geographic area where a service or offered item is provided.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/AdministrativeArea
	// - http://schema.org/GeoShape
	// - http://schema.org/Place
	//
	AreaServed interface{} `json:"areaServed,omitempty" jsonld:"http://schema.org/areaServed"`

	// Physical address of the item.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/PostalAddress
	//
	Address interface{} `json:"address,omitempty" jsonld:"http://schema.org/address"`

	// Email address.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Email string `json:"email,omitempty" jsonld:"http://schema.org/email"`

	// The fax number.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	FaxNumber string `json:"faxNumber,omitempty" jsonld:"http://schema.org/faxNumber"`

	// The geographic area where the service is provided.
	//
	// RangeIncludes:
	// - http://schema.org/GeoShape
	// - http://schema.org/AdministrativeArea
	// - http://schema.org/Place
	//
	ServiceArea interface{} `json:"serviceArea,omitempty" jsonld:"http://schema.org/serviceArea"`

	// The telephone number.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Telephone string `json:"telephone,omitempty" jsonld:"http://schema.org/telephone"`

	// Upcoming or past event associated with this place, organization, or action.
	//
	// RangeIncludes:
	// - http://schema.org/Event
	//
	Event *Event `json:"event,omitempty" jsonld:"http://schema.org/event"`

	// Upcoming or past events associated with this place or organization.
	//
	// RangeIncludes:
	// - http://schema.org/Event
	//
	Events *Event `json:"events,omitempty" jsonld:"http://schema.org/events"`

	// The <a href=\"http://www.gs1.org/gln\">Global Location Number</a> (GLN, sometimes also referred to as International Location Number or ILN) of the respective organization, person, or place. The GLN is a 13-digit number used to identify parties and physical locations.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	GlobalLocationNumber string `json:"globalLocationNumber,omitempty" jsonld:"http://schema.org/globalLocationNumber"`

	// The International Standard of Industrial Classification of All Economic Activities (ISIC), Revision 4 code for a particular organization, business person, or place.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	IsicV4 string `json:"isicV4,omitempty" jsonld:"http://schema.org/isicV4"`

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

	// A contact point for a person or organization.
	//
	// RangeIncludes:
	// - http://schema.org/ContactPoint
	//
	ContactPoint *ContactPoint `json:"contactPoint,omitempty" jsonld:"http://schema.org/contactPoint"`

	// A contact point for a person or organization.
	//
	// RangeIncludes:
	// - http://schema.org/ContactPoint
	//
	ContactPoints *ContactPoint `json:"contactPoints,omitempty" jsonld:"http://schema.org/contactPoints"`

	// The Dun &amp; Bradstreet DUNS number for identifying an organization or business person.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Duns string `json:"duns,omitempty" jsonld:"http://schema.org/duns"`

	// Indicates an OfferCatalog listing for this Organization, Person, or Service.
	//
	// RangeIncludes:
	// - http://schema.org/OfferCatalog
	//
	HasOfferCatalog *OfferCatalog `json:"hasOfferCatalog,omitempty" jsonld:"http://schema.org/hasOfferCatalog"`

	// Points-of-Sales operated by the organization or person.
	//
	// RangeIncludes:
	// - http://schema.org/Place
	//
	HasPOS *Place `json:"hasPOS,omitempty" jsonld:"http://schema.org/hasPOS"`

	// Of a <a class=\"localLink\" href=\"http://schema.org/Person\">Person</a>, and less typically of an <a class=\"localLink\" href=\"http://schema.org/Organization\">Organization</a>, to indicate a topic that is known about - suggesting possible expertise but not implying it. We do not distinguish skill levels here, or yet relate this to educational content, events, objectives or <a class=\"localLink\" href=\"http://schema.org/JobPosting\">JobPosting</a> descriptions.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Thing
	// - http://schema.org/URL
	//
	KnowsAbout interface{} `json:"knowsAbout,omitempty" jsonld:"http://schema.org/knowsAbout"`

	// Of a <a class=\"localLink\" href=\"http://schema.org/Person\">Person</a>, and less typically of an <a class=\"localLink\" href=\"http://schema.org/Organization\">Organization</a>, to indicate a known language. We do not distinguish skill levels or reading/writing/speaking/signing here. Use language codes from the <a href=\"http://tools.ietf.org/html/bcp47\">IETF BCP 47 standard</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Language
	//
	KnowsLanguage interface{} `json:"knowsLanguage,omitempty" jsonld:"http://schema.org/knowsLanguage"`

	// A pointer to products or services offered by the organization or person.
	//
	// RangeIncludes:
	// - http://schema.org/Offer
	//
	MakesOffer *Offer `json:"makesOffer,omitempty" jsonld:"http://schema.org/makesOffer"`

	// An Organization (or ProgramMembership) to which this Person or Organization belongs.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/ProgramMembership
	//
	MemberOf interface{} `json:"memberOf,omitempty" jsonld:"http://schema.org/memberOf"`

	// The North American Industry Classification System (NAICS) code for a particular organization or business person.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Naics string `json:"naics,omitempty" jsonld:"http://schema.org/naics"`

	// Products owned by the organization or person.
	//
	// RangeIncludes:
	// - http://schema.org/OwnershipInfo
	// - http://schema.org/Product
	//
	Owns interface{} `json:"owns,omitempty" jsonld:"http://schema.org/owns"`

	// A pointer to products or services sought by the organization or person (demand).
	//
	// RangeIncludes:
	// - http://schema.org/Demand
	//
	Seeks *Demand `json:"seeks,omitempty" jsonld:"http://schema.org/seeks"`

	// The Tax / Fiscal ID of the organization or person, e.g. the TIN in the US or the CIF/NIF in Spain.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	TaxID string `json:"taxID,omitempty" jsonld:"http://schema.org/taxID"`

	// The Value-added Tax ID of the organization or person.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	VatID string `json:"vatID,omitempty" jsonld:"http://schema.org/vatID"`

	// For a <a class=\"localLink\" href=\"http://schema.org/NewsMediaOrganization\">NewsMediaOrganization</a> or other news-related <a class=\"localLink\" href=\"http://schema.org/Organization\">Organization</a>, a statement about public engagement activities (for news media, the newsroom’s), including involving the public - digitally or otherwise -- in coverage decisions, reporting and activities after publication.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	// - http://schema.org/URL
	//
	ActionableFeedbackPolicy interface{} `json:"actionableFeedbackPolicy,omitempty" jsonld:"http://schema.org/actionableFeedbackPolicy"`

	// Alumni of an organization.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Alumni *Person `json:"alumni,omitempty" jsonld:"http://schema.org/alumni"`

	// For an <a class=\"localLink\" href=\"http://schema.org/Organization\">Organization</a> (e.g. <a class=\"localLink\" href=\"http://schema.org/NewsMediaOrganization\">NewsMediaOrganization</a>), a statement describing (in news media, the newsroom’s) disclosure and correction policy for errors.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	// - http://schema.org/URL
	//
	CorrectionsPolicy interface{} `json:"correctionsPolicy,omitempty" jsonld:"http://schema.org/correctionsPolicy"`

	// A relationship between an organization and a department of that organization, also described as an organization (allowing different urls, logos, opening hours). For example: a store with a pharmacy, or a bakery with a cafe.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	//
	Department *Organization `json:"department,omitempty" jsonld:"http://schema.org/department"`

	// The date that this organization was dissolved.
	//
	// RangeIncludes:
	// - http://schema.org/Date
	//
	DissolutionDate *Date `json:"dissolutionDate,omitempty" jsonld:"http://schema.org/dissolutionDate"`

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

	// Someone working for this organization.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Employee *Person `json:"employee,omitempty" jsonld:"http://schema.org/employee"`

	// People working for this organization.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Employees *Person `json:"employees,omitempty" jsonld:"http://schema.org/employees"`

	// Statement about ethics policy, e.g. of a <a class=\"localLink\" href=\"http://schema.org/NewsMediaOrganization\">NewsMediaOrganization</a> regarding journalistic and publishing practices, or of a <a class=\"localLink\" href=\"http://schema.org/Restaurant\">Restaurant</a>, a page describing food source policies. In the case of a <a class=\"localLink\" href=\"http://schema.org/NewsMediaOrganization\">NewsMediaOrganization</a>, an ethicsPolicy is typically a statement describing the personal, organizational, and corporate standards of behavior expected by the organization.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	// - http://schema.org/URL
	//
	EthicsPolicy interface{} `json:"ethicsPolicy,omitempty" jsonld:"http://schema.org/ethicsPolicy"`

	// A person who founded this organization.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Founder *Person `json:"founder,omitempty" jsonld:"http://schema.org/founder"`

	// A person who founded this organization.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Founders *Person `json:"founders,omitempty" jsonld:"http://schema.org/founders"`

	// The date that this organization was founded.
	//
	// RangeIncludes:
	// - http://schema.org/Date
	//
	FoundingDate *Date `json:"foundingDate,omitempty" jsonld:"http://schema.org/foundingDate"`

	// The place where the Organization was founded.
	//
	// RangeIncludes:
	// - http://schema.org/Place
	//
	FoundingLocation *Place `json:"foundingLocation,omitempty" jsonld:"http://schema.org/foundingLocation"`

	// The official name of the organization, e.g. the registered company name.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	LegalName string `json:"legalName,omitempty" jsonld:"http://schema.org/legalName"`

	// An organization identifier that uniquely identifies a legal entity as defined in ISO 17442.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	LeiCode string `json:"leiCode,omitempty" jsonld:"http://schema.org/leiCode"`

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

	// The number of employees in an organization e.g. business.
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	//
	NumberOfEmployees *QuantitativeValue `json:"numberOfEmployees,omitempty" jsonld:"http://schema.org/numberOfEmployees"`

	// For an <a class=\"localLink\" href=\"http://schema.org/Organization\">Organization</a> (often but not necessarily a <a class=\"localLink\" href=\"http://schema.org/NewsMediaOrganization\">NewsMediaOrganization</a>), a description of organizational ownership structure; funding and grants. In a news/media setting, this is with particular reference to editorial independence.   Note that the <a class=\"localLink\" href=\"http://schema.org/funder\">funder</a> is also available and can be used to make basic funder information machine-readable.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/CreativeWork
	// - http://schema.org/URL
	// - http://schema.org/AboutPage
	//
	OwnershipFundingInfo interface{} `json:"ownershipFundingInfo,omitempty" jsonld:"http://schema.org/ownershipFundingInfo"`

	// The larger organization that this organization is a <a class=\"localLink\" href=\"http://schema.org/subOrganization\">subOrganization</a> of, if any.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	//
	ParentOrganization *Organization `json:"parentOrganization,omitempty" jsonld:"http://schema.org/parentOrganization"`

	// A relationship between two organizations where the first includes the second, e.g., as a subsidiary. See also: the more specific 'department' property.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	//
	SubOrganization *Organization `json:"subOrganization,omitempty" jsonld:"http://schema.org/subOrganization"`

	// For an <a class=\"localLink\" href=\"http://schema.org/Organization\">Organization</a> (typically a <a class=\"localLink\" href=\"http://schema.org/NewsMediaOrganization\">NewsMediaOrganization</a>), a statement about policy on use of unnamed sources and the decision process required.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	// - http://schema.org/URL
	//
	UnnamedSourcesPolicy interface{} `json:"unnamedSourcesPolicy,omitempty" jsonld:"http://schema.org/unnamedSourcesPolicy"`

}

type Organization struct {
	MetaTrait
	OrganizationTrait
	ThingTrait
	AdditionalTrait
}

func NewOrganization() (x Organization) {
	x.Type = "http://schema.org/Organization"

	return
}
