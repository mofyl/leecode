package v1

import (
	"fmt"
	"strings"
	"testing"
)

func TestSimplifyPath(t *testing.T) {
	//path := "/home//foo/"
	//path := "/home/../"

	path := "/a/./b/../../c/"
	//path := "/a/../../b/../c//.//"

	res := simplifyPath(path)
	fmt.Println(res)

}

func simplifyPath(path string) string {

	if len(path) < 1 {
		return path
	}

	part := strings.Split(path, "/")
	stack := make([]string, 0, len(part))
	for i := 0; i < len(part); i++ {

		if part[i] == "" || part[i] == "." {
			continue
		} else if part[i] == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, part[i])
		}

	}
	return "/" + strings.Join(stack, "/")
}
