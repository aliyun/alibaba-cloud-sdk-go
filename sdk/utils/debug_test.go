package utils

import (
  "testing"
  // "fmt"
)

func TestMain(t *testing.T) {
  debug := Init("sdk")
  debug("%s", "testing")
}
