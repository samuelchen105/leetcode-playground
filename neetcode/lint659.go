package problem

import (
	"fmt"
	"strconv"
	"strings"
)

func encode(strs []string) string {
	var bd strings.Builder
	for _, s := range strs {
		fmt.Fprintf(&bd, "%d:%s", len(s), s)
	}
	return bd.String()
}

func decode(str string) []string {
	rs := []rune(str)
	ret := []string{}
	i := 0
	for i < len(rs) {
		j := i
		for rs[j] != ':' {
			j++
		}
		length, _ := strconv.Atoi(string(rs[i:j]))
		i = j + 1 + length
		sub := string(rs[j+1 : i])
		ret = append(ret, sub)
	}

	return ret
}
