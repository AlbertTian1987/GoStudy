package strings

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"unicode"
)

func Test_StringFmt(t *testing.T) {
	t.Log(StringsFmt1("a", "b", "c", "d", 1, false, 5.754))
	t.Log(StringsFmt2("a", "b", "c", "d", 1, false, 5.754))

	t.Log(StringsFmt1("aaaaaaa"))
	t.Log(StringsFmt2("aaaaaaa"))
}

func Test_Strings(t *testing.T) {

	s := []string{"a", "b", "c", "d", "e", "f", "g"}

	fmt.Println("strings.Join = ", strings.Join(s, ","))

	fmt.Println("strings.Index = ", strings.Index("Child", "ild"))
	fmt.Println("strings.Index = ", strings.Index("Child", "ildx"))

	fmt.Println("strings.Repeat = ", strings.Repeat("hello ", 5))

	fmt.Println("strings.Replace = ", strings.Replace("old old old", "ol", "xx", 2))
	fmt.Println("strings.Replace = ", strings.Replace("old old old", "ol", "i'm", -1))

	fmt.Println("strings.Split = ", strings.Split("a,b,c", ","))
	fmt.Println("strings.Split = ", strings.Split("a,b,c", "xxxx"), "len=", len(strings.Split("a,b,c", "xxxx")))

	fmt.Println("strings.TrimFunc = ", strings.TrimFunc(" !!! Achtung !!! ", func(i rune) bool {
		return i == rune(' ') || i == rune('!')
	}))

	fmt.Println("strings.Trim = ", strings.Trim(" !!! Achtung !!! ", "! "))

	fmt.Println("strings.Fields = ", strings.Fields("  foo bar  baz he"))
	fmt.Println("strings.FieldsFunc = ", strings.FieldsFunc("  foo bar  baz he", func(i rune) bool {
		return unicode.IsSpace(i) || i == rune('b')
	}))

}

func Test_strconv(t *testing.T) {
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10)
	t.Log(string(str))
	str = strconv.AppendBool(str, false)
	t.Log(string(str))
	str = strconv.AppendBool(str, true)
	t.Log(string(str))
	str = strconv.AppendQuote(str, "abcdefg")
	t.Log(string(str))
	str = strconv.AppendQuoteRune(str, '单')
	t.Log(string(str))
}

func Test_format(t *testing.T) {
	a := strconv.FormatBool(false)
	b := strconv.FormatFloat(123.23, 'g', 12, 64)
	c := strconv.FormatInt(1234, 10)
	d := strconv.FormatUint(12345, 10)
	e := strconv.Itoa(1023)
	fmt.Println(a, b, c, d, e)
}
