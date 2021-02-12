input {
  arg x {
    default = 1
  }
  arg y {
    default = 2
  }
}

scenario "scenario fails" {
  continueOnError=true
  when "sum two numbers"{
    set result {
      value = x + y
    }
  }
  then "the result is the expected"{
    assert {
      assertion = result == 2
    }
  }
}

scenario "scenario success" {
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