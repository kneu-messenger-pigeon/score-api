package scoreApi

type DisciplineScore struct {
	Discipline Discipline `json:"discipline"`
	Score      Score      `json:"scores,omitempty"`
}
