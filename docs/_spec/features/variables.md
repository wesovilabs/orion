---
layout: default
title: Input variables
nav_order: 1
parent: Feature
---

<link rel="stylesheet" href="../../../assets/css/custom.css">

# Input Variables

It's used to run acceptance tests with data provided by the user. Variables are defined inside a block `input`. 
Only one block `input` is permitted per file. 

**Example** [download](https://raw.githubusercontent.com/wesovilabs/orion-examples/master/site/feature002.hcl)

```hcl
input {
  arg people {
    default = [
      { firstname = "John", lastname = "Doe"},
      { firstname = "Jane", lastname = "Doe"},
    ]
  }
  arg company {
    default = "Wesovilabs"
  }
}
scenario "print variables" {
  when "iterate over the people"{
    block{
      set person {
        value = people[_.index]
      }
      print {
        msg = "${person.firstname} ${person.lastname} is hired at ${company}"
      }
      while = _.index < len(people)
    }
  }
  then "verify the postconditions" {
    assert {
      assertion = eqIgnoreCase(company,"wesovilabs")
    }
  }
}
```

Providing variables is not rquired since both, variables `people` and `company` have a default value

So far, variables can be provided by passing a HCL file. This file contains one entre per variable.

**Example**: [download](https://raw.githubusercontent.com/wesovilabs/orion-examples/master/site/variables002.hcl)

```hcl
people = [
    { firstname = "Tim", lastname = "Doe" },
    { firstname = "Loe", lastname = "Roe" }
]
company = "wesoviLabs"
```

*In upcoming releases, variables could be passed by a flag or being taken by the environment.*

Variables are passed using argument `--vars`.

```bash  
>> orion-cli run --input feature.hcl --vars variables.hcl
# Tim Doe is hired at wesoviLabs
# Loe Roe is hired at wesoviLabs
```

---

## input

It acts like a container that groups a set of blocks `arg`. As It was mentioned on the above, there will be one block input per file at maximum.


### arg

It is used to define an input variable. It's declared as `arg <var>`, where var is the name of the variable. 


**Example**
```hcl
input {
  arg person{}
  arg firstname{
    default = "Doe" 
  }
  arg cars{
    description = "list of cars"
  }
}
```

Two optional attributes can be provided in the block `arg`. 

- **description**: A string value that provides a brief description of the variable.
- **default**: Value that will be used in case of the variable is not provided.

```hcl
arg firstname {
    description = "firstname of the person to be searched for in the scenarios" 
    default = "John"
}
```