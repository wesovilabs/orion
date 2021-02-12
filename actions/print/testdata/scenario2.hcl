print {
  msg = "Hello Mr Robot"
}

print {
  msg = "Hello Mr Robot"
  prefix = "DEBUG"
}

print {
  msg = "Hello Mr ${lastname}"
  prefix = "DEBUG"
}

print {
  msg = lastname
}

print {
  msg = true
}

print {
  msg = "${lastname}_${lastname}"
  format="json"
}

print {
  msg = "${lastname}_${lastname}"
  format="plain"
}

print {
  msg = "Hello Mr ${lastname}"
  showTimestamp = true
  timestampFormat = "yyyy-MM-dd hh:mm"
  format = "json"
}