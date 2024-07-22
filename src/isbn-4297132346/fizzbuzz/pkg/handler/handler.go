package handler

import (
	"github.com/sakuyacatcat/fizzbuzz/pkg/controller"
	"github.com/sakuyacatcat/fizzbuzz/pkg/domain/rules"
	"github.com/sakuyacatcat/fizzbuzz/pkg/domain/services"
	"github.com/sakuyacatcat/fizzbuzz/pkg/views"
)

type HandlerInterface interface {
	Handle()
}

type Handler struct {}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Handle() {
	fizzRule := rules.NewCyclicNumberRule(3, "Fizz")
	buzzRule := rules.NewCyclicNumberRule(5, "Buzz")
	passThroughRule := rules.NewPassThroughRule()

	fizzBuzzRule := []rules.ReplaceRule{
		fizzRule,
		buzzRule,
		passThroughRule,
	}

	converter := services.NewNumberConverter(fizzBuzzRule)
	output := views.NewStandardOutput()

	rangePrinter := controller.NewRangePrinter(converter, output)
	rangePrinter.PrintRange(1, 100)
}
