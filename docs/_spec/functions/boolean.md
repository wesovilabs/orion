---
layout: default
title: Boolean functions
parent: Functions
nav_order: 2
---
<link rel="stylesheet" href="../../../assets/css/custom.css">

# Boolean functions

## Comparisons

| Operation     | Description                                 | Signature  |
|:--------------|:--------------------------------------------|-----------:|
| ==            | It returns true if `val1`is equal to `val2` |   val1 == val2 |
| !=            | It returns true if `val1`is equal to `val2` and false otherwise|   val1 != val2 |
| \|\|  | or comparator|val1 \|\| val2 | 
| &&  | and comparator|val1 && val2 | 
| !  | not | !val1 | 
| ()  | group conditions | (val1 && val2) | 


**Example** [download](https://raw.githubusercontent.com/wesovilabs-tools/orion-examples/master/site/feature013.hcl)
```hcl
scenario "check boolean operations" {
  given "set value" {
    set val1{
         value = true
     }
  }
  when "do the operations" {
     set val2{
         value = !val1
     }
  }
  then "the result is the exapected"{
      assert {
          assertion = val1 && !val2 && (val1 || val2) && !val2 && val2==false 
      }
  }
}
```