scenario "scenario demo" {
  examples = [
    {x=1,   y=2,  out=3},
    {x=1,   y=20, out=21},
    {x=-1,  y=20, out=19},
  ]
  when "add x to y" {
    set result {
      value = x + y
    }
  }
  then "check the output" {
    assert {
      assertion = result==out
    }
  }
}

scenario "scenario demo II" {
  examples = [
    {data = ["o","n","e"], word ="one"},
    {data = ["b","y","e"], word ="bye"},
    {data = ["t","w","o"], word ="two"},
  ]
  given "the output array"{
    set output {
      value = ""
    }
  }
  when "print the letters in the array" {
    block {
      set output{
        value = "${output}${data[_.index]}"
      }
      count = len(data)
    }
  }
  then "check the output" {
    assert {
      assertion = word==output
    }
  }
}