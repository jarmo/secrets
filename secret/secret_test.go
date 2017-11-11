package secret

import (
  "testing"
  "fmt"
)

func TestNew(t *testing.T) {
  result := New("secret-name", "secret-value")
  if result.Id.String() == "" {
    t.Fatal(fmt.Sprintf("Expected Id to be present: %v", result))
  }
  expectedSecret := Secret{result.Id, "secret-name", "secret-value"}

  if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", expectedSecret) {
    t.Fatal(fmt.Sprintf("Expected secret to be '%v', but got '%v'", expectedSecret, result))
  }
}

func TestString(t *testing.T) {
  result := New("secret-name", "secret-value")
  expectedResult := fmt.Sprintf(`
[%s]
%s
%s`, result.Id, result.Name, result.Value)

  if expectedResult != result.String() {
    t.Fatal(fmt.Sprintf("Expected output to be '%s', but got '%s'", expectedResult, result))
  }
}
