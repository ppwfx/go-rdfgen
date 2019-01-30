package schema

type VehicleTrait struct {

	// The type of fuel suitable for the engine or engines of the vehicle. If the vehicle has only one engine, this property can be attached directly to the vehicle.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	// - http://schema.org/QualitativeValue
	//
	FuelType interface{} `json:"fuelType,omitempty" jsonld:"http://schema.org/fuelType"`

	// The date of production of the item, e.g. vehicle.
	//
	// RangeIncludes:
	// - http://schema.org/Date
	//
	ProductionDate *Date `json:"productionDate,omitempty" jsonld:"http://schema.org/productionDate"`

	// The date the item e.g. vehicle was purchased by the current owner.
	//
	// RangeIncludes:
	// - http://schema.org/Date
	//
	PurchaseDate *Date `json:"purchaseDate,omitempty" jsonld:"http://schema.org/purchaseDate"`

	// The drive wheel configuration, i.e. which roadwheels will receive torque from the vehicle's engine via the drivetrain.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/DriveWheelConfigurationValue
	//
	DriveWheelConfiguration interface{} `json:"driveWheelConfiguration,omitempty" jsonld:"http://schema.org/driveWheelConfiguration"`

	// A textual description of known damages, both repaired and unrepaired.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	KnownVehicleDamages string `json:"knownVehicleDamages,omitempty" jsonld:"http://schema.org/knownVehicleDamages"`

	// The number or type of airbags in the vehicle.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Number
	//
	NumberOfAirbags interface{} `json:"numberOfAirbags,omitempty" jsonld:"http://schema.org/numberOfAirbags"`

	// The type of component used for transmitting the power from a rotating power source to the wheels or other relevant component(s) (\"gearbox\" for cars).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	// - http://schema.org/QualitativeValue
	//
	VehicleTransmission interface{} `json:"vehicleTransmission,omitempty" jsonld:"http://schema.org/vehicleTransmission"`

	// The Vehicle Identification Number (VIN) is a unique serial number used by the automotive industry to identify individual motor vehicles.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	VehicleIdentificationNumber string `json:"vehicleIdentificationNumber,omitempty" jsonld:"http://schema.org/vehicleIdentificationNumber"`

	// A short text indicating the configuration of the vehicle, e.g. '5dr hatchback ST 2.5 MT 225 hp' or 'limited edition'.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	VehicleConfiguration string `json:"vehicleConfiguration,omitempty" jsonld:"http://schema.org/vehicleConfiguration"`

	// Indicates the design and body style of the vehicle (e.g. station wagon, hatchback, etc.).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	// - http://schema.org/QualitativeValue
	//
	BodyType interface{} `json:"bodyType,omitempty" jsonld:"http://schema.org/bodyType"`

	// Indicates whether the vehicle has been used for special purposes, like commercial rental, driving school, or as a taxi. The legislation in many countries requires this information to be revealed when offering a car for sale.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/CarUsageType
	//
	VehicleSpecialUsage interface{} `json:"vehicleSpecialUsage,omitempty" jsonld:"http://schema.org/vehicleSpecialUsage"`

	// The type or material of the interior of the vehicle (e.g. synthetic fabric, leather, wood, etc.). While most interior types are characterized by the material used, an interior type can also be based on vehicle usage or target audience.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	VehicleInteriorType string `json:"vehicleInteriorType,omitempty" jsonld:"http://schema.org/vehicleInteriorType"`

	// The color or color combination of the interior of the vehicle.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	VehicleInteriorColor string `json:"vehicleInteriorColor,omitempty" jsonld:"http://schema.org/vehicleInteriorColor"`

	// Indicates that the vehicle meets the respective emission standard.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	// - http://schema.org/QualitativeValue
	//
	MeetsEmissionStandard interface{} `json:"meetsEmissionStandard,omitempty" jsonld:"http://schema.org/meetsEmissionStandard"`

	// The number of doors.<br/><br/>\n\nTypical unit code(s): C62
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	// - http://schema.org/Number
	//
	NumberOfDoors interface{} `json:"numberOfDoors,omitempty" jsonld:"http://schema.org/numberOfDoors"`

	// The number of axles.<br/><br/>\n\nTypical unit code(s): C62
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	// - http://schema.org/Number
	//
	NumberOfAxles interface{} `json:"numberOfAxles,omitempty" jsonld:"http://schema.org/numberOfAxles"`

	// The capacity of the fuel tank or in the case of electric cars, the battery. If there are multiple components for storage, this should indicate the total of all storage of the same type.<br/><br/>\n\nTypical unit code(s): LTR for liters, GLL of US gallons, GLI for UK / imperial gallons, AMH for ampere-hours (for electrical vehicles).
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	//
	FuelCapacity *QuantitativeValue `json:"fuelCapacity,omitempty" jsonld:"http://schema.org/fuelCapacity"`

	// The permitted weight of passengers and cargo, EXCLUDING the weight of the empty vehicle.<br/><br/>\n\nTypical unit code(s): KGM for kilogram, LBR for pound<br/><br/>\n\n<ul>\n<li>Note 1: Many databases specify the permitted TOTAL weight instead, which is the sum of <a class=\"localLink\" href=\"http://schema.org/weight\">weight</a> and <a class=\"localLink\" href=\"http://schema.org/payload\">payload</a></li>\n<li>Note 2: You can indicate additional information in the <a class=\"localLink\" href=\"http://schema.org/name\">name</a> of the <a class=\"localLink\" href=\"http://schema.org/QuantitativeValue\">QuantitativeValue</a> node.</li>\n<li>Note 3: You may also link to a <a class=\"localLink\" href=\"http://schema.org/QualitativeValue\">QualitativeValue</a> node that provides additional information using <a class=\"localLink\" href=\"http://schema.org/valueReference\">valueReference</a>.</li>\n<li>Note 4: Note that you can use <a class=\"localLink\" href=\"http://schema.org/minValue\">minValue</a> and <a class=\"localLink\" href=\"http://schema.org/maxValue\">maxValue</a> to indicate ranges.</li>\n</ul>\n
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	//
	Payload *QuantitativeValue `json:"payload,omitempty" jsonld:"http://schema.org/payload"`

	// The distance between the centers of the front and rear wheels.<br/><br/>\n\nTypical unit code(s): CMT for centimeters, MTR for meters, INH for inches, FOT for foot/feet
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	//
	Wheelbase *QuantitativeValue `json:"wheelbase,omitempty" jsonld:"http://schema.org/wheelbase"`

	// The speed range of the vehicle. If the vehicle is powered by an engine, the upper limit of the speed range (indicated by <a class=\"localLink\" href=\"http://schema.org/maxValue\">maxValue</a> should be the maximum speed achievable under regular conditions.<br/><br/>\n\nTypical unit code(s): KMH for km/h, HM for mile per hour (0.447 04 m/s), KNT for knot<br/><br/>\n\n*Note 1: Use <a class=\"localLink\" href=\"http://schema.org/minValue\">minValue</a> and <a class=\"localLink\" href=\"http://schema.org/maxValue\">maxValue</a> to indicate the range. Typically, the minimal value is zero.\n* Note 2: There are many different ways of measuring the speed range. You can link to information about how the given value has been determined using the <a class=\"localLink\" href=\"http://schema.org/valueReference\">valueReference</a> property.
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	//
	Speed *QuantitativeValue `json:"speed,omitempty" jsonld:"http://schema.org/speed"`

	// The permitted vertical load (TWR) of a trailer attached to the vehicle. Also referred to as Tongue Load Rating (TLR) or Vertical Load Rating (VLR)<br/><br/>\n\nTypical unit code(s): KGM for kilogram, LBR for pound<br/><br/>\n\n<ul>\n<li>Note 1: You can indicate additional information in the <a class=\"localLink\" href=\"http://schema.org/name\">name</a> of the <a class=\"localLink\" href=\"http://schema.org/QuantitativeValue\">QuantitativeValue</a> node.</li>\n<li>Note 2: You may also link to a <a class=\"localLink\" href=\"http://schema.org/QualitativeValue\">QualitativeValue</a> node that provides additional information using <a class=\"localLink\" href=\"http://schema.org/valueReference\">valueReference</a>.</li>\n<li>Note 3: Note that you can use <a class=\"localLink\" href=\"http://schema.org/minValue\">minValue</a> and <a class=\"localLink\" href=\"http://schema.org/maxValue\">maxValue</a> to indicate ranges.</li>\n</ul>\n
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	//
	TongueWeight *QuantitativeValue `json:"tongueWeight,omitempty" jsonld:"http://schema.org/tongueWeight"`

	// The time needed to accelerate the vehicle from a given start velocity to a given target velocity.<br/><br/>\n\nTypical unit code(s): SEC for seconds<br/><br/>\n\n<ul>\n<li>Note: There are unfortunately no standard unit codes for seconds/0..100 km/h or seconds/0..60 mph. Simply use \"SEC\" for seconds and indicate the velocities in the <a class=\"localLink\" href=\"http://schema.org/name\">name</a> of the <a class=\"localLink\" href=\"http://schema.org/QuantitativeValue\">QuantitativeValue</a>, or use <a class=\"localLink\" href=\"http://schema.org/valueReference\">valueReference</a> with a <a class=\"localLink\" href=\"http://schema.org/QuantitativeValue\">QuantitativeValue</a> of 0..60 mph or 0..100 km/h to specify the reference speeds.</li>\n</ul>\n
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	//
	AccelerationTime *QuantitativeValue `json:"accelerationTime,omitempty" jsonld:"http://schema.org/accelerationTime"`

	// The total distance travelled by the particular vehicle since its initial production, as read from its odometer.<br/><br/>\n\nTypical unit code(s): KMT for kilometers, SMI for statute miles
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	//
	MileageFromOdometer *QuantitativeValue `json:"mileageFromOdometer,omitempty" jsonld:"http://schema.org/mileageFromOdometer"`

	// The amount of fuel consumed for traveling a particular distance or temporal duration with the given vehicle (e.g. liters per 100 km).<br/><br/>\n\n<ul>\n<li>Note 1: There are unfortunately no standard unit codes for liters per 100 km.  Use <a class=\"localLink\" href=\"http://schema.org/unitText\">unitText</a> to indicate the unit of measurement, e.g. L/100 km.</li>\n<li>Note 2: There are two ways of indicating the fuel consumption, <a class=\"localLink\" href=\"http://schema.org/fuelConsumption\">fuelConsumption</a> (e.g. 8 liters per 100 km) and <a class=\"localLink\" href=\"http://schema.org/fuelEfficiency\">fuelEfficiency</a> (e.g. 30 miles per gallon). They are reciprocal.</li>\n<li>Note 3: Often, the absolute value is useful only when related to driving speed (\"at 80 km/h\") or usage pattern (\"city traffic\"). You can use <a class=\"localLink\" href=\"http://schema.org/valueReference\">valueReference</a> to link the value for the fuel consumption to another value.</li>\n</ul>\n
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	//
	FuelConsumption *QuantitativeValue `json:"fuelConsumption,omitempty" jsonld:"http://schema.org/fuelConsumption"`

	// The permitted weight of a trailer attached to the vehicle.<br/><br/>\n\nTypical unit code(s): KGM for kilogram, LBR for pound\n* Note 1: You can indicate additional information in the <a class=\"localLink\" href=\"http://schema.org/name\">name</a> of the <a class=\"localLink\" href=\"http://schema.org/QuantitativeValue\">QuantitativeValue</a> node.\n* Note 2: You may also link to a <a class=\"localLink\" href=\"http://schema.org/QualitativeValue\">QualitativeValue</a> node that provides additional information using <a class=\"localLink\" href=\"http://schema.org/valueReference\">valueReference</a>.\n* Note 3: Note that you can use <a class=\"localLink\" href=\"http://schema.org/minValue\">minValue</a> and <a class=\"localLink\" href=\"http://schema.org/maxValue\">maxValue</a> to indicate ranges.
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	//
	TrailerWeight *QuantitativeValue `json:"trailerWeight,omitempty" jsonld:"http://schema.org/trailerWeight"`

	// The distance traveled per unit of fuel used; most commonly miles per gallon (mpg) or kilometers per liter (km/L).<br/><br/>\n\n<ul>\n<li>Note 1: There are unfortunately no standard unit codes for miles per gallon or kilometers per liter. Use <a class=\"localLink\" href=\"http://schema.org/unitText\">unitText</a> to indicate the unit of measurement, e.g. mpg or km/L.</li>\n<li>Note 2: There are two ways of indicating the fuel consumption, <a class=\"localLink\" href=\"http://schema.org/fuelConsumption\">fuelConsumption</a> (e.g. 8 liters per 100 km) and <a class=\"localLink\" href=\"http://schema.org/fuelEfficiency\">fuelEfficiency</a> (e.g. 30 miles per gallon). They are reciprocal.</li>\n<li>Note 3: Often, the absolute value is useful only when related to driving speed (\"at 80 km/h\") or usage pattern (\"city traffic\"). You can use <a class=\"localLink\" href=\"http://schema.org/valueReference\">valueReference</a> to link the value for the fuel economy to another value.</li>\n</ul>\n
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	//
	FuelEfficiency *QuantitativeValue `json:"fuelEfficiency,omitempty" jsonld:"http://schema.org/fuelEfficiency"`

	// The available volume for cargo or luggage. For automobiles, this is usually the trunk volume.<br/><br/>\n\nTypical unit code(s): LTR for liters, FTQ for cubic foot/feet<br/><br/>\n\nNote: You can use <a class=\"localLink\" href=\"http://schema.org/minValue\">minValue</a> and <a class=\"localLink\" href=\"http://schema.org/maxValue\">maxValue</a> to indicate ranges.
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	//
	CargoVolume *QuantitativeValue `json:"cargoVolume,omitempty" jsonld:"http://schema.org/cargoVolume"`

	// The number of owners of the vehicle, including the current one.<br/><br/>\n\nTypical unit code(s): C62
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	// - http://schema.org/Number
	//
	NumberOfPreviousOwners interface{} `json:"numberOfPreviousOwners,omitempty" jsonld:"http://schema.org/numberOfPreviousOwners"`

	// The total number of forward gears available for the transmission system of the vehicle.<br/><br/>\n\nTypical unit code(s): C62
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	// - http://schema.org/Number
	//
	NumberOfForwardGears interface{} `json:"numberOfForwardGears,omitempty" jsonld:"http://schema.org/numberOfForwardGears"`

	// The number of passengers that can be seated in the vehicle, both in terms of the physical space available, and in terms of limitations set by law.<br/><br/>\n\nTypical unit code(s): C62 for persons.
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	// - http://schema.org/Number
	//
	VehicleSeatingCapacity interface{} `json:"vehicleSeatingCapacity,omitempty" jsonld:"http://schema.org/vehicleSeatingCapacity"`

	// The number of persons that can be seated (e.g. in a vehicle), both in terms of the physical space available, and in terms of limitations set by law.<br/><br/>\n\nTypical unit code(s): C62 for persons
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	// - http://schema.org/Number
	//
	SeatingCapacity interface{} `json:"seatingCapacity,omitempty" jsonld:"http://schema.org/seatingCapacity"`

	// The permitted total weight of the loaded vehicle, including passengers and cargo and the weight of the empty vehicle.<br/><br/>\n\nTypical unit code(s): KGM for kilogram, LBR for pound<br/><br/>\n\n<ul>\n<li>Note 1: You can indicate additional information in the <a class=\"localLink\" href=\"http://schema.org/name\">name</a> of the <a class=\"localLink\" href=\"http://schema.org/QuantitativeValue\">QuantitativeValue</a> node.</li>\n<li>Note 2: You may also link to a <a class=\"localLink\" href=\"http://schema.org/QualitativeValue\">QualitativeValue</a> node that provides additional information using <a class=\"localLink\" href=\"http://schema.org/valueReference\">valueReference</a>.</li>\n<li>Note 3: Note that you can use <a class=\"localLink\" href=\"http://schema.org/minValue\">minValue</a> and <a class=\"localLink\" href=\"http://schema.org/maxValue\">maxValue</a> to indicate ranges.</li>\n</ul>\n
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	//
	WeightTotal *QuantitativeValue `json:"weightTotal,omitempty" jsonld:"http://schema.org/weightTotal"`

	// Information about the engine or engines of the vehicle.
	//
	// RangeIncludes:
	// - http://schema.org/EngineSpecification
	//
	VehicleEngine *EngineSpecification `json:"vehicleEngine,omitempty" jsonld:"http://schema.org/vehicleEngine"`

	// The date of the first registration of the vehicle with the respective public authorities.
	//
	// RangeIncludes:
	// - http://schema.org/Date
	//
	DateVehicleFirstRegistered *Date `json:"dateVehicleFirstRegistered,omitempty" jsonld:"http://schema.org/dateVehicleFirstRegistered"`

	// The CO2 emissions in g/km. When used in combination with a QuantitativeValue, put \"g/km\" into the unitText property of that value, since there is no UN/CEFACT Common Code for \"g/km\".
	//
	// RangeIncludes:
	// - http://schema.org/Number
	//
	EmissionsCO2 float64 `json:"emissionsCO2,omitempty" jsonld:"http://schema.org/emissionsCO2"`

	// The release date of a vehicle model (often used to differentiate versions of the same make and model).
	//
	// RangeIncludes:
	// - http://schema.org/Date
	//
	ModelDate *Date `json:"modelDate,omitempty" jsonld:"http://schema.org/modelDate"`

	// The position of the steering wheel or similar device (mostly for cars).
	//
	// RangeIncludes:
	// - http://schema.org/SteeringPositionValue
	//
	SteeringPosition *SteeringPositionValue `json:"steeringPosition,omitempty" jsonld:"http://schema.org/steeringPosition"`

	// The release date of a vehicle model (often used to differentiate versions of the same make and model).
	//
	// RangeIncludes:
	// - http://schema.org/Date
	//
	VehicleModelDate *Date `json:"vehicleModelDate,omitempty" jsonld:"http://schema.org/vehicleModelDate"`

}

type Vehicle struct {
	MetaTrait
	VehicleTrait
	ProductTrait
	ThingTrait
	AdditionalTrait
}

func NewVehicle() (x Vehicle) {
	x.Type = "http://schema.org/Vehicle"

	return
}
