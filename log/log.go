package log

import (
	"fmt"
	"os"
	"time"
)

const layout = "2006-01-02 15:04:05"

func Fatal(s string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, "[%s] gotunnel: %s\n", time.Now().Format(layout), fmt.Sprintf(s, a...))
	os.Exit(2)
}

func Log(msg string, r ...interface{}) {
	fmt.Printf("[%s] %s\n", time.Now().Format(layout), fmt.Sprintf(msg, r...))
}

func Info(msg string, r ...interface{}) {
	fmt.Printf(msg, r...)
}
