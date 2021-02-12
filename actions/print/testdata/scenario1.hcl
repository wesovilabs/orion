
print {
  msg = "Hello Mr Robot"
  prefix = "DEBUG"
}

print {
  msg = "Hello Mr ${lastname}"
  prefix = "DEBUG"
}

print {
  msg = "Hello Mr ${lastname}"
  prefix = "DEBUG"
  showTimestamp = true
}

print {
  msg = "Hello Mr ${lastname}"
  prefix = "DEBUG"
  showTimestamp = showTimestamp
}

print {
  msg = "Hello Mr ${lastname}"
  prefix = "DEBUG"
  showTimestamp = showTimestamp
  when = true
}