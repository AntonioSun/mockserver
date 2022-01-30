package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
)

var respMap RespMap

func ParseMockJson(file string) error {
	jsonFile, err := os.Open(filepath.Clean(file))
	if err != nil {
		return err
	}
	fmt.Println("✔ Successfully opened:", file)

	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(jsonFile)

	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &respMap)
	if err != nil {
		return err
	}
	fmt.Println("✔ Successfully parsed:", file)
	//fmt.Printf("] %+v\n", respMap)

	return nil
}

func VerifyMockJson() error {
	return nil
}

func checkFieldPresent(i interface{}, key, fName string) error {
	dict := reflect.ValueOf(i)
	val := dict.MapIndex(reflect.ValueOf(fName))
	if val == reflect.ValueOf(nil) {
		return errors.New(fmt.Sprintf("Missing `%v` field in %v", fName, key))
	}
	return nil
}

func printPaths() {
	if e.Verbose < 1 {
		return
	}
	fmt.Println("Available paths: ")
	for _, k := range respMap {
		fmt.Println("=>", k.HTTPRequest.Path)
	}
}
