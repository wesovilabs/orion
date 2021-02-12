
input {
  arg factor {
    default = 10
  }
  arg elements {
    default = [2,3,4,5,6]
  }
}

scenario "do an operation over the elements in the collection" {
  given "initialize output array" {
    set output {
      value = []
    }
  }
  when "process elements in the collection" {
    block {
      set output {
        arrayIndex = _.index
        value =  factor * elements[_.index]
      }
      count = len(elements)
    }
  }

  then "check the output array contains the expected items" {
    assert {
      description = "It check the len of the output array is the expected"
      assertion = len(elements)==len(output)
    }
    assert {
      description = "It verifies that each element in the array is the expected."
      assertion = elements[_.index]*factor==output[_.index]
      count = len(output)
    }
  }
}
