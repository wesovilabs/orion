assert {
  assertion = true
  when = true
}

assert {
  assertion = false
}

assert {
  assertion = true && false
}

assert {
  assertion = person.lastname == "Robot"
}

assert {
  assertion = person.lastname == "Robot" || person.age <15
}

assert {
  assertion = person.lastname == "Robot" && person.age >15
}

assert {
  assertion = person.lastname == "Robotic" || person.age <15
}

assert {
  assertion = true
}

assert {
  assertion = unknown
}

