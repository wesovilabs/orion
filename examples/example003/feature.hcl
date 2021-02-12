scenario "basic usage" {
  when "input values are sum" {
    set c {
      value = 2 + 3
    }
  }
  then "the result of variable c is correct" {
    assert {
      assertion = c == 5
    }
  }
}