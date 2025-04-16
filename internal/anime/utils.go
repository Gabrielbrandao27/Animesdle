package anime

import (
	"fmt"
	"strings"
)

func compareStrings(expected, actual string) FieldComparison {
	expectedSlice := splitAndTrim(expected)
	actualSlice := splitAndTrim(actual)

	if slicesEqual(expectedSlice, actualSlice) {
		return FieldComparison{Value: actual, Status: "correct"}
	}

	if hasIntersection(expectedSlice, actualSlice) {
		return FieldComparison{Value: actual, Status: "partial"}
	}

	return FieldComparison{Value: actual, Status: "wrong"}
}

func splitAndTrim(s string) []string {
	parts := strings.Split(s, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}

func slicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	m := make(map[string]bool)
	for _, v := range a {
		m[v] = true
	}
	for _, v := range b {
		if !m[v] {
			return false
		}
	}
	return true
}

func hasIntersection(a, b []string) bool {
	set := make(map[string]struct{})
	for _, v := range a {
		set[v] = struct{}{}
	}
	for _, v := range b {
		if _, found := set[v]; found {
			return true
		}
	}
	return false
}

func compareInts(expected, actual int) FieldComparison {
	switch {
	case expected == actual:
		return FieldComparison{Value: fmt.Sprintf("%d", actual), Status: "correct"}
	case actual < expected:
		return FieldComparison{Value: fmt.Sprintf("%d", actual), Status: "less"}
	default:
		return FieldComparison{Value: fmt.Sprintf("%d", actual), Status: "greater"}
	}
}
