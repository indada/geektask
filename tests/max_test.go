package tests

import "testing"

func TestInt(t *testing.T)  {
	tests := []struct{
		name string
		a int
		b int
		big int
	}{
		{"haha",2,3,3},
		{"hala",20,30,30},
		{"hala",40,33,33},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int(tt.a,tt.b);got !=tt.big {
				t.Errorf("exp-big:%d,got:%d",tt.big,got)
			}
		})
	}


}