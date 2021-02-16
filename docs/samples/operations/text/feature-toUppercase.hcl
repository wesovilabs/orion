input {
    arg input {
        default ="hello world"
    }
}
scenario "check eq funcion" {
    when "evaluate a variable" {
        set res1 {
            value = toUppercase(input)
        }
    }
    then "check results"{
        assert {
            assertion = eq(res1,"HELLO WORLD")
        }
    }
}