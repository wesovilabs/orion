scenario "operation substract" {
    ignore = opSub.a > 10
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