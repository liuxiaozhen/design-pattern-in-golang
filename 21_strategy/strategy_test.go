package strategy

import (
	"fmt"
	"testing"
)

func Test_strategy(t *testing.T) {
	sc := NewStrategyContext(5, 3, new(AddStrategy))
	v := sc.executeStrategy()
	expect := 8
	if v != expect {
		fmt.Printf("expect=%v,got=%v", expect, v)
	}
	sc1 := NewStrategyContext(5, 3, new(SubStrategy))
	v = sc1.executeStrategy()
	expect = 2
	if v != expect {
		fmt.Printf("expect=%v,got=%v", expect, v)
	}
	sc2 := NewStrategyContext(5, 3, new(MulStrategy))
	v = sc2.executeStrategy()
	expect = 15
	if v != expect {
		fmt.Printf("expect=%v,got=%v", expect, v)
	}
}
