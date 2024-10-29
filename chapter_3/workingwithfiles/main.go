package main

import (
	"archive/zip"
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	task2()
}

func readDirExample(dirName string) {
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		fmt.Printf("Error while reading the directory: %s: \n", err)
		return
	}
	for _, file := range files {
		fmt.Printf("Name: %s, Size: %d bytes\n", file.Name(), file.Size())
	}
}

// os package methods practice

func createExample() {
	path := "chapter_3/workingwithfiles/example.txt"
	file, err := os.Create(path)
	if err != nil {
		fmt.Print("Error while create a file", err)
		return
	}
	defer file.Close()

	file.WriteString("this is simple text file.\n")
	file.WriteString("Second line in the file.\n")
}

func openExample() {
	path := "chapter_3/workingwithfiles/example.txt"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error while opening the file", err)
		return
	}
	content, _ := os.ReadFile(file.Name())
	fmt.Printf("File name: %s. Content: \n %s", file.Name(), string(content))
	file.Close()
}

func statExample() {
	path := "chapter_3/workingwithfiles/exampe.txt"
	info, err := os.Stat(path)
	if err != nil {
		panic("File not fount")
	}
	fmt.Printf("File name: %s\nSize: %d\nIsDir: %v\nMode: %v\nSystem: %v\nModification time: %v\n",
		info.Name(),
		info.Size(),
		info.IsDir(),
		info.Mode(),
		info.Sys(),
		info.ModTime())
}

//	bufio package example in practice

func bufioReaderTest() {
	path := "chapter_3/workingwithfiles/bufiotest.txt"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Failed to open file", err)
		return
	}
	defer file.Close()

	rd := bufio.NewReader(file)

	buf := make([]byte, 10)
	n, err := rd.Read(buf)
	if err != nil {
		fmt.Println("Error while reading file")
		return
	}

	fmt.Printf("Read %d bytes: %s\n", n, buf)
	// Read until newline character
	line, err := rd.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading line:", err)
		return
	}
	fmt.Printf("Read line: %s\n", line)
}

func osExample() {
	// Open the source file for reading
	sourceFile, err := os.Open("chapter_3/workingwithfiles/example.txt")
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return
	}
	defer sourceFile.Close()

	// Create the destination file for writing
	destFile, err := os.Create("chapter_3/workingwithfiles/destination.txt")
	if err != nil {
		fmt.Println("Error creating destination file:", err)
		return
	}
	defer destFile.Close()

	// Create a scanner to read the source file line by line
	scanner := bufio.NewScanner(sourceFile)
	// Create a writer to write to the destination file
	writer := bufio.NewWriter(destFile)

	// Process each line, transform it, and write it to the new file
	for scanner.Scan() {
		line := scanner.Text()
		// Modify the line in some way, for example by converting it to uppercase
		modifiedLine := strings.ToUpper(line)
		// Write the modified line to the destination file
		writer.WriteString(modifiedLine + "\n")
	}

	// Check for errors in scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading source file:", err)
	}

	// Flush the writer to ensure all data is written to the file
	writer.Flush()
}

func task() {
	scanner := bufio.NewScanner(os.Stdin)
	res := 0
	for scanner.Scan() {
		line := scanner.Text()
		if temp, err := strconv.Atoi(line); err == nil {
			res += temp
		}
	}
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(strconv.Itoa(res) + "\n")
	writer.Flush()
}

// encoding/csv package
func testEncodingCsv() {
	buf := bytes.NewBuffer(nil)
	w := csv.NewWriter(buf)

	for i := 1; i <= 3; i++ {
		val1 := fmt.Sprintf("row %d col 1", i)
		val2 := fmt.Sprintf("row %d col 2", i)
		val3 := fmt.Sprintf("row %d col 3", i)
		err := w.Write([]string{val1, val2, val3})
		if err != nil {
			fmt.Println("Error writing row: ", err)
			return
		}
	}

	w.Flush()

	w.WriteAll([][]string{
		{"row 4 col 1", "row 4 col 2", "row 4 col 3"},
		{"row 5 col 1", "row 5 col 2", "row 5 col 3"},
	})

	w.Flush()
	err := os.WriteFile("chapter_3/workingwithfiles/csvtest.txt", buf.Bytes(), 0644)
	if err != nil {
		fmt.Println("Error while writing csv format data to file", err)
	}
	fmt.Println(buf.String())
}
func task1() {
	zipPath := "/Users/orazboyevolmosbek/Downloads/task.zip"
	zipFile, err := zip.OpenReader(zipPath)
	if err != nil {
		log.Fatal("Failed to open ZIP archive:", err)
	}
	defer zipFile.Close()

	var targetData string

	// Step 2: Iterate over each file in the ZIP archive
	for _, file := range zipFile.File {
		// Check if the file has a .txt extension
		if filepath.Ext(file.Name) == ".txt" {
			// Open the file for reading
			f, err := file.Open()
			if err != nil {
				log.Printf("Could not open file %s: %v", file.Name, err)
				continue
			}
			defer f.Close()

			// Try parsing the file as a CSV
			csvReader := csv.NewReader(f)
			csvReader.Comma = ',' // Ensure comma is used as separator

			// Attempt to read all rows
			rows, err := csvReader.ReadAll()
			if err == nil && len(rows) == 10 && len(rows[0]) == 10 {
				// Check if it is the 10x10 table we are looking for
				targetData = rows[4][2] // 5th line (index 4), 3rd column (index 2)
				break                   // Stop searching once we find the correct file
			}
		}
	}

	// Step 3: Check if the target data was found and print it
	if targetData != "" {
		fmt.Println("Data at 5th line, 3rd column:", targetData)
	} else {
		log.Fatal("No valid CSV data found in any .txt file")
	}
}

func task2() {
	path := "/Users/orazboyevolmosbek/Downloads/task.data"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error while opening file", err)
		return
	}
	pos := 1
	reader := bufio.NewReader(file)
	for {
		elem, err := reader.ReadString(';')
		fmt.Println("element: ", elem)
		if err != nil {
			break
		} else if strings.ReplaceAll(elem, ";", "") == "0" {
			break
		}
		pos++
	}
	fmt.Println(pos)
}
