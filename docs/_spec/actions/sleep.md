---
layout: default
title: sleep
parent: Actions
nav_order: 7
---
<link rel="stylesheet" href="../../../assets/css/custom.css">
# sleep

The action **sleep** is used to pause the execution for a specified duration.

## Specification

### Arguments 

|                 | Type      | Required?| Vars supported? |
|:----------------|:----------|---------:|----------------:|
| **duration**    | duration  | yes      | yes             |
| **description** | string    | no       | no              |
| **while**       | boolean   | no       | yes             |
| **when**        | boolean   | no       | yes             |
| **count**       | numeric   | no       | yes             |


**duration** ( duration \| required ) : It specified the `time.duration` to pause the execution.

*Example 1: Basic use of argument duration*

```hcl
sleep {
  duration = "2s"
}
```
---
**description** ( string \| optional )  It is used to apply descriptive text to  the action.

*Example 1: Basic use of argument*

```hcl
sleep {
  description = "this statement is used to pause the execution for 2 seconds"
  duration = "2s"
}
```
---
**when** ( bool | optional ) It is used to control if the action must be executed.

```hcl
var {
  evalJobStatus = false
}

sleep {
  duration = "2s"
  when = evalJobStatus && exists(job)
}
```
---
**count** ( number || optional ) It determines the number of times the action is executed.

*Example 1: Basic use of the argument count*
```hcl
sleep {
  duration = "2s"
  count = 3
}
```
---
**while** ( boolean \| optional )  The action is executed repeatedly as long as the value of this argument is met. Additionally, the variable **_.index** is increased in each iteration. The value of _.index starts with 0 and increase in 1 in each iteration.

*Example 1: Basic use of argument while*
```hcl
assert {
  duration = "2s"
  while = _.index<=2
}
```
