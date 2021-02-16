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
              eq(res1,"hello world") &&  eq(res2,"HELLO WORLD") && 
              eq(res3,"world") &&  eq(res4,"Hello") && 
              eq(res5,"Bye world") &&  eq(res6,"Hella warld")
          )
      }
  }
}

// go run ./cmd/orion.go run --input ${ORION_SAMPLES}/functions/text-converter.hcl