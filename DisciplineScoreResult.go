package scoreApi

type DisciplineScoreResult struct {
	Discipline  Discipline  `json:"discipline"`
	ScoreRating ScoreRating `json:"scoreRating"`
	Scores      []Score     `json:"scores,omitempty"`
}
