input {
  arg s {
    default ="Hello world"
  }
}
scenario "check string search oeprations" {
  when "do the operations" {
    set res1 {
      value = contains(s,"world")
    }
    set res2 {
      value = hasPrefix(s,"Hello")
    }
    set res3 {
      value = hasSuffix(s,"world")
    }
    set res4 {
      value = count(s,"l")
    }
    set res5 {
      value = indexOf(s,"l")
    }
    set res6 {
      value = lastIndexOf(s,"l")
    }
  }
  then "the result is the exapected"{
    assert {
      assertion = res1 && res2 && res3 && res4==3 && res5==2 && res6==9
    }
  }
}