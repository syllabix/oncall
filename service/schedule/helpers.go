package schedule

func nextShift(activeShift string, shifts []string) string {
	if len(shifts) < 1 {
		return ""
	}

	if activeShift == "" {
		return shifts[0]
	}

	for i, shift := range shifts {
		if shift == activeShift {
			if len(shifts) == i+1 {
				return shifts[0]
			}
			return shifts[i+1]
		}
	}

	return ""
}
