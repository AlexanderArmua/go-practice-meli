package main

import "testing"

func TestRaizCuadrada(t *testing.T) {
	type args struct {
		numero    float64
		precision int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RaizCuadrada(tt.args.numero, tt.args.precision); got != tt.want {
				t.Errorf("RaizCuadrada() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkRaizCuadrada(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RaizCuadrada(10, 1)
	}
}

func BenchmarkRaizGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RaizGo(10)
	}
}
