package utils

// Function to check if string slice contains an element
func Contains(slice []string, element string) bool {
	for _, s := range slice {
		if s == element {
			return true
		}
	}
	return false
}
