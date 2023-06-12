package scoreApi

import (
	"testing"
)

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

func TestScore_IsEqual(t *testing.T) {
	t.Run("same_scores", func(t *testing.T) {
		expectedIsEqual := true

		t.Run("with_both_values", func(t *testing.T) {
			leftScore := &Score{
				FirstScore:  floatPointer(1.2),
				SecondScore: floatPointer(9),
			}

			rightScore := &Score{
				FirstScore:  floatPointer(1.2),
				SecondScore: floatPointer(9),
			}

			if leftScore.IsEqual(rightScore) != expectedIsEqual {
				t.Errorf("Expected to be equal, score left: %v == score right %v ", leftScore, rightScore)
			}
		})

		t.Run("with_first_value", func(t *testing.T) {
			leftScore := &Score{
				FirstScore: floatPointer(1.2),
			}

			rightScore := &Score{
				FirstScore: floatPointer(1.2),
			}

			if leftScore.IsEqual(rightScore) != expectedIsEqual {
				t.Errorf("Expected to be equal, score left: %v == score right %v ", leftScore, rightScore)
			}
		})

		t.Run("with_second_values", func(t *testing.T) {
			leftScore := &Score{
				SecondScore: floatPointer(1.2),
			}

			rightScore := &Score{
				SecondScore: floatPointer(1.2),
			}

			if leftScore.IsEqual(rightScore) != expectedIsEqual {
				t.Errorf("Expected to be equal, score left: %v == score right %v ", leftScore, rightScore)
			}
		})

		t.Run("with_absents", func(t *testing.T) {
			leftScore := &Score{
				IsAbsent: true,
			}

			rightScore := &Score{
				IsAbsent: true,
			}

			if leftScore.IsEqual(rightScore) != expectedIsEqual {
				t.Errorf("Expected to be equal, score left: %v == score right %v ", leftScore, rightScore)
			}
		})

		t.Run("deleted", func(t *testing.T) {
			leftScore := &Score{}
			rightScore := &Score{}
			if leftScore.IsEqual(rightScore) != expectedIsEqual {
				t.Errorf("Expected to be equal, score left: %v == score right %v ", leftScore, rightScore)
			}
		})
	})

	t.Run("different", func(t *testing.T) {
		expectedIsEqual := false

		t.Run("with_both_values", func(t *testing.T) {
			leftScore := &Score{
				FirstScore:  floatPointer(5.2),
				SecondScore: floatPointer(9),
			}

			rightScore := &Score{
				FirstScore:  floatPointer(9),
				SecondScore: floatPointer(9),
			}

			if leftScore.IsEqual(rightScore) != expectedIsEqual {
				t.Errorf("Expected to be diff, score left: %v == score right %v ", leftScore, rightScore)
			}
		})

		t.Run("with_both_values", func(t *testing.T) {
			leftScore := &Score{
				FirstScore:  floatPointer(9),
				SecondScore: floatPointer(5.2),
			}

			rightScore := &Score{
				FirstScore:  floatPointer(9),
				SecondScore: floatPointer(9),
			}

			if leftScore.IsEqual(rightScore) != expectedIsEqual {
				t.Errorf("Expected to be diff, score left: %v == score right %v ", leftScore, rightScore)
			}
		})

		t.Run("with_first_value", func(t *testing.T) {
			leftScore := &Score{
				FirstScore: floatPointer(1.2),
			}

			rightScore := &Score{
				FirstScore: floatPointer(4.2),
			}

			if leftScore.IsEqual(rightScore) != expectedIsEqual {
				t.Errorf("Expected to be diff, score left: %v == score right %v ", leftScore, rightScore)
			}
		})

		t.Run("with_second_values", func(t *testing.T) {
			leftScore := &Score{
				SecondScore: floatPointer(4.2),
			}

			rightScore := &Score{
				SecondScore: floatPointer(1.2),
			}

			if leftScore.IsEqual(rightScore) != expectedIsEqual {
				t.Errorf("Expected to be diff, score left: %v == score right %v ", leftScore, rightScore)
			}
		})

		t.Run("with_different_keys", func(t *testing.T) {
			leftScore := &Score{
				FirstScore: floatPointer(1.2),
			}

			rightScore := &Score{
				SecondScore: floatPointer(1.2),
			}

			if leftScore.IsEqual(rightScore) != expectedIsEqual {
				t.Errorf("Expected to be diff, score left: %v == score right %v ", leftScore, rightScore)
			}
		})

		t.Run("with_absents", func(t *testing.T) {
			leftScore := &Score{
				IsAbsent: true,
			}

			rightScore := &Score{
				IsAbsent: false,
			}

			if leftScore.IsEqual(rightScore) != expectedIsEqual {
				t.Errorf("Expected to be diff, score left: %v == score right %v ", leftScore, rightScore)
			}
		})

		t.Run("deleted", func(t *testing.T) {
			leftScore := &Score{}
			rightScore := &Score{
				SecondScore: floatPointer(2),
			}
			if leftScore.IsEqual(rightScore) != expectedIsEqual {
				t.Errorf("Expected to be diff, score left: %v == score right %v ", leftScore, rightScore)
			}
		})
	})

}

func floatPointer(value float32) *float32 {
	return &value
}
