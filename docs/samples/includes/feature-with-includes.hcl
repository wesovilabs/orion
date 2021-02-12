includes = [
    "hooks.hcl"
]

scenario "scenario with tag smoke-test" {
    tags =  ["smoke-test","other"]
    when "assign true to a local variable" {
        set a {
            value = true
        }
    }
    then "the value of the variable is true" {
        assert {
            assertion = a
        }
    }
}
scenario "scenario with tag smoke-test" {
    tags =  ["performance-test","other"]
    when "assign true to a local variable" {
        set a {
            value = true
        }
    }
    then "the value of the variable is true" {
        assert {
            assertion = a
        }
    }
}