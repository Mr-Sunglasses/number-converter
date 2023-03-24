package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
    "strings"
	"net/http"
)

type RomanNumeral struct {
	Roman string `json:"roman,omitempty"`
}

type romanError struct {
	char string
}

func (e *romanError) Error() string {
	return fmt.Sprintf("illegal character %s in roman", e.char)
}

// POST /convert
// Convert Roman Numeral
func ConvertRoman(c *gin.Context) {
	// Validate input
	var input RomanNumeral
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    err := CheckValidRoman(input.Roman)
    if err != nil {
        c.JSON(http.StatusOK, gin.H{"roman": input.Roman, "error": err.Error()})
    } else {
        decimal, err := RomanToInt(input.Roman)
        if err != nil {
            c.JSON(http.StatusOK, gin.H{"roman": input.Roman, "error": err.Error()})
        } else {
            c.JSON(http.StatusOK, gin.H{"roman": input.Roman, "integer": decimal})
        }
    }
}

func RomanToInt(roman string) (int, error) {
	romanSymbols := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	total := 0
	for index := 0; index < len(roman); index++ {
		if value, ok := romanSymbols[string(roman[index])]; !ok {
			return value, &romanError{string(roman[index])}
		}
		// prevent index out of bound panic
		if index < len(roman)-1 {
			if romanSymbols[string(roman[index])] >= romanSymbols[string(roman[index+1])] {
				total += romanSymbols[string(roman[index])]
			} else {
				total -= romanSymbols[string(roman[index])]
			}
		} else {
			total += romanSymbols[string(roman[index])]
		}
	}
	return total, nil
}

func CheckValidRoman(roman string) error {
    // D, L, V can only appear once in roman numeral
    onlyOnce := []string{"D", "L", "V"}

    for _, ch := range onlyOnce {
        if strings.Count(roman, ch) > 1 {
            return fmt.Errorf("%s cannot only appear once", ch)
        }
    }
    return nil
}
