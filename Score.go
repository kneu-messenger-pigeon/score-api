package scoreApi

type Score struct {
	Lesson     Lesson  `json:"lesson,omitempty"`
	LessonHalf int     `json:"lessonHalf"`
	Score      float32 `json:"score"`
	IsAbsent   bool    `json:"isAbsent"`
}
