package part_one

import (
	"bytes"
	"unicode"
)

func readInput(input []byte) int {
	lines := bytes.Split(input, []byte("\n"))
	res := 0
	for _, line := range lines {
		var record []byte
		var contigous []int
		tmp := 0
		for _, char := range line {
			if char == '.' || char == '#' || char == '?' {
				record = append(record, char)
			}
			if unicode.IsDigit(rune(char)) {
				tmp = tmp*10 + int(rune(char)-'0')
			}
			if char == ',' {
				contigous = append(contigous, tmp)
				tmp = 0
			}
		}
		if tmp != 0 {
			contigous = append(contigous, tmp)
		}
		res += isContigous(record, contigous, 0)
	}
	return res
}

func isContigous(record []byte, contigous []int, lecture int) int {
	copy_of_record := make([]byte, len(record))
	copy(copy_of_record, record)
	// fmt.Println(copy_of_record)
	if lecture == len(record) {
		if possiblePattern(record, contigous) {
			return 1
		} else {
			return 0
		}
	} else {
		if copy_of_record[lecture] == '?' {
			copy_of_record[lecture] = '.'
			lecture++
			a := isContigous(copy_of_record, contigous, lecture)
			lecture--
			copy_of_record[lecture] = '#'
			lecture++
			b := isContigous(copy_of_record, contigous, lecture)
			return a + b
		} else {
			lecture++
			return isContigous(copy_of_record, contigous, lecture)
		}
	}
}

func possiblePattern(record []byte, contigous []int) bool {
	count := 0
	var patternOfRecord []int
	for i := 0; i < len(record); i++ {
		if record[i] == '#' {
			count++
		} else if record[i] != '#' && count != 0 {
			patternOfRecord = append(patternOfRecord, count)
			count = 0
		}
	}
	if count != 0 {
		patternOfRecord = append(patternOfRecord, count)
	}

	// fmt.Println(patternOfRecord)

	if len(patternOfRecord) != len(contigous) {
		return false
	} else {
		for i := 0; i < len(patternOfRecord); i++ {
			if patternOfRecord[i] != contigous[i] {
				return false
			}
		}
		return true
	}
}

func Part_One(input []byte) int {

	return readInput(input)
}
