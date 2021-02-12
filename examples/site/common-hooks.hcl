before sample {
  set a {
    value = 1
  }
}
before common {
  set b {
    value = 2
  }
}
after common {
  print {
    msg = "operation completed with result ${result}"
  }
}