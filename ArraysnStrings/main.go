package arraysnstrings

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// # 1.1
//
// Is Unique: Implement an algorithm to determine if a string has all unique characters. What if you
// cannot use additional data structures?
//
// IsUnique returns true if a string has all unique characters, and versa
func IsUnique(s string) bool {
	m := make(map[byte]struct{})

	for i := 0; i < len(s); i++ {
		c := s[i]
		if _, ok := m[c]; ok {
			return false
		}
		m[c] = struct{}{}
	}
	return true
}

// We can use a bit vector to store the appearance of characters, we assume that this string contains
// only lowercase alphabet characters.
func IsUniqueV2(s string) bool {
	if len(s) <= 1 {
		return true
	}
	if len(s) > 26 {
		return false
	}

	vector := 0

	for i := 0; i < len(s); i++ {
		c := s[i] - 'a'
		if vector&(2<<c) > 0 {
			return false
		}
		vector |= 2 << c
	}
	return true
}

type SortString []byte

func (s *SortString) Len() int {
	return len(*s)
}

func (s *SortString) Less(i, j int) bool {
	return (*s)[i] < (*s)[j]
}

func (s *SortString) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

// # 1.2
//
// Check Permutation: Given two strings, write a method to decide if one is a permutation of the
// other
//
// CheckPermutation return true if one is a permutation of the other, and versa
func CheckPermutation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	v1 := SortString(s1)
	v2 := SortString(s2)
	sort.Sort(&v1)
	sort.Sort(&v2)
	for i, v := range v1 {
		if v != v2[i] {
			return false
		}
	}
	return true
}

func CheckPermutationV2(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	l := len(s1)
	arr := [256]int{}

	for i := 0; i < l; i++ {
		arr[s1[i]]++
	}

	for i := 0; i < l; i++ {
		c := s2[i]
		arr[c]--
		if arr[c] < 0 {
			return false
		}
	}

	return true
}

// # 1.3
//
// URLify: Write a method to replace all spaces in a string with '%20'. You may assume that the string
// has sufficient space at the end to hold the additional characters, and that you are given the "true"
// length of the string. (Note: If implementing in Java, please use a character array so that you can
// perform this operation in place.)
//
// # EXAMPLE
//
// Input: "Mr John Smith     ", 13
//
// Output: "Mr%20John%20Smith"
func URLify(s string, l int) string {
	result := strings.Builder{}
	for i, c := range s {
		if i == l {
			break
		}
		if c != ' ' {
			_, err := result.WriteRune(c)
			if err != nil {
				return ""
			}
		} else if c == ' ' {
			_, err := result.WriteString("%20")
			if err != nil {
				return ""
			}
		}
	}
	return result.String()
}

// # 1.4
//
// Palindrome Permutation: Given a string, write a function to check if it is a permutation of a palindrome.
// A palindrome is a word or phrase that is the same forwards and backwards. A permutation
// is a rearrangement of letters. The palindrome does not need to be limited to just dictionary words.
//
// # EXAMPLE
//
// Input: Tact Coa
//
// Output: True (permutations: "taco cat", "atco eta", etc.)
func PalindromePermutation(s string) bool {
	m := make(map[rune]int)
	for _, c := range s {
		if _, ok := m[c]; ok {
			m[c]++
		} else {
			m[c] = 1
		}
	}

	sign := false
	for _, v := range m {
		if v%2 != 0 {
			if sign == true {
				return false
			}
			sign = true
		}
	}
	return true
}

// # 1.5
//
// One Away: There are three types of edits that can be performed on strings: insert a character,
// remove a character, or replace a character. Given two strings, write a function to check if they are
// one edit (or zero edits) away.
//
// # EXAMPLE
//
// pale, ple -> true
//
// pales, pale -> true
//
// pale, bale -> true
//
// pale, bake -> false
func OneAway(s1, s2 string) bool {
	l1, l2 := len(s1), len(s2)
	if math.Abs(float64(l1-l2)) > 1 {
		return false
	}

	check := func(short, long string) bool {
		l1, l2 := len(short), len(long)
		diff := 0
		equal := l1 == l2
		for i := 0; i < l1 && i+diff < l2; i++ {
			if equal {
				if short[i] != long[i] {
					diff++
				}
			} else if short[i] != long[i+diff] {
				diff++
			}
			if diff > 1 {
				return false
			}
		}
		return true
	}

	if l1 == l2 {
		return check(s1, s2)
	} else if l1 < l2 {
		return check(s1, s2)
	} else {
		return check(s2, s1)
	}
}

// # 1.6
//
// String Compression: Implement a method to perform basic string compression using the counts
// of repeated characters. For example, the string aabcccccaaa would become a2b1c5a3. If the
// "compressed" string would not become smaller than the original string, your method should return
// the original string. You can assume the string has only uppercase and lowercase letters (a - z).
func StringCompression(s string) string {
	result := strings.Builder{}
	if len(s) == 0 {
		return ""
	}

	current := rune(s[0])
	count := 1
	err := result.WriteByte(s[0])
	if err != nil {
		return result.String()
	}
	for _, c := range s[1:] {
		if c != current {
			_, err := result.WriteString(fmt.Sprint(count))
			if err != nil {
				return result.String()
			}
			_, err = result.WriteRune(c)
			if err != nil {
				return result.String()
			}
			current = c
			count = 1
		} else {
			count++
		}
	}
	_, err = result.WriteString(fmt.Sprint(count))
	if err != nil {
		return result.String()
	}

	if result.Len() >= len(s) {
		return s
	}

	return result.String()
}

// # 1.7
//
// Rotate Matrix: Given an image represented by an NxN matrix, where each pixel in the image is 4
// bytes, write a method to rotate the image by 90 degrees. Can you do this in place?
func RotateMatrix(matrix [][]int, n int, direction int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}

	maxIndex := n - 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			var iNew, jNew int
			if direction == 1 {
				iNew = j
				jNew = maxIndex - i
			} else if direction == -1 {
				iNew = maxIndex - j
				jNew = i
			}
			result[iNew][jNew] = matrix[i][j]
		}
	}
	return result
}

// # 1.8
//
// Zero Matrix: Write an algorithm such that if an element in an MxN matrix is 0, its entire row and
// column are set to 0.
func ZeroMatrix(matrix [][]int, m, n int) [][]int {
	is := make(map[int]struct{})
	js := make(map[int]struct{})
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				if _, exist := is[i]; !exist {
					is[i] = struct{}{}
				}
				if _, exist := js[j]; !exist {
					js[j] = struct{}{}
				}
			}
		}
	}

	for i := range is {
		for j := 0; j < n; j++ {
			matrix[i][j] = 0
		}
	}

	for j := range js {
		for i := 0; i < m; i++ {
			matrix[i][j] = 0
		}
	}
	return matrix
}

// # 1.9
//
// String Rotation: Assume you have a method isSubstring which checks if one word is a substring
// of another. Given two strings, s1 and s2, write code to check if s2 is a rotation of s1 using only one
// call to isSubstring (e.g., "waterbottle" is a rotation of "erbottlewat").
func StringRotation(s1, s2 string) bool {
	if len(s1) != len(s2) || len(s1) < 0 {
		return false
	}
	if len(s1) == 0 {
		return true
	}
	dupS1 := s1 + s1
	return strings.Contains(dupS1, s2)
}
