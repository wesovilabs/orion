before each {
  mongo drop {
    description = "drop collection"
    connection {
      uri = "mongodb://localhost:27017"
      auth scram-sha-1{
        username = "root"
        password = "secret"
      }
    }
    query {
      database = "test"
      collection = "users"
    }
  }

  mongo create {
    connection {
      uri = "mongodb://localhost:27017"
      auth scram-sha-1{
        username = "root"
        password = "secret"
      }
    }
    query {
      database = "test"
      collection = "users"
    }
  }
}

scenario "test mongo connection" {
  when "check connection" {



    mongo insertOne {
      connection {
        uri = "mongodb://localhost:27017"
        auth scram-sha-1{
          username = "root"
          password = "secret"
        }
      }
      query {
        database = "test"
        collection = "users"
        document {
          firstname = "John"
          lastname = "Doe"
          address = {
            city = "Mostoles"
            zip = "28938"
          }
          role = "admin"
        }
      }
    }

    mongo findOne {
      connection {
        uri = "mongodb://localhost:27017"
        auth scram-sha-1{
          username = "root"
          password = "secret"
        }
      }
      query {
        database = "test"
        collection = "users"
      }
      response {
        personId = _.mongo.document._id
        fullname = "${_.mongo.document.firstname} ${_.mongo.document.lastname}"
        role = _.mongo.document.role
      }
    }
    print {
      msg= "${fullname} owns role ${role}"
    }

    mongo find {
      connection {
        uri = "mongodb://localhost:27017"
        auth scram-sha-1{
          username = "root"
          password = "secret"
        }
      }
      query {
        database = "test"
        collection = "users"
      }
      response {
        users = _.mongo.documents
      }
    }



    mongo updateOne {
      connection {
        uri = "mongodb://localhost:27017"
        auth scram-sha-1{
          username = "root"
          password = "secret"
        }
      }
      query {
        database = "test"
        collection = "users"
        filter {
          _id = personId
        }
        set {
          firstname = "Tom"
        }
      }
      response {
        users = _.mongo.document
      }
    }
    mongo updateMany {
      connection {
        uri = "mongodb://localhost:27017"
        auth scram-sha-1{
          username = "root"
          password = "secret"
        }
      }
      query {
        database = "test"
        collection = "users"
        set{
          firstname = "Jim"
        }
      }
      response {
        users = _.mongo.document
      }
    }

    print {
      msg= "${fullname} owns role ${role}"
    }
  }

  then "connection is established"{
    assert {
      assertion = role == "admin"
    }
    assert {
      assertion = users[_.index].lastname == "Doe"
      count = len(users)
    }
  }
}

// go run ./cmd/orion.go run --input ${ORION_SAMPLES}/plugins/mongo/sample.hcl
