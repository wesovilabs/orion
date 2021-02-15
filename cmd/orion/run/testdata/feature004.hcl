input {
  arg x {
    default = 10
  }
  arg y {
    default = 5
  }
  arg out {
    default = 50
  }
}

scenario "scenario demo" {
  when "multiply x * y" {
    set result {
      value = x * y
    }
  }
  then "check the output" {
    assert {
      assertion = result==out
    }
  }
}
