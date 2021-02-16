---
# Feel free to add content and custom Front Matter to this file.
# To modify the layout, see https://jekyllrb.com/docs/themes/#overriding-theme-defaults

layout: home
---
<link rel="stylesheet" href="assets/css/custom.css">

# Hello, I'm Orion

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

## Goals

- **Non-technical** people can design but also **implement the acceptance tests**.
- It's implemented under a **pluggable architecture**. New actions can be implemented easily.
- **Reusable** functionality can be shared between **features**.

## License

Software is completely free, and It will be distributed as opensource soon.

## Contributing

Orion will be opensource soon. In the meanwhile any feedback, suggestion is welcome! [Contact me](/contact/index/)
