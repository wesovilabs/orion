---
layout: default
title: mongo
parent: Actions
nav_order: 5
---
<link rel="stylesheet" href="../../../assets/css/custom.css">
# mongo 

work in progress
{: .label .label-yellow }

It's used to deal with mongo operations. It's represented with keyword `mongo` followed by a label.

**example.hcl**
```hcl
mongo <operation>{}
```
 The label can be any of the below values:

- **findOne**: It finds  a document in the collection.
- **find**: It finds a list of documents in the collection.
- **insertOne**: It inserts a single document into the collection.
- **insertMany**: It inserts a list od documents into the collection.
- **updateOne**: It set attributes in a single document.
- **updateMany**: It set attributes in many documents.
- **deleteOne**: It deletes one document from the collection.
- **deleteMany**: It deletes many documents from the collection.
- **create**: It creates a collection into the database.
- **drop**: It drops the collection from the database.
- **count**: It returns the number of found documents.

The arguments for the action are:

|                 | Type      | Required?| Vars supported? |
|:----------------|:----------|---------:|----------------:|
| **description** | string    | no       | no              |
| **while**       | boolean   | no       | yes             |
| **when**        | boolean   | no       | yes             |
| **count**       | numeric   | no       | yes             |
| **connection**  | Connection| yes      | yes             |
| **query**       | Query     | yes      | yes             |
| **response**    | Response  | yes      | yes             |


**description** ( string \| optional )  It is used to apply descriptive text to  the action.

*Example 1: Basic use of argument*

```hcl
mongo find{
  description = "it find all the users with the below role"
  ...
}
```
---
**when** ( bool | optional ) It is used to control if the action must be executed.

```hcl
var {
  userDoesNotExist = false
}

mongo insertOne{
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

## Connection

|               | Type       | Required?| Vars supported? |
|:--------------|:-----------|---------:|----------------:|
| **uri**       | string     | yes      | yes             |
| **timeout**   | timeout     | no       | yes             |
| **auth**      | Auth       | no       | yes             |



**uri**: The connection uri. 
```hcl
 mongo drop {
    ...
    connection {
        uri = "mongodb://localhost:27017"
        ...
    }
}
```
--- 
**timeout** It is used to establish the client timeout. Its value must be represented in duration format. (1m | 1s).

```hcl
 mongo drop {
    ...
    connection {
        timeout = "10s"
        ...
    }
}
```
--- 
**Auth**:  

It's a block with keyword `auth` followed by one of these labels: `scram-sha-1`, `scram-sha-256`, `mongodb-cr`, `plain`, `gssapi`, `mongodb-x509` or `mongodb-aws`

We can provide both `username` and `password` arguments inside this block.

```hcl
 mongo drop {
    ...
    connection {
        ...
        auth scram-sha-1{
            username = "admin"
            password = "secret"
        }
    }
}
```

*This block will be enriched in upcoming releases*


## Query

It's a block with name `query`. It contains details to build the query.

|               | Type      | Required?| Vars supported? |
|:--------------|:----------|---------:|----------------:|
| **database**  | string    | yes      | yes             |
| **collection**| string    | no       | yes             |
| **limit**     | numeric   | no       | yes             |


**database**: Name of the database to be connected to. 

**collection**: Name of the collection used to run the queries.

**limit**: Optional attribute used to establish the maximum number of returned documents.


```hcl
 mongo find {
    ...   
    query {
        database = "test"
        collection = "users"
        limit = 5
    }
}
```

## Response

This block is used to take data from the mongo database. So far, only operations `findOne`and `find` return data.

Actually depending on the operation the variables will be


- **_.mongo.document**: Found document (findOne)
- **_.mongo.documents**: Found documents (find)

```hcl
mongo findOne {
    response {
        user = _.mongo.document
    }
}

mongo find {
    response {
        users = _.mongo.documents
    }
}
```


