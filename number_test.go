package fake

import (
	"math"
	"testing"
)

const testTryCount = 100

func TestFloat32(t *testing.T) {
	TestFloat32Range(t)
}

func TestFloat32Range(t *testing.T) {
	type args struct {
		min float32
		max float32
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "full range", args: args{math.MinInt64, math.MaxInt64}},
		{name: "negative range", args: args{math.MinInt64, 0}},
		{name: "negative range", args: args{0, math.MaxInt64}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Float32Range(tt.args.min, tt.args.max)
			if got < tt.args.min || got > tt.args.max {
				t.Errorf("Float32Range() = %v, range %v ~ %v", got, tt.args.min, tt.args.max)
			}
		})
	}
}

func TestFloat64(t *testing.T) {
	TestFloat64Range(t)
}

func TestFloat64Range(t *testing.T) {
	type args struct {
		min float64
		max float64
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "full range", args: args{math.MinInt64, math.MaxInt64}},
		{name: "negative range", args: args{math.MinInt64, 0}},
		{name: "negative range", args: args{0, math.MaxInt64}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Float64Range(tt.args.min, tt.args.max)
			if got < tt.args.min || got > tt.args.max {
				t.Errorf("Float64Range() = %v, range %v ~ %v", got, tt.args.min, tt.args.max)
			}
		})
	}
}

func TestInt16(t *testing.T) {
	TestInt16Range(t)
}

func TestInt16Range(t *testing.T) {
	type args struct {
		min int16
		max int16
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "full range", args: args{math.MinInt16, math.MaxInt16}},
		{name: "negative range", args: args{math.MinInt16, 0}},
		{name: "positive range", args: args{0, math.MaxInt16}},
		{name: "negative equals", args: args{math.MinInt16, math.MinInt16 + 1}},
		{name: "positive equals", args: args{math.MaxInt16 - 1, math.MaxInt16}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Int16Range(tt.args.min, tt.args.max)
			if got < tt.args.min || got > tt.args.max {
				t.Errorf("Int16Range() = %v, range %v ~ %v", got, tt.args.min, tt.args.max)
			}
		})
	}
}

func TestInt32(t *testing.T) {
	TestInt32Range(t)
}

func TestInt32Range(t *testing.T) {
	type args struct {
		min int32
		max int32
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "full range", args: args{math.MinInt32, math.MaxInt32}},
		{name: "negative range", args: args{math.MinInt32, 0}},
		{name: "positive range", args: args{0, math.MaxInt32}},
		{name: "negative equals", args: args{math.MinInt32, math.MinInt32 + 1}},
		{name: "positive equals", args: args{math.MaxInt32 - 1, math.MaxInt32}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < testTryCount; i++ {
				got := Int32Range(tt.args.min, tt.args.max)
				if got < tt.args.min || got > tt.args.max {
					t.Errorf("Int32Range() = %v, range %v ~ %v", got, tt.args.min, tt.args.max)
				}
			}
		})
	}
}

func TestInt64(t *testing.T) {
	TestInt64Range(t)
}

func TestInt64Range(t *testing.T) {
	type args struct {
		min int64
		max int64
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "full range", args: args{math.MinInt64, math.MaxInt64}},
		{name: "negative range", args: args{math.MinInt64, 0}},
		{name: "positive range", args: args{0, math.MaxInt32}},
		{name: "negative equals", args: args{math.MinInt64, math.MinInt64 + 1}},
		{name: "positive equals", args: args{math.MaxInt64 - 1, math.MaxInt64}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Int64Range(tt.args.min, tt.args.max)
			if got < tt.args.min || got > tt.args.max {
				t.Errorf("Int64Range() = %v, range %v ~ %v", got, tt.args.min, tt.args.max)
			}
		})
	}
}

func TestInt8(t *testing.T) {
	TestInt8Range(t)
}

func TestInt8Range(t *testing.T) {
	type args struct {
		min int8
		max int8
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "full range", args: args{math.MinInt8, math.MaxInt8}},
		{name: "negative range", args: args{math.MinInt8, 0}},
		{name: "positive range", args: args{0, math.MaxInt8}},
		{name: "negative equals", args: args{math.MinInt8, math.MinInt8 + 1}},
		{name: "positive equals", args: args{math.MaxInt8 - 1, math.MaxInt8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < testTryCount; i++ {
				got := Int8Range(tt.args.min, tt.args.max)
				if got < tt.args.min || got > tt.args.max {
					t.Errorf("Int8Range() = %v, range %v ~ %v", got, tt.args.min, tt.args.max)
				}
			}
		})
	}
}

func TestUint8(t *testing.T) {
	TestUint8Range(t)
}

func TestUint8Range(t *testing.T) {
	type args struct {
		min uint8
		max uint8
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "full range", args: args{0, math.MaxUint8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint8Range(tt.args.min, tt.args.max)
			if got < tt.args.min || got > tt.args.max {
				t.Errorf("Uint8Range() = %v, range %v ~ %v", got, tt.args.min, tt.args.max)
			}
		})
	}
}

func TestUint16(t *testing.T) {
	TestUint16Range(t)
}

func TestUint16Range(t *testing.T) {
	type args struct {
		min uint16
		max uint16
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "full range", args: args{0, math.MaxUint16}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint16Range(tt.args.min, tt.args.max)
			if got < tt.args.min || got > tt.args.max {
				t.Errorf("Uint16Range() = %v, range %v ~ %v", got, tt.args.min, tt.args.max)
			}
		})
	}
}

func TestUint32(t *testing.T) {
	TestUint32Range(t)
}

func TestUint32Range(t *testing.T) {
	type args struct {
		min uint32
		max uint32
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "full range", args: args{0, math.MaxUint32}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint32Range(tt.args.min, tt.args.max)
			if got < tt.args.min || got > tt.args.max {
				t.Errorf("Uint32Range() = %v, range %v ~ %v", got, tt.args.min, tt.args.max)
			}
		})
	}
}

func TestUint64(t *testing.T) {
	TestUint64Range(t)
}

func TestUint64Range(t *testing.T) {
	type args struct {
		min uint64
		max uint64
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "full range", args: args{0, math.MaxUint64}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint64Range(tt.args.min, tt.args.max)
			if got < tt.args.min || got > tt.args.max {
				t.Errorf("Uint64Range() = %v, range %v ~ %v", got, tt.args.min, tt.args.max)
			}
		})
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
	}{
		{name: "full range", args: args{math.MinInt64 / 2, (math.MaxInt64 - 3) / 2}},
		{name: "negative range", args: args{math.MinInt64 + 2, 0}},
		{name: "positive range", args: args{0, math.MaxInt64 - 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Num(tt.args.min, tt.args.max)
			if got < tt.args.min || got > tt.args.max {
				t.Errorf("Num() = %v, range %v ~ %v", got, tt.args.min, tt.args.max)
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
	}{
		{name: "full range", args: args{math.MinInt64 / 2, (math.MaxInt64 - 3) / 2}},
		{name: "negative range", args: args{math.MinInt64 + 2, 0}},
		{name: "positive range", args: args{0, math.MaxInt64 - 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Num64(tt.args.min, tt.args.max)
			if got < tt.args.min || got > tt.args.max {
				t.Errorf("Num64() = %v, range %v ~ %v", got, tt.args.min, tt.args.max)
			}
		})
	}
}

// Benchmark

func BenchmarkInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int8()
	}
}

func BenchmarkInt16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int16()
	}
}

func BenchmarkInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int32()
	}
}

func BenchmarkInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int64()
	}
}

func BenchmarkUInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uint8()
	}
}

func BenchmarkUInt16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uint16()
	}
}

func BenchmarkUInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uint32()
	}
}

func BenchmarkUInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uint64()
	}
}

func BenchmarkFloat32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Float32()
	}
}

func BenchmarkFloat642(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Float64()
	}
}
