scenario "basic usage" {
  tags = ["basic", "usage"]
  given "a couple of numbers" {
    set a {
      value = 1
    }
    set b {
      value = 2
    }
  }
  when "input values are sum" {
    set c {
      value = a + b
    }
  }
  then "the result of variable c is correct" {
    assert {
      assertion = c == 3
    }
  }
}
