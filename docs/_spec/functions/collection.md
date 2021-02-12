---
layout: default
title: Collection functions
parent: Functions
nav_order: 4
---
<link rel="stylesheet" href="../../../assets/css/custom.css">

# Collection functions

## Items

| Operation     | Description                                 | Signature  |
|:--------------|:--------------------------------------------|-----------:|
| at            | It returns element in position `pos`in the collection `coll`|   at(coll,post)|
| first  | It returns the first element in the collection `coll` | first(coll) | 
| last  | It returns the last element in the collection `coll` | last(coll) | 
| len  | It returns the numer of elements in the collection `coll` | len(coll) | 


**Example** [download](https://raw.githubusercontent.com/wesovilabs-tools/orion-examples/master/site/feature016.hcl)
```hcl
scenario "check collection functions" {
  given "set value" {
    set elements{
      value = [
        {
          name = "Sally"
          specie = "dog"
        },
        {
          name = "Molly"
          specie = "dog"
        },
        {
          name = "Coco"
          specie = "dog"
        }
      ]
    }
  }
  when "obtain the element at position 1" {
    set element {
      value = elements[1]
    }

  }
  then "check the element"{
    assert {
      assertion = element.name == "Molly"
    }
  }
  when "obtain the first element" {
    set element {
      value = first(elements)
    }

  }
  then "check the element"{
    assert {
      assertion = element.name == "Sally"
    }
  }
  when "obtain the last element" {
    set element {
      value = last(elements)
    }

  }
  then "check the element"{
    assert {
      assertion = element.name == "Coco"
    }
  }
  when "obtain the number of elements in the list" {
    set totalElement {
      value = len(elements)
    }

  }
  then "check the element"{
    assert {
      assertion = totalElement==3
    }
  }
}
```