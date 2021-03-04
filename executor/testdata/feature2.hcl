# feature-math-operations.hcl
description = <<EOF
   This feature is used to demonstrate that both add and subs
   operations work as expected.
EOF
input {
  arg x {
    default = 10
  }
  arg y {
    default = 5
  }
  arg sumResult {
    default = 15
  }
  arg subResult {
    default = 5
  }
}
after each {
  print {
    msg = "the output of this operation is ${result}"
  }
}
includes = [
  "scenario-sum.hcl",
  "scenario-sub.hcl",
  "scenario-mult.hcl"
]
