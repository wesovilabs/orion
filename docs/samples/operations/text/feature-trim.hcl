input {
    arg input {
        default =" hello world "
    }
}
scenario "check trim funcion" {
    when "evaluate a variable" {
        set res1 {
            value = trim(input)
        }
    }
    then "check results"{
        assert {
            assertion = eq(res1,"helloworld")
        }
    }
}