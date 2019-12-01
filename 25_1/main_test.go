package main

import (
	"reflect"
	"testing"
)

//func TestReadfile(t *testing.T) {
//	testTable := map[string]struct {
//		File string
//		Out  []Star
//	}{
//		"test1": {
//			File: "golden/test1.txt",
//			Out: []Star{
//				Star{x: 0, y: 0, z: 0, z1: 0},
//				Star{x: 3, y: 0, z: 0, z1: 0},
//				Star{x: 0, y: 3, z: 0, z1: 0},
//				Star{x: 0, y: 0, z: 3, z1: 0},
//				Star{x: 0, y: 0, z: 0, z1: 3},
//				Star{x: 0, y: 0, z: 0, z1: 6},
//				Star{x: 9, y: 0, z: 0, z1: 0},
//				Star{x: 12, y: 0, z: 0, z1: 0},
//			},
//		},
//	}
//
//	for testName, test := range testTable {
//		testRes, _ := readFile(test.File)
//		if !reflect.DeepEqual(testRes, test.Out) {
//			t.Fatalf("%s: want %#v, got %#v", testName, test.Out, testRes)
//		}
//	}
//}

//func TestDistance(t *testing.T) {
//	testTable := map[string]struct {
//		Star1    Star
//		Star2    Star
//		Distance float64
//	}{
//		"test1": {
//			Star1:    Star{0, 0, 0, 0},
//			Star2:    Star{0, 0, 0, 1},
//			Distance: 1,
//		},
//	}
//
//	for testName, test := range testTable {
//		testRes := starsDistance(test.Star1, test.Star2)
//		if !reflect.DeepEqual(testRes, test.Distance) {
//			t.Fatalf("%s: want %#v, got %#v", testName, test.Distance, testRes)
//		}
//	}
//
//}

func TestCountConstellations(t *testing.T) {
	testTable := map[string]struct {
		File                string
		ConstellationsCount int
	}{
		"test1": {
			File:                "golden/test1.txt",
			ConstellationsCount: 2,
		},
		"test2": {
			File:                "golden/test2.txt",
			ConstellationsCount: 4,
		},
		"test3": {
			File:                "golden/test3.txt",
			ConstellationsCount: 3,
		},
		"test4": {
			File:                "golden/test4.txt",
			ConstellationsCount: 8,
		},
		"task": {
			File:                "golden/task.txt",
			ConstellationsCount: 100,
		},
	}

	for testName, test := range testTable {
		_, testCount, _ := readFile(test.File)
		if !reflect.DeepEqual(testCount, test.ConstellationsCount) {
			t.Fatalf("%s: want %#v, got %#v", testName, test.ConstellationsCount, testCount)
		}
	}
}
