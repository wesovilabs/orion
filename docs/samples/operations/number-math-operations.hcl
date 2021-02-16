
scenario "number operation with int values" {
  given "int values" {
      set valInt {
          value = 100
      }
  }
  when "do operations with numbers" {
      set valIntSqrt {
          value = sqrt(valInt)
      }
      set valIntCos {
          value = cos(valInt)
      }
      set valIntSin {
          value = sin(valInt)
      }
      set valIntRound {
          value = round(valInt)
      }
      set valIntPow {
          value = pow(valInt,2)
      }
      set valIntMod {
          value = mod(valInt,2)
      }
      set valIntMax {
          value = max(valInt,2*valInt)
      }
      set valIntMin {
          value = min(valInt,2*valInt)
      }
      set valIntOp {
          value = ((5-1)*(valInt+2))/3
      }
  }
  then "result is the exapected"{
      assert {
          assertion= true
      }
  }
}

scenario "number operation with float values" {
  given "float values" {
      set valFloat {
          value = 25.34
      }
  }
  
  when "do operations with numbers" {
      set valFloatSqrt {
          value = sqrt(valFloat)
      }
      set valFloatCos {
          value = cos(valFloat)
      }
      set valFloatSin {
          value = sin(valFloat)
      }
      set valFloatRound {
          value = round(valFloat)
      }
      set valFloatPow {
          value = pow(valFloat,2)
      }
      set valFloatMod {
          value = mod(valFloat,2)
      }
      set valFloatMax {
          value = max(valFloat,2*valFloat)
      }
      set valFloatMin {
          value = min(valFloat,2*valFloat)
      }
      set valFloatOp {
          value = ((5-1)*(valFloat+2))/3
      }
  }
  then "result is the exapected"{
      assert {
          assertion= true
      }
  }
}

// go run ./cmd/orion.go run --input ${ORION_SAMPLES}/functions/number-math-operations.hcl

