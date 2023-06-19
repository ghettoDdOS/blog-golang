package utils

func RemoveDuplicates(slice []string) []string {
	uniqueMap := make(map[string]bool)

	for _, element := range slice {
		if !uniqueMap[element] {
			uniqueMap[element] = true
		}
	}

	uniqueSlice := make([]string, 0, len(uniqueMap))

	for element := range uniqueMap {
		uniqueSlice = append(uniqueSlice, element)
	}

	return uniqueSlice
}
