---
layout: default
title: Hooks
nav_order: 3
parent: Feature
---

<link rel="stylesheet" href="../../../assets/css/custom.css">

# Hooks

A hook represents a set of actions that will be executed for one or more scenarios.
There are two types of hooks:

- **before hooks**: It's launched before a scenario starts.
- **after hooks**: It's launched after a scenario ends.

The hooks are represented with keyword `before` or `after` and followed by a label 
This label is a tag, and the hook will be executed for those scenarios with this tag.


**Example** [download](https://raw.githubusercontent.com/wesovilabs-tools/orion-examples/master/site/feature006.hcl)
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

There are two special hooks that will be executed before and after each scenario. In this case
the tag must be `each`.  

**Example** [download](https://raw.githubusercontent.com/wesovilabs-tools/orion-examples/master/site/feature007.hcl)
```hcl
before each {
  description = "common hook to be executed before each scenario"
  set a {
    value = 1
  }
}
after each {
  description = "common hook to be executed after each scenario"
  print {
    msg = "value of b is ${b}"
  }
}
scenario "scenario3" {
  tags = ["common"]
  when "do a sum" {
    set b {
      value = 2 * a
    }
  }
  then "verify the result"{
    assert {
      assertion= b==2
    }
  }
}
```

**Supported actions:**

- [set](../../actions/set)
- [print](../../actions/print)
- [assert](../../actions/assert)
- [http](../../actions/http)
- [mongo](../../actions/http)
- [block](../../actions/block)