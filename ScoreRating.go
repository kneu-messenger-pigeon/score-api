package scoreApi

type ScoreRating struct {
	Total         float32 `json:"total"`
	MinTotal      float32 `json:"minTotal"`
	MaxTotal      float32 `json:"maxTotal"`
	Rating        int     `json:"rating"`
	StudentsCount int     `json:"studentsCount"`
}
