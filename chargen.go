package chargen

import "math/rand"

// Character is a character
type Character struct {
	HairColor           string
	HairStyle           string
	EyeColor            string
	FaceShape           string
	Height              int
	Weight              int
	Race                string
	Gender              string
	Orientation         string
	Attitude            string
	Hobby               string
	PsychologicalTraits []string
	Motivation          string
}

func randomHeight(race string, gender string) int {
	minHeight := 0
	maxHeight := 1
  genderRatio := 1.0
  if gender == "male" {
    genderRatio = 1.08
  }

	if race == "human" {
		minHeight = 65
		maxHeight = 78
	} else if race == "elf" {
		minHeight = 60
		maxHeight = 70
	} else if race == "dwarf" {
		minHeight = 40
		maxHeight = 55
	}

  height := rand.Intn(maxHeight-minHeight) + minHeight

  return int(float64(height) * genderRatio)
}

func randomWeight(race string, gender string) int {
	minWeight := 0
	maxWeight := 1
  genderRatio := 1.0
  if gender == "male" {
    genderRatio = 1.16
  }

	if race == "human" {
		minWeight = 140
		maxWeight = 230
	} else if race == "elf" {
		minWeight = 70
		maxWeight = 160
	} else if race == "dwarf" {
		minWeight = 180
		maxWeight = 300
	}

  weight := rand.Intn(maxWeight-minWeight) + minWeight

  return int(float64(weight) * genderRatio)
}

// Generate generates a random character
func Generate() Character {
	char := Character{}

	char.HairColor = randomItem(hairColors)
	char.HairStyle = randomItem(hairStyles)
	char.EyeColor = randomItem(eyeColors)
	char.FaceShape = randomItem(faceShapes)
	char.Race = randomItem(races)
	char.Gender = randomItem(genders)
	char.Orientation = randomItem(orientations)
	char.Attitude = randomItem(attitudes)
	char.Hobby = randomItem(hobbies)
	char.Motivation = randomItem(motivations)

	for i := 0; i < 2; i++ {
		trait := randomItem(traits)
		for itemInCollection(trait, char.PsychologicalTraits) {
			trait = randomItem(traits)
		}
		char.PsychologicalTraits = append(char.PsychologicalTraits, trait)
	}

	char.Height = randomHeight(char.Race, char.Gender)
	char.Weight = randomWeight(char.Race, char.Gender)

	return char
}
