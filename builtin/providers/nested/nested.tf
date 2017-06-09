resource "nested" "example" {
  name = "Test"
  foo {
    name = "Foo Test"
    bar {
      name = "Bar Test"
      baz {
        name = "Baz Test"
      }
    }
  }
}