# fake

Fake data generator in Go

[![Build Status](https://travis-ci.org/devplayg/fake.svg?branch=master)](https://travis-ci.org/devplayg/fake)

## Data types

### Number

Normal works

```go
Num(min, max int) int 
Num64(min, max int64) int64
``` 

Security-sensitive works

- [X] Int8() int8 
- [X] Int16() int16 
- [X] Int32() int32 
- [X] Int64() int64 
- [X] Int8Range(min, max int8) int8 
- [X] Int16Range(min, max int16) int16 
- [X] Int32Range(min, max int32) int32 
- [X] Int64Range(min, max int64) int64 
- [X] Uint8() uint8 
- [X] Uint16() uint16 
- [X] Uint32() uint32 
- [X] Uint64() uint64 
- [X] Uint8Range(min, max uint8) uint8 
- [X] Uint16Range(min, max uint16) uint16 
- [X] Uint32Range(min, max uint32) uint32 
- [X] Uint64Range(min, max uint64) uint64 
- [X] Float32() float32 
- [X] Float32Range(min, max float32) float32 
- [X] Float64() float64 
- [X] Float64Range(min, max float64) float64 

### String 

- [ ] Digit(min, max int) string // "a", "ab", "abc", "abcd"
- [ ] DigitN(n int) string
- [ ] Char() string // "a"
- [ ] CharN() string "aBc"
- [ ] LowerCase() string // "a"
- [ ] LowerCaseN() string // "abc"
- [ ] UpperCase() string // "A"
- [ ] UpperCaseN() string // "ABC"

### Byte

- [ ] Byte() // size: 0..2147483647
- [ ] Byte(5) // size: 0..5
- [ ] Byte(10, 100) // size: 10..100

### Network

- [ ] IPv4() //
- [ ] IPv4(192) // A Class
- [ ] IPv4(192, 168) // B Class
- [ ] IPv4(192, 168, 0) // C Class
- [ ] IPv4But()


### Date

- [ ] Date

## Benchmark

    $ go test -bench=. -benchmem
    goos: windows
    goarch: amd64
    pkg: github.com/devplayg/fake
    BenchmarkInt8-8         60005700                20.5 ns/op             0 B/op          0 allocs/op
    BenchmarkInt16-8        57141495                20.5 ns/op             0 B/op          0 allocs/op
    BenchmarkInt32-8        54393153                22.3 ns/op             0 B/op          0 allocs/op
    BenchmarkInt64-8         2123890               563 ns/op             168 B/op          8 allocs/op
    BenchmarkUInt8-8        54545702                22.3 ns/op             0 B/op          0 allocs/op
    BenchmarkUInt16-8       54543223                22.3 ns/op             0 B/op          0 allocs/op
    BenchmarkUInt32-8       52179811                22.3 ns/op             0 B/op          0 allocs/op
    BenchmarkUInt64-8        2325564               514 ns/op             120 B/op          7 allocs/op
    BenchmarkFloat32-8      63062967                19.5 ns/op             0 B/op          0 allocs/op
    BenchmarkFloat642-8     66609308                17.9 ns/op             0 B/op          0 allocs/op
    PASS
    ok      github.com/devplayg/fake        13.496s