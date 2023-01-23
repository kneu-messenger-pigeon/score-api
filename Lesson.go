package scoreApi

import "time"

type Lesson struct {
	Id   int        `json:"id"`
	Date time.Time  `json:"date"`
	Type LessonType `json:"type"`
}
