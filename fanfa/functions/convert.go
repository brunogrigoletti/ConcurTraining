package functions

import (
	"fmt"
	"strings"
)

func Functionsfunc() {
	author := "thiagão da massa"
	course := "go basics for dummies"

	fmt.Println(converter(author, course))
}

func converter(author, course string) (a, c string) {
	author = strings.ToUpper(author)
	course = strings.Title(course)

	return author, course
}
