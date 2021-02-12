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

scenario "operation add" {
    when "values are added" {
        set result {
            value = opSum.a + opSum.b 
        }
    } 
    then "the result of the operation is the expected" {
        assert {
            assertion = result==opSum.expected
        }
    }
}

scenario "operation substract" {
    when "values are subtracted" {
        set result {
            value = opSub.a - opSub.b 
        }
    } 
    then "the result of the operation is the expected" {
        assert {
            assertion = result==opSub.expected
        }
    }
}