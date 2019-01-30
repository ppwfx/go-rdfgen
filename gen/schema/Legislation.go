package schema

type LegislationTrait struct {

	// An identifier for the legislation. This can be either a string-based identifier, like the CELEX at EU level or the NOR in France, or a web-based, URL/URI identifier, like an ELI (European Legislation Identifier) or an URN-Lex.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	LegislationIdentifier interface{} `json:"legislationIdentifier,omitempty" jsonld:"http://schema.org/legislationIdentifier"`

	// The jurisdiction from which the legislation originates.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/AdministrativeArea
	//
	LegislationJurisdiction interface{} `json:"legislationJurisdiction,omitempty" jsonld:"http://schema.org/legislationJurisdiction"`

	// The type of the legislation. Examples of values are \"law\", \"act\", \"directive\", \"decree\", \"regulation\", \"statutory instrument\", \"loi organique\", \"règlement grand-ducal\", etc., depending on the country.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/CategoryCode
	//
	LegislationType interface{} `json:"legislationType,omitempty" jsonld:"http://schema.org/legislationType"`

	// Indicates that this legislation (or part of a legislation) somehow transfers another legislation in a different legislative context. This is an informative link, and it has no legal value. For legally-binding links of transposition, use the <a href=\"/legislationTransposes\">legislationTransposes</a> property. For example an informative consolidated law of a European Union's member state \"applies\" the consolidated version of the European Directive implemented in it.
	//
	// RangeIncludes:
	// - http://schema.org/Legislation
	//
	LegislationApplies *Legislation `json:"legislationApplies,omitempty" jsonld:"http://schema.org/legislationApplies"`

	// Another legislation that this legislation changes. This encompasses the notions of amendment, replacement, correction, repeal, or other types of change. This may be a direct change (textual or non-textual amendment) or a consequential or indirect change. The property is to be used to express the existence of a change relationship between two acts rather than the existence of a consolidated version of the text that shows the result of the change. For consolidation relationships, use the <a href=\"/legislationConsolidates\">legislationConsolidates</a> property.
	//
	// RangeIncludes:
	// - http://schema.org/Legislation
	//
	LegislationChanges *Legislation `json:"legislationChanges,omitempty" jsonld:"http://schema.org/legislationChanges"`

	// Indicates another legislation taken into account in this consolidated legislation (which is usually the product of an editorial process that revises the legislation). This property should be used multiple times to refer to both the original version or the previous consolidated version, and to the legislations making the change.
	//
	// RangeIncludes:
	// - http://schema.org/Legislation
	//
	LegislationConsolidates *Legislation `json:"legislationConsolidates,omitempty" jsonld:"http://schema.org/legislationConsolidates"`

	// The date of adoption or signature of the legislation. This is the date at which the text is officially aknowledged to be a legislation, even though it might not even be published or in force.
	//
	// RangeIncludes:
	// - http://schema.org/Date
	//
	LegislationDate *Date `json:"legislationDate,omitempty" jsonld:"http://schema.org/legislationDate"`

	// The point-in-time at which the provided description of the legislation is valid (e.g. : when looking at the law on the 2016-04-07 (= dateVersion), I get the consolidation of 2015-04-12 of the \"National Insurance Contributions Act 2015\")
	//
	// RangeIncludes:
	// - http://schema.org/Date
	//
	LegislationDateVersion *Date `json:"legislationDateVersion,omitempty" jsonld:"http://schema.org/legislationDateVersion"`

	// Whether the legislation is currently in force, not in force, or partially in force.
	//
	// RangeIncludes:
	// - http://schema.org/LegalForceStatus
	//
	LegislationLegalForce *LegalForceStatus `json:"legislationLegalForce,omitempty" jsonld:"http://schema.org/legislationLegalForce"`

	// The person or organization that originally passed or made the law : typically parliament (for primary legislation) or government (for secondary legislation). This indicates the \"legal author\" of the law, as opposed to its physical author.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	// - http://schema.org/Organization
	//
	LegislationPassedBy interface{} `json:"legislationPassedBy,omitempty" jsonld:"http://schema.org/legislationPassedBy"`

	// An individual or organization that has some kind of responsibility for the legislation. Typically the ministry who is/was in charge of elaborating the legislation, or the adressee for potential questions about the legislation once it is published.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	// - http://schema.org/Organization
	//
	LegislationResponsible interface{} `json:"legislationResponsible,omitempty" jsonld:"http://schema.org/legislationResponsible"`

	// Indicates that this legislation (or part of legislation) fulfills the objectives set by another legislation, by passing appropriate implementation measures. Typically, some legislations of European Union's member states or regions transpose European Directives. This indicates a legally binding link between the 2 legislations.
	//
	// RangeIncludes:
	// - http://schema.org/Legislation
	//
	LegislationTransposes *Legislation `json:"legislationTransposes,omitempty" jsonld:"http://schema.org/legislationTransposes"`

}

type Legislation struct {
	MetaTrait
	LegislationTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewLegislation() (x Legislation) {
	x.Type = "http://schema.org/Legislation"

	return
}
