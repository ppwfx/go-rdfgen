package schema

type BroadcastChannelTrait struct {

	// Genre of the creative work, broadcast channel or group.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	Genre interface{} `json:"genre,omitempty" jsonld:"http://schema.org/genre"`

	// The unique address by which the BroadcastService can be identified in a provider lineup. In US, this is typically a number.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	BroadcastChannelId string `json:"broadcastChannelId,omitempty" jsonld:"http://schema.org/broadcastChannelId"`

	// The frequency used for over-the-air broadcasts. Numeric values or simple ranges e.g. 87-99. In addition a shortcut idiom is supported for frequences of AM and FM radio channels, e.g. \"87 FM\".
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/BroadcastFrequencySpecification
	//
	BroadcastFrequency interface{} `json:"broadcastFrequency,omitempty" jsonld:"http://schema.org/broadcastFrequency"`

	// The type of service required to have access to the channel (e.g. Standard or Premium).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	BroadcastServiceTier string `json:"broadcastServiceTier,omitempty" jsonld:"http://schema.org/broadcastServiceTier"`

	// The CableOrSatelliteService offering the channel.
	//
	// RangeIncludes:
	// - http://schema.org/CableOrSatelliteService
	//
	InBroadcastLineup *CableOrSatelliteService `json:"inBroadcastLineup,omitempty" jsonld:"http://schema.org/inBroadcastLineup"`

	// The BroadcastService offered on this channel.
	//
	// RangeIncludes:
	// - http://schema.org/BroadcastService
	//
	ProvidesBroadcastService *BroadcastService `json:"providesBroadcastService,omitempty" jsonld:"http://schema.org/providesBroadcastService"`

}

type BroadcastChannel struct {
	MetaTrait
	BroadcastChannelTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewBroadcastChannel() (x BroadcastChannel) {
	x.Type = "http://schema.org/BroadcastChannel"

	return
}
