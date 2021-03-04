# feature-math-operations.hcl
description = <<EOF
   This feature is used to demonstrate that both add and subs
   operations work as expected.
EOF

scenario "operation add" {
  given "the variables x and y" {
    set x {
      value = 10
    }
    set y{
      value = 5
    }
  }
  when "values are added" {
    set result {
      value = x + y
    }
    print {
      msg = "${x} + ${y} is ${result}"
    }
  }
  then "the result of the operation is the expected" {
    assert {
      assertion = result==15
    }
  }
}
scenario "operation substract" {
  given "variables x and y" {
    set x {
      value = 10
    }
    set y{
      value = 5
    }
  }
  when "subtract y to x" {
    set result {
      value = x - y
    }
    print {
      msg = "${x} - ${y} is ${result}"
    }
  }
  then "the result of the operation is the expected" {
    assert {
      assertion = result==5
    }
  }
}
