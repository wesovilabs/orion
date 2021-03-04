---
layout: default
title: print
parent: Actions
nav_order: 2
---
<link rel="stylesheet" href="../../../assets/css/custom.css">
# print

The action **print** prints messages to the stdout.

## Specification

### Arguments 

|                 | Type      | Required?| Vars supported? |
|:----------------|:----------|---------:|----------------:|
| **msg**         | string    | yes      | yes             |
| **prefix**      | string    | no       | yes             |
| **showTimestamp** | boolean | no       | yes             |
| **timestampFormat** | string| no       | yes             |
| **format**      | string| no       | yes             |
| **description** | string    | no       | no              |
| **while**       | boolean   | no       | yes             |
| **when**        | boolean   | no       | yes             |
| **count**       | numeric   | no       | yes             |

**msg** ( string \| required ): It contains the text to be printed out. 

*Example 1: Basic use of argument*

```hcl
print { 
  msg:"Hello Mr Doe"
}
```

```bash
>> Hello Mr Doe
```

*Example 2: Argument msg support variables*

```hcl
vars {
  lastname = "Doe"
}

print { 
  msg:"Hello Mr ${lastname}"
}
print { 
  msg:"Hello Mr ${toUppercase(lastname)}"
}
```
```bash
>> Hello Mr Doe
>> Hello Mr DOE
```
--- 
**prefix** ( string \| optional ): Text which is placed before the message. In case of being provided a whitespace is printed between `prefix` and `msg`.

*Example 1: Basic use of argument prefix* 

```hcl
print {
  prefix = "[DEBUG]"
  msg = "There are 20 cars."
}
```
```bash
>> [DEBUG] There are 20 cars.
```

*Example 2: Argument prefix supports variables*


```hcl
vars {
  logLevel = DEBUG
}

print {
  prefix = substring(logLevel,0,2)
  msg = "There are ${totalCars} cars."
}
```
```bash
>> DEB There are 20 cars.
```
--- 
**showTimestamp** ( boolean \| optional ): If this argument is true, a timestamps is placed before the msg (and the prefix if it's provided). Default timestamp format is `YYYY-MM-DD HH:mm:ss` but it can be changed with argument `timestampFormat`.

*Example 1: Basic use of argument* 

```hcl
print {
  msg = "The pub is closed."
  showTimestamp = true
}
```
```bash
>> 2013-02-23 12::05:12 There pub is closed 
```
--- 
**timestampFormat** ( string \| optional ):  It allows you to determine the format of the timestamp. Argument showTimestamp can be omitted when timestampFormat is provided.

| Symbol|      Desc     |
|-------|:-------------:|
| yyyy  | Year (2020-XXXX) | 
| yy    | Year (00-99)   | 
| MM    | Month of the year (1-12)   | 
| dd    | Day of the month (1-31)  | 
| HH    | Hour (0-23)   | 
| hh    | Hour (1-12)   | 
| mm    | Minutes (0-59)   | 
| ss    | Seconds (0-59)   | 


*Example 1: Basic use of argument timestampFormat* 

```hcl
print {
  msg = "Hello my friend"
  timestampFormat = "HH:mm:ss"
}
``` 
```bash
>> 20:10:02 Hello my friend
```
---
**format** ( string \| optional ) It specifies the format in which the message is  printed out.  Supported values for this argument are: **plain** or **json**. If
argument is not provided default value is plain.

*Example 1: Basic use of argument format* 

```hcl
print {
  msg = "Hello Mr Doe"  
  format = "json"
}
```

```bash
>> {"msg": "Hello Mr Doe"}
```
*Example 2: Usage of argument when timestamp & prefix* are provided too. 

```hcl
print {
  msg           = "Hello Mr Doe"
  showTimestamp = true  
  prefix        = "[DEBUG]"  
  format        = "json"
}
```

```bash
>> {"msg": "[DEBUG] Hello Mr Doe", "timestamp"="23/12/2020 12:22:23"}
```


--- 
**description** ( string \| optional )  It is used to apply descriptive text to  the action.

*Example 1: Basic use of argument* 

```hcl
print {
  desc = "this action is used to print out the user details"
  msg = "Hello Mr Doe"  
}
```
--- 
**when** ( bool | optional ) It is used to control if the action must be executed.

```hcl
var {
  open = false
}

print {
  msg = "The airport is closed"
  when = !open
}
```  
---
**count** ( number || optional ) It determines the number of times the action is executed. Additionally, the variable **_.index** is increased in each iteration. 
The value of _.index starts with 0 and it ends with count-1.

*Example 1: Basic use of the argument *
```hcl
print {
  msg = "Hello my friend"
  count = 3
}
```
```bash
>> Hello my friend
>> Hello my friend
>> Hello my friend
``` 

*Example 2: Use of variable _.index*

```hcl
print {
  msg = "${_.index} Hello my friend"
  count = 3
}
```
```bash
>> 0 Hello my friend
>> 1 Hello my friend
>> 2 Hello my friend
```
--- 
**while** ( boolean \| optional )  The action is executed repeatedly as long as the value of this argument is met. Additionally, the variable **_.index** is increased in each iteration. The value of _.index starts with 0 and increase in 1 in each iteration.
 
*Example 1: Basic use of argument while*
```hcl
print {
  msg = "Hello my friend"
  while = _.index<=2
}
```
```bash
>> Hello my friend
>> Hello my friend
>> Hello my friend
``` 

*Example 2: Use of variable _.index*

```hcl
print {
  msg = "${_.index} Hello my friend"
  count = 3
  while = _.index<=4
}

print {
  msg = "${_.index} Bye my friend"
  count = 3
  while = _.index<=1
}
```
```bash
>> 0 Hello my friend
>> 1 Hello my friend
>> 2 Hello my friend
>> 3 Hello my friend
>> 0 Bye my friend
>> 1 Bye my friend
```
