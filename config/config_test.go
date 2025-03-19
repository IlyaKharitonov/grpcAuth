package config

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	Parse()

	fmt.Println(GetConfig())
}
