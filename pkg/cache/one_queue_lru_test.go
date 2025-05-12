package cache

import (
	"fmt"
	"testing"
)

func TestOneQueueLru_Set(t *testing.T) {
	c := NewOneQueueLru[string, string, string](3)
	c.Set("a", "1", "a1")
	c.Set("a", "2", "a2")
	c.Set("a", "3", "a3")
	c.Set("a", "4", "a4")
	c.Set("a", "5", "a5")
	c.Set("b", "1", "a1")
	c.Set("b", "5", "a5")
	c.Set("b", "8", "a8")
	c.Set("b", "3", "a3")
	res1 := c.ValuesByFirstKey("a")
	res2 := c.ValuesByFirstKey("b")
	fmt.Println(res1, res2)
}
