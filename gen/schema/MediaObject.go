package schema

type MediaObjectTrait struct {

	// A NewsArticle associated with the Media Object.
	//
	// RangeIncludes:
	// - http://schema.org/NewsArticle
	//
	AssociatedArticle *NewsArticle `json:"associatedArticle,omitempty" jsonld:"http://schema.org/associatedArticle"`

	// The bitrate of the media object.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Bitrate string `json:"bitrate,omitempty" jsonld:"http://schema.org/bitrate"`

	// File size in (mega/kilo) bytes.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	ContentSize string `json:"contentSize,omitempty" jsonld:"http://schema.org/contentSize"`

	// Actual bytes of the media object, for example the image file or video file.
	//
	// RangeIncludes:
	// - http://schema.org/URL
	//
	ContentUrl string `json:"contentUrl,omitempty" jsonld:"http://schema.org/contentUrl"`

	// The duration of the item (movie, audio recording, event, etc.) in <a href=\"http://en.wikipedia.org/wiki/ISO_8601\">ISO 8601 date format</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Duration
	//
	Duration *Duration `json:"duration,omitempty" jsonld:"http://schema.org/duration"`

	// A URL pointing to a player for a specific video. In general, this is the information in the <code>src</code> element of an <code>embed</code> tag and should not be the same as the content of the <code>loc</code> tag.
	//
	// RangeIncludes:
	// - http://schema.org/URL
	//
	EmbedUrl string `json:"embedUrl,omitempty" jsonld:"http://schema.org/embedUrl"`

	// The CreativeWork encoded by this media object.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	//
	EncodesCreativeWork *CreativeWork `json:"encodesCreativeWork,omitempty" jsonld:"http://schema.org/encodesCreativeWork"`

	// Media type typically expressed using a MIME format (see <a href=\"http://www.iana.org/assignments/media-types/media-types.xhtml\">IANA site</a> and <a href=\"https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types\">MDN reference</a>) e.g. application/zip for a SoftwareApplication binary, audio/mpeg for .mp3 etc.).<br/><br/>\n\nIn cases where a <a class=\"localLink\" href=\"http://schema.org/CreativeWork\">CreativeWork</a> has several media type representations, <a class=\"localLink\" href=\"http://schema.org/encoding\">encoding</a> can be used to indicate each <a class=\"localLink\" href=\"http://schema.org/MediaObject\">MediaObject</a> alongside particular <a class=\"localLink\" href=\"http://schema.org/encodingFormat\">encodingFormat</a> information.<br/><br/>\n\nUnregistered or niche encoding and file formats can be indicated instead via the most appropriate URL, e.g. defining Web page or a Wikipedia/Wikidata entry.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	EncodingFormat interface{} `json:"encodingFormat,omitempty" jsonld:"http://schema.org/encodingFormat"`

	// The height of the item.
	//
	// RangeIncludes:
	// - http://schema.org/Distance
	// - http://schema.org/QuantitativeValue
	//
	Height interface{} `json:"height,omitempty" jsonld:"http://schema.org/height"`

	// Player type required&#x2014;for example, Flash or Silverlight.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	PlayerType string `json:"playerType,omitempty" jsonld:"http://schema.org/playerType"`

	// The production company or studio responsible for the item e.g. series, video game, episode etc.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	//
	ProductionCompany *Organization `json:"productionCompany,omitempty" jsonld:"http://schema.org/productionCompany"`

	// The regions where the media is allowed. If not specified, then it's assumed to be allowed everywhere. Specify the countries in <a href=\"http://en.wikipedia.org/wiki/ISO_3166\">ISO 3166 format</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Place
	//
	RegionsAllowed *Place `json:"regionsAllowed,omitempty" jsonld:"http://schema.org/regionsAllowed"`

	// Indicates if use of the media require a subscription  (either paid or free). Allowed values are <code>true</code> or <code>false</code> (note that an earlier version had 'yes', 'no').
	//
	// RangeIncludes:
	// - http://schema.org/Boolean
	// - http://schema.org/MediaSubscription
	//
	RequiresSubscription interface{} `json:"requiresSubscription,omitempty" jsonld:"http://schema.org/requiresSubscription"`

	// Date when this media object was uploaded to this site.
	//
	// RangeIncludes:
	// - http://schema.org/Date
	//
	UploadDate *Date `json:"uploadDate,omitempty" jsonld:"http://schema.org/uploadDate"`

	// The width of the item.
	//
	// RangeIncludes:
	// - http://schema.org/Distance
	// - http://schema.org/QuantitativeValue
	//
	Width interface{} `json:"width,omitempty" jsonld:"http://schema.org/width"`

}

type MediaObject struct {
	MetaTrait
	MediaObjectTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewMediaObject() (x MediaObject) {
	x.Type = "http://schema.org/MediaObject"

	return
}
