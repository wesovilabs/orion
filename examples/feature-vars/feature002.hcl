
input {
  arg factor {
    default = 10
  }
  arg elements {
    default = [2,3,4,5,6]
  }
}

scenario "do an operation over the elements in the collection" {
  given "initialize the accumulator"{
    set accumulator {
      value = 0
    }
  }
  when "process elements in the collection" {
    block {
      print {
        msg = "add ${elements[_.index] * factor} to ${accumulator}"
      }
      set accumulator {
        value = accumulator + (factor * elements[_.index])
      }
      count = len(elements)
    }
  }
  then "check the output array" {
    assert {
      assertion = accumulator==200
    }
  }
}