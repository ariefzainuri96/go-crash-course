package l10nested

/*
example data [billy, billy, bob, joe]

The data from nested map likely look like this

b: {
    billy: 2,
    bob: 1
},
j: {
    joe: 1
}
*/

func getNameCounts(names []string) map[rune]map[string]int {
	testMap := make(map[rune]map[string]int)

	for _, value := range names {

		// create slice with type data rune from value
		// if the value is "Arief"
		// then the runeArray is [A r i e f]
		runeArray := []rune(value)

		if _, isExist := testMap[runeArray[0]]; isExist {
			testMap[runeArray[0]][value] += 1
		} else {
			testMap[runeArray[0]] = map[string]int{
				value: 1,
			}
		}
	}

	return testMap
}
