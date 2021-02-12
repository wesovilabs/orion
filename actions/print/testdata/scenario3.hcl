print {
  msg = "Hello Mr ${lastname}"
  prefix = prefixes[0]
}

print {
  msg = "Hello Mr ${lastname}"
}

print {
  prefix = prefixes[0]
  msg = "Hello Mr ${lastname}"
  showTimestamp = ""
  timestampFormat = timestampFmt
  format = "json"
}