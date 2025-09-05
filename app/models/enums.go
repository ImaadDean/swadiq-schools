package models

// AttendanceStatus defines the possible status values for attendance.
type AttendanceStatus string

const (
	Present AttendanceStatus = "present"
	Absent  AttendanceStatus = "absent"
	Late    AttendanceStatus = "late"
)

// RecipientType defines the possible recipient types for notifications.
type RecipientType string

const (
	StudentRecipient RecipientType = "student"
	ParentRecipient  RecipientType = "parent"
	TeacherRecipient RecipientType = "teacher"
)

// DayOfWeek defines the days of the week for schedules.
type DayOfWeek string

const (
	Monday    DayOfWeek = "Monday"
	Tuesday   DayOfWeek = "Tuesday"
	Wednesday DayOfWeek = "Wednesday"
	Thursday  DayOfWeek = "Thursday"
	Friday    DayOfWeek = "Friday"
	Saturday  DayOfWeek = "Saturday"
	Sunday    DayOfWeek = "Sunday"
)

// Gender defines the possible gender values for a student.
type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
	Other  Gender = "other"
)

// RelationshipType defines the relationship of a parent/guardian to a student.
type RelationshipType string

const (
	Father   RelationshipType = "Father"
	Mother   RelationshipType = "Mother"
	Guardian RelationshipType = "Guardian"
	Brother  RelationshipType = "Brother"
	Sister   RelationshipType = "Sister"
	OtherRel RelationshipType = "Other"
)
