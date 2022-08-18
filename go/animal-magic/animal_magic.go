package chance

import (
	"math/rand"
	"time"
)

// SeedWithTime seeds math/rand with the current computer time.
func SeedWithTime() {
	rand.Seed(time.Now().UnixNano())
}

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	return rand.Intn(20) + 1
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	rand_float := rand.Float64() + float64(rand.Intn(12))
	if rand_float >= 12.0 {
		return rand_float - 1.0
	}
	return rand_float
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	animals := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}

	sequences := rand.Perm(len(animals))
	shuffled := make([]string, len(animals))

	for old_index, new_index := range sequences {
		shuffled[new_index] = animals[old_index]
	}

	return shuffled
}
