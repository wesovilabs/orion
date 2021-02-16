vars {
  items = [
    { product = "tofu",     vegan  = true   },
    { product = "meat",     vegan  = false  },
    { product = "fish",     vegan  = false  },
    { product = "avocado",  vegan  = true   },
    { product = "sheep",   vegan  = false   },
    { product = "rabbit",  vegan  = false   },
  ]
}

func calculateKarma {
  input {
    arg ingredients {}
  }
  body {
    set karma {
      value = 0
    }
    block {
      set karma {
        value = karma + 1
        when = _.item.vegan
      }
      set karma {
        value = karma - 1
        when = !_.item.vegan
      }
      items = ingredients
    }
  }
  return {
    value = karma
  }
}

scenario "calculate karma" {
  when "calculate the karma"  {
    call calculateKarma {
      with{
        ingredients = items
      }
      as = "karma"
    }
  }
  then "the karma is -2" {
    assert {
      assertion = _
    }
  }
}
