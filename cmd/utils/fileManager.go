package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func ValidateFile(fileHeader *multipart.FileHeader) bool {
	file, err := fileHeader.Open()
	defer file.Close()
	// Create a new ZIP reader.
	zipReader, err := zip.NewReader(file, fileHeader.Size)
	if err != nil {
		//http.Error(w, "Not a valid ZIP file", http.StatusBadRequest)
		fmt.Println("1")
		return false
	}

	// Verify that the ZIP file contains at least one file.
	if len(zipReader.File) < 1 {
		//http.Error(w, "ZIP file is empty", http.StatusBadRequest)
		fmt.Println("2")
		return false
	}

	// Check if the first file in the ZIP file has a .txt extension.
	firstFile := zipReader.File[0]
	if filepath.Ext(firstFile.Name) != ".txt" {
		//http.Error(w, "First file in ZIP file does not have a .txt extension", http.StatusBadRequest)
		//	fmt.Println("3")
		//	return false
	}

	/*	// Open the first file in the ZIP file.
		fileReader, err := firstFile.Open()
		if err != nil {
			//http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println("5")
			//	return false
		}
		defer fileReader.Close()*/
	return true
}

func SaveFile(fileHeader *multipart.FileHeader) (bool, string) {
	file, err := fileHeader.Open()
	defer file.Close()
	uploadTimeStamp := time.Now()
	year, month, day := uploadTimeStamp.Date()

	outputPath := fmt.Sprintf("statics/saves/%d/%d/%d/", year, month, day)
	//create directory
	err = os.MkdirAll(filepath.Dir(outputPath), 0755)
	if err != nil {
		fmt.Println("Failed to create directory:", err)
		return false, ""
	}

	// Create the output file.
	outputFile, err := os.Create(outputPath + strconv.Itoa(int(uploadTimeStamp.Unix())) + "_" + fileHeader.Filename)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err.Error())
		return false, ""
	}
	defer outputFile.Close()

	// Copy the contents of the first file to the output file.
	_, err = io.Copy(outputFile, file)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("6")
		return false, ""
	}
	return true, outputFile.Name()
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}
