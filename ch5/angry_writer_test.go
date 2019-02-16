package ch5

import (
	"bytes"
	"testing"
)

func TestNewAngryReader(t *testing.T) {
	input := bytes.NewBufferString("foo")
	output := make([]byte, 3)

	angry := NewAngryReader(input)
	_, err := angry.Read(output)
	if err != nil {
		t.Errorf("failed with err: %s", err)
	}

	if "FOO" != string(output) {
		t.Errorf("expected FOO but was: %s", string(output))
	}
}
