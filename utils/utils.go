package utils

import "strings"
import "github.com/nu7hatch/gouuid"


// method to generate UUID

func GenerateUniqueUUID () string {
	deck_uuid, _ := uuid.NewV4()
	var deck_id string = deck_uuid.String()

	return deck_id
}


// method to check if a `target` /item exists in an array

func StringInSlice(array []string, target string) bool {
	// search for `target` in `array`

	for _, sample := range array {
		sample = strings.TrimSpace(sample) // trim sample
		target = strings.TrimSpace(target) // trim target

		if (sample == target) {
			// `target` found -> `target` is in `array`
			return true
		}
	}

	// `target` not found -> `targte is not in array`
	return false
}
