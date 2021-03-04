scenario "operation subtract" {
    when "values are subtracted" {
        set result {
            value = opSub.a - opSub.b
            when = opSub.a > opSub.b
        }
        set result {
            value = opSub.b - opSub.a
            when = opSub.a <=  opSub.b
        }
    }
    then "the result of the operation is the expected" {
        assert {
            assertion = result==opSub.expected
        }
    }
}
