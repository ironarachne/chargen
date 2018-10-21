package chargen

var (
	ageCategories = map[string]int{
		"adult":       12,
		"elderly":     1,
		"young adult": 2,
		"child":       1,
	}

	orientations = map[string]int{
		"straight": 10,
		"gay":      1,
		"bi":       1,
	}

	genders = []string{
		"female",
		"male",
	}

	professions = []string{
		"alchemist",
		"baker",
		"bartender",
		"blacksmith",
		"bowyer",
		"carpenter",
		"cook",
		"farmer",
		"ferrier",
		"fletcher",
		"guard",
		"guildmaster",
		"healer",
		"mage",
		"merchant",
	}

	hobbies = []string{
		"whittling",
		"carving",
		"darts",
		"tournaments",
		"fishing",
		"hunting",
		"painting",
		"drinking",
	}

	traits = []string{
		"addict",
		"anxious",
		"arrogant",
		"brave",
		"cautious",
		"confident",
		"contemptuous",
		"devious",
		"drunkard",
		"encouraging",
		"envious",
		"fair",
		"fierce",
		"just",
		"lusty",
		"meticulous",
		"optimistic",
		"persistent",
		"pessimistic",
		"proud",
		"reckless",
		"religious",
		"resilient",
		"sickly",
		"stable",
		"strong",
		"stubborn",
		"trusting",
		"wise",
	}

	motivations = []string{
		"debauchery",
		"fame",
		"family",
		"friends",
		"justice",
		"knowledge",
		"money",
		"power",
		"revenge",
	}
)
