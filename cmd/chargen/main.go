package main

import (
	"fmt"

	"github.com/ironarachne/chargen"
)

func displayCharacter(char chargen.Character) {
	fmt.Println("Hair: " + char.HairColor + ", " + char.HairStyle)
	fmt.Println("Eyes: " + char.EyeColor)
	fmt.Println("Face Shape: " + char.FaceShape)
	fmt.Println("Race: " + char.Race)
	fmt.Println("Gender: " + char.Gender)
	fmt.Println("Orientation: " + char.Orientation)
	fmt.Println("Attitude: " + char.Attitude)
	fmt.Println("Traits: ")
	for _, trait := range char.PsychologicalTraits {
		fmt.Println("- " + trait)
	}
	fmt.Println("Hobby: " + char.Hobby)
	fmt.Println("Motivation: " + char.Motivation)
}

func main() {
	character := chargen.Generate()

	displayCharacter(character)
}
