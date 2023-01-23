package scoreApi

type Score struct {
	Lesson      Lesson  `json:"lesson,omitempty"`
	FirstScore  float32 `json:"firstScore"`
	SecondScore float32 `json:"secondScore"`
	IsAbsent    bool    `json:"isAbsent"`
}
