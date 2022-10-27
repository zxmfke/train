package list

import (
	"fmt"
	"testing"
)

func TestNewListRoot(t *testing.T) {
	root := NewListRoot()
	root.Set("a", 1)
	root.Set("b", 2)
	root.Set("a", 3)

	fmt.Println(fmt.Sprintf("%s", root))
	root.Delete("a")

	fmt.Println(fmt.Sprintf("%s", root))

	root.Delete("b")

	fmt.Println(fmt.Sprintf("%s", root))
}
