package schema

type BroadcastServiceTrait struct {

	// The frequency used for over-the-air broadcasts. Numeric values or simple ranges e.g. 87-99. In addition a shortcut idiom is supported for frequences of AM and FM radio channels, e.g. \"87 FM\".
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/BroadcastFrequencySpecification
	//
	BroadcastFrequency interface{} `json:"broadcastFrequency,omitempty" jsonld:"http://schema.org/broadcastFrequency"`

	// The type of screening or video broadcast used (e.g. IMAX, 3D, SD, HD, etc.).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	VideoFormat string `json:"videoFormat,omitempty" jsonld:"http://schema.org/videoFormat"`

	// The name displayed in the channel guide. For many US affiliates, it is the network name.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	BroadcastDisplayName string `json:"broadcastDisplayName,omitempty" jsonld:"http://schema.org/broadcastDisplayName"`

	// The timezone in <a href=\"http://en.wikipedia.org/wiki/ISO_8601\">ISO 8601 format</a> for which the service bases its broadcasts
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	BroadcastTimezone string `json:"broadcastTimezone,omitempty" jsonld:"http://schema.org/broadcastTimezone"`

	// The area within which users can expect to reach the broadcast service.
	//
	// RangeIncludes:
	// - http://schema.org/Place
	//
	Area *Place `json:"area,omitempty" jsonld:"http://schema.org/area"`

	// A broadcast channel of a broadcast service.
	//
	// RangeIncludes:
	// - http://schema.org/BroadcastChannel
	//
	HasBroadcastChannel *BroadcastChannel `json:"hasBroadcastChannel,omitempty" jsonld:"http://schema.org/hasBroadcastChannel"`

	// The media network(s) whose content is broadcast on this station.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	//
	BroadcastAffiliateOf *Organization `json:"broadcastAffiliateOf,omitempty" jsonld:"http://schema.org/broadcastAffiliateOf"`

	// The organization owning or operating the broadcast service.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	//
	Broadcaster *Organization `json:"broadcaster,omitempty" jsonld:"http://schema.org/broadcaster"`

	// A broadcast service to which the broadcast service may belong to such as regional variations of a national channel.
	//
	// RangeIncludes:
	// - http://schema.org/BroadcastService
	//
	ParentService *BroadcastService `json:"parentService,omitempty" jsonld:"http://schema.org/parentService"`

}

type BroadcastService struct {
	MetaTrait
	BroadcastServiceTrait
	ServiceTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewBroadcastService() (x BroadcastService) {
	x.Type = "http://schema.org/BroadcastService"

	return
}
