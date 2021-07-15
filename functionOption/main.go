package main

import "fmt"

type Option struct {
	A string
	B string
	C int
}
type OptionFunc func(*Option)

var (
	defaultOption = &Option{
		A: "A",
		B: "B",
		C: 100,
	}
)

func newOption1(a, b string, c int) *Option {
	return &Option{
		A: a,
		B: b,
		C: c,
	}
}
func WithA(a string) OptionFunc {
	return func(o *Option) {
		o.A = a
	}
}
func WithB(b string) OptionFunc {
	return func(o *Option) {
		o.B = b
	}
}
func WithC(c int) OptionFunc {
	return func(o *Option) {
		o.C = c
	}
}
func newOption2(opts ...OptionFunc) (opt *Option) {
	opt = defaultOption
	for _, o := range opts {
		o(opt)
	}
	return
}
func main() {

	x := newOption2(
		WithA("沙河娜扎"),
		WithC(250),
	)
	fmt.Println(x)
}
