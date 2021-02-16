[![Build Status](https://travis-ci.org/wesovilabs/orion.svg?branch=master)](https://travis-ci.org/wesovilabs/orion)
[![Go Report Card](https://goreportcard.com/badge/github.com/wesovilabs/orion)](https://goreportcard.com/report/github.com/wesovilabs/orion)
[![GoDoc](https://godoc.org/github.com/wesovilabs/orion?status.svg)](https://godoc.org/github.com/wesovilabs/orion)

 
# Orion

Orion is born to simplify the way we write our acceptance tests. 

```hcl
input {
  arg ghBase {
    default = "hattps://api.github.com"
  }
  arg expect {}
  arg orgId {}
}

before each {
  set mediaType {
    value = "application/vnd.github.v3+json"
  }
}

scenario "username details" {
  when "call Github api" {
    http get {
      request {
        baseUrl = ghBase
        path = "/users/${username}"
        headers {
          Accept = mediaType
        }
      }
      response {
        user = json(_.http.body)
      }
    }
    print {
      msg = "user ${username} has ${user.followers} followers."
    }
  }
  then "the user details are the expected" {
    assert {
      assertion = user!=null && user.followers>expect.user.followers
    }
  }
}
```

## About the project

Orion is born to change the way we implement our acceptance tests. It takes advantage of HCL from Hashicorp to provide a **simple DSL to write the acceptance tests**. The syntax is inspired in Gherkin.




    
