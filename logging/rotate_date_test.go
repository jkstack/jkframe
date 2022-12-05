package logging

import (
	"testing"
)

func TestLog(t *testing.T) {
	SetDateRotate(DateRotateConfig{
		Dir:         "./logs",
		Name:        "test",
		Rotate:      7,
		WriteStdout: true,
		WriteFile:   true,
	})
	for i := 0; i < 10000; i++ {
		Info("i=%d", i)
	}
}

func TestGetLogOutput(t *testing.T) {
	str := Info("foo %s", "bar")
	if str != "[INFO]foo bar" {
		t.Fatal("unexpected info log")
	}
}
