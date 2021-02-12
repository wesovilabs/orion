vars {
  x = 1
  elements = [
    { product = "tofu",     vegan  = true},
    { product = "meat",     vegan  = false},
    { product = "fish",     vegan  = false},
    { product = "avocado",  vegan  = true},
  ]
}


scenario "feature with plain var" {
  when "multiply by 2" {
    set result{
      value = x * 2
    }
  }
  then "check output" {
    assert {
      assertion = result==2
    }
  }
}

scenario "feature with array var" {
  given "initial karma"{
    set karma {
      value = 0
    }
  }
  when "calculate karma" {
    block {
      set karma {
        value = karma + 1
        when = elements[_.index].vegan
      }
      set karma {
        value = karma - 1
        when = !elements[_.index].vegan
      }
      count = len(elements)
    }
  }
  then "check output" {
    assert {
      assertion = karma==0
    }
  }
  when "delete dead life" {
    set goodElements {
      value = []
    }
    block {
      set goodElements {
        arrayIndex = len(goodElements)
        value = elements[_.index]
        when = elements[_.index].vegan
      }
      count = len(elements)
    }
    set elements {
      value = goodElements
    }
  }
  then "check the good elements" {
    assert {
      assertion = len(goodElements)==2
    }
  }
  when "calculate karma" {
    block {
      set karma {
        value = karma + 1
        when = elements[_.index].vegan
      }
      set karma {
        value = karma - 1
        when = !elements[_.index].vegan
      }
      count = len(elements)
    }
  }
  then "check output" {
    assert {
      assertion = karma==2
    }
  }
}