---
layout: default
title: Number functions
parent: Functions
nav_order: 3
---

<link rel="stylesheet" href="../../../assets/css/custom.css">

# Number functions

## Operations

| Operation     | Description                                 | Signature  |
|:--------------|:--------------------------------------------|-----------:|
| +      | It adds  `x` to `y`| x+y|
| -      | It subtracts `y` to `x`| x-y |
| *      | It multiplies `x` by `y`| x*y |
| /      | It divides `x` by `y`| x/y 
| sqrt          | It returns the square root of `x`|   sqrt(x) |
| cos          | It returns the cosine of the radian argument `x`|   cos(x) |
| sin          | It returns the sine of the radian argument `x`|   sin(x) |
| round          | It returns the nearest integer to `x`, rounding half away from zero|   round(x) |
| pow          | It returns the base-`x` exponential of `y`| pow(x,y) |
| mod          | It returns the floating-point remainder of `x`/`y`| mod(x,y) |
| max          | It returns the larger of `x` or `y`| max(x,y) |
| min          | It returns the shorter of `x` or `y`| min(x,y) |



**Example** [download](https://raw.githubusercontent.com/wesovilabs/orion-examples/master/site/feature014.hcl)
```hcl
scenario "number operation with int values" {
  given "int values" {
    set val {
      value = 100
    }
  }
  when "do operations with numbers" {
    set valSqrt {
      value = sqrt(val)
    }
    set valCos {
      value = cos(val)
    }
    set valSin {
      value = sin(val)
    }
    set valRound {
      value = round(val)
    }
    set valPow {
      value = pow(val, 2)
    }
    set valMod {
      value = mod(val, 2)
    }
    set valMax {
      value = max(val, 2*val)
    }
    set valMin {
      value = min(val, 2*val)
    }
    set valOp {
      value = ((5 - 1)*(val+2))/3
    }
    print {
      msg = valSin
    }
  }
  then "result is the exapected" {
    assert {
      assertion = (
        valSqrt==10 && valMax==200 && valOp==136 &&  val==100 &&
        valCos==0.8623188723 && valSin==-0.5063656411 && valRound==100 &&
        valMod==0 && valMin==100
      )
    }
  }
}

```

## Comparisons

| Operation     | Description                                 | Signature  |
|:--------------|:--------------------------------------------|-----------:|
| >  | It returns true if `x` is higher than `y` and false otherwise| x>y|
| >= | It returns true if `x` is higher than or equal to  `y` and false otherwise| x>=y|
| <  | It returns true if `x` is lower than `y` and false otherwise| x<y|
| <>= | It returns true if `x` is lower than or equal to  `y` and false otherwise| x<>=y|

**Example** [download](https://raw.githubusercontent.com/wesovilabs/orion-examples/master/site/feature015.hcl)
```hcl
scenario "number comparisons" {
  given "input values" {
      set val1 {
          value = 100
      }
      set val2 {
          value = 30.213
      }
  }
  when "do operations with numbers" {
     print {
         msg="Input values are val1:${val1} and val2:${val2}"
     }
  }
  then "result is the exapected"{
      assert {
          assertion= val1>val2 && val1>=val2 && val2<val1 && val2<=val1
      }
  }
}
```
