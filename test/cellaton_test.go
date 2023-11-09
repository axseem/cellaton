package cellaton_test

import (
	"testing"

	"github.com/axseem/cellaton"
	"github.com/axseem/cellaton/neighborhood"
	"github.com/axseem/cellaton/rules"
)

func BenchmarkGof128x128(b *testing.B) {
	field := cellaton.NewField(256, 256, rules.Gof, neighborhood.Moore)
	field.FillRand(50, 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		field.NextGen()
	}
}

func BenchmarkGof256x256(b *testing.B) {
	field := cellaton.NewField(256, 256, rules.Gof, neighborhood.Moore)
	field.FillRand(50, 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		field.NextGen()
	}
}

func BenchmarkGof512x512(b *testing.B) {
	field := cellaton.NewField(512, 512, rules.Gof, neighborhood.Moore)
	field.FillRand(50, 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		field.NextGen()
	}
}

func BenchmarkGof1024x1024(b *testing.B) {
	field := cellaton.NewField(1024, 1024, rules.Gof, neighborhood.Moore)
	field.FillRand(50, 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		field.NextGen()
	}
}
