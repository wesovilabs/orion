---
layout: default
title: Scenarios
parent: Feature
nav_order: 1
---

<link rel="stylesheet" href="../../../assets/css/custom.css">

# Scenarios

A scenario corresponds to the definition of an acceptance test.  

**Example** [download](https://raw.githubusercontent.com/wesovilabs/orion-examples/master/site/feature003.hcl)

```hcl
scenario "square root operation work as expected" {
  tags = ["maths"]
  ignore = false
  continueOnError = true
  given "a number higher than zero" {
    set number {
      value  = 5
    }
  }
  when "the number is multiplied  by itself" {
    set result {
      value = 5 * 5
    }
  }
  then "the square root of the result is input number"{
    assert {
      assertion = number==sqrt(result)
    }
  }
}
```

A scenario has three optional arguments:

- **tags**: A list of text values. These are used to identify if a hook needs to be executed.
- **ignore**: Boolean value, by default It is false. If we set true to this value, the scenario won't be executed.
- **continueOnError**: Boolean value, by default is false. If It is true other feature execution will continue even though this scenario fails.
- **examples**: It contains a list of dataset. At the bottom of this page we can find a more detailed explanation.


A scenario is composed by 3 different types of blocks:

- **given**: It defines variables used in our scenario.
- **when**:  It contains a set of actions to be executed.
- **then**: It validates the outcome.

To write valid scenarios must meet the following standards:
- The first section in a scenario must be `given` or `when`.
- The last section in a scenario must be `then`.
- Section `given` must be followed of section `when`.
- Section `when` must be followed of section `then`.

As you could realize the structure of a valid sceneario will be G-W-T, W-T, but also G-W-T-W-T ....  


**Example** [download](https://raw.githubusercontent.com/wesovilabs/orion-examples/master/site/feature004.hcl)
```hcl
scenario "do some math operations" {
  given "a number higher than zero" {
    set number {
      value  = 5
    }
  }
  when "the number is multiplied  by itself" {
    set result {
      value = 5 * 5
    }
  }
  then "the square root of the result is input number"{
    assert {
      assertion = number==sqrt(result)
    }
  }
  when "the number is added to the current value" {
    set result {
      value = result + 5
    }
  }
  then "the result is the expected"{
    assert {
      assertion = number==sqrt(result - number)
    }
  }
}
```

--- 
## given

It contains the definition of variables used by the scenario. 

```hcl 
given "the person details"{
   set person {
       value = {
           firstname = "John"
           lastname = "Doe"
       }
   }
}
```

**Supported actions:**

- [set](../../actions/set)

--- 
## when

Set of actions to be executed.

```hcl 
when "we fetch the list of universities in a given country"{
    http get {
      request {
        baseUrl = universitiesApi
        path = "/search"
        queryParams {
          country = country
        }
      }
      response {
        universities = json(_body)
        statusCode = _statusCode
      }
    }
    print {
      msg = "http status code is ${statusCode} and there's ${len(universities)} univeristies in ${country}"
    }
}
```

**Supported actions:**

- [set](../../actions/set)
- [print](../../actions/print)
- [assert](../../actions/assert)
- [http](../../actions/http)
- [mongo](../../actions/mongo)
- [block](../../actions/block)


--- 
## then

List of asserts to be performed. More than one block assert can be defined in section then.

```hcl 
then "the person is created successfully"{
   assert {
       assertion = person.firstname == "Tom" && person.age>=30
   }
}
```
**Supported actions:**

- [assert](../../actions/assert)


## Scenarios with examples

Sometimes we want to run our scenarios wih different set pf data. We can do it
just defining attribute `examples` and providing a collection of data.

In the above example,  the scenario will be executed four times. 

**Example** [download](https://raw.githubusercontent.com/wesovilabs/orion-examples/master/site/feature005.hcl)
```
scenario "operation add" {
    examples = [
        { a = 3,  b = 2,  c = 5},
        { a = 4,  b = 2,  c = 6},
        { a = 8,  b = 12, c = 20},
        { a = 13, b = 12, c = 25},
    ]
    when "values are added" {
        set result {
            value = a + b 
        }
    } 
    then "the result of the operation is the expected" {
        assert {
            assertion = result==expected
        }
    }
}
```
