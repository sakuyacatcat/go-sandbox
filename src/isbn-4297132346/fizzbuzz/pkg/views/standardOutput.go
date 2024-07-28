package views

import "fmt"

type StandardOutput struct{}

func NewStandardOutput() *StandardOutput {
	return &StandardOutput{}
}

func (so *StandardOutput) Write(s string) {
	fmt.Println(s)
}
