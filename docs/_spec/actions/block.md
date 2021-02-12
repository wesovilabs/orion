---
layout: default
title: block
parent: Actions
nav_order: 7
---
<link rel="stylesheet" href="../../../assets/css/custom.css">
# block
It  is used to group a set of actions.

```hcl
block {
    set carOwner {
        value = "John Doe"
    }
    print {
        msg = "the owner of the car is ${carOwner}"
    }
    when = car.color=="red"
}
```

**Supported actions:**
- [set](../set)
- [print](../print)
- [assert](../assert)
- [http](../http)
- [mongo](../mongo)

## Specification

### Arguments 

|                 | Tpye      | Required?| Vars supported? |
|:----------------|:----------|---------:|----------------:|
| **description** | string    | no       | no              |
| **while**       | boolean   | no       | yes             |
| **when**        | boolean   | no       | yes             |
| **count**       | numeric   | no       | yes             |


**description** ( string \| optional )  It is used to apply descriptive text to  the action.

*Example 1: Basic use of argument*

```hcl
block{
  description = "a new person is created and registered."
}
```
---
**when** ( bool | optional ) It is used to control if the action must be executed.

```hcl
var {
  evalJobStatus = false
}

block{
  when = evalJobStatus
}
```
---
**count** ( number || optional ) It determines the number of times the action is executed. Additionally, the variable **_.index** is increased in each iteration. 
The value of _.index starts with 0 and it ends with count-1.

*Example 1: Basic use of the argument*
```hcl
block {

  count = 3
}
```

---
**while** ( boolean \| optional )  The action is executed repeatedly as long as the value of this argumentis met. Additionally, the variable **_.index** is increased in each iteration. The value of _.index starts with 0 and increase in 1 in each iteration.

*Example 1: Basic use of argument while*
```hcl
block {
  while = _.index<=2
}
```