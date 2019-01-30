package schema

type CourseInstanceTrait struct {

	// The medium or means of delivery of the course instance or the mode of study, either as a text label (e.g. \"online\", \"onsite\" or \"blended\"; \"synchronous\" or \"asynchronous\"; \"full-time\" or \"part-time\") or as a URL reference to a term from a controlled vocabulary (e.g. https://ceds.ed.gov/element/001311#Asynchronous ).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	CourseMode interface{} `json:"courseMode,omitempty" jsonld:"http://schema.org/courseMode"`

	// A person assigned to instruct or provide instructional assistance for the <a class=\"localLink\" href=\"http://schema.org/CourseInstance\">CourseInstance</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Instructor *Person `json:"instructor,omitempty" jsonld:"http://schema.org/instructor"`

}

type CourseInstance struct {
	MetaTrait
	CourseInstanceTrait
	EventTrait
	ThingTrait
	AdditionalTrait
}

func NewCourseInstance() (x CourseInstance) {
	x.Type = "http://schema.org/CourseInstance"

	return
}
