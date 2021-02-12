input {
  arg s {
      default ="tofu,tempeh,kale,edamame"
  }
}
scenario "test functions with strings" {
  when "do the operations" {
      set elements {
          value = split(s,",")
      }
      print {
          msg = elements[_.index]
          while = _.index < len(elements)
      }
  }
  then "check results"{
      assert {
          assertion = (
              len(elements)==4 && 
              elements[0]=="tofu" && elements[1]=="tempeh" &&
              elements[2]=="kale" && elements[3]=="edamame"   
          )
      }
  }
}

// go run ./cmd/orion.go run --input ${ORION_SAMPLES}/functions/text-others.hcl