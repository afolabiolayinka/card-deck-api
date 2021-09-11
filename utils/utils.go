package utils

import "strings"
import "github.com/nu7hatch/gouuid"


func GenerateUniqueUUID () string {
	deck_uuid, _ := uuid.NewV4()
	var deck_id string = deck_uuid.String()

	return deck_id
}


func StringInSlice(array []string, target string) bool {
	for _, sample := range array {
		sample = strings.TrimSpace(sample)
		target = strings.TrimSpace(target)

		if (sample == target) {
			return true
		}
	}

	return false
}


func ArrayStrComp(array []string, target string) bool {
	for i := 0; i < len(array); i++ {
		if (strings.Compare(array[i], target) == 0) {
			return true
		}
	}

	return false
}
