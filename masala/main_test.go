package main

import "testing"

func TestTriangle(t *testing.T) {
	testCases := []struct{
		a int
		b int
		c int
		output bool
	}{	
		{2,3,4,false},
		{6,8,10,true},
	}
	for _, testCase := range testCases {
		res := Triangle(testCase.a,testCase.b,testCase.c)
		if res != testCase.output {
			t.Errorf("Uchburchak yasab bomidi")
		}
	} 
}