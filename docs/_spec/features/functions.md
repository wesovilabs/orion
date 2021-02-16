---
layout: default
title: Functions
nav_order: 7
parent: Feature
---

<link rel="stylesheet" href="../../../assets/css/custom.css">

# Functions

Writing functions, we can define reusable set of actions to be invoked from different
scenarios.  A function looks like this:

```hcl
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
```

A function is described with a name (`calculateKarma`), and the following three blocks:

- **input**: Input arguments of the function. It's an optional  block.
- **body** : Set of actions to be executed. It's required block.
- **return**: The output value if is provided. It's an optional block.

Additionally, we make use of action `call` to invoke a function.

In the below example, th function `calculateKarma` receives a list of ingredients.
Per each vegan ingredient the karma increases in one if not the karma reduces one.

```hcl
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
      ingredients = items
    }
  }
  then "the karma is -2" {
    assert {
      assertion = karma
    }
  }
}
```
