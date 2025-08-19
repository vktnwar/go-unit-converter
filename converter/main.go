package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// Converted Functions

func convertTemperature(from, to string, value float64) (float64, error) {
	switch from {
	case "celsius":
		switch to {
		case "fahrenheit":
			return value*9/5 + 32, nil
		case "kelvin":
			return value + 273.15, nil
		case "celsius":
			return value, nil
		default:
			return 0, fmt.Errorf("unidade de destino '%s' inválida para temperatura", to)
		}
	case "fahrenheit":
		switch to {
		case "celsius":
			return (value - 32) * 5 / 9, nil
		case "kelvin":
			return (value-32)*5/9 + 273.15, nil
		case "fahrenheit":
			return value, nil
		default:
			return 0, fmt.Errorf("unidade de destino '%s' inválida para temperatura", to)
		}
	case "kelvin":
		switch to {
		case "celsius":
			return value - 273.15, nil
		case "fahrenheit":
			return (value-273.15)*9/5 + 32, nil
		case "kelvin":
			return value, nil
		default:
			return 0, fmt.Errorf("unidade de destino '%s' inválida para temperatura", to)
		}
	default:
		return 0, fmt.Errorf("unidade de origem '%s' inválida para temperatura", from)
	}
}

func main() {
	category := flag.String("category", "", "Categories (temperature, weight, distance)")
	from := flag.String("from", "", "Origin unit")
	to := flag.String("to", "", "Destination unit")
	value := flag.Float64("value", 0, "Amount to be converted")

	flag.Parse()

	*category = strings.ToLower((*category))
	*from = strings.ToLower(*from)
	*to = strings.ToLower(*to)

	if *category == "" {
		fmt.Println("Erro: categoria é obrigatória")
		os.Exit(1)
	}
	if *from == "" {
		fmt.Println("Erro: Origem é obrigatória")
		os.Exit(1)
	}
	if *to == "" {
		fmt.Println("Erro: destino é obrigatória")
		os.Exit(1)
	}

	var result float64
	var err error

	switch *category {
	case "temperature":
		result, err = convertTemperature(*from, *to, *value)
	case "weight":
		result, err = convertWeight(*from, *to, *value)
	case "distance":
		result, err = convertDistance(*from, *to, *value)
	default:
		fmt.Printf("Error: Categoria '%s' não suportada\n", *category)
		fmt.Println("Categorias disponíveis: Temperature, Weight, Distance")
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("Error %v\n", err)
		os.Exit(1)
	}

	fmt.Println("%.2f %s = %.2f %s\n", *value, *from, result, *to)
}
