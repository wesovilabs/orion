
# scenario-mult.hcl
scenario "operation multiplication" {
  examples = [
    { x = 20, y = 10, multResult= 200},
    { x = -1, y = -2, multResult= 2},
    { x = 5, y = 5, multResult= 25},
    { x = 5, y = 0, multResult= 0},
  ]
  given "initialie result" {
    set result {
      value = 0
    }
  }
  when "multiply y by x" {
    block {
      set result {
        value = result + x
      }
      print {
        msg = "${x} * ${_.index+1} is ${result}"
      }
      count = y
      when = x>0 && y>0
    }
    set result {
      value = x * y
      when = x<0 || y<0
    }
  }
  then "the result of the operation is the expected" {
    assert {
      assertion = result==multResult
    }
  }
}
