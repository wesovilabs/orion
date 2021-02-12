scenario "number operation with int values" {
  given "int values" {
    set val {
      value = 100
    }
  }
  when "do operations with numbers" {
    set valSqrt {
      value = sqrt(val)
    }
    set valCos {
      value = cos(val)
    }
    set valSin {
      value = sin(val)
    }
    set valRound {
      value = round(val)
    }
    set valPow {
      value = pow(val, 2)
    }
    set valMod {
      value = mod(val, 2)
    }
    set valMax {
      value = max(val, 2*val)
    }
    set valMin {
      value = min(val, 2*val)
    }
    set valOp {
      value = ((5 - 1)*(val+2))/3
    }
    print {
      msg = valSin
    }
  }
  then "result is the exapected" {
    assert {
      assertion = (
        valSqrt==10 && valMax==200 && valOp==136 &&  val==100 &&
        valCos==0.8623188723 && valSin==-0.5063656411 && valRound==100 &&
        valMod==0 && valMin==100
      )
    }
  }
}
