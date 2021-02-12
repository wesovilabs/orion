scenario "check collection functions" {
  given "set value" {
    set elements{
      value = [
        {
          name = "Sally"
          specie = "dog"
        },
        {
          name = "Molly"
          specie = "dog"
        },
        {
          name = "Coco"
          specie = "dog"
        }
      ]
    }
  }
  when "obtain the element at position 1" {
    set element {
      value = elements[1]
    }

  }
  then "check the element"{
    assert {
      assertion = element.name == "Molly"
    }
  }
  when "obtain the first element" {
    set element {
      value = first(elements)
    }

  }
  then "check the element"{
    assert {
      assertion = element.name == "Sally"
    }
  }
  when "obtain the last element" {
    set element {
      value = last(elements)
    }

  }
  then "check the element"{
    assert {
      assertion = element.name == "Coco"
    }
  }
  when "obtain the number of elements in the list" {
    set totalElement {
      value = len(elements)
    }

  }
  then "check the element"{
    assert {
      assertion = totalElement==3
    }
  }
}