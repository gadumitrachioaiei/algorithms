package steadygene

import (
	"io/ioutil"
	"strconv"
	"testing"
)

func TestSteadyGene(t *testing.T) {
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			"case1",
			"GAAATAAA",
			5,
		},
		{
			"case2",
			"ACTGAAAG",
			2,
		},
		{
			"case3",
			"TGATGCCGTCCCCTCAACTTGAGTGCTCCTAATGCGTTGC",
			5,
		},
		{
			"case4",
			"ACGT",
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SteadyGene(tt.args); got != tt.want {
				t.Errorf("SteadyGene() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSteadyGeneFromInput(t *testing.T) {
	inputPath := "input.txt"
	input, err := ioutil.ReadFile(inputPath)
	if err != nil {
		t.Fatal(err)
	}
	outputPath := "output.txt"
	output, err := ioutil.ReadFile(outputPath)
	if err != nil {
		t.Fatal(err)
	}
	result, err := strconv.Atoi(string(output))
	if err != nil {
		t.Fatal(err)
	}
	if got := SteadyGene(string(input)); got != result {
		t.Errorf("SteadyGene() = %v, want %v", got, result)
	}
}
