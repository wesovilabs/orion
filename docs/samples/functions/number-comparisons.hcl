
scenario "number comparisons" {
  given "input values" {
      set val1 {
          value = 100
      }
      set val2 {
          value = 30.213
      }
  }
  when "do operations with numbers" {
     print {
         msg="Input values are val1:${val1} and val2:${val2}"
     }
  }
  then "result is the exapected"{
      assert {
          assertion= val1>val2 && val1>=val2 && val2<val1 && val2<=val1
      }
  }
}


// go run ./cmd/orion.go run --input ${ORION_SAMPLES}/functions/number-comparisons.hcl

