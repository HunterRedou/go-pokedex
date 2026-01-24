package main

import(
	"testing"
	"github.com/google/go-cmp/cmp"
)

func TestCleanInput(t *testing.T){
	tests := map[string]struct{
		input string
		want []string
	}{
		"simple": 	{input: "Pikachu Abor Kicklee", want: []string{"pikachu", "abor", "kicklee"}},
		"more": 	{input: "  ABOR KickLee   TUrtok", want: []string{"abor", "kicklee", "turtok"}},
		"all":		{input: "pikachuKickleeTurtok", want: []string{"pikachukickleeturtok"}},
		"sepdiff":	{input: "Abor   Pikachu  Turtok ", want: []string{"abor", "pikachu", "turtok"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := cleanInput(tc.input)
			diff := cmp.Diff(tc.want, got)
			if diff != ""{
				t.Fatalf("description here: %v", diff)
			}
		})
	}
}