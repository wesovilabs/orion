---
layout: default
title: http
parent: Actions
nav_order: 4
---
<link rel="stylesheet" href="../../../assets/css/custom.css">

# http 

work in progress
{: .label .label-yellow }

The action **http** is used to make http request. This action requires a label that must be a valid http verb.

```hcl
http <method>{}
// method must be a valid http verb
http post {}
http get {}
...
```
Supported http verbs are: `get`, `post`, `put`, `patch`, `delete`, `options` and `head`.

The arguments for the action are:

|                 | Type      | Required?| Vars supported? |
|:----------------|:----------|---------:|----------------:|
| **description** | string    | no       | no              |
| **while**       | boolean   | no       | yes             |
| **when**        | boolean   | no       | yes             |
| **count**       | numeric   | no       | yes             |
| **request**     | Request   | yes      | yes             |
| **response**    | Response  | yes      | yes             |


**description** ( string \| optional )  It is used to apply descriptive text to  the action.

*Example 1: Basic use of argument*

```hcl
http post{
  description = "it creates a new person"
  ...
}
```
---
**when** ( bool | optional ) It is used to control if the action must be executed.

```hcl
var {
  evalJobStatus = false
}

http get{
  ...
  when = evalJobStatus
}
```
---
**count** ( number || optional ) It determines the number of times the action is executed. Additionally, the variable **_.index** is increased in each iteration. 
The value of _.index starts with 0 and it ends with count-1.

*Example 1: Basic use of the argument*
```hcl
http list {
  ...
  count = 3
}
```

---
**while** ( boolean \| optional )  The action is executed repeatedly as long as the value of this argument is met. Additionally, the variable **_.index** is increased in each iteration. The value of _.index starts with 0 and increase in 1 in each iteration.

*Example 1: Basic use of argument while*
```hcl
http get {
  ...
  while = _.index<=2
}
```

## Request

|                   | Type       | Required?| Vars supported? |
|:------------------|:-----------|---------:|----------------:|
| **baseUrl**       | string     | yes      | yes             |
| **path**          | string     | yes      | yes             |
| **quueryParams**  | QueryParams| no       | yes             |
| **headers**       | Headers    | no       | yes             |
| **payload**       | Payload    | no       | yes             |
| **cookies**       | []Cookie   | no       | yes             |


**baseUrl**: It contains the base url to the full url 
```hcl
http post {
    ...
    request {
        baseUrl = "http://api.hostname.com"
    }
}
```
--- 
**path**: It contains the path of the url to be invoked. 
```hcl
http post {
    ...
    request {
        baseUrl = "http://api.hostname.com"
        path = "/v1/users"
    }
}
// POST http://api.hostname.com/v1/users
```
--- 
**QueryParams**:  

They're represneted by a block named `queryParams` that contains the list of params to be appended to the path.

```hcl
http get {
    ...
    request {
        baseUrl = "http://api.hostname.com"
        path = "/v1/users"
        queryParams {
            orderBy = ["firstname","age"]
            total = 5
        }
    }
}
// GET http://api.hostname.com/v1/users?prderBy=firstname,age&total=5
```
--- 
**Headers**
They're represneted by a block named `headers` that contains the list of headers to be sent. The value of each header can be a `string` or a `list of string` values.

```hcl
http get {
    ...
    request {
        baseUrl = "http://api.hostname.com"
        path = "/v1/users"
        headers {
            Authorization = "Bearer ${Token}"
            Content-Type = "application/json"
            Tags = ["service","custom-tag"]
        }
    }
}
```
--- 
**Payload**
It's represented with a block named `payload`. This block requires a label whose value must be one of these: `json`, `xml`, `form` or `raw`. The label means in which format the body will be serialized. 

Additional the payload content must be set to an argument named data.

|             | Type      | Required?| Vars supported? |
|:------------|:----------|---------:|----------------:|
| **data**    | string    | yes      | yes             |

```hcl
...
set johnDetails {
    value = {
        firstname = "John"
        lastname = "Doe"
        job = {
            company = "wesovilabs"
            role = "developer"
        }
    }
}
...
http post {
    ...
    request {
        baseUrl = "http://api.hostname.com"
        path = "/v1/users"
        payload json {
            data = johnDetails
        }
    }
}
```
--- 
**Connection**:

|             | Type      | Required?| Vars supported? |
|:------------|:----------|---------:|----------------:|
| **timeout** | duration  | yes      | yes             |
| **proxy**   | string    | yes      | yes             |


**timeout** It represents the established timeout for client. Its value must be represented in duration format. (1m | 1s )

```hcl
http get {
    ...
    request {
        baseUrl = "http://api.hostname.com"
        path = "/v1/users"
        connection {
            timeout = 15s
        } 
    }
}
```

**proxy** It value of proxy used to establish the connection. 

```hcl
http get {
    ...
    request {
        baseUrl = "http://api.hostname.com"
        path = "/v1/users"
        connection {
            proxy = "http://proxy.myhostname.com:9090"
        } 
    }
}
```

## response

This block is used to take data from the http response and put into variables which will be used in the scenario. 

There are three special variables that we can use in block response:

- **_.http.body**: Returned body from the server.
- **_.http.headers**: The response headers.
- **_.http.statusCode**: The status code.

```hcl
http get {
  request {
    baseUrl = ghBase
    path = "/orgs/${orgId}"
    headers {
      Accept = mediaType
    }
  }
  response {
    org = json(_.http.body)
    statusCode = _.http.statusCode
    server = _.http.headers["Server"]
  }
}
```
 

