package main

import (
	"fmt"

	"github.com/sgtest/go15vendor/bye"
	"github.com/sourcegraph/john-test/hi"
)

func main() {
	fmt.Println(hi.GetHi())
	fmt.Println(bye.GetBye())
}
