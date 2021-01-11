package main

import (
  "testing"
  "fmt"
)

func TestSum(t *testing.T){
  total := Sum(50,5)
  if total != 55 {
    t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
  }
}

func TestSumTableDriven(t *testing.T) {
    var tests = []struct {
        a, b int
        sum int
    }{
        {0, 1, 1},
        {2, 2, 4},
        {2, -2, 0},
        {10, 10, 20},
        {15, 2, 17},
    }

    for _, tt := range tests {
        testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
        t.Run(testname, func(t *testing.T) {
            ans := Sum(tt.a, tt.b)
            if ans != tt.sum {
                t.Errorf("got %d, want %d", ans, tt.sum)
            }
        })
    }
}
