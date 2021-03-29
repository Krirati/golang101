package grade

// Calculate grade
func Calculate(score float64) string {
	if score >= 80 {
		return "A"
	} else if score == 79.0 {
		return "B+"
	} else if score == 70.0 {
		return "B"
	} else if score >= 65.0 {
		return "C+"
	} else if score == 60.0 {
		return "C"
	} else if score == 59.0 {
		return "D+"
	} else {
		return "D"
	}
}
