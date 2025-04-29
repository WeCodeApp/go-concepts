package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	// string conversion
	num := 18
	str1 := strconv.Itoa(num)
	fmt.Println("num length:", len(str1))

	// string splitting
	fruits1 := "apple,banana,orange"
	parts1 := strings.Split(fruits1, ",")
	fmt.Println("Fruit parts:", parts1)

	fruits2 := "apple-banana-orange"
	parts2 := strings.Split(fruits2, "-")
	fmt.Println("Fruit parts:", parts2)

	//string joining
	countries := []string{"India", "USA", "China"}
	joined := strings.Join(countries, ", ")
	fmt.Println("countries joined:", joined)

	str := "Hello, Go!"
	fmt.Println("Does str have Go?:", strings.Contains(str, "Go"))

	replaced := strings.Replace(str, "Go", "Golang", 1)
	fmt.Println("Replaced Go with Golang:", replaced)

	strwspace := "   Hello, Everyone!   "
	fmt.Println("String with leading and trailing spaces:", strwspace)
	fmt.Println("Removed leading and trailing spaces:", strings.TrimSpace(strwspace))

	fmt.Println("Convert string to lowercase:", strings.ToLower(strwspace))
	fmt.Println("Convert string to uppercase:", strings.ToUpper(strwspace))

	fmt.Println("repeat foo 3 times:", strings.Repeat("foo", 3))

	fmt.Println("Count no of l in hello:", strings.Count("hello", "l"))
	fmt.Println("Is Hello preffic He?:", strings.HasPrefix("Hello", "He"))
	fmt.Println("Is Hello suffix lo?:", strings.HasSuffix("Hello", "lo"))
	fmt.Println("Is Hello suffix la?:", strings.HasSuffix("Hello", "la"))

	str2 := "Hel1lo, 123 Go 11!"
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(str2, -1)
	fmt.Println(matches)

	str3 := "Hello 你好\n"
	fmt.Println(utf8.RuneCountInString(str3))
}
