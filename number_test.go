package fake

import (
	"github.com/davecgh/go-spew/spew"
	"math"
	"testing"
)

const testTryCount = 10

func TestFloat32(t *testing.T) {
	for i := 0; i < testTryCount; i++ {
		n := Float32()
		if n < math.SmallestNonzeroFloat32 || n > math.MaxFloat32 {
			t.Errorf("%v overflows %v ~ %v", n, math.SmallestNonzeroFloat32, math.MaxFloat32)
		}
	}
}

func TestFloat32Range(t *testing.T) {

	a := float32Range(-3.1, -3.0)
	spew.Dump(a)
	type args struct {
		min float32
		max float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float32Range(tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("Float32Range() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64(t *testing.T) {
	tests := []struct {
		name string
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64(); got != tt.want {
				t.Errorf("Float64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64Range(t *testing.T) {
	type args struct {
		min float64
		max float64
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
			if got := Float64Range(tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("Float64Range() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16(t *testing.T) {
	for i := 0; i < testTryCount; i++ {
		n := Int16()
		if n < math.MinInt16 || n > math.MaxInt16 {
			t.Errorf("%v overflows %v ~ %v", n, math.MinInt16, math.MaxInt16)
		}
	}
}

func TestInt32(t *testing.T) {
	for i := 0; i < testTryCount; i++ {
		n := Int32()
		if n < math.MinInt32 || n > math.MaxInt32 {
			t.Errorf("%v overflows %v ~ %v", n, math.MinInt32, math.MaxInt32)
		}
	}
}

func TestInt64(t *testing.T) {
	for i := 0; i < testTryCount; i++ {
		n := Int64()
		if n < math.MinInt64 || n > math.MaxInt64 {
			t.Errorf("%v overflows %v ~ %v", n, math.MinInt64, math.MaxInt64)
		}
	}
}

func TestInt8(t *testing.T) {
	for i := 0; i < testTryCount; i++ {
		n := Int8()
		if n < math.MinInt8 || n > math.MaxInt8 {
			t.Errorf("%v overflows %v ~ %v", n, math.MinInt8, math.MaxInt8)
		}
	}
}

func TestNum(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Num(tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("Num() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNum64(t *testing.T) {
	type args struct {
		min int64
		max int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Num64(tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("Num64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16(t *testing.T) {
	tests := []struct {
		name string
		want uint16
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint16(); got != tt.want {
				t.Errorf("Uint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32(t *testing.T) {
	tests := []struct {
		name string
		want uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint32(); got != tt.want {
				t.Errorf("Uint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64(t *testing.T) {
	tests := []struct {
		name string
		want uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint64(); got != tt.want {
				t.Errorf("Uint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8(t *testing.T) {
	tests := []struct {
		name string
		want uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint8(); got != tt.want {
				t.Errorf("Uint8() = %v, want %v", got, tt.want)
			}
		})
	}
}
