input {
  arg s {
    default ="Hello world"
  }
}
scenario "check string converter oeprations" {
  when "convert the input" {
    set res1 {
      value = toLowercase(s)
    }
    set res2 {
      value = toUppercase(s)
    }
    set res3 {
      value = trimPrefix(s,"Hello ")
    }
    set res4 {
      value = trimSuffix(s," world")
    }
    set res5 {
      value = replaceOne(s,"Hello", "Bye")
    }
    set res6 {
      value = replaceAll(s,"o", "a")
    }
  }
  then "check results"{
    assert {
      assertion = (
        res1=="hello world" &&  res2=="HELLO WORLD" &&
        res3 == "world" &&  res4 == "Hello" &&
        res5 == "Bye world" &&  res6 == "Hella warld"
      )
    }
  }
}