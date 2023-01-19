package scoreApi

type Lesson struct {
	Id     int        `json:"id"`
	Date   string     `json:"date"`
	TypeId int        `json:"typeId"`
	Type   LessonType `json:"type"`
}
