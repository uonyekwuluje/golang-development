package main

import (
  "testing"
  "reflect"
)

func TestTwoSum1(t *testing.T) {
  numset1 := []int{2,5,5,11}
  target := 10
  result := twoSum(numset1, target)
  expected := []int{1,2}

  if !(reflect.DeepEqual(result, expected)) {
    t.Errorf("twoSum() test returned an unexpected result: got %v want %v", result, expected)
  }
}

func TestTwoSum2(t *testing.T) {
  numset1 := []int{2,7,11,15}
  target := 9
  result := twoSum(numset1, target)
  expected := []int{0,1}

  if !(reflect.DeepEqual(result, expected)) {
    t.Errorf("twoSum() test returned an unexpected result: got %v want %v", result, expected)
  }
}
