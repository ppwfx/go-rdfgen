package schema

type OccupationTrait struct {

	// Educational background needed for the position or Occupation.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	EducationRequirements string `json:"educationRequirements,omitempty" jsonld:"http://schema.org/educationRequirements"`

	// A property describing the estimated salary for a job posting based on a variety of variables including, but not limited to industry, job title, and location. The estimated salary is usually computed by outside organizations and therefore the hiring organization is not bound to this estimated salary.
	//
	// RangeIncludes:
	// - http://schema.org/MonetaryAmount
	// - http://schema.org/MonetaryAmountDistribution
	// - http://schema.org/Number
	// - http://schema.org/PriceSpecification
	//
	EstimatedSalary interface{} `json:"estimatedSalary,omitempty" jsonld:"http://schema.org/estimatedSalary"`

	// Description of skills and experience needed for the position or Occupation.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	ExperienceRequirements string `json:"experienceRequirements,omitempty" jsonld:"http://schema.org/experienceRequirements"`

	// Category or categories describing the job. Use BLS O*NET-SOC taxonomy: http://www.onetcenter.org/taxonomy.html. Ideally includes textual label and formal code, with the property repeated for each applicable value.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	OccupationalCategory string `json:"occupationalCategory,omitempty" jsonld:"http://schema.org/occupationalCategory"`

	// Specific qualifications required for this role or Occupation.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Qualifications string `json:"qualifications,omitempty" jsonld:"http://schema.org/qualifications"`

	// Responsibilities associated with this role or Occupation.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Responsibilities string `json:"responsibilities,omitempty" jsonld:"http://schema.org/responsibilities"`

	// Skills required to fulfill this role or in this Occupation.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Skills string `json:"skills,omitempty" jsonld:"http://schema.org/skills"`

	// The region/country for which this occupational description is appropriate. Note that educational requirements and qualifications can vary between jurisdictions.
	//
	// RangeIncludes:
	// - http://schema.org/AdministrativeArea
	//
	OccupationLocation *AdministrativeArea `json:"occupationLocation,omitempty" jsonld:"http://schema.org/occupationLocation"`

}

type Occupation struct {
	MetaTrait
	OccupationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewOccupation() (x Occupation) {
	x.Type = "http://schema.org/Occupation"

	return
}
