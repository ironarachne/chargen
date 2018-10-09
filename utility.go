package chargen

import "math/rand"

func randomItem(items []string) string {
	return items[rand.Intn(len(items))]
}
