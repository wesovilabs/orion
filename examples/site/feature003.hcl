scenario "square root operation work as expected" {
  tags = ["maths"]
  ignore = false
  continueOnError = true
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
}