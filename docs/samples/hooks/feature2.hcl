before each {
  description = "common hook to be executed before each scenario"
  set a {
    value = 1
  }
}

after each {
  description = "common hook to be executed after each scenario"
  print {
    msg = "value of b is ${b}"
  }
}

scenario "scenario3" {
  tags = ["common"]
  when "do a sum" {
    set b {
      value = 2 * a
    }
  }
  then "verify the result"{
    assert {
      assertion= b==2
    }
  }
}