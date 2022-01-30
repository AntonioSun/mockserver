package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseMockJson(t *testing.T) {

	/*
		Happy path
	*/
	if respMap != nil {
		t.Errorf("respMap should be empty")
	}
	err := ParseMockJson("mock.json")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if respMap == nil {
		t.Errorf("respMap should not be empty")
	}

	/*
		Wrong file
	*/
	// err = ParseMockJson("haha.json")
	// if err == nil {
	// 	t.Errorf("expected error but nil")
	// }
	// if err.Error() != "open haha.json: no such file or directory" {
	// 	t.Errorf("unexpected error %v", err)
	// }

	/*
		Invalid json file
	*/
	err = ParseMockJson("tests/wrong_json.json")
	if err == nil {
		t.Errorf("expected error but nil")
	}
	if err.Error() != "invalid character 'I' looking for beginning of value" {
		t.Errorf("unexpected error %v", err)
	}

}

func TestVerifyMockJson(t *testing.T) {

	/*
		Happy path
	*/
	err := VerifyMockJson()
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}


	/*
		Missing status code
	*/
	// err = ParseMockJson("tests/no_statuscode.json")
	// err = VerifyMockJson()
	// if err == nil {
	// 	t.Errorf("expected error but nil")
	// }
	// if err.Error() != "Missing `statusCode` field in /hello/worlds" {
	// 	t.Errorf("unexpected error %v", err)
	// }


	/*
		Missing response body
	*/
	// err = ParseMockJson("tests/no_responsebody.json")
	// err = VerifyMockJson()
	// if err == nil {
	// 	t.Errorf("expected error but nil")
	// }
	// if err.Error() != "Missing `responseBody` field in /hello/worlds" {
	// }
}
