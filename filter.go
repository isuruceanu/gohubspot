package gohubspot

const (
	// Contact Operators:
	IsNotEmpty ContactOperator = "IS_NOT_EMPTY"
	IsEmpty    ContactOperator = "IS_EMPTY"
	Eq         ContactOperator = "EQ"
	NotEq      ContactOperator = "NEQ"
	Conatains  ContactOperator = "CONTAINS"
	Lt         ContactOperator = "LT"
	Lte        ContactOperator = "LTE"
	Gt         ContactOperator = "GT"
	SetAny     ContactOperator = "SET_ANY"
	SetNotAny  ContactOperator = "SET_NOT_ANY"
	SetEq      ContactOperator = "SET_EQ"
	SetNotEq   ContactOperator = "SET_NEQ"
	SetAll     ContactOperator = "SET_ALL"

	// ComputedProperty
	DateOfLastFormSubmission ComputedProperty = "DATE_OF_LAST_FORM_SUBMISSION"
	NumberOfFormsFilledOut   ComputedProperty = "NUMBER_OF_FORMS_FILLED_OUT"

	// ListMembership operator
	InList    ListMembershipOperator = "IN_LIST"
	NotInList ListMembershipOperator = "NOT_IN_LIST"

	// Form Submission
	HasFilledOutForm    FormSubmissionOperator = "HAS_FILLED_OUT_FORM"
	HasNotFilledOutForm FormSubmissionOperator = "HAS_NOT_FILLED_OUT_FORM"

	// Events
	HasEvent    EventsOpeerator = "HAS_EVENT"
	NotHasEvent EventsOpeerator = "NOT_HAS_EVENT"

	// PageViews

	PageviewEq              PageviewsOperator = "HAS_PAGEVIEW_EQ"
	PageviewContains        PageviewsOperator = "HAS_PAGEVIEW_CONTAINS"
	PageviewMatchesRegexp   PageviewsOperator = "HAS_PAGEVIEW_MATCHES_REGEX"
	PageviewNotEq           PageviewsOperator = "NOT_HAS_PAGEVIEW_EQ"
	PageviewNotContains     PageviewsOperator = "NOT_HAS_PAGEVIEW_CONTAINS"
	PageviewNotMatchRegexpt PageviewsOperator = "NOT_HAS_PAGEVIEW_MATCHES_REGEX"
)

type Filters struct {
}

type Filter struct {
}

type FilterItem struct {
	Operator FilterOperator
}

type ContactFilterItem struct {
	Operator         ContactOperator  `json:"operator"`
	Property         string           `json:"property"`
	ComputedProperty ComputedProperty `json:"computed_property"`
	Value            interface{}      `json:"value"`
	Type             string           `json:"type"`
}

type ListMembershipFilterItem struct {
	Operator ListMembershipOperator `json:"operator"`
	ListID   int                    `json:"list"`
}

type FormSubmissionFilterItem struct {
	Operator        FormSubmissionOperator `json:"operator"`
	FormID          int                    `json:"form"`
	PageID          int                    `json:"page"`
	AfterTimestamp  UnixTime               `json:"afterTimestamp"`
	Beforetimestamp UnixTime               `json:"beforeTimestamp"`
}

type EventsFilterItem struct {
	Operator             EventsOpeerator `json:"operator"`
	EventID              int             `json:"value"`
	FirstOccurenceAfter  UnixTime        `json:"firstOccurrenceAfterTimestamp"`
	FirstOccurenceBefore UnixTime        `json:"firstOccurrenceBeforeTimestamp"`
	LastOccurenceAfter   UnixTime        `json:"lastOccurrenceAfterTimestamp"`
	LastOccurenceBefore  UnixTime        `json:"lastOccurrenceBeforeTimestamp"`
	MinOccurence         int             `json:"minOccurrences"`
	MaxOccurence         int             `json:"maxOccurrences"`
}

type PageViewFilterItem struct {
	Operator             PageviewsOperator `json:"operator"`
	Value                string            `json:"value"`
	FirstOccurenceAfter  UnixTime          `json:"firstOccurrenceAfterTimestamp"`
	FirstOccurenceBefore UnixTime          `json:"firstOccurrenceBeforeTimestamp"`
	LastOccurenceAfter   UnixTime          `json:"lastOccurrenceAfterTimestamp"`
	LastOccurenceBefore  UnixTime          `json:"lastOccurrenceBeforeTimestamp"`
	MinOccurence         int               `json:"minOccurrences"`
	MaxOccurence         int               `json:"maxOccurrences"`
}

type FilterOperator string

type ComputedProperty string

type ContactOperator FilterOperator

type ListMembershipOperator FilterOperator

type FormSubmissionOperator FilterOperator

type EventsOpeerator FilterOperator

type PageviewsOperator FilterOperator
