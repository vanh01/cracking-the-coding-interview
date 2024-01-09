package arraysnstrings

import (
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
func OneAway() {}

// # 1.6
//
// String Compression: Implement a method to perform basic string compression using the counts
// of repeated characters. For example, the string aabcccccaaa would become a2blc5a3. If the
// "compressed" string would not become smaller than the original string, your method should return
// the original string. You can assume the string has only uppercase and lowercase letters (a - z).
func StringCompression() {}

// # 1.7
//
// Rotate Matrix: Given an image represented by an NxN matrix, where each pixel in the image is 4
// bytes, write a method to rotate the image by 90 degrees. Can you do this in place?
func RotateMatrix() {}

// # 1.8
//
// Zero Matrix: Write an algorithm such that if an element in an MxN matrix is 0, its entire row and
// column are set to 0.
func ZeroMatrix() {}

// # 1.9
//
// String Rotation:Assumeyou have a method isSubstringwhich checks if oneword is a substring
// of another. Given two strings, sl and s2, write code to check if s2 is a rotation of sl using only one
// call to isSubstring (e.g., "waterbottle" is a rotation of"erbottlewat").
func StringRotation() {}
