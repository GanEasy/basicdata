package serve_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// Test_IntMapToString []int to []string
func Test_IntMapToString(t *testing.T) {
	welfare := []int{1, 2, 3}

	var paramSlice []string
	for _, param := range welfare {
		paramSlice = append(paramSlice, strconv.Itoa(param))
	}

	fmt.Println(len(welfare), strings.Join(paramSlice, ","))
}
