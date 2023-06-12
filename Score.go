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

func (score *Score) IsEqual(other *Score) bool {
	return score.IsDeleted() == other.IsDeleted() &&
		score.IsAbsent == other.IsAbsent &&
		compareScoreValue(score.FirstScore, other.FirstScore) &&
		compareScoreValue(score.SecondScore, other.SecondScore)
}

func compareScoreValue(this *float32, other *float32) bool {
	return (this == nil) == (other == nil) && // one of them is not nil
		(this == other || *this == *other) // both are nil or have same value
}
