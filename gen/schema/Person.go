package schema

type PersonTrait struct {

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

	// The height of the item.
	//
	// RangeIncludes:
	// - http://schema.org/Distance
	// - http://schema.org/QuantitativeValue
	//
	Height interface{} `json:"height,omitempty" jsonld:"http://schema.org/height"`

	// The publishingPrinciples property indicates (typically via <a class=\"localLink\" href=\"http://schema.org/URL\">URL</a>) a document describing the editorial principles of an <a class=\"localLink\" href=\"http://schema.org/Organization\">Organization</a> (or individual e.g. a <a class=\"localLink\" href=\"http://schema.org/Person\">Person</a> writing a blog) that relate to their activities as a publisher, e.g. ethics or diversity policies. When applied to a <a class=\"localLink\" href=\"http://schema.org/CreativeWork\">CreativeWork</a> (e.g. <a class=\"localLink\" href=\"http://schema.org/NewsArticle\">NewsArticle</a>) the principles are those of the party primarily responsible for the creation of the <a class=\"localLink\" href=\"http://schema.org/CreativeWork\">CreativeWork</a>.<br/><br/>\n\nWhile such policies are most typically expressed in natural language, sometimes related information (e.g. indicating a <a class=\"localLink\" href=\"http://schema.org/funder\">funder</a>) can be expressed using schema.org terminology.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	// - http://schema.org/URL
	//
	PublishingPrinciples interface{} `json:"publishingPrinciples,omitempty" jsonld:"http://schema.org/publishingPrinciples"`

	// A person or organization that supports a thing through a pledge, promise, or financial contribution. e.g. a sponsor of a Medical Study or a corporate sponsor of an event.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/Person
	//
	Sponsor interface{} `json:"sponsor,omitempty" jsonld:"http://schema.org/sponsor"`

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

	// The telephone number.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Telephone string `json:"telephone,omitempty" jsonld:"http://schema.org/telephone"`

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

	// An additional name for a Person, can be used for a middle name.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	AdditionalName string `json:"additionalName,omitempty" jsonld:"http://schema.org/additionalName"`

	// An organization that this person is affiliated with. For example, a school/university, a club, or a team.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	//
	Affiliation *Organization `json:"affiliation,omitempty" jsonld:"http://schema.org/affiliation"`

	// An organization that the person is an alumni of.
	//
	// RangeIncludes:
	// - http://schema.org/EducationalOrganization
	// - http://schema.org/Organization
	//
	AlumniOf interface{} `json:"alumniOf,omitempty" jsonld:"http://schema.org/alumniOf"`

	// Date of birth.
	//
	// RangeIncludes:
	// - http://schema.org/Date
	//
	BirthDate *Date `json:"birthDate,omitempty" jsonld:"http://schema.org/birthDate"`

	// The place where the person was born.
	//
	// RangeIncludes:
	// - http://schema.org/Place
	//
	BirthPlace *Place `json:"birthPlace,omitempty" jsonld:"http://schema.org/birthPlace"`

	// The brand(s) associated with a product or service, or the brand(s) maintained by an organization or business person.
	//
	// RangeIncludes:
	// - http://schema.org/Brand
	// - http://schema.org/Organization
	//
	Brand interface{} `json:"brand,omitempty" jsonld:"http://schema.org/brand"`

	// A child of the person.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Children *Person `json:"children,omitempty" jsonld:"http://schema.org/children"`

	// A colleague of the person.
	//
	// RangeIncludes:
	// - http://schema.org/URL
	// - http://schema.org/Person
	//
	Colleague interface{} `json:"colleague,omitempty" jsonld:"http://schema.org/colleague"`

	// A colleague of the person.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Colleagues *Person `json:"colleagues,omitempty" jsonld:"http://schema.org/colleagues"`

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

	// Date of death.
	//
	// RangeIncludes:
	// - http://schema.org/Date
	//
	DeathDate *Date `json:"deathDate,omitempty" jsonld:"http://schema.org/deathDate"`

	// The place where the person died.
	//
	// RangeIncludes:
	// - http://schema.org/Place
	//
	DeathPlace *Place `json:"deathPlace,omitempty" jsonld:"http://schema.org/deathPlace"`

	// The Dun &amp; Bradstreet DUNS number for identifying an organization or business person.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Duns string `json:"duns,omitempty" jsonld:"http://schema.org/duns"`

	// Family name. In the U.S., the last name of an Person. This can be used along with givenName instead of the name property.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	FamilyName string `json:"familyName,omitempty" jsonld:"http://schema.org/familyName"`

	// The most generic uni-directional social relation.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Follows *Person `json:"follows,omitempty" jsonld:"http://schema.org/follows"`

	// Gender of the person. While http://schema.org/Male and http://schema.org/Female may be used, text strings are also acceptable for people who do not identify as a binary gender.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/GenderType
	//
	Gender interface{} `json:"gender,omitempty" jsonld:"http://schema.org/gender"`

	// Given name. In the U.S., the first name of a Person. This can be used along with familyName instead of the name property.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	GivenName string `json:"givenName,omitempty" jsonld:"http://schema.org/givenName"`

	// The Person's occupation. For past professions, use Role for expressing dates.
	//
	// RangeIncludes:
	// - http://schema.org/Occupation
	//
	HasOccupation *Occupation `json:"hasOccupation,omitempty" jsonld:"http://schema.org/hasOccupation"`

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

	// A contact location for a person's residence.
	//
	// RangeIncludes:
	// - http://schema.org/ContactPoint
	// - http://schema.org/Place
	//
	HomeLocation interface{} `json:"homeLocation,omitempty" jsonld:"http://schema.org/homeLocation"`

	// An honorific prefix preceding a Person's name such as Dr/Mrs/Mr.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	HonorificPrefix string `json:"honorificPrefix,omitempty" jsonld:"http://schema.org/honorificPrefix"`

	// An honorific suffix preceding a Person's name such as M.D. /PhD/MSCSW.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	HonorificSuffix string `json:"honorificSuffix,omitempty" jsonld:"http://schema.org/honorificSuffix"`

	// The job title of the person (for example, Financial Manager).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	JobTitle string `json:"jobTitle,omitempty" jsonld:"http://schema.org/jobTitle"`

	// The most generic bi-directional social/work relation.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Knows *Person `json:"knows,omitempty" jsonld:"http://schema.org/knows"`

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

	// Nationality of the person.
	//
	// RangeIncludes:
	// - http://schema.org/Country
	//
	Nationality *Country `json:"nationality,omitempty" jsonld:"http://schema.org/nationality"`

	// The total financial value of the person as calculated by subtracting assets from liabilities.
	//
	// RangeIncludes:
	// - http://schema.org/MonetaryAmount
	// - http://schema.org/PriceSpecification
	//
	NetWorth interface{} `json:"netWorth,omitempty" jsonld:"http://schema.org/netWorth"`

	// Products owned by the organization or person.
	//
	// RangeIncludes:
	// - http://schema.org/OwnershipInfo
	// - http://schema.org/Product
	//
	Owns interface{} `json:"owns,omitempty" jsonld:"http://schema.org/owns"`

	// A parent of this person.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Parent *Person `json:"parent,omitempty" jsonld:"http://schema.org/parent"`

	// A parents of the person.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Parents *Person `json:"parents,omitempty" jsonld:"http://schema.org/parents"`

	// Event that this person is a performer or participant in.
	//
	// RangeIncludes:
	// - http://schema.org/Event
	//
	PerformerIn *Event `json:"performerIn,omitempty" jsonld:"http://schema.org/performerIn"`

	// The most generic familial relation.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	RelatedTo *Person `json:"relatedTo,omitempty" jsonld:"http://schema.org/relatedTo"`

	// A pointer to products or services sought by the organization or person (demand).
	//
	// RangeIncludes:
	// - http://schema.org/Demand
	//
	Seeks *Demand `json:"seeks,omitempty" jsonld:"http://schema.org/seeks"`

	// A sibling of the person.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Sibling *Person `json:"sibling,omitempty" jsonld:"http://schema.org/sibling"`

	// A sibling of the person.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Siblings *Person `json:"siblings,omitempty" jsonld:"http://schema.org/siblings"`

	// The person's spouse.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Spouse *Person `json:"spouse,omitempty" jsonld:"http://schema.org/spouse"`

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

	// The weight of the product or person.
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	//
	Weight *QuantitativeValue `json:"weight,omitempty" jsonld:"http://schema.org/weight"`

	// A contact location for a person's place of work.
	//
	// RangeIncludes:
	// - http://schema.org/ContactPoint
	// - http://schema.org/Place
	//
	WorkLocation interface{} `json:"workLocation,omitempty" jsonld:"http://schema.org/workLocation"`

	// Organizations that the person works for.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	//
	WorksFor *Organization `json:"worksFor,omitempty" jsonld:"http://schema.org/worksFor"`

}

type Person struct {
	MetaTrait
	PersonTrait
	ThingTrait
	AdditionalTrait
}

func NewPerson() (x Person) {
	x.Type = "http://schema.org/Person"

	return
}
