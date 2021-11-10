package main

import (
	"fmt"

	"github.com/cardil/deviate/pkg/metadata"
)

func main() {
	fmt.Println("Version: ", metadata.Version) //nolint:forbidigo
}
