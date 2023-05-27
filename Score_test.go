package scoreApi

import "testing"

func TestScore_IsDeleted(t *testing.T) {
	t.Run("deleted", func(t *testing.T) {
		score := Score{}
		if score.IsDeleted() == false {
			t.Errorf("Expected deleted score, score: %v ", score)
		}
	})

	t.Run("First_score_present", func(t *testing.T) {
		scoreValue := float32(9)
		score := Score{
			FirstScore: &scoreValue,
		}
		if score.IsDeleted() == true {
			t.Errorf("Expected not deleted score, score: %v ", score)
		}
	})

	t.Run("Seconds_score_present", func(t *testing.T) {
		scoreValue := float32(9)
		score := Score{
			SecondScore: &scoreValue,
		}
		if score.IsDeleted() == true {
			t.Errorf("Expected not deleted score, score: %v ", score)
		}
	})

	t.Run("Absent_score", func(t *testing.T) {
		score := Score{
			IsAbsent: true,
		}
		if score.IsDeleted() == true {
			t.Errorf("Expected not deleted score, score: %v ", score)
		}
	})
}
