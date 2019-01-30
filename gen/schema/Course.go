package schema

type CourseTrait struct {

	// The identifier for the <a class=\"localLink\" href=\"http://schema.org/Course\">Course</a> used by the course <a class=\"localLink\" href=\"http://schema.org/provider\">provider</a> (e.g. CS101 or 6.001).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	CourseCode string `json:"courseCode,omitempty" jsonld:"http://schema.org/courseCode"`

	// Requirements for taking the Course. May be completion of another <a class=\"localLink\" href=\"http://schema.org/Course\">Course</a> or a textual description like \"permission of instructor\". Requirements may be a pre-requisite competency, referenced using <a class=\"localLink\" href=\"http://schema.org/AlignmentObject\">AlignmentObject</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/AlignmentObject
	// - http://schema.org/Course
	//
	CoursePrerequisites interface{} `json:"coursePrerequisites,omitempty" jsonld:"http://schema.org/coursePrerequisites"`

	// A description of the qualification, award, certificate, diploma or other educational credential awarded as a consequence of successful completion of this course.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	EducationalCredentialAwarded interface{} `json:"educationalCredentialAwarded,omitempty" jsonld:"http://schema.org/educationalCredentialAwarded"`

	// An offering of the course at a specific time and place or through specific media or mode of study or to a specific section of students.
	//
	// RangeIncludes:
	// - http://schema.org/CourseInstance
	//
	HasCourseInstance *CourseInstance `json:"hasCourseInstance,omitempty" jsonld:"http://schema.org/hasCourseInstance"`

}

type Course struct {
	MetaTrait
	CourseTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewCourse() (x Course) {
	x.Type = "http://schema.org/Course"

	return
}
