package main

import (
	"fmt"
	"log"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
)

func main() {
	if err := f1(); err != nil {
		fmt.Println(err)
	}
}

func f1() error {
	if err := f2(); err != nil {
		return errors.Wrap(err, "f1\n")
	}
	return nil
}
func f2() error {
	if err := f3(); err != nil {
		return errors.Wrap(err, "f2\n")
	}
	return nil
}
func f3() error {
	return errors.New("f3: error")
}

// backoff
func backOff() {
	operation := func() error {
		log.Println("tick...")
		return errors.New("error") // or an error
	}

	err := backoff.Retry(operation, &backoff.ExponentialBackOff{
		InitialInterval: 1 * time.Second,
		Multiplier:      2.0,
		MaxInterval:     60 * time.Second,
		Clock:           backoff.SystemClock,
	})
	if err != nil {
		log.Println(err)
		return
	}
}

// error wrapping
func errWrapping() {
	var err error
	fmt.Println(errors.Wrap(err, "this is an error"))

}

// time
func timeTest() {
	t := time.Now()
	fmt.Println("current", t)

	fmt.Println(t.AddDate(0, 0, 2))
	fmt.Println(t.Add(time.Hour * 1))
}

// diffTime
func diffTime() {
	t1 := time.Now()
	t2 := time.Now()
	date1 := time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.UTC)
	date2 := time.Date(t2.Year(), t2.Month(), t2.Day()-2, 0, 0, 0, 0, time.UTC)
	diff := date1.Sub(date2).Hours() / 24
	fmt.Println(diff)
}

// validator
func validation() {
	type test struct {
		Amount float64 `validate:"min=0.00,max=100.00,required"`
		Name   string  `validate:"required"`
		Opt    string
		Dev    string
	}

	t := test{Amount: 1, Name: "", Opt: "1", Dev: "1"}
	val := validator.New()

	if val.Var(t.Dev, "gte=1") != nil || val.Var(t.Opt, "gte=1") != nil {
		fmt.Println("error")
	}
	// if err := val.Struct(t); err != nil {
	// 	fmt.Println(err)
	// }
}

// largest prime factor
func problem3() {
	n := 600851475143

	_ = n
}

// even fibonancci
func problem2() {
	fibo := func() func() int {
		x, y := 1, 1
		return func() int {
			defer func() { x, y = y, x+y }()
			return y
		}
	}
	f := fibo()

	sum := 0

	n := 0
	for {
		if n > 4000000 {
			break
		}
		n = f()
		if n%2 == 0 {
			sum += n
		}
	}
	fmt.Println(sum)
}

// multiple of 3 and 5
func Problem1() {
	sum := 0
	for i := 3; i < 1000; i++ {
		if i%3 == 0 ||
			i%5 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
