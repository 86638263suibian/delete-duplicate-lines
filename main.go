package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputPtr := flag.String("input", "", "Input file path")
	outputPtr := flag.String("output", "", "Output file path")
	seperatorPtr := flag.String("seperator", "", "Seperator for each lines")
	compareIndexPtr := flag.Int("index", 0, "Index of seperated slice to compare")
	skipLinesPtr := flag.Bool("skip", false, "Skip lines that cannot be seperated")
	flag.Parse()
	if *inputPtr == "" {
		fmt.Println("Please provide input file path")
		return
	}
	inputFile, err := os.Open(*inputPtr)
	if err != nil {
		fmt.Printf("There was a problem with input file: %s", err)
		return
	}
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	var outputBuff bytes.Buffer
	w := bufio.NewWriter(&outputBuff)
	var rows []string
	var seperatedRow []string
ScannerLoop:
	for scanner.Scan() {
		row := scanner.Text()
		if *seperatorPtr != "" {
			seperatedRow = strings.Split(row, *seperatorPtr)
			if len(seperatedRow) <= *compareIndexPtr {
				if *skipLinesPtr == false {
					rows = append(rows, row)
				}
				continue ScannerLoop
			}
		}
		for _, r := range rows {
			if *seperatorPtr != "" {
				seperatedR := strings.Split(r, *seperatorPtr)
				if seperatedRow[*compareIndexPtr] == seperatedR[*compareIndexPtr] {
					continue ScannerLoop
				}
			} else if r == row {
				continue ScannerLoop
			}
		}
		rows = append(rows, row)
		if _, err := w.WriteString(row + "\n"); err != nil {
			fmt.Printf("error writing: %s", err)
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
		fmt.Printf("There was a problem with output file: %s", err)
		return
	}
	defer outputFile.Close()
	_, err = outputFile.Write(outputBuff.Bytes())
	if err != nil {
		fmt.Printf("There was a problem writing in output file: %s", err)
		return
	}
	fmt.Println("output file create: ", *outputPtr)
}
