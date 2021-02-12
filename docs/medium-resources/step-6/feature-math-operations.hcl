description = <<EOF
    This feature is used to demonstrate that both add and subs operations 
    work as expected.
EOF

input {
    arg opSum {
        default = { a = 10,  b = 5, expected = 15 }
    }
    arg opSub {
        default = { a = 10,  b = 5, expected = 5 }
    }
}

after each {
    print {
        msg = "the output of this operation is ${result}"
    }
}


includes = [
    "scenario-sum.hcl",
    "scenario-sub.hcl"
]

