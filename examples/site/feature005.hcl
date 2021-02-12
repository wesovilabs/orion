scenario "operation add" {
  examples = [
    { a = 3,  b = 2,  c = 5},
    { a = 4,  b = 2,  c = 6},
    { a = 8,  b = 12, c = 20},
    { a = 13, b = 12, c = 25},
  ]
  when "values are added" {
    set result {
      value = a + b
    }
  }
  then "the result of the operation is the expected" {
    assert {
      assertion = result==c
    }
  }
}