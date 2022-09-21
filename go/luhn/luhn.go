package luhn

import (
	"regexp"
	"strconv"
	"strings"
)

func string_to_int_slice(s string) []int {
	id_int := []int{}
	for _, char := range s {

		char_s := string(char)
		num, _ := strconv.Atoi(char_s)
		id_int = append(id_int, num)

	}
	return id_int
}

func reverse_slice(slice []int) []int {
	slice_len := len(slice)
	reversed_slice := []int{}

	for i := slice_len - 1; i >= 0; i-- {
		reversed_slice = append(reversed_slice, slice[i])
	}
	return reversed_slice
}

func is_a_valid_id(id string) bool {
	re, _ := regexp.Compile(`^\d+$`)
	match := re.MatchString(id)
	return match
}

func calc_luhn(slice []int) int {
	total := 0
	for idx, num := range slice {
		if idx%2 != 1 {
			total += num
			continue
		}
		num_x2 := num * 2
		if num_x2 > 9 {
			total += num_x2 - 9
		} else {
			total += num_x2
		}
	}

	return total
}

func Valid(id string) bool {
	id = strings.Replace(id, " ", "", -1)
	if len(id) <= 1 {
		return false
	}

	if !is_a_valid_id(id) {
		return false
	}

	id_int := reverse_slice(string_to_int_slice(id))
	total := calc_luhn(id_int)

	return total%10 == 0
}
