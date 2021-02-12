input {
    arg people {
        default = [
          { firstname = "John", lastname = "Doe"},
          { firstname = "Jane", lastname = "Doe"},
        ]
    }
    arg company {
        default = "Wesovilabs"
    }
}
scenario "print variables" {
  when "iterate over the people"{
    block{
        set person {
            value = people[_.index]
        }
        print {
            msg = "${person.firstname} ${person.lastname} is hiread at ${company}"  
        }
        while = _.index < len(people)
    }
  }
  then "verify the postconditions" {
    assert {
      assertion = eqIgnoreCase(company,"wesovilabs")
    }
  }
}