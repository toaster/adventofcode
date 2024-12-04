package wordsearch

// WordSearch represents a word search puzzle.
type WordSearch struct {
	height  int
	letters [][]rune
	width   int
}

// ParseWordSearch parses a new WordSearch from a couple of lines of equal length.
func ParseWordSearch(lines []string) WordSearch {
	height := len(lines)
	width := len(lines[0])
	letters := make([][]rune, height)
	for i, line := range lines {
		letters[i] = []rune(line)
	}
	return WordSearch{
		height:  height,
		letters: letters,
		width:   width,
	}
}

// Count counts the occurrences of the given word.
func (s WordSearch) Count(word string) int {
	count := 0
	wordLetters := []rune(word)
	wordLen := len(word)
	for x := 0; x < s.width; x++ {
		for y := 0; y < s.height; y++ {
			if s.letters[y][x] == wordLetters[0] {
				if x <= s.width-wordLen {
					if s.matches(wordLetters, x, y, 1, 0) {
						count++
					}
					if y <= s.height-wordLen {
						if s.matches(wordLetters, x, y, 1, 1) {
							count++
						}
					}
					if y >= wordLen-1 {
						if s.matches(wordLetters, x, y, 1, -1) {
							count++
						}
					}
				}
				if y <= s.height-wordLen {
					if s.matches(wordLetters, x, y, 0, 1) {
						count++
					}
				}
				if y >= wordLen-1 {
					if s.matches(wordLetters, x, y, 0, -1) {
						count++
					}
				}
				if x >= wordLen-1 {
					if s.matches(wordLetters, x, y, -1, 0) {
						count++
					}
					if y <= s.height-wordLen {
						if s.matches(wordLetters, x, y, -1, 1) {
							count++
						}
					}
					if y >= wordLen-1 {
						if s.matches(wordLetters, x, y, -1, -1) {
							count++
						}
					}
				}
			}
		}
	}
	return count
}

// CountX counts the x-wise occurrences of the given word.
func (s WordSearch) CountX(word string) int {
	count := 0
	wordLetters := []rune(word)
	wordLen := len(word)
	legLength := (wordLen - 1) / 2
	for x := legLength; x < s.width-legLength; x++ {
		for y := legLength; y < s.height-legLength; y++ {
			if (s.matches(wordLetters, x-legLength, y-legLength, 1, 1) || s.matches(wordLetters, x+legLength, y+legLength, -1, -1)) &&
				(s.matches(wordLetters, x-legLength, y+legLength, 1, -1) || s.matches(wordLetters, x+legLength, y-legLength, -1, 1)) {
				count++
			}
		}
	}
	return count
}

func (s WordSearch) matches(letters []rune, x, y, stepX, stepY int) bool {
	for i := 0; i < len(letters); i++ {
		if s.letters[y+stepY*i][x+stepX*i] != letters[i] {
			return false
		}
	}
	return true
}
