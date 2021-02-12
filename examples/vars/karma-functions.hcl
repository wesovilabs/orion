func filterElements {
  set out {
    value = []
  }
  block {
    set item {
      value = data[_.index]
    }
    set out {
      arrayIndex = len(out)
      value = item
      when = call(filter,item)
    }
    count = len(elements)
  }
  return {
    value = out
  }
}

func calculateKarma {
  arg karma {
    default = 0
  }
  body {

  }
  return {

  }
}


func calculateKarma {
  vars {
    karma  = 0
  }
  block {
    set ingredient {
      value = ingredients[_.index]
    }
    set karma {
      value = karma + 1
      when = call(isCrueltyFree,ingredient)
    }
    set karma {
      value = karma - 1
      when = !call(isCrueltyFree,ingredient)
    }
    count = len(ingredients)
  }
  return {
    value = karma
  }
}

func isCrueltyFree {
  return {
    value = ingredient.vegan
  }
}