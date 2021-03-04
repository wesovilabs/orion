---
layout: default
title: call
parent: Actions
nav_order: 6
---

<link rel="stylesheet" href="../../../assets/css/custom.css">

# call

The action **call** is used to invoke defined functions.

## Specification

### Arguments 

|                 | Type      | Required?| Vars supported? |
|:----------------|:----------|---------:|----------------:|
| **as**          | string    | no       | no              |
| **description** | string    | no       | no              |
| **while**       | boolean   | no       | yes             |
| **when**        | boolean   | no       | yes             |
| **count**       | numeric   | no       | yes             |

**as** ( string \| optional )  The name of the variable used to persist the output of the invoked function.

*Example 1: Basic use of argument*

```hcl
call listPeople{
  as = people
}
```
---
**description** ( string \| optional )  It is used to apply descriptive text to  the action.

*Example 1: Basic use of argument*

```hcl
call createPerson{
  description = "it invokes function createPerson"
}
```
---
**when** ( bool | optional ) It is used to control if the action must be executed.

```hcl
var {
  mustCreatePerson = false
}

call createPerson{
  when = mustCreatePerson
}
```
---
**count** ( number || optional ) It determines the number of times the action is executed. Additionally, the variable **_.index** is increased in each iteration. 
The value of _.index starts with 0 and it ends with count-1.

*Example 1: Basic use of the argument*
```hcl
call createPerson{
  with {
    id = "user_${_.index}"
  }
  count = 3
}
// After this statement the value of variable counter will be 2
```

---
**while** ( boolean \| optional )  The action is executed repeatedly as long as the value of this argument is met. Additionally, the variable **_.index** is increased in each iteration. The value of _.index starts with 0 and increase in 1 in each iteration.

*Example 1: Basic use of argument while*
```hcl
call createPerson{
  with {
    id = "user_${_.index}"
  }
  while = _.index<=2
}
```

### Blocks

A function could require arguments. The arguments are provided within a block `with`. 

```hcl
call createPerson {
    with {
        firstname = user.firstname
        role = "ADMIN"
    }
}
```
