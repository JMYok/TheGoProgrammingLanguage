package main

import (
	"strings"
	"testing"
)

func TestConcat(t *testing.T) {
	a := []string{"a", "b", "c"}
	var sb strings.Builder
	sb.WriteString(a[0])
	sb.WriteString(a[1])
	sb.WriteString(a[2])
	ret := sb.String()
	t.Log(ret)
}
