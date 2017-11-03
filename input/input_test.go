package input

import (
  "testing"
  "strings"
  "fmt"
)

func TestReplaceUnprintableCharacters(t *testing.T) {
  ctrlD, ctrlX, ctrlZ := "\x04", "\x18", "\x1A"
  expectedInput := "foo"
  input := []string{ctrlD, ctrlX, ctrlZ, expectedInput, ctrlD, ctrlX, ctrlZ}

  result := replaceUnprintableCharacters(strings.Join(input, ""))

  if result != expectedInput {
    t.Fatal(fmt.Sprintf("Expected result to be '%v', but got '%v'", []byte(expectedInput), []byte(result)))
  }
}
