package ascii_art_web

// validBannerName checks if the provided banner name is valid.
func validBannerName(str string) bool {
	banners := []string{"standard", "shadow", "thinkertoy", "graffiti"}
	for _, banner := range banners {
		if banner == str {
			return true
		}
	}
	return false
}

// filterInput - a function that filters the given string
// from characters that not in the ascii table.
func filterInput(input string) string {
	new_input := ""
	for _, c := range input {
		if (c >= ' ' && c <= '~') || c == '\r' || c == '\n' {
			new_input += string(c)
		}
	}
	return new_input
}
