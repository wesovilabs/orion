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