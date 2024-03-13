package strs

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"text/template"

	"golang.org/x/exp/constraints"
)

func ToInts(ss []string) []int {
	r := make([]int, len(ss))
	for i, v := range ss {
		r[i], _ = strconv.Atoi(v)
	}
	return r
}
func ToInt64s(ss []string) []int64 {
	r := make([]int64, len(ss))
	for i, v := range ss {
		r[i], _ = strconv.ParseInt(v, 10, 64)
	}
	return r
}

func ToIntsDefault(s string, dv ...int) int {
	v, err := strconv.Atoi(s)
	if nil != err {
		if len(dv) > 0 {
			return dv[0]
		}
		return 0
	}
	return v
}
func ToString[T constraints.Integer](ints []T) []string {
	r := make([]string, len(ints))
	for i, v := range ints {
		r[i] = strconv.FormatInt(int64(v), 10)
	}
	return r
}

// FormatTpl format string with text/template
func FormatTpl(s string, args map[string]any) string {
	buf := bytes.NewBuffer(nil)
	template.Must(template.New("").Parse(s)).Execute(buf, args)
	return buf.String()
}

// Format: tpl: "Hello ${name}" {name: "World"} -> "Hello World"
func Format(tpl string, args map[string]any) string {
	return regexp.MustCompile(`\${([^}]*)}`).ReplaceAllStringFunc(tpl, func(s string) string {
		key := s[2 : len(s)-1]
		if v, ok := args[key]; ok {
			if r, ok := v.(string); ok {
				return r
			} else {
				return fmt.Sprintf("%v", v)
			}
		}
		return s
	})
}
