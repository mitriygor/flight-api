package main

import (
	"reflect"
	"testing"
)

func Test_GetItinerary(t *testing.T) {
	cases := []struct {
		data     [][]string
		expected [][]string
	}{
		{[][]string{[]string{"IND", "EWR"}, []string{"SFO", "ATL"}, []string{"GSO", "IND"}, []string{"ATL", "GSO"}},
			[][]string{[]string{"SFO", "ATL"}, []string{"ATL", "GSO"}, []string{"GSO", "IND"}, []string{"IND", "EWR"}}},
	}
	for _, c := range cases {
		result := GetItinerary(c.data)

		if !reflect.DeepEqual(result, c.expected) {
			t.Errorf("Expected %v; nonetheless, but got %v", c.expected, result)
		}
	}
}

func Test_GetTransfers(t *testing.T) {
	cases := []struct {
		data     []string
		expected [][]string
	}{
		{[]string{"IND", "EWR", "SFO", "ATL", "GSO", "IND", "ATL", "GSO"}, [][]string{[]string{"IND", "EWR"}, []string{"SFO", "ATL"}, []string{"GSO", "IND"}, []string{"ATL", "GSO"}}},
	}
	for _, c := range cases {
		result := GetTransfers(c.data)

		if !reflect.DeepEqual(result, c.expected) {
			t.Errorf("Expected %v; nonetheless, but got %v", c.expected, result)
		}
	}
}

func Test_GetCount(t *testing.T) {
	cases := []struct {
		data     [][]string
		expected int
	}{
		{[][]string{[]string{"IND", "EWR"}, []string{"SFO", "ATL"}, []string{"GSO", "IND"}, []string{"ATL", "GSO"}}, 4},
	}
	for _, c := range cases {
		result := GetCount(c.data)

		if result != c.expected {
			t.Errorf("Expected %v; nonetheless, but got %v", c.expected, result)
		}
	}
}

func Test_GetDestination(t *testing.T) {
	cases := []struct {
		data     [][]string
		expected []string
	}{
		{[][]string{[]string{"IND", "EWR"}, []string{"SFO", "ATL"}, []string{"GSO", "IND"}, []string{"ATL", "GSO"}}, []string{"SFO", "EWR"}},
	}
	for _, c := range cases {
		var reversedResult []string
		result := GetDestination(c.data)

		for i := len(result) - 1; i >= 0; i-- {
			reversedResult = append(reversedResult, result[i])
		}

		if !(reflect.DeepEqual(result, c.expected) || reflect.DeepEqual(reversedResult, c.expected)) {
			t.Errorf("Expected %v; nonetheless, got %v", c.expected, result)
		}
	}
}
