package serve_test

import (
	"fmt"
	"net/url"
	"testing"
)

// Test_IntMapToString []int to []string
func Test_URLAddParam(t *testing.T) {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	//解析这个 URL 并确保解析没有出错。
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	m, _ := url.ParseQuery(u.RawQuery)
	m.Add("open_id", "xxxx")
	u.RawQuery = m.Encode()
	fmt.Println(u.String())

}
