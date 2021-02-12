scenario "demo include" {
  when "sum two numbers"{
    set result {
      value = x + y
    }
  }
  then "the result is the expected"{
    assert {
      assertion = result == 3
    }
  }
}