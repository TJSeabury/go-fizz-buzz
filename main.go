package main

import (
	"fmt"
)

type Rule[T comparable] struct {
	Condition func(T) bool
	Transform func(T) string
}

func fizzBuzz[T comparable](rules []Rule[T]) func(T) string {
	return func(n T) string {
		for _, rule := range rules {
			if rule.Condition(n) {
				return rule.Transform(n)
			}
		}
		return fmt.Sprintf("%v", n)
	}
}

func main() {
	rules := []Rule[int]{
		{
			func(n int) bool { return n == 69 },
			func(n int) string { return "Nice" },
		},
		{
			func(n int) bool { return n%15 == 0 },
			func(n int) string { return "FizzBuzz" },
		},
		{
			func(n int) bool { return n%3 == 0 },
			func(n int) string { return "Fizz" },
		},
		{
			func(n int) bool { return n%5 == 0 },
			func(n int) string { return "Buzz" },
		},
		{
			func(n int) bool { return n%7 == 0 },
			func(n int) string { return "Bang" },
		},
	}
	fizzbuzz := fizzBuzz(rules)

	for i := 1; i <= 100; i++ {
		fmt.Println(fizzbuzz(i))
	}
}
