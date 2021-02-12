description = <<EOF
    This feature is used to demonstrate that both add and subs operations 
    work as expected.
EOF

scenario "operation add" {
    given "the input of the operation" {
        set a {
            value = 10
        }    
        set b {
            value = 5
        }
    }
    when "values are added" {
        set result {
            value = a + b 
        }
        print {
            msg = "${a} + ${b} is ${result}"
        }
    } 
    then "the result of the operation is the expected" {
        assert {
            assertion = result==15
        }
    }
}

scenario "operation substract" {
    given "the input of the operation" {
        set a {
            value = 10
        }    
        set b {
            value = 5
        }
    }
    when "values are subtracted" {
        set result {
            value = a - b 
        }
        print {
            msg = "${a} - ${b} is ${result}"
        }
    } 
    then "the result of the operation is the expected" {
        assert {
            assertion = result==5
        }
    }
}