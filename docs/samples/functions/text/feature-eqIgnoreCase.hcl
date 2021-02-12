input {
    arg input {
        default ="hello world"
    }
}
scenario "check eqIgnoreCase funcion" {
    when "evaluate a variable" {
        set res1 {
            value = eqIgnoreCase(input,"Hello World")
        }
        set res2 {
            value = eqIgnoreCase(input,"hello world")
        }
    }
    then "check results"{
        assert {
            assertion = res1 && res2
        }
    }
}