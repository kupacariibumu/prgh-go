package hello

func HelloTest(t *testing.T) {
    expected := "Hello Hai"
    if ret := Hello(); ret != expected {
      t.Errorf("hello() = %q", ret)
    }
}
