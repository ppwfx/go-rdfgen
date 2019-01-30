package gen

type Field struct {
	Comment string
	Name string
	IsBasicType bool
	Type string
	JsonTag string
	JsonLdTag string
}

type Struct struct {
	Extends []string
	Package string
	Type string
	Name   string
	Fields []Field
}
