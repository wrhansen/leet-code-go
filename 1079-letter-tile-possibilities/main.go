package main

func generateCombinations(runes []rune, used []bool, current string, possibilities map[string]bool) {
	if len(current) > 0 {
		possibilities[current] = true
	}

	for i := 0; i < len(runes); i++ {
		if used[i] {
			continue
		}
		used[i] = true
		generateCombinations(runes, used, current+string(runes[i]), possibilities)
		used[i] = false
	}
}

func numTilePossibilities(tiles string) int {
	possibilities := make(map[string]bool)
	runeArray := []rune(tiles)
	used := make([]bool, len(runeArray))
	generateCombinations(runeArray, used, "", possibilities)
	return len(possibilities)
}
