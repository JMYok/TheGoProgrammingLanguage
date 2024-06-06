package trychannel

import (
	"testing"
)

func TestClosedChannel(t *testing.T) {
	ch := make(chan int)
	close(ch)
	t.Log(<-ch)
}
