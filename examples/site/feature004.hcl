scenario "do some math operations" {
  given "a number higher than zero" {
    set number {
      value  = 5
    }
  }
  when "the number is multiplied  by itself" {
    set result {
      value = 5 * 5
    }
  }
  then "the square root of the result is input number"{
    assert {
      assertion = number==sqrt(result)
    }
  }
  when "the number is added to the current value" {
    set result {
      value = result + 5
    }
  }
  then "the result is the expected"{
    assert {
      assertion = number==sqrt(result - number)
    }
  }
}