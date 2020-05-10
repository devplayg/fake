# fake

## Number

- [ ] Num(min, max int) int
- [ ] Num64(min, max int64) int64
- [ ] Big(min, max int64) *big.Int

## String 

- [ ] Digit(min, max int) string // "a", "ab", "abc", "abcd"
- [ ] DigitN(n int) string
- [ ] Char() string // "a"
- [ ] CharN() string "aBc"
- [ ] LowerCase() string // "a"
- [ ] LowerCaseN() string // "abc"
- [ ] UpperCase() string // "A"
- [ ] UpperCaseN() string // "ABC"

## Byte

- [ ] Byte() // size: 0..2147483647
- [ ] Byte(5) // size: 0..5
- [ ] Byte(10, 100) // size: 10..100

## Network

- [ ] IPv4() //
- [ ] IPv4(192) // A Class
- [ ] IPv4(192, 168) // B Class
- [ ] IPv4(192, 168, 0) // C Class
- [ ] IPv4But()


## Date

- [ ] Date