package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println


func main() {

	p("Contains: ", s.Contains("test", "es")) // true
	p("Count: ", s.Count("test", "t")) // 2
	p("HasPrefix: ", s.HasPrefix("test", "te")) // true
	p("HasSuffix: ", s.HasSuffix("test", "st")) // true
	p("Index: ", s.Index("test", "e")) // 1
	p("Join: ", s.Join([]string{"a", "b"}, "-")) // a-b
	p("Repeat: ", s.Repeat("a", 5)) // aaaaa
	p("Replace: ", s.Replace("foo", "o", "0", -1)) // f00
	p("Replace: ", s.Replace("foo", "o", "0", 1)) // f0o
	p("Split: ", s.Split("a-b-c-d-e", "-")) // [a b c d e]
	p("ToLower: ", s.ToLower("TEST")) // test
	p("ToUpper: ", s.ToUpper("test")) // TEST
}