before operation {
  set x {
    value = 1
  }
  set y {
    value = 2
  }
}

after sum {
  print {
    msg = "${x} + ${y} = ${result}"
  }
}

after sub {
  print {
    msg = "${x} - ${y} = ${result}"
  }
}

scenario "example of default hooks I" {
  tags = ["operation", "sum"]
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
  tags = ["operation", "sub"]
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