package dueta

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

const WHITESPACE = " \t\n\r\v\f"
const ASCII_LOWERCASE = "abcdefghijklmnopqrstuvwxyz"
const ASCII_UPPERCASE = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const ASCII_LETTERS = ASCII_LOWERCASE + ASCII_UPPERCASE
const DIGITS = "0123456789"
const HEXDIGITS = DIGITS + "abcdef" + "ABCDEF"
const OCTDIGITS = "01234567"
const PUNCTUATION = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
const PRINTTABLE = DIGITS + ASCII_LETTERS + PUNCTUATION + WHITESPACE

func str(value any) string {
	return fmt.Sprintf("%v", value)
}
func split(value string) []string {
	return strings.Split(value, "")
}

func split_by_sep(value string, sep string) []string {
	return strings.Split(value, sep)
}

func split_by_sep_max(value string, sep string, max int) []string {
	return strings.SplitN(value, sep, max)
}

func strip(value string, chars ...string) string {
	if len(chars) == 0 {
		return strings.TrimSpace(value)
	} else {
		for _, v := range chars {
			value = strings.Trim(value, v)
		}
		return value
	}

}

func integer(value any) int {
	value = str(value)
	value = strip(value.(string))
	intValue, err := strconv.Atoi(str(value))
	if err != nil {
		panic(err)
	} else {
		return intValue
	}
}

func print(values ...any) {
	for _, value := range values {
		value = str(value) + " "
		fmt.Print(value)
	}
	fmt.Println()
}

func input(__prompt ...string) string {
	var prompt string
	var value string
	if len(__prompt) == 0 {
		prompt = ""
	} else {
		prompt = __prompt[0]
	}
	fmt.Print(prompt)
	fmt.Scanf("%s\n", &value)
	return value
}

func includes[T comparable](findItem T, slice []T) bool {
	for _, el := range slice {
		if el == findItem {
			return true
		}
	}
	return false
}

func vtype(value any) string {
	return reflect.TypeOf(value).String()
}

func float(value any) float64 {
	value = str(value)
	value = strip(value.(string))
	float64Value, err := strconv.ParseFloat(str(value), 64)
	if err != nil {
		panic(err)
	} else {
		return float64Value
	}
}

func lower(value string) string {
	return strings.ToLower(value)
}

func upper(value string) string {
	return strings.ToUpper(value)
}

func join(slice ...any) string {
	result := ""
	for _, el := range slice {
		result = result + str(el)
	}
	return result
}

func encode(value string) []byte {
	return []byte(value)
}

func decode(value []byte) string {
	return string(value)
}

