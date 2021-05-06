package helper

import "testing"

func Test_Sum(t *testing.T) {
	actual := Sum(1, 2, 3, 4)
	expect := 10
	if actual != expect {
		t.Errorf("Actual %v is not equal to expect %v", actual, expect)
	}
}
