package contactproperty

type DataType string

type FieldType string

const (
	TextAreaField        FieldType = "textarea"
	TextField            FieldType = "text"
	DateField            FieldType = "date"
	FileField            FieldType = "file"
	NumberField          FieldType = "number"
	SelectField          FieldType = "select"
	RadioField           FieldType = "radio"
	CheckboxField        FieldType = "checkbox"
	BooleanCheckboxField FieldType = "booleancheckbox"
)

const (
	String      DataType = "string"
	Number      DataType = "number"
	Date        DataType = "date"
	DateTime    DataType = "datetime"
	Enumeration DataType = "enumeration"
)
