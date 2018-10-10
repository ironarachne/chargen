package chargen

var (
	orientations = map[string]int{
		"straight": 10,
		"gay":      1,
		"bi":       1,
	}

	genders = []string{
		"male",
		"female",
	}

	attitudes = []string{
		"confident",
		"contemptuous",
		"devious",
		"proud",
		"optimistic",
		"pessimistic",
		"encouraging",
		"fair",
		"just",
		"religious",
	}

	hobbies = []string{
		"whittling",
		"carving",
		"darts",
		"tournaments",
		"fishing",
		"hunting",
		"painting",
	}

	traits = []string{
		"arrogant",
		"anxious",
		"envious",
		"stable",
		"brave",
		"trusting",
		"meticulous",
		"fierce",
	}

	motivations = []string{
		"money",
		"power",
		"fame",
		"family",
		"friends",
		"revenge",
	}
)
