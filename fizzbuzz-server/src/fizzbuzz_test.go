package main

import (
	"strings"
	"fmt"
	"testing"
)

func TestFizzBuzz15(t *testing.T) {
	have := FizzBuzz(int64(15))

	have = strings.Replace(have, "\n", "", -1)

	fmt.Println(have)
	want := "12Fizz4BuzzFizz78FizzBuzz11Fizz1314FizzBuzz"
	
	fmt.Println(want)

	if have != want {
		t.Errorf("TestFizzBuzz15 wanted: \n %v \n but got \n %v \n", want, have)
	}
}
