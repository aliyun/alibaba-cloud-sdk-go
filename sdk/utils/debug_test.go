package utils

import (
  "testing"
  // "fmt"
)

func TestMain(m *testing.M) {
  debug := Init("sdk")
  debug("%s", "testing")
}
