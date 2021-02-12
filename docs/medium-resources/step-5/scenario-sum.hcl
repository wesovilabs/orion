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