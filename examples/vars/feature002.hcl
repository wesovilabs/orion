includes = [
  "karma-functions.hcl"
]

vars {
  x = 1
  items = [
    { product = "tofu",     vegan  = true   },
    { product = "meat",     vegan  = false  },
    { product = "fish",     vegan  = false  },
    { product = "avocado",  vegan  = true   },
  ]
}

func calculateKarma {
  arg ingredients {
    default = []
  }
  arg karma {
    default = 0
  }
  body {
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
  given "initial karma" {
    set karma {
      value = 0
    }
  }
  when "calculate the karma"  {
    call calculateKarma {
      ingredients = items
    }
  }
  then "the karma is 0" {

  }
}
