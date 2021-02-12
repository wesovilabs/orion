input {
    arg input {
        default ="Hello World"
    }
}
scenario "check eq funcion" {
    when "evaluate a variable" {
        set res1 {
            value = toLowercase(input)
        }
    }
    then "check results"{
        assert {
            assertion = eq(res1,"hello world")
        }
    }
}