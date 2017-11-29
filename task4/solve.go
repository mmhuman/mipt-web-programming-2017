package main

import "unicode"

func RemoveEven(a []int) []int {
	var res []int
	for _, val := range(a) {
		if val % 2 == 1 {
			res = append(res, val)
		}
	}
	return res
}

func PowerGenerator(x int) (func() int) {
    res := 1
    return func() int {
        res *= x
        return res
    }
}

func DifferentWordsCount(text string) int {
	m := make(map[string]int, 0)
	val := ""
	for _, c := range(text) {
		if unicode.IsLetter(c) {
			c = unicode.ToLower(c)
			val += string(c)
		} else {
			m[val] = 1
			val = ""
		}
	}
	m[val] = 1
	res := len(m)
	_, ok := m[""]
	if ok {
		res--
	}
	return res
}