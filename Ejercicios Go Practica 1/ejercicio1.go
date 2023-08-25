// Descripción: Este programa en Go cuenta el número de caracteres, palabras y líneas en un texto ingresado por el usuario.
// Autor: Fabián Fernández

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Ingrese el texto:")
	text := readMultilineInput()

	charCount := len(text)
	wordCount := countWords(text)
	lineCount := strings.Count(text, "\n") + 1

	fmt.Printf("Número de caracteres: %d\n", charCount)
	fmt.Printf("Número de palabras: %d\n", wordCount)
	fmt.Printf("Número de líneas: %d\n", lineCount)
}

func readMultilineInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		lines = append(lines, line)
	}

	return strings.Join(lines, "\n")
}

func countWords(text string) int {
	words := strings.Fields(text)
	return len(words)
}
