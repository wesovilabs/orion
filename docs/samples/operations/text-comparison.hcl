scenario "text functions" {
  given {
      set input {
          value = "hello word"
      }
  }
  when "do text comparisons" {
      set res1 {
          value = input=="Hello World"
      }
      set res2 {
          value = input=="hello world"
      }
      set res3 {
          value = input!="Hello World"
      }
      set res4 {
          value = input!="hello world"
      }
      set res5 {
          value = eqIgnoreCase(input,"Hello World")
      }
      set res6 {
          value = eqIgnoreCase(input,"hello world")
      }
  }
  then "the results are the expected"{
      assert {
          assertion = !res1 && res2 && res3 && !res4 && res5 && res6
      }
  }
}

// go run ./cmd/orion.go run --input ${ORION_SAMPLES}/functions/text-comparison.hcl