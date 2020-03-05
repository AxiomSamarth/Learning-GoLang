package stringutils

// FullName returns the full name by concatenating the first and last names with a space
func FullName(firstName, lastName string) (string, int) {
	full := firstName + " " + lastName
	return full, len(full)
}
