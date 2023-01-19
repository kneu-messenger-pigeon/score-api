package scoreApi

type Lesson struct {
	Id   int        `json:"id"`
	Date string     `json:"date"`
	Type LessonType `json:"type"`
}
