scenario "square root operation work as expected" {
    tags = ["maths"]
    ignore = false
    continueOnError = true
    given "a number higher than zero" {
        set number {
            value  = 5
        }
    }
    when "the number is multiplied  by itself" {
        set result {
            value = 5 * 5
        }
        print {
            msg  = "The square root of ${number} is ${result}"
        }
    }
    then "the result is the square root of the number"{
        assert {
            assertion = result==sqr(number)
        }
    }
}