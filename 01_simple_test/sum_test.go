package sum

import "testing"

func TestInts(t *testing.T)  {
  tt := []struct{
      name    string
      numbers []int
      sum     int
  }{
    {"one to five",[]int{1,2,3,4,5},15},
    {"no numbers",nil,0},
    {"one and minus one",[]int{1,-1},0},
  }

  for _, tc := range tt {
    s := Ints(tc.numbers...)
    if s != tc.sum {
      t.Errorf("Sum of %v should be %v; got %v",tc.numbers, tc.sum, s)
    }
  }
}
