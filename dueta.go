package duata

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"

	// "path/filepath"
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

func Str(value any) string {
	return fmt.Sprintf("%v", value)
}
func Split(value string) []string {
	return strings.Split(value, "")
}

func Split_by_sep(value string, sep string) []string {
	return strings.Split(value, sep)
}

func Split_by_sep_max(value string, sep string, max int) []string {
	return strings.SplitN(value, sep, max)
}

func Strip(value string, chars ...string) string {
	if len(chars) == 0 {
		return strings.TrimSpace(value)
	} else {
		for _, v := range chars {
			value = strings.Trim(value, v)
		}
		return value
	}

}

func Integer(value any) int {
	value = Str(value)
	value = Strip(value.(string))
	intValue, err := strconv.Atoi(Str(value))
	if err != nil {
		panic(err)
	} else {
		return intValue
	}
}

func Print(values ...any) {
	for _, value := range values {
		value = Str(value) + " "
		fmt.Print(value)
	}
	fmt.Println()
}

func Input(__prompt ...string) string {
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

func Includes(findItem any, slice any) bool {
	if Type(slice) == "string" {
		findItem = Str(findItem)
		return strings.Contains(slice.(string), findItem.(string))
	} else {
		return slices.Contains(slice.([]any), findItem)
	}

}

func Type(value any) string {
	return reflect.TypeOf(value).String()
}

func Float(value any) float64 {
	value = Str(value)
	value = Strip(value.(string))
	float64Value, err := strconv.ParseFloat(Str(value), 64)
	if err != nil {
		panic(err)
	} else {
		return float64Value
	}
}

func Lower(value string) string {
	return strings.ToLower(value)
}

func Upper(value string) string {
	return strings.ToUpper(value)
}

func Join(step string, slice ...any) string {
	result := ""
	for _, el := range slice {
		result = result + step + Str(el)
	}
	return result
}

func Encode(value string) []byte {
	return []byte(value)
}

func Decode(value []byte) string {
	return string(value)
}

func Exists_path(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		return true
	}
}

func Makedirs(path string) {

	path_tmp := Get_root_dir(path)
	for _, dir := range Split_path(path)[1:] {
		os.MkdirAll(path, os.ModePerm)
		path_tmp = path_tmp + "/" + dir
	}
}

func Get_root_dir(path string) string {
	path = filepath.FromSlash(path)
	sep := string(filepath.Separator)
	if Includes("://", path) {
		root_2 := Split_by_sep(path, sep)[1:][1]
		root_1 := Split_by_sep(path, "://")[0]
		root := root_1 + "://" + root_2
		return root
	} else {
		root_1 := Split_by_sep(path, "/")[0]
		root_2 := Split_by_sep(path, "/")[1]
		root := root_1 + "/" + root_2
		return root
	}
}

func Split_path(path string) []string {
	path = filepath.FromSlash(path)
	sep := string(filepath.Separator)
	if Includes("://", path) {
		root_2 := Split_by_sep(path, sep)[1:][1]
		root_1 := Split_by_sep(path, "://")[0]
		root := root_1 + "://" + root_2
		paths := Split_by_sep(path, "/")[3:]
		data := []string{root}
		data = append(data, paths...)
		return data
	} else {
		root_1 := Split_by_sep(path, "/")[0]
		root_2 := Split_by_sep(path, "/")[1]
		root := root_1 + "/" + root_2
		paths := Split_by_sep(path, "/")[2:]
		data := []string{root}
		data = append(data, paths...)
		return data
	}

}

func Startswith(findValue string, value string) bool {
	return strings.HasPrefix(value, findValue)
}
func Endswith(findValue string, value string) bool {
	return strings.HasSuffix(value, findValue)
}
func Open(path string, mode string, custom ...any) *os.File {
	var flags = os.O_CREATE

	if Startswith("a", mode) {
		flags = flags | os.O_APPEND
	} else if Startswith("w", mode) {
		flags = flags | os.O_WRONLY
	} else if Startswith("r", mode) {
		flags = flags | os.O_RDONLY
	}

	fp, _ := os.OpenFile(path, flags, 755)
	return fp

}
