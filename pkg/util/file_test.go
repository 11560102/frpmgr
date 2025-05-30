package util

import (
	"os"
	"reflect"
	"testing"
	"time"
)

func TestSplitExt(t *testing.T) {
	tests := []struct {
		input        string
		expectedName string
		expectedExt  string
	}{
		{input: "C:\\test\\a.ini", expectedName: "a", expectedExt: ".ini"},
		{input: "b.exe", expectedName: "b", expectedExt: ".exe"},
		{input: "c", expectedName: "c", expectedExt: ""},
		{input: "", expectedName: "", expectedExt: ""},
	}
	for i, test := range tests {
		name, ext := SplitExt(test.input)
		if name != test.expectedName {
			t.Errorf("Test %d: expected: %v, got: %v", i, test.expectedName, name)
		}
		if ext != test.expectedExt {
			t.Errorf("Test %d: expected: %v, got: %v", i, test.expectedExt, ext)
		}
	}
}

func TestFindLogFiles(t *testing.T) {
	tests := []struct {
		create        []string
		expectedFiles []string
		expectedDates []time.Time
	}{
		{
			create:        []string{"example.log", "example.20230320-000000.log", "example.20230321-010203.log", "example.2023-03-21.log"},
			expectedFiles: []string{"example.log", "example.20230320-000000.log", "example.20230321-010203.log"},
			expectedDates: []time.Time{
				{},
				time.Date(2023, 3, 20, 0, 0, 0, 0, time.Local),
				time.Date(2023, 3, 21, 1, 2, 3, 0, time.Local),
			},
		},
	}
	if err := os.MkdirAll("testdata", 0750); err != nil {
		t.Fatal(err)
	}
	os.Chdir("testdata")
	for i, test := range tests {
		for _, f := range test.create {
			os.WriteFile(f, []byte("test"), 0666)
		}
		logs, dates, err := FindLogFiles(test.create[0])
		if err != nil {
			t.Error(err)
			continue
		}
		if !reflect.DeepEqual(logs, test.expectedFiles) {
			t.Errorf("Test %d: expected: %v, got: %v", i, test.expectedFiles, logs)
		}
		if !reflect.DeepEqual(dates, test.expectedDates) {
			t.Errorf("Test %d: expected: %v, got: %v", i, test.expectedDates, dates)
		}
	}
}
