package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// solveQuadratic обчислює корені квадратного рівняння a*x^2 + b*x + c = 0.
// Якщо a == 0, повертається помилка, оскільки рівняння не є квадратним.
// Повертає зріз коренів: 0, 1 або 2 значень.
func solveQuadratic(a, b, c float64) ([]float64, error) {
	if a == 0 {
		return nil, fmt.Errorf("a cannot be 0 (this is not a quadratic equation)")
	}

	d := b*b - 4*a*c

	if d > 0 {
		sqrtD := math.Sqrt(d)
		x1 := (-b + sqrtD) / (2 * a)
		x2 := (-b - sqrtD) / (2 * a)
		return []float64{x1, x2}, nil
	} else if d == 0 {
		x := -b / (2 * a)
		return []float64{x}, nil
	} else {
		return []float64{}, nil
	}
}

// readCoefficientsInteractive зчитує коефіцієнти з консолі.
func readCoefficientsInteractive() (float64, float64, float64, error) {
	reader := bufio.NewReader(os.Stdin)
	var a, b, c float64
	var err error

	readFloat := func(prompt string) (float64, error) {
		for {
			fmt.Print(prompt)
			line, err := reader.ReadString('\n')
			if err != nil {
				return 0, err
			}
			line = strings.TrimSpace(line)
			val, convErr := strconv.ParseFloat(line, 64)
			if convErr != nil {
				fmt.Fprintf(os.Stderr, "Error. Expected a valid real number, got '%s'. Please try again.\n", line)
				continue
			}
			return val, nil
		}
	}

	a, err = readFloat("a = ")
	if err != nil {
		return 0, 0, 0, err
	}
	b, err = readFloat("b = ")
	if err != nil {
		return 0, 0, 0, err
	}
	c, err = readFloat("c = ")
	if err != nil {
		return 0, 0, 0, err
	}

	return a, b, c, nil
}

// readCoefficientsFromFile зчитує коефіцієнти з файлу.
// Файл повинен містити принаймні три числа, розділених пробілами або новими рядками.
func readCoefficientsFromFile(filename string) (float64, float64, float64, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return 0, 0, 0, err
	}

	parts := strings.Fields(string(data))
	if len(parts) < 3 {
		return 0, 0, 0, fmt.Errorf("file does not contain 3 coefficients")
	}

	a, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("cannot parse a: %v", err)
	}
	b, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("cannot parse b: %v", err)
	}
	c, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("cannot parse c: %v", err)
	}

	return a, b, c, nil
}

func main() {
	args := os.Args[1:]

	// Якщо аргументів немає, працюємо в інтерактивному режимі.
	if len(args) == 0 {
		fmt.Println("Enter coefficients for the quadratic equation a*x^2 + b*x + c = 0")
		a, b, c, err := readCoefficientsInteractive()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading coefficients: %v\n", err)
			os.Exit(1)
		}

		roots, err := solveQuadratic(a, b, c)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error solving equation: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Equation is: (%.2f) x^2 + (%.2f) x + (%.2f) = 0\n", a, b, c)
		switch len(roots) {
		case 0:
			fmt.Println("There are 0 roots")
		case 1:
			fmt.Println("There is 1 root")
			fmt.Printf("x1 = %.4f\n", roots[0])
		case 2:
			fmt.Println("There are 2 roots")
			fmt.Printf("x1 = %.4f\n", roots[0])
			fmt.Printf("x2 = %.4f\n", roots[1])
		}
	} else if len(args) == 1 {
		// Неінтерактивний режим: зчитування з файлу.
		filename := args[0]
		a, b, c, err := readCoefficientsFromFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading coefficients from file: %v\n", err)
			os.Exit(1)
		}

		roots, err := solveQuadratic(a, b, c)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error solving equation: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Equation is: (%.2f) x^2 + (%.2f) x + (%.2f) = 0\n", a, b, c)
		switch len(roots) {
		case 0:
			fmt.Println("There are 0 roots")
		case 1:
			fmt.Println("There is 1 root")
			fmt.Printf("x1 = %.4f\n", roots[0])
		case 2:
			fmt.Println("There are 2 roots")
			fmt.Printf("x1 = %.4f\n", roots[0])
			fmt.Printf("x2 = %.4f\n", roots[1])
		}
	} else {
		// Якщо аргументів більше одного, виводимо підказку.
		fmt.Printf("Usage:\n  %s              (interactive mode)\n  %s <file>       (non-interactive file mode)\n", os.Args[0], os.Args[0])
		os.Exit(1)
	}
}
