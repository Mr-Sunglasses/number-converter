package controllers

import (
	"strings"
	"testing"
)

// AssertError checks if the error message in out contains the text in want
func AssertError(out error, want string) bool {
	if out == nil {
		return want == ""
	}
	if want == "" {
		return false
	}
	return strings.Contains(out.Error(), want)
}

func TestRomanToInteger(t *testing.T) {
	t.Run("Check Boundary Values (Basic Symbols)", func(t *testing.T) {
		validRomans := map[string]int{
			"I": 1,
			"V": 5,
			"X": 10,
			"L": 50,
			"C": 100,
			"D": 500,
			"M": 1000,
		}
		for roman := range validRomans {
			ans, _ := RomanToInt(roman)
			if ans != validRomans[roman] {
				t.Errorf("RomanToInt: want %d got %d", validRomans[roman], ans)
			}
		}
	})
	t.Run("Check Random Roman Values", func(t *testing.T) {
		validRomans := map[string]int{
			"LXIX":    69,
			"CDXX":    420,
			"MCMXCIV": 1994,
		}
		for roman := range validRomans {
			ans, _ := RomanToInt(roman)
			if ans != validRomans[roman] {
				t.Errorf("RomanToInt: want %d got %d", validRomans[roman], ans)
			}
		}
	})
}

func TestRomanToIntegerIncorrect(t *testing.T) {
	t.Run("Check Error Messages for Inorrect Romans", func(t *testing.T) {
		invalidRomans := []string{
			"LXIG",
			"CVXQ",
		}
		for _, roman := range invalidRomans {
			_, err := RomanToInt(roman)
			if !AssertError(err, "illegal character") {
				t.Errorf("Want error got %v", err)
			}
		}
	})
}
