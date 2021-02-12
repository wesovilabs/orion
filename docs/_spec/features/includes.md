---
layout: default
title: Includes
nav_order: 4
parent: Feature
---
# Includes

<link rel="stylesheet" href="../../../assets/css/custom.css">

The purpose of use includes is to reuse any feature part between multiple executions. 
The includes are defined in attribute `includes`,  of type array, in the root of the feature.

Let's go through the below example. We could define a file with common hooks.

**Example: common-hooks.hcl** [download](https://raw.githubusercontent.com/wesovilabs/orion-examples/master/site/common-hooks.hcl)
```hcl
before sample {
  set a {
    value = 1
  }
}
before common {
  set b {
    value = 2
  }
}
after common {
  print {
    msg = "operation completed with result ${result}"
  }
}
```

The common-hooks.hcl file could be included in any feature:

**Example** [download](https://raw.githubusercontent.com/wesovilabs/orion-examples/master/site/feature008.hcl)
```hcl
includes = [
  "common-hooks.hcl"
]

scenario "scenario1" {
  tags = ["common", "sample"]
  when "a and b are added" {
    set result {
      value = a + b
    }
  }
  then "the result is the expected"{
    assert {
      assertion =  result==3
    }
  }
}
scenario "scenario2" {
  tags = ["common"]
  when "b is duplicated" {
    set result {
      value = 2 * b
    }
  }
  then "verify the result"{
    assert {
      assertion = result==4
    }
  }
}
```
