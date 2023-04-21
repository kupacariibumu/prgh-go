package hello

import "testing"

func TestHello(t *testing.T) {
    expected := "Hello Hai"
    if ret := Hello(); ret != expected {
      t.Errorf("hello() = %q want %q", ret, expected)
    }
}
