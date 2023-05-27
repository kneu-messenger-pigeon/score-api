package scoreApi

type Score struct {
	Lesson      Lesson   `json:"lesson,omitempty"`
	FirstScore  *float32 `json:"firstScore"`
	SecondScore *float32 `json:"secondScore"`
	IsAbsent    bool     `json:"isAbsent"`
}

func (score *Score) IsDeleted() bool {
	return score.IsAbsent == false && score.FirstScore == nil && score.SecondScore == nil
}
