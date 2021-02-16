input {
    arg input {
        default ="hello world"
    }
}
scenario "check eq funcion" {
    when "evaluate a variable" {
        set res1 {
            value = eq(input,"Hello World")
        }
        set res2 {
            value = eq(input,"hello world")
        }
    }
    then "check results"{
        assert {
            assertion = !res1 && res2
        }
    }
}