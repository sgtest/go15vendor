package main

import (
	"fmt"

	"github.com/sgtest/go-vendored-lib/hi"
	"github.com/sgtest/go15vendor/bye"
)

func main() {
	fmt.Println(hi.GetHi())
	fmt.Println(bye.GetBye())
}
