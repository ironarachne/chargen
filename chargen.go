package chargen

import (
	"math"
	"math/rand"

	"github.com/ironarachne/random"
	"github.com/ironarachne/utility"

	"github.com/ironarachne/culturegen"
)

// Character is a character
type Character struct {
	FirstName           string
	LastName            string
	HairColor           string
	HairStyle           string
	EyeColor            string
	FaceShape           string
	Culture             culturegen.Culture
	Height              int
	Weight              int
	Race                Race
	Age                 int
	AgeCategory         string
	Gender              string
	Orientation         string
	Profession          string
	Hobby               string
	PsychologicalTraits []string
	Motivation          string
}

// Couple is a pair of partners
type Couple struct {
	Partner1        Character
	Partner2        Character
	CanHaveChildren bool
}

// Family is a family of characters
type Family struct {
	FamilyName string
	Parents    Couple
	Children   []Character
}

func getOppositeGender(gender string) string {
	if gender == "male" {
		return "female"
	}

	return "male"
}

func getAppropriateName(gender string, culture culturegen.Culture) (string, string) {
	firstName := culture.Language.RandomGenderedName(gender)
	lastName := culture.Language.RandomName()

	return firstName, lastName
}

func getAgeCategoryRange(race Race, ageCategory string) (int, int) {
	minAge := 1
	maxAge := 1

	if ageCategory == "child" {
		minAge = 1
		maxAge = race.FertilityAge - 1
	} else if ageCategory == "young adult" {
		minAge = race.FertilityAge
		maxAge = race.AdultAge - 1
	} else if ageCategory == "adult" {
		minAge = race.AdultAge
		maxAge = int(math.Floor(float64(race.LifeSpan)*0.85)) - 1
	} else if ageCategory == "elderly" {
		minAge = int(math.Floor(float64(race.LifeSpan) * 0.85))
		maxAge = race.LifeSpan
	}

	return minAge, maxAge
}

func getRandomAge(race Race, ageCategory string) int {
	minAge, maxAge := getAgeCategoryRange(race, ageCategory)

	return rand.Intn(maxAge-minAge) + minAge
}

func getAgeCategoryFromAge(age int, race Race) string {
	ageCategory := ""
	minAge := 0
	maxAge := 0

	for category := range ageCategories {
		minAge, maxAge = getAgeCategoryRange(race, category)
		if age >= minAge && age <= maxAge {
			ageCategory = category
		}
	}

	return ageCategory
}

func getAgeFromParents(parents Couple, race Race) (int, string) {
	lowestAge := utility.Min(parents.Partner1.Age, parents.Partner2.Age)

	childAge := rand.Intn(lowestAge-race.FertilityAge) + 1
	childAgeCategory := getAgeCategoryFromAge(childAge, race)

	return childAge, childAgeCategory
}

func getRaceFromParents(parents Couple) Race {
	if parents.Partner1.Race.Name == parents.Partner2.Race.Name {
		return parents.Partner1.Race
	}

	possibleRaces := []string{parents.Partner1.Race.Name, parents.Partner2.Race.Name}

	if utility.ItemInCollection("elf", possibleRaces) && utility.ItemInCollection("human", possibleRaces) {
		return races["half-elf"]
	}

	if utility.ItemInCollection("dwarf", possibleRaces) && utility.ItemInCollection("human", possibleRaces) {
		return races["half-dwarf"]
	}

	if utility.ItemInCollection("half-elf", possibleRaces) {
		return races["half-elf"]
	}

	if utility.ItemInCollection("half-dwarf", possibleRaces) {
		return races["half-dwarf"]
	}

	return races["human"]
}

func randomOrientation() string {
	return random.ItemFromThresholdMap(orientations)
}

func randomHeight(race Race, gender string) int {
	minHeight := race.MinHeight
	maxHeight := race.MaxHeight
	genderRatio := 1.0
	if gender == "male" {
		genderRatio = 1.08
	}

	height := rand.Intn(maxHeight-minHeight) + minHeight

	return int(float64(height) * genderRatio)
}

func randomRace() Race {
	raceName := random.ItemFromThresholdMap(raceData)

	race := races[raceName]

	return race
}

func randomWeight(race Race, gender string) int {
	minWeight := race.MinWeight
	maxWeight := race.MaxWeight
	genderRatio := 1.0
	if gender == "male" {
		genderRatio = 1.16
	}

	weight := rand.Intn(maxWeight-minWeight) + minWeight

	return int(float64(weight) * genderRatio)
}

// GenerateCharacter generates a random character
func GenerateCharacter() Character {
	char := Character{}

	char.Race = randomRace()
	char.Gender = random.Item(genders)
	char.Culture = culturegen.GenerateCulture()

	char.FirstName, char.LastName = getAppropriateName(char.Gender, char.Culture)

	char.AgeCategory = random.ItemFromThresholdMap(ageCategories)
	char.Age = getRandomAge(char.Race, char.AgeCategory)

	char.HairColor = random.ItemFromThresholdMap(char.Race.HairColors)
	char.HairStyle = random.ItemFromThresholdMap(char.Race.HairStyles)
	char.EyeColor = random.ItemFromThresholdMap(char.Race.EyeColors)
	char.FaceShape = random.ItemFromThresholdMap(char.Race.FaceShapes)

	char.Orientation = randomOrientation()
	char.Profession = random.Item(professions)
	char.Hobby = random.Item(hobbies)
	char.Motivation = random.Item(motivations)

	for i := 0; i < 2; i++ {
		trait := random.Item(traits)
		for utility.ItemInCollection(trait, char.PsychologicalTraits) {
			trait = random.Item(traits)
		}
		char.PsychologicalTraits = append(char.PsychologicalTraits, trait)
	}

	char.Height = randomHeight(char.Race, char.Gender)
	char.Weight = randomWeight(char.Race, char.Gender)

	return char
}

// GenerateCharacterOfCulture generates a random character with a given culture
func GenerateCharacterOfCulture(culture culturegen.Culture) Character {
	character := GenerateCharacter()
	character = character.SetCulture(culture)

	return character
}

// GenerateCouple generates a couple
func GenerateCouple() Couple {
	char1 := GenerateCharacter()
	char2 := GenerateCharacter()
	canHaveChildren := false

	if char1.AgeCategory == "child" {
		char1.AgeCategory = "adult"
		char1.Age = getRandomAge(char1.Race, char1.AgeCategory)
	}

	if char2.AgeCategory == "child" {
		char2.AgeCategory = "adult"
		char2.Age = getRandomAge(char2.Race, char2.AgeCategory)
	}

	raceNames := []string{char1.Race.Name, char2.Race.Name}
	orientations := []string{char1.Orientation, char2.Orientation}

	if (char1.Orientation == "gay" && char2.Orientation == "straight") || (char1.Orientation == "straight" && char2.Orientation == "gay") {
		char2.Orientation = char1.Orientation
		orientations = []string{char1.Orientation}
	}

	if char1.Gender == char2.Gender && utility.ItemInCollection("straight", orientations) {
		char2.Gender = getOppositeGender(char1.Gender)
		char2.FirstName, _ = getAppropriateName(char2.Gender, char2.Culture)
	} else if char1.Gender != char2.Gender && utility.ItemInCollection("gay", orientations) {
		char2.Gender = char1.Gender
		char2.FirstName, _ = getAppropriateName(char2.Gender, char2.Culture)
	}

	if utility.ItemInCollection("dwarf", raceNames) && utility.ItemInCollection("elf", raceNames) {
		char1.Race = char2.Race
		char1.FirstName, _ = getAppropriateName(char1.Gender, char1.Culture)
	}

	if utility.ItemInCollection("halfling", raceNames) {
		char1.Race = races["halfling"]
		char2.Race = races["halfling"]
	}

	if char1.Gender != char2.Gender {
		canHaveChildren = true
	}

	couple := Couple{char1, char2, canHaveChildren}

	return couple
}

// GenerateAdultDescendent generates an adult character based on a couple
func GenerateAdultDescendent(couple Couple) Character {
	descendent := GenerateCharacter()

	descendent.LastName = couple.Partner1.LastName
	descendent.Race = getRaceFromParents(couple)

	descendent.Age = getRandomAge(descendent.Race, "adult")
	descendent.AgeCategory = getAgeCategoryFromAge(descendent.Age, descendent.Race)

	return descendent
}

// GenerateChild generates a child character for a couple
func GenerateChild(couple Couple) Character {
	child := GenerateCharacter()

	child.LastName = couple.Partner1.LastName
	child.Race = getRaceFromParents(couple)
	child.Age, child.AgeCategory = getAgeFromParents(couple, child.Race)

	if child.AgeCategory == "child" {
		child.Profession = "none"
	}

	return child
}

// GenerateCompatibleMate generates a character appropriate as a mate for another
func GenerateCompatibleMate(char Character) Character {
	mate := GenerateCharacter()

	mate.Age = getRandomAge(mate.Race, char.AgeCategory)
	mate.AgeCategory = getAgeCategoryFromAge(mate.Age, mate.Race)

	if char.Orientation == "straight" {
		mate.Gender = getOppositeGender(char.Gender)
		mate.Orientation = "straight"
	} else {
		mate.Gender = char.Gender
		mate.Orientation = "gay"
	}

	return mate
}

// GenerateFamily generates a random family
func GenerateFamily() Family {
	child := Character{}
	children := []Character{}

	parents := GenerateCouple()

	familyName := parents.Partner1.LastName

	parents.Partner2.LastName = familyName

	if parents.CanHaveChildren {
		for i := 0; i < rand.Intn(6); i++ {
			child = GenerateChild(parents)
			child.LastName = familyName
			children = append(children, child)
		}
	}

	return Family{familyName, parents, children}
}

// MarryCouple returns a couple from two characters
func MarryCouple(partner1 Character, partner2 Character) Couple {
	canHaveChildren := false

	if partner1.Gender != partner2.Gender {
		canHaveChildren = true
	}

	return Couple{partner1, partner2, canHaveChildren}
}

// SetCulture sets the culture of the character
func (character Character) SetCulture(culture culturegen.Culture) Character {
	newCharacter := character

	newCharacter.Culture = culture
	newCharacter.FirstName, newCharacter.LastName = getAppropriateName(newCharacter.Gender, newCharacter.Culture)

	return newCharacter
}
