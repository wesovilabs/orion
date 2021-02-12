---
layout: default
title: Feature
nav_order: 0
has_children: true
has_toc: false
---
<link rel="stylesheet" href="../../../assets/css/custom.css">

# Feature

A feature is described in a file and it contains one or more scenarios. The extension of the file doesn't matter at all. On the other hand,  we recommend you to use extension `hcl` so far. 

**Example** [download](https://raw.githubusercontent.com/wesovilabs/orion-examples/master/site/feature001.hcl)
```hcl
description = <<EOF
    This feature is used to demonstrate that both add and subs operations 
    work as expected.
EOF

after each {
    print {
        msg = "the output of this operation is ${result}"
    }
}

scenario "operation add" {
    given "the input of the operation" {
        set a {
            value = 10
        }    
        set b {
            value = 5
        }
    }
    when "values are added" {
        set result {
            value = a + b 
        }
    } 
    then "the result of the operation is the expected" {
        assert {
            assertion = result==15
        }
    }
}
```

We can make use of the attribute `description` to describe what the feature does. 
In a feature we can find the below sections::

- [Input variables](../variables)
- [Scenarios](../scenarios)
- [Hooks](../hooks)
- [Includes](../includes)

