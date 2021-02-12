vars {
  gender = "masculine"
  people = [
    {
      firstname = "John"
      lastname = "Doe"
    },
    {
      firstname = "Jane"
      lastname = "Doe"
    },
    {
      firstname = "Jimmie"
      lastname = "Loe"
    }
  ]
  letters = ["a","b","c"]
}

print {
  msg = "Hello Mr ${firstname} ${lastname}"
  when = eq(gender,"masculine")
}

print {
  msg = "Hello Ms ${firstname} ${lastname}"
  when = eq(gender,"female")
}

print {
  msg ="iteration number ${_.index}"
  while = _.index<10
}

print {
  msg ="iteration number ${_.index}"
  count = 10
}

print {
  msg ="iteration number ${_.index}"
  count = 10
  while = _.index < 5
}

print {
  msg = "There are ${len(people)} people."
}

print {
  msg = "${people[_.index].firstname} ${people[_.index].lastname}"
  while = _.index < len(people)
}

loop {
  items = people
  as = person
  filter = person.lastname == "Doe"

  print {
    msg = "${person.firstname} ${person.lastname}"
  }

  set {
    variable = 2 * _.index
  }
}

loop {
  items = range(1,20)
  print {
    msg = "value ${_.item}"
  }
}

loop {
  items = ["a","b","c"]
  as = letter
  print {
    msg = letter
  }
}

print {
  with logger{}
  msg = "There are 20 cars"
}

vars {
  apiBaseUrl = "http://mycompany.com"
}











print {
  trait logger{}
  msg = "Hello Mr Robinson"
}

print {
  trait logger{}
  msg = "Hello Mr Robinson"
}


