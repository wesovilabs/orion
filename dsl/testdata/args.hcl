arg firstname {
  description = "firstname of the person"
  default = "John"
}
arg lastname {}
arg age {
  default = 2 * 3
}

arg test{
  unknown=""
}

arg elements {
  default = ["a","b","c"]
}

arg city{}

arg country{
  description = "${home}"
}

arg gender{}

arg team{
  default = "team-${unknown}"
}
