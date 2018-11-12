package utils

import (
	"testing"
)

func TestMain(t *testing.T) {
	debug := Init("sdk")
	debug("%s", "testing")
}
