---
layout: default
title: set
parent: Actions
nav_order: 1
---

<link rel="stylesheet" href="../../../assets/css/custom.css">

# set

The action **set** is used to add a new (or update) value for a varibale.

## Specification

### Arguments 

|                 | Tpye      | Required?| Vars supported? |
|:----------------|:----------|---------:|----------------:|
| **value**       | string    | yes      | yes             |
| **arrayIndex**  | numeric   | no       | yes             |
| **description** | string    | no       | no              |
| **while**       | boolean   | no       | yes             |
| **when**        | boolean   | no       | yes             |
| **count**       | numeric   | no       | yes             |


**value** ( any \| required ) : Expression to be evaluated. 

*Example 1: Basic use of argument setion*

```hcl
set firstname{
  value = Jhon
}
// value of variable firstname is John
```

*Example 2: Complex variables*

```hcl
set countryName {
    value = "spain"
}
set john{
  value = {
    firstname = John
    lastname = Doe
    country = toUppercase(countryName)
  }
}
// john.firstanem is John, john.lastname is Doe and country is SPAIN
```
---

**arrayIndex** ( numeric \| optional ) : If variable is an array, we can modify it partially. 

*Example 1: Basic use of argument setion*

```hcl
set heroes {
  value = [{firstname=Laika},{firstname=Valto}]
}

set heroes{
  arrayIndex=1
  value = {
     firstname = Balto
  }
}
// value of heroes is [{firstname=Laika},{firstname=Balto}]
```
---

**description** ( string \| optional )  It is used to apply descriptive text to  the action.

*Example 1: Basic use of argument*

```hcl
set firstname{
  description = "it firstname of the person to be created"
  value = John
}
```
---
**when** ( bool | optional ) It is used to control if the action must be executed.

```hcl
var {
  evalJobStatus = false
}

set firstname{
  value = John
  when = evalJobStatus
}
```
---
**count** ( number || optional ) It determines the number of times the action is executed. Additionally, the variable **_.index** is increased in each iteration. 
The value of _.index starts with 0 and it ends with count-1.

*Example 1: Basic use of the argument*
```hcl
set counter {
  value = _.index
  count = 3
}
// After this statement the value of variable counter will be 2
```

---
**while** ( boolean \| optional )  The action is executed repeatedly as long as the value of this argumentis met. Additionally, the variable **_.index** is increased in each iteration. The value of _.index starts with 0 and increase in 1 in each iteration.

*Example 1: Basic use of argument while*
```hcl
set counter {
  value = _.index
  while = _.index<=2
}
```