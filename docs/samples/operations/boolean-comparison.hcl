scenario "check boolean operations" {
  given "set value" {
    set val1{
         value = true
     }
  }
  when "do the operations" {
     set val2{
         value = !val1
     }
  }
  then "the result is the expected"{
      assert {
          assertion = val1 && val2 && (val1 || val2) && !val2 && val2==false 
      }
  }
}

// go run ./cmd/orion.go run --input ${ORION_SAMPLES}/functions/boolean-comparison.hcl