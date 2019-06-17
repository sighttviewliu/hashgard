package gov

import "strings"

// camel string, xx_yy to XxYy
func CamelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}
func GetParamKey(keystr string) string {
	strs := strings.Split(keystr, "/")
	if len(strs) != 2 {
		return ""
	}
	return strs[1]
}

func GetParamSpaceFromKey(keystr string) string {
	strs := strings.Split(keystr, "/")
	if len(strs) != 2 {
		return ""
	}
	return strs[0]
}
