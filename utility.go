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
