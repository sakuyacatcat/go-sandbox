package controller

import (
	"github.com/sakuyacatcat/fizzbuzz/pkg/domain/services"
	"github.com/sakuyacatcat/fizzbuzz/pkg/views"
)

type RangePrinter struct {
	converter services.NumberConverterInterface
	output   views.OutputInterface
}

func NewRangePrinter(nc services.NumberConverterInterface, so views.OutputInterface) *RangePrinter {
	return &RangePrinter{
		converter: nc,
		output:   so,
	}
}

func (rp *RangePrinter) PrintRange(begin int, end int) {
	for i := begin; i <= end; i++ {
		rp.output.Write(rp.converter.Convert(i))
	}
}
