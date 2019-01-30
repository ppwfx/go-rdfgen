package schema

type CreativeWorkTrait struct {

	// The subject matter of the content.
	//
	// RangeIncludes:
	// - http://schema.org/Thing
	//
	About *Thing `json:"about,omitempty" jsonld:"http://schema.org/about"`

	// The human sensory perceptual system or cognitive faculty through which a person may process or perceive information. Expected values include: auditory, tactile, textual, visual, colorDependent, chartOnVisual, chemOnVisual, diagramOnVisual, mathOnVisual, musicOnVisual, textOnVisual.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	AccessMode string `json:"accessMode,omitempty" jsonld:"http://schema.org/accessMode"`

	// A list of single or combined accessModes that are sufficient to understand all the intellectual content of a resource. Expected values include:  auditory, tactile, textual, visual.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	AccessModeSufficient string `json:"accessModeSufficient,omitempty" jsonld:"http://schema.org/accessModeSufficient"`

	// Indicates that the resource is compatible with the referenced accessibility API (<a href=\"http://www.w3.org/wiki/WebSchemas/Accessibility\">WebSchemas wiki lists possible values</a>).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	AccessibilityAPI string `json:"accessibilityAPI,omitempty" jsonld:"http://schema.org/accessibilityAPI"`

	// Identifies input methods that are sufficient to fully control the described resource (<a href=\"http://www.w3.org/wiki/WebSchemas/Accessibility\">WebSchemas wiki lists possible values</a>).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	AccessibilityControl string `json:"accessibilityControl,omitempty" jsonld:"http://schema.org/accessibilityControl"`

	// Content features of the resource, such as accessible media, alternatives and supported enhancements for accessibility (<a href=\"http://www.w3.org/wiki/WebSchemas/Accessibility\">WebSchemas wiki lists possible values</a>).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	AccessibilityFeature string `json:"accessibilityFeature,omitempty" jsonld:"http://schema.org/accessibilityFeature"`

	// A characteristic of the described resource that is physiologically dangerous to some users. Related to WCAG 2.0 guideline 2.3 (<a href=\"http://www.w3.org/wiki/WebSchemas/Accessibility\">WebSchemas wiki lists possible values</a>).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	AccessibilityHazard string `json:"accessibilityHazard,omitempty" jsonld:"http://schema.org/accessibilityHazard"`

	// A human-readable summary of specific accessibility features or deficiencies, consistent with the other accessibility metadata but expressing subtleties such as \"short descriptions are present but long descriptions will be needed for non-visual users\" or \"short descriptions are present and no long descriptions are needed.\"
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	AccessibilitySummary string `json:"accessibilitySummary,omitempty" jsonld:"http://schema.org/accessibilitySummary"`

	// Specifies the Person that is legally accountable for the CreativeWork.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	AccountablePerson *Person `json:"accountablePerson,omitempty" jsonld:"http://schema.org/accountablePerson"`

	// The overall rating, based on a collection of reviews or ratings, of the item.
	//
	// RangeIncludes:
	// - http://schema.org/AggregateRating
	//
	AggregateRating *AggregateRating `json:"aggregateRating,omitempty" jsonld:"http://schema.org/aggregateRating"`

	// A secondary title of the CreativeWork.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	AlternativeHeadline string `json:"alternativeHeadline,omitempty" jsonld:"http://schema.org/alternativeHeadline"`

	// A media object that encodes this CreativeWork. This property is a synonym for encoding.
	//
	// RangeIncludes:
	// - http://schema.org/MediaObject
	//
	AssociatedMedia *MediaObject `json:"associatedMedia,omitempty" jsonld:"http://schema.org/associatedMedia"`

	// An intended audience, i.e. a group for whom something was created.
	//
	// RangeIncludes:
	// - http://schema.org/Audience
	//
	Audience *Audience `json:"audience,omitempty" jsonld:"http://schema.org/audience"`

	// An embedded audio object.
	//
	// RangeIncludes:
	// - http://schema.org/AudioObject
	//
	Audio *AudioObject `json:"audio,omitempty" jsonld:"http://schema.org/audio"`

	// The author of this content or rating. Please note that author is special in that HTML 5 provides a special mechanism for indicating authorship via the rel tag. That is equivalent to this and may be used interchangeably.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/Person
	//
	Author interface{} `json:"author,omitempty" jsonld:"http://schema.org/author"`

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

	// Fictional person connected with a creative work.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Character *Person `json:"character,omitempty" jsonld:"http://schema.org/character"`

	// A citation or reference to another creative work, such as another publication, web page, scholarly article, etc.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	// - http://schema.org/Text
	//
	Citation interface{} `json:"citation,omitempty" jsonld:"http://schema.org/citation"`

	// Comments, typically from users.
	//
	// RangeIncludes:
	// - http://schema.org/Comment
	//
	Comment *Comment `json:"comment,omitempty" jsonld:"http://schema.org/comment"`

	// The number of comments this CreativeWork (e.g. Article, Question or Answer) has received. This is most applicable to works published in Web sites with commenting system; additional comments may exist elsewhere.
	//
	// RangeIncludes:
	// - http://schema.org/Integer
	//
	CommentCount *Integer `json:"commentCount,omitempty" jsonld:"http://schema.org/commentCount"`

	// The location depicted or described in the content. For example, the location in a photograph or painting.
	//
	// RangeIncludes:
	// - http://schema.org/Place
	//
	ContentLocation *Place `json:"contentLocation,omitempty" jsonld:"http://schema.org/contentLocation"`

	// Official rating of a piece of content&#x2014;for example,'MPAA PG-13'.
	//
	// RangeIncludes:
	// - http://schema.org/Rating
	// - http://schema.org/Text
	//
	ContentRating interface{} `json:"contentRating,omitempty" jsonld:"http://schema.org/contentRating"`

	// The specific time described by a creative work, for works (e.g. articles, video objects etc.) that emphasise a particular moment within an Event.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	ContentReferenceTime *DateTime `json:"contentReferenceTime,omitempty" jsonld:"http://schema.org/contentReferenceTime"`

	// A secondary contributor to the CreativeWork or Event.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/Person
	//
	Contributor interface{} `json:"contributor,omitempty" jsonld:"http://schema.org/contributor"`

	// The party holding the legal copyright to the CreativeWork.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/Person
	//
	CopyrightHolder interface{} `json:"copyrightHolder,omitempty" jsonld:"http://schema.org/copyrightHolder"`

	// The year during which the claimed copyright for the CreativeWork was first asserted.
	//
	// RangeIncludes:
	// - http://schema.org/Number
	//
	CopyrightYear float64 `json:"copyrightYear,omitempty" jsonld:"http://schema.org/copyrightYear"`

	// Indicates a correction to a <a class=\"localLink\" href=\"http://schema.org/CreativeWork\">CreativeWork</a>, either via a <a class=\"localLink\" href=\"http://schema.org/CorrectionComment\">CorrectionComment</a>, textually or in another document.
	//
	// RangeIncludes:
	// - http://schema.org/CorrectionComment
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	Correction interface{} `json:"correction,omitempty" jsonld:"http://schema.org/correction"`

	// The creator/author of this CreativeWork. This is the same as the Author property for CreativeWork.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/Person
	//
	Creator interface{} `json:"creator,omitempty" jsonld:"http://schema.org/creator"`

	// The date on which the CreativeWork was created or the item was added to a DataFeed.
	//
	// RangeIncludes:
	// - http://schema.org/Date
	// - http://schema.org/DateTime
	//
	DateCreated interface{} `json:"dateCreated,omitempty" jsonld:"http://schema.org/dateCreated"`

	// The date on which the CreativeWork was most recently modified or when the item's entry was modified within a DataFeed.
	//
	// RangeIncludes:
	// - http://schema.org/Date
	// - http://schema.org/DateTime
	//
	DateModified interface{} `json:"dateModified,omitempty" jsonld:"http://schema.org/dateModified"`

	// Date of first broadcast/publication.
	//
	// RangeIncludes:
	// - http://schema.org/Date
	//
	DatePublished *Date `json:"datePublished,omitempty" jsonld:"http://schema.org/datePublished"`

	// A link to the page containing the comments of the CreativeWork.
	//
	// RangeIncludes:
	// - http://schema.org/URL
	//
	DiscussionUrl string `json:"discussionUrl,omitempty" jsonld:"http://schema.org/discussionUrl"`

	// Specifies the Person who edited the CreativeWork.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Editor *Person `json:"editor,omitempty" jsonld:"http://schema.org/editor"`

	// An alignment to an established educational framework.
	//
	// RangeIncludes:
	// - http://schema.org/AlignmentObject
	//
	EducationalAlignment *AlignmentObject `json:"educationalAlignment,omitempty" jsonld:"http://schema.org/educationalAlignment"`

	// The purpose of a work in the context of education; for example, 'assignment', 'group work'.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	EducationalUse string `json:"educationalUse,omitempty" jsonld:"http://schema.org/educationalUse"`

	// A media object that encodes this CreativeWork. This property is a synonym for associatedMedia.
	//
	// RangeIncludes:
	// - http://schema.org/MediaObject
	//
	Encoding *MediaObject `json:"encoding,omitempty" jsonld:"http://schema.org/encoding"`

	// Media type typically expressed using a MIME format (see <a href=\"http://www.iana.org/assignments/media-types/media-types.xhtml\">IANA site</a> and <a href=\"https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types\">MDN reference</a>) e.g. application/zip for a SoftwareApplication binary, audio/mpeg for .mp3 etc.).<br/><br/>\n\nIn cases where a <a class=\"localLink\" href=\"http://schema.org/CreativeWork\">CreativeWork</a> has several media type representations, <a class=\"localLink\" href=\"http://schema.org/encoding\">encoding</a> can be used to indicate each <a class=\"localLink\" href=\"http://schema.org/MediaObject\">MediaObject</a> alongside particular <a class=\"localLink\" href=\"http://schema.org/encodingFormat\">encodingFormat</a> information.<br/><br/>\n\nUnregistered or niche encoding and file formats can be indicated instead via the most appropriate URL, e.g. defining Web page or a Wikipedia/Wikidata entry.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	EncodingFormat interface{} `json:"encodingFormat,omitempty" jsonld:"http://schema.org/encodingFormat"`

	// A media object that encodes this CreativeWork.
	//
	// RangeIncludes:
	// - http://schema.org/MediaObject
	//
	Encodings *MediaObject `json:"encodings,omitempty" jsonld:"http://schema.org/encodings"`

	// A creative work that this work is an example/instance/realization/derivation of.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	//
	ExampleOfWork *CreativeWork `json:"exampleOfWork,omitempty" jsonld:"http://schema.org/exampleOfWork"`

	// Date the content expires and is no longer useful or available. For example a <a class=\"localLink\" href=\"http://schema.org/VideoObject\">VideoObject</a> or <a class=\"localLink\" href=\"http://schema.org/NewsArticle\">NewsArticle</a> whose availability or relevance is time-limited, or a <a class=\"localLink\" href=\"http://schema.org/ClaimReview\">ClaimReview</a> fact check whose publisher wants to indicate that it may no longer be relevant (or helpful to highlight) after some date.
	//
	// RangeIncludes:
	// - http://schema.org/Date
	//
	Expires *Date `json:"expires,omitempty" jsonld:"http://schema.org/expires"`

	// Media type, typically MIME format (see <a href=\"http://www.iana.org/assignments/media-types/media-types.xhtml\">IANA site</a>) of the content e.g. application/zip of a SoftwareApplication binary. In cases where a CreativeWork has several media type representations, 'encoding' can be used to indicate each MediaObject alongside particular fileFormat information. Unregistered or niche file formats can be indicated instead via the most appropriate URL, e.g. defining Web page or a Wikipedia entry.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	FileFormat interface{} `json:"fileFormat,omitempty" jsonld:"http://schema.org/fileFormat"`

	// A person or organization that supports (sponsors) something through some kind of financial contribution.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/Person
	//
	Funder interface{} `json:"funder,omitempty" jsonld:"http://schema.org/funder"`

	// Genre of the creative work, broadcast channel or group.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	Genre interface{} `json:"genre,omitempty" jsonld:"http://schema.org/genre"`

	// Indicates an item or CreativeWork that is part of this item, or CreativeWork (in some sense).
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	// - http://schema.org/Trip
	//
	HasPart interface{} `json:"hasPart,omitempty" jsonld:"http://schema.org/hasPart"`

	// Headline of the article.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Headline string `json:"headline,omitempty" jsonld:"http://schema.org/headline"`

	// The language of the content or performance or used in an action. Please use one of the language codes from the <a href=\"http://tools.ietf.org/html/bcp47\">IETF BCP 47 standard</a>. See also <a class=\"localLink\" href=\"http://schema.org/availableLanguage\">availableLanguage</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Language
	// - http://schema.org/Text
	//
	InLanguage interface{} `json:"inLanguage,omitempty" jsonld:"http://schema.org/inLanguage"`

	// The number of interactions for the CreativeWork using the WebSite or SoftwareApplication. The most specific child type of InteractionCounter should be used.
	//
	// RangeIncludes:
	// - http://schema.org/InteractionCounter
	//
	InteractionStatistic *InteractionCounter `json:"interactionStatistic,omitempty" jsonld:"http://schema.org/interactionStatistic"`

	// The predominant mode of learning supported by the learning resource. Acceptable values are 'active', 'expositive', or 'mixed'.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	InteractivityType string `json:"interactivityType,omitempty" jsonld:"http://schema.org/interactivityType"`

	// A flag to signal that the item, event, or place is accessible for free.
	//
	// RangeIncludes:
	// - http://schema.org/Boolean
	//
	IsAccessibleForFree bool `json:"isAccessibleForFree,omitempty" jsonld:"http://schema.org/isAccessibleForFree"`

	// A resource that was used in the creation of this resource. This term can be repeated for multiple sources. For example, http://example.com/great-multiplication-intro.html.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	// - http://schema.org/Product
	// - http://schema.org/URL
	//
	IsBasedOn interface{} `json:"isBasedOn,omitempty" jsonld:"http://schema.org/isBasedOn"`

	// A resource that was used in the creation of this resource. This term can be repeated for multiple sources. For example, http://example.com/great-multiplication-intro.html.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	// - http://schema.org/Product
	// - http://schema.org/URL
	//
	IsBasedOnUrl interface{} `json:"isBasedOnUrl,omitempty" jsonld:"http://schema.org/isBasedOnUrl"`

	// Indicates whether this content is family friendly.
	//
	// RangeIncludes:
	// - http://schema.org/Boolean
	//
	IsFamilyFriendly bool `json:"isFamilyFriendly,omitempty" jsonld:"http://schema.org/isFamilyFriendly"`

	// Indicates an item or CreativeWork that this item, or CreativeWork (in some sense), is part of.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	// - http://schema.org/Trip
	//
	IsPartOf interface{} `json:"isPartOf,omitempty" jsonld:"http://schema.org/isPartOf"`

	// Keywords or tags used to describe this content. Multiple entries in a keywords list are typically delimited by commas.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Keywords string `json:"keywords,omitempty" jsonld:"http://schema.org/keywords"`

	// The predominant type or kind characterizing the learning resource. For example, 'presentation', 'handout'.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	LearningResourceType string `json:"learningResourceType,omitempty" jsonld:"http://schema.org/learningResourceType"`

	// A license document that applies to this content, typically indicated by URL.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	// - http://schema.org/URL
	//
	License interface{} `json:"license,omitempty" jsonld:"http://schema.org/license"`

	// The location where the CreativeWork was created, which may not be the same as the location depicted in the CreativeWork.
	//
	// RangeIncludes:
	// - http://schema.org/Place
	//
	LocationCreated *Place `json:"locationCreated,omitempty" jsonld:"http://schema.org/locationCreated"`

	// Indicates the primary entity described in some page or other CreativeWork.
	//
	// RangeIncludes:
	// - http://schema.org/Thing
	//
	MainEntity *Thing `json:"mainEntity,omitempty" jsonld:"http://schema.org/mainEntity"`

	// A material that something is made from, e.g. leather, wool, cotton, paper.
	//
	// RangeIncludes:
	// - http://schema.org/Product
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	Material interface{} `json:"material,omitempty" jsonld:"http://schema.org/material"`

	// Indicates that the CreativeWork contains a reference to, but is not necessarily about a concept.
	//
	// RangeIncludes:
	// - http://schema.org/Thing
	//
	Mentions *Thing `json:"mentions,omitempty" jsonld:"http://schema.org/mentions"`

	// An offer to provide this item&#x2014;for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
	//
	// RangeIncludes:
	// - http://schema.org/Offer
	//
	Offers *Offer `json:"offers,omitempty" jsonld:"http://schema.org/offers"`

	// The position of an item in a series or sequence of items.
	//
	// RangeIncludes:
	// - http://schema.org/Integer
	// - http://schema.org/Text
	//
	Position interface{} `json:"position,omitempty" jsonld:"http://schema.org/position"`

	// The person or organization who produced the work (e.g. music album, movie, tv/radio series etc.).
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/Person
	//
	Producer interface{} `json:"producer,omitempty" jsonld:"http://schema.org/producer"`

	// The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/Person
	//
	Provider interface{} `json:"provider,omitempty" jsonld:"http://schema.org/provider"`

	// A publication event associated with the item.
	//
	// RangeIncludes:
	// - http://schema.org/PublicationEvent
	//
	Publication *PublicationEvent `json:"publication,omitempty" jsonld:"http://schema.org/publication"`

	// The publisher of the creative work.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/Person
	//
	Publisher interface{} `json:"publisher,omitempty" jsonld:"http://schema.org/publisher"`

	// The publishing division which published the comic.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	//
	PublisherImprint *Organization `json:"publisherImprint,omitempty" jsonld:"http://schema.org/publisherImprint"`

	// The publishingPrinciples property indicates (typically via <a class=\"localLink\" href=\"http://schema.org/URL\">URL</a>) a document describing the editorial principles of an <a class=\"localLink\" href=\"http://schema.org/Organization\">Organization</a> (or individual e.g. a <a class=\"localLink\" href=\"http://schema.org/Person\">Person</a> writing a blog) that relate to their activities as a publisher, e.g. ethics or diversity policies. When applied to a <a class=\"localLink\" href=\"http://schema.org/CreativeWork\">CreativeWork</a> (e.g. <a class=\"localLink\" href=\"http://schema.org/NewsArticle\">NewsArticle</a>) the principles are those of the party primarily responsible for the creation of the <a class=\"localLink\" href=\"http://schema.org/CreativeWork\">CreativeWork</a>.<br/><br/>\n\nWhile such policies are most typically expressed in natural language, sometimes related information (e.g. indicating a <a class=\"localLink\" href=\"http://schema.org/funder\">funder</a>) can be expressed using schema.org terminology.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	// - http://schema.org/URL
	//
	PublishingPrinciples interface{} `json:"publishingPrinciples,omitempty" jsonld:"http://schema.org/publishingPrinciples"`

	// The Event where the CreativeWork was recorded. The CreativeWork may capture all or part of the event.
	//
	// RangeIncludes:
	// - http://schema.org/Event
	//
	RecordedAt *Event `json:"recordedAt,omitempty" jsonld:"http://schema.org/recordedAt"`

	// The place and time the release was issued, expressed as a PublicationEvent.
	//
	// RangeIncludes:
	// - http://schema.org/PublicationEvent
	//
	ReleasedEvent *PublicationEvent `json:"releasedEvent,omitempty" jsonld:"http://schema.org/releasedEvent"`

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

	// Indicates (by URL or string) a particular version of a schema used in some CreativeWork. For example, a document could declare a schemaVersion using an URL such as http://schema.org/version/2.0/ if precise indication of schema version was required by some application.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	SchemaVersion interface{} `json:"schemaVersion,omitempty" jsonld:"http://schema.org/schemaVersion"`

	// Indicates the date on which the current structured data was generated / published. Typically used alongside <a class=\"localLink\" href=\"http://schema.org/sdPublisher\">sdPublisher</a>
	//
	// RangeIncludes:
	// - http://schema.org/Date
	//
	SdDatePublished *Date `json:"sdDatePublished,omitempty" jsonld:"http://schema.org/sdDatePublished"`

	// A license document that applies to this structured data, typically indicated by URL.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	// - http://schema.org/URL
	//
	SdLicense interface{} `json:"sdLicense,omitempty" jsonld:"http://schema.org/sdLicense"`

	// Indicates the party responsible for generating and publishing the current structured data markup, typically in cases where the structured data is derived automatically from existing published content but published on a different site. For example, student projects and open data initiatives often re-publish existing content with more explicitly structured metadata. The\n<a class=\"localLink\" href=\"http://schema.org/sdPublisher\">sdPublisher</a> property helps make such practices more explicit.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/Person
	//
	SdPublisher interface{} `json:"sdPublisher,omitempty" jsonld:"http://schema.org/sdPublisher"`

	// The Organization on whose behalf the creator was working.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	//
	SourceOrganization *Organization `json:"sourceOrganization,omitempty" jsonld:"http://schema.org/sourceOrganization"`

	// The spatialCoverage of a CreativeWork indicates the place(s) which are the focus of the content. It is a subproperty of\n      contentLocation intended primarily for more technical and detailed materials. For example with a Dataset, it indicates\n      areas that the dataset describes: a dataset of New York weather would have spatialCoverage which was the place: the state of New York.
	//
	// RangeIncludes:
	// - http://schema.org/Place
	//
	SpatialCoverage *Place `json:"spatialCoverage,omitempty" jsonld:"http://schema.org/spatialCoverage"`

	// A person or organization that supports a thing through a pledge, promise, or financial contribution. e.g. a sponsor of a Medical Study or a corporate sponsor of an event.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/Person
	//
	Sponsor interface{} `json:"sponsor,omitempty" jsonld:"http://schema.org/sponsor"`

	// The temporalCoverage of a CreativeWork indicates the period that the content applies to, i.e. that it describes, either as a DateTime or as a textual string indicating a time period in <a href=\"https://en.wikipedia.org/wiki/ISO_8601#Time_intervals\">ISO 8601 time interval format</a>. In\n      the case of a Dataset it will typically indicate the relevant time period in a precise notation (e.g. for a 2011 census dataset, the year 2011 would be written \"2011/2012\"). Other forms of content e.g. ScholarlyArticle, Book, TVSeries or TVEpisode may indicate their temporalCoverage in broader terms - textually or via well-known URL.\n      Written works such as books may sometimes have precise temporal coverage too, e.g. a work set in 1939 - 1945 can be indicated in ISO 8601 interval format format via \"1939/1945\".
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	TemporalCoverage interface{} `json:"temporalCoverage,omitempty" jsonld:"http://schema.org/temporalCoverage"`

	// The textual content of this CreativeWork.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Text string `json:"text,omitempty" jsonld:"http://schema.org/text"`

	// A thumbnail image relevant to the Thing.
	//
	// RangeIncludes:
	// - http://schema.org/URL
	//
	ThumbnailUrl string `json:"thumbnailUrl,omitempty" jsonld:"http://schema.org/thumbnailUrl"`

	// Approximate or typical time it takes to work with or through this learning resource for the typical intended target audience, e.g. 'P30M', 'P1H25M'.
	//
	// RangeIncludes:
	// - http://schema.org/Duration
	//
	TimeRequired *Duration `json:"timeRequired,omitempty" jsonld:"http://schema.org/timeRequired"`

	// The work that this work has been translated from. e.g. 物种起源 is a translationOf “On the Origin of Species”
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	//
	TranslationOfWork *CreativeWork `json:"translationOfWork,omitempty" jsonld:"http://schema.org/translationOfWork"`

	// Organization or person who adapts a creative work to different languages, regional differences and technical requirements of a target market, or that translates during some event.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/Person
	//
	Translator interface{} `json:"translator,omitempty" jsonld:"http://schema.org/translator"`

	// The typical expected age range, e.g. '7-9', '11-'.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	TypicalAgeRange string `json:"typicalAgeRange,omitempty" jsonld:"http://schema.org/typicalAgeRange"`

	// The version of the CreativeWork embodied by a specified resource.
	//
	// RangeIncludes:
	// - http://schema.org/Number
	// - http://schema.org/Text
	//
	Version interface{} `json:"version,omitempty" jsonld:"http://schema.org/version"`

	// An embedded video object.
	//
	// RangeIncludes:
	// - http://schema.org/VideoObject
	//
	Video *VideoObject `json:"video,omitempty" jsonld:"http://schema.org/video"`

	// Example/instance/realization/derivation of the concept of this creative work. eg. The paperback edition, first edition, or eBook.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	//
	WorkExample *CreativeWork `json:"workExample,omitempty" jsonld:"http://schema.org/workExample"`

	// A work that is a translation of the content of this work. e.g. 西遊記 has an English workTranslation “Journey to the West”,a German workTranslation “Monkeys Pilgerfahrt” and a Vietnamese  translation Tây du ký bình khảo.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	//
	WorkTranslation *CreativeWork `json:"workTranslation,omitempty" jsonld:"http://schema.org/workTranslation"`

}

type CreativeWork struct {
	MetaTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewCreativeWork() (x CreativeWork) {
	x.Type = "http://schema.org/CreativeWork"

	return
}
