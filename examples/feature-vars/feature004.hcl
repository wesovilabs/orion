
input {
  arg cities {
    default = [
      {name="Madrid", population=4563212},
      {name="Málaga", population=515023},
    ]
  }
}

scenario "do an operation over the elements in the collection" {

  when "it modifies the name of the cities" {
    block {
      set cities {
        arrayIndex = _.index
        value = {
          name = toUppercase(cities[_.index].name)
        }
      }
      count=len(cities)
    }
  }

  then "check the output array" {
    assert {
      description = "It checks the len of the array has not been modified"
      assertion = len(cities)==2
    }
    assert {
      description = "It verifies that name of the cities is in uppercase."
      assertion = (
        cities[0].name=="MADRID" && cities[1].name=="MÁLAGA" &&
        cities[0].population==4563212 && cities[1].population==515023
      )
    }
  }
}
