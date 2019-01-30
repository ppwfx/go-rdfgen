package schema

type JobPostingTrait struct {

	// The date after when the item is not valid. For example the end of an offer, salary period, or a period of opening hours.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	ValidThrough *DateTime `json:"validThrough,omitempty" jsonld:"http://schema.org/validThrough"`

	// The base salary of the job or of an employee in an EmployeeRole.
	//
	// RangeIncludes:
	// - http://schema.org/MonetaryAmount
	// - http://schema.org/Number
	// - http://schema.org/PriceSpecification
	//
	BaseSalary interface{} `json:"baseSalary,omitempty" jsonld:"http://schema.org/baseSalary"`

	// Description of benefits associated with the job.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Benefits string `json:"benefits,omitempty" jsonld:"http://schema.org/benefits"`

	// Publication date for the job posting.
	//
	// RangeIncludes:
	// - http://schema.org/Date
	//
	DatePosted *Date `json:"datePosted,omitempty" jsonld:"http://schema.org/datePosted"`

	// Educational background needed for the position or Occupation.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	EducationRequirements string `json:"educationRequirements,omitempty" jsonld:"http://schema.org/educationRequirements"`

	// Type of employment (e.g. full-time, part-time, contract, temporary, seasonal, internship).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	EmploymentType string `json:"employmentType,omitempty" jsonld:"http://schema.org/employmentType"`

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

	// Organization offering the job position.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	//
	HiringOrganization *Organization `json:"hiringOrganization,omitempty" jsonld:"http://schema.org/hiringOrganization"`

	// Description of bonus and commission compensation aspects of the job.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	IncentiveCompensation string `json:"incentiveCompensation,omitempty" jsonld:"http://schema.org/incentiveCompensation"`

	// Description of bonus and commission compensation aspects of the job.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Incentives string `json:"incentives,omitempty" jsonld:"http://schema.org/incentives"`

	// The industry associated with the job position.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Industry string `json:"industry,omitempty" jsonld:"http://schema.org/industry"`

	// Description of benefits associated with the job.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	JobBenefits string `json:"jobBenefits,omitempty" jsonld:"http://schema.org/jobBenefits"`

	// A (typically single) geographic location associated with the job position.
	//
	// RangeIncludes:
	// - http://schema.org/Place
	//
	JobLocation *Place `json:"jobLocation,omitempty" jsonld:"http://schema.org/jobLocation"`

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

	// The Occupation for the JobPosting.
	//
	// RangeIncludes:
	// - http://schema.org/Occupation
	//
	RelevantOccupation *Occupation `json:"relevantOccupation,omitempty" jsonld:"http://schema.org/relevantOccupation"`

	// Responsibilities associated with this role or Occupation.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Responsibilities string `json:"responsibilities,omitempty" jsonld:"http://schema.org/responsibilities"`

	// The currency (coded using <a href=\"http://en.wikipedia.org/wiki/ISO_4217\">ISO 4217</a> ) used for the main salary information in this job posting or for this employee.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	SalaryCurrency string `json:"salaryCurrency,omitempty" jsonld:"http://schema.org/salaryCurrency"`

	// Skills required to fulfill this role or in this Occupation.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Skills string `json:"skills,omitempty" jsonld:"http://schema.org/skills"`

	// Any special commitments associated with this job posting. Valid entries include VeteranCommit, MilitarySpouseCommit, etc.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	SpecialCommitments string `json:"specialCommitments,omitempty" jsonld:"http://schema.org/specialCommitments"`

	// The title of the job.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Title string `json:"title,omitempty" jsonld:"http://schema.org/title"`

	// The typical working hours for this job (e.g. 1st shift, night shift, 8am-5pm).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	WorkHours string `json:"workHours,omitempty" jsonld:"http://schema.org/workHours"`

}

type JobPosting struct {
	MetaTrait
	JobPostingTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewJobPosting() (x JobPosting) {
	x.Type = "http://schema.org/JobPosting"

	return
}
