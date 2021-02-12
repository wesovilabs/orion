input {
  arg ghBase {
    default = "https://api.github.com"
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

scenario "list Github organizations"{
  given "establish expectations" {
    set expectation {
      value = expect.org
    }
  }
  when "obtain the list of organizations from Github"{
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
    print {
      msg = "Server is ${server}"
    }
  }
  then "the list of organizations is the expected" {
    assert {
      description = "It verifies the status code is the expected"
      assertion = statusCode == 200 && eqIgnoreCase(server,"github.com")
    }
    assert {
      assertion = (
        org!=null && org.public_repos==expectation.pubRepos &&
        org.email==expectation.email
      )
    }
  }
}
