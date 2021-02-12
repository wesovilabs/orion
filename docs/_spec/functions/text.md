---
layout: default
title: Text functions
parent: Functions
nav_order: 1
---

<link rel="stylesheet" href="../../../assets/css/custom.css">

# Text functions

## Comparisons 

| Operation     | Description                                        | Signature  |
|:--------------|:---------------------------------------------------|-----------:|
| ==            | It returns true if `s1` and `s2` are equal (case sensitive) and false otherwise |   s1==s2 |
| !=            | It returns true if `s1` and `s2`  are not equal and false otherwise | s1!=s2 |
| eqIgnorecase  | Like operation `==` but case insensitive| eqIgnoreCase(s1,s2)| 

**Example** [download](https://raw.githubusercontent.com/wesovilabs-tools/orion-examples/master/site/feature009.hcl)
```hcl
scenario "text functions" {
  given "the an input sentence"{
    set input {
      value = "hello world"
    }
  }
  when "do text comparisons" {
    set res1 {
      value = input=="Hello World"
    }
    set res2 {
      value = input=="hello world"
    }
    set res3 {
      value = input!="Hello World"
    }
    set res4 {
      value = input!="hello world"
    }
    set res5 {
      value = eqIgnoreCase(input,"Hello World")
    }
    set res6 {
      value = eqIgnoreCase(input,"hello world")
    }
  }
  then "the results are the expected"{
    assert {
      assertion = !res1 && res2 && res3 && !res4 && res5 && res6
    }
  }
}
```

## Transformations

| Operation     | Description                                        | Signature  |
|:--------------|:---------------------------------------------------|-----------:|
| toUppercase   | It transforms `s` into uppercase | toUppercase(s)|
| toLowercase   | It transforms `s` into lowercase | toLowercase(s)|
| trimPrefix    | It removes the `prefix` from `s` if it exists | trimPrefix(s,prefix)|
| trimSuffix    | It removes the `suffix` from `s` if it exists | trimSuffix(s,suffix) |
| replaceOne    | It replaces `old` by `new` in `s`  once | replaceOne(s,old,new) |
| replaceAll    | It replaces `old` by `new` in `s`  as many time as `ld` appears  | replaceAll(s,old,new) |


**Example** [download](https://raw.githubusercontent.com/wesovilabs-tools/orion-examples/master/site/feature010.hcl)
```hcl
input {
  arg s {
    default ="Hello world"
  }
}
scenario "check string converter oeprations" {
  when "convert the input" {
    set res1 {
      value = toLowercase(s)
    }
    set res2 {
      value = toUppercase(s)
    }
    set res3 {
      value = trimPrefix(s,"Hello ")
    }
    set res4 {
      value = trimSuffix(s," world")
    }
    set res5 {
      value = replaceOne(s,"Hello", "Bye")
    }
    set res6 {
      value = replaceAll(s,"o", "a")
    }
  }
  then "check results"{
    assert {
      assertion = (
        res1=="hello world" &&  res2=="HELLO WORLD" &&
        res3 == "world" &&  res4 == "Hello" &&
        res5 == "Bye world" &&  res6 == "Hella warld"
      )
    }
  }
}
```

## Search

| Operation     | Description                                        | Signature  |
|:--------------|:---------------------------------------------------|-----------:|
| contains   | It returns true if  `s` contains `text`; fals in other case lowercase | contains(s,text)|
| hasPrefix   | It tests whether the string `s` begins with `prefix` | hasPrefix(s,prefix)|
| hasSuffix   | It tests whether the string `s` ends with `suffix` | hasSuffix(s,suffix)|
| count | It counts the number of instances of `substr` in `s` | count(s,substr) |
| indexOf   | It returns the index of the first instance of `substr` in `s`, or -1 if `substr` is not present in `s` | indexOf(s,substr)|
| lastIndexOf   | It returns the index of the last instance of `substr` in `s`, or -1 if `substr` is not present in `s`. | lastIndexOf(s,substr)|# 



**Example** [download](https://raw.githubusercontent.com/wesovilabs-tools/orion-examples/master/site/feature011.hcl)
```hcl
input {
  arg s {
      default ="Hello world"
  }
}
scenario "check string search oeprations" {
  when "do the operations" {
      set res1 {
          value = contains(s,"world")
      }
      set res2 {
          value = hasPrefix(s,"Hello")
      }
      set res3 {
          value = hasSuffix(s,"world")
      }
      set res4 {
          value = count(s,"l")
      }
      set res5 {
          value = indexOf(s,"l")
      }
      set res6 {
          value = lastIndexOf(s,"l")
      }
  }
  then "the result is the exapected"{
      assert {
          assertion = res1 && res2 && res3 && res4==3 && res5==2 && res6==9
      }
  }
}
```

## Others

| Operation     | Description                                        | Signature  |
|:--------------|:---------------------------------------------------|-----------:|
| split    | It slices `s` into all substrings separated by `sep` and returns a slice of the substrings between those separators.  | split(s,sep) |
| len   | It returns the number of characters in `s` | len(s)|


**Example** [download](https://raw.githubusercontent.com/wesovilabs-tools/orion-examples/master/site/feature012.hcl)
```hcl
input {
  arg s {
      default ="tofu,tempeh,kale,edamame"
  }
}
scenario "test functions with strings" {
  when "do the operations" {
      set elements {
          value = split(s,",")
      }
      print {
          msg = elements[_.index]
          while = _.index < len(elements)
      }
  }
  then "check results"{
      assert {
          assertion = (
              len(elements)==4 && 
              elements[0]=="tofu" && elements[1]=="tempeh" &&
              elements[2]=="kale" && elements[3]=="edamame"   
          )
      }
  }
}
```

* To concat strings we just need to do this: `"${value},"${value2}"`
