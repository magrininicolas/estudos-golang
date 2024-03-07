package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	ErrDivideByZero = MathError("cannot divide by zero")
)

type Product struct {
	Name string
	Code int
}

type ProductErr struct {
	Product
}

func (pe ProductErr) Error() string {
	return fmt.Sprintf("%s: %d", pe.Name, pe.Code)
}

func showProduct(product Product) error {
	return ProductErr{
		Product: product,
	}
}

type MathError string

func (me MathError) Error() string {
	return string(me)
}

func divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return float64(a / b), nil
}

func verifiesArchive(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("Error opening %s: %w", name, err)
	}
	f.Close()
	fmt.Println("Archive verified")
	return nil
}

func main() {
	// Sentinel Error
	fmt.Printf("\nSentinel Error\n")
	r, err := divide(10, 2)
	if err != nil {
		if err == ErrDivideByZero {
			fmt.Println(err)
			os.Exit(0)
		}
	}
	fmt.Println(r)

	// Package errors
	fmt.Printf("\nPackage Errors\n")
	archiveErr := verifiesArchive("archive.txt")
	if errors.Is(archiveErr, os.ErrNotExist) {
		fmt.Println(archiveErr)
	}

	fmt.Printf("\nErrors\n")
	pErr := showProduct(Product{"Mouse", 1})
	var pe ProductErr
	if errors.As(pErr, &pe) {
		fmt.Println(pe)
	}
}
