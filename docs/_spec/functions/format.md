---
layout: default
title: Formatter functions
parent: Functions
nav_order: 45
---

<link rel="stylesheet" href="../../../assets/css/custom.css">

# Formatter functions

## Converter

| Operation     | Description                                 | Signature  |
|:--------------|:--------------------------------------------|-----------:|
| json          | It unmarshals the input `text` as a json|   json(text)|


**formatter.hcl**
```hcl
scenario "test formatter functions" {
    given "string values" {
        set valueJSON {
            value = "{\"firstname\":\"John\"}"
        }
    }
    when "convert json into map" {
        set person {
            value = json(valueJSON)
        }
    }
    then "the json has been formatted successfully" {
        assert {
            assertion = person.firstname == "John"
        }
    }
}
```