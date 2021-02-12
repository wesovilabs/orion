before each {
  set x {
    value = 1
  }
  set y {
    value = 2
  }
}

after each {
  print {
    prefix=">>"
    msg = "the result is ${result}"
  }
}

scenario "example of default hooks I" {
  when "do a sum" {
    set result {
      value = x + y
    }
  }
  then "the result of variable is the expected" {
    assert {
      assertion = result == 3
    }
  }
}

scenario "example of default hooks II" {
  when "do a subtract" {
    set result {
      value = x - y
    }
  }
  then "the result of variable is the expected" {
    assert {
      assertion = result == -1
    }
  }
}