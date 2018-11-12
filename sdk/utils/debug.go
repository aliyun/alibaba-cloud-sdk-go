package utils

import (
	"fmt"
	"os"
	"strings"
)

type Debug func(format string, v ...interface{})

func Init(flag string) Debug {
	enable := false

	env := os.Getenv("DEBUG")
	parts := strings.Split(env, ",")
	for _, part := range parts {
		if part == flag {
			enable = true
			break
		}
	}

	return func(format string, v ...interface{}) {
		if enable {
			fmt.Println(fmt.Sprintf(format, v...))
		}
	}
}
