package gen

import (
	"fmt"
	"testing"
)

func TestImportHandle(t *testing.T) {
	fmt.Println(importHandle("a.a"))
	fmt.Println(importHandle("a./a"))
}