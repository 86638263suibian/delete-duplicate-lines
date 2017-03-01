package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
)

func main() {
	inputPtr := flag.String("input", "", "Input file path")
	outputPtr := flag.String("output", "", "Output file path")
	flag.Parse()
	if *inputPtr == "" {
		fmt.Println("Please provide input file path")
		return
	}
	inputFile, err := os.Open(*inputPtr)
	if err != nil {
		fmt.Errorf("There was a problem with input file: %s", err)
		return
	}
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	var outputBuff bytes.Buffer
	w := bufio.NewWriter(&outputBuff)
	var rows []string
ScannerLoop:
	for scanner.Scan() {
		row := scanner.Text()
		for _, r := range rows {
			if r == row {
				continue ScannerLoop
			}
		}
		rows = append(rows, row)
		if _, err := w.WriteString(row + "\n"); err != nil {
			fmt.Errorf("error writing: %s", err)
			return
		}
	}
	w.Flush()
	if *outputPtr == "" {
		fmt.Printf("output: %s", outputBuff.String())
		return
	}
	outputFile, err := os.Create(*outputPtr)
	if err != nil {
		fmt.Errorf("There was a problem with output file: %s", err)
		return
	}
	defer outputFile.Close()
	_, err = outputFile.Write(outputBuff.Bytes())
	if err != nil {
		fmt.Errorf("There was a problem writing in output file: %s", err)
		return
	}
	fmt.Println("output file create: %s", *outputPtr)
}
