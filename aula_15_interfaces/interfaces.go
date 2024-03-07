package main

import "fmt"

type runner interface {
	run() string
}

type turbo interface {
	turbo() string
}

type turboRunner interface {
	runner
	turbo
}

type cat struct{}

func (c cat) run() string {
	return "Running like a tiger"
}

type dog struct{}

func (d dog) run() string {
	return "Running like a wolf"
}

type dogRobot struct {
	battery float64
	dog
}

func (dr dogRobot) turbo() string {
	return "DogRobot says SUTUTUTUTU"
}

func run(r runner) {
	fmt.Println(r.run())
}

func twinTurbo(t turbo) {
	fmt.Println(t.turbo())
	fmt.Println(t.turbo())
}

func main() {
	thor := dog{}
	fluffy := cat{}
	thor3000 := dogRobot{}

	run(thor)
	run(fluffy)
	run(thor3000)

	twinTurbo(thor3000)
}
