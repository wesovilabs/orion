---
layout: default
title: assert
parent: Actions
nav_order: 3
---
<link rel="stylesheet" href="../../../assets/css/custom.css">
# assert

The action **assert** is used to verify the condition is satisfied. If condition
was not satisfied the scenario ends with an error and next actions are not executed.

## Specification

### Arguments 

|                 | Tpye      | Required?| Vars supported? |
|:----------------|:----------|---------:|----------------:|
| **assertion**   | boolean   | yes      | yes             |
| **description** | string    | no       | no              |
| **while**       | boolean   | no       | yes             |
| **when**        | boolean   | no       | yes             |
| **count**       | numeric   | no       | yes             |


**assertion** ( boolean \| required ) : Ic contains the condition to be evaluated.

*Example 1: Basic use of argument assertion*

```hcl
assert {
  assertion = eq(firstname,"John") && person.age<=20
}
```
---
**description** ( string \| optional )  It is used to apply descriptive text to  the action.

*Example 1: Basic use of argument*

```hcl
assert {
  description = "this statement is used to verify the job status is the expected"
  assertion = eqIgnoreCase(job.build,"success")
}
```
---
**when** ( bool | optional ) It is used to control if the action must be executed.

```hcl
var {
  evalJobStatus = false
}

assert {
  assertion = eqIgnoreCase(job.build,"success")
  when = evalJobStatus && exists(job)
}
```
---
**count** ( number || optional ) It determines the number of times the action is executed. Additionally, the variable **_.index** is increased in each iteration. 
The value of _.index starts with 0 and it ends with count-1.

*Example 1: Basic use of the argument count*
```hcl
assert {
  assertion = eqIgnoreCase(job.build,"success") && _.index<5
  count = 3
}
```
---
**while** ( boolean \| optional )  The action is executed repeatedly as long as the value of this argument is met. Additionally, the variable **_.index** is increased in each iteration. The value of _.index starts with 0 and increase in 1 in each iteration.

*Example 1: Basic use of argument while*
```hcl
assert {
  assertion = eqIgnoreCase(job.build,"success") && _.index<5
  while = _.index<=2
}
```
