scenario "do some math opertions" {
    given "a number higher than zero" {
        set number {
            value  = 5
        }
    }
    when "the number is multiplied  by itself" {
        set result {
            value = 5 * 5
        }
    }
    then "the result is the square root of the number"{
        assert {
            assertion = result==sqr(number)
        }
    }
    when "the numer is added to the current value" {
        set result {
            value = result + 5
        }
        print {
            msg = "the result is ${result}"
        }
    }
    then "the result is the expected"{
        assert {
            assertion = result==sum(sqr(number),number)
        }
    }
}