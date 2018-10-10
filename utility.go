package chargen

import "math/rand"

func randomItem(items []string) string {
	return items[rand.Intn(len(items))]
}

func itemInCollection(item string, collection []string) bool {
	for _, element := range collection {
		if item == element {
			return true
		}
	}
	return false
}

func randomItemFromThresholdMap(items map[string]int) string {
	result := ""
	ceiling := 0
	start := 0
	var thresholds = make(map[string]int)

	for item, weight := range items {
		ceiling += weight
		thresholds[item] = start
		start += weight
	}

	randomValue := rand.Intn(ceiling)

	for item, threshold := range thresholds {
		if threshold <= randomValue {
			result = item
		}
	}

	return result
}
