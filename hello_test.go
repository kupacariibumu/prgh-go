package hello

import "teating"

func HelloTest(t *testing.T) {
    expected := "Hello Hai"
    if ret := hello(); ret != expected {
      t.Errorf("hello() = %q", ret)
    }
}
