package util

import (
	"fmt"
	"hash/fnv"
)

func HashName(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	return fmt.Sprint(h.Sum32())
}

func InList(l []string, n string) bool {
	for _, v := range l {
		if v == n {
			return true
		}
	}
	return false
}
