package pkg

func Contains(word string, list []string) bool {
	match := false
	for _, w := range list {
		if word == w {
			match = true
			break
		}
	}
	return match
}
