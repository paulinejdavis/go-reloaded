package main 

import ( 
// "os" 
"testing" 
) 

func TestStringremove(t *testing.T) { 
slice := []string{"foo", "bar", "baz"} 
expectedSlice := []string{"foo", "baz"} 

result := stringremove(slice, 1) 

if len(result) != len(expectedSlice) { 
t.Errorf("expected slice: %v, got slice: %v", expectedSlice, result) 
} 

for i := range expectedSlice { 
if expectedSlice[i] != result[i] { 
t.Errorf("expected slice: %v, got slice: %v", expectedSlice, result) 
} 
} 
} 

func TestRuneremove(t *testing.T) { 
slice := []rune{'f', 'o', 'o', 'b', 'a', 'r', 'b', 'a', 'z'} 
expectedSlice := []rune{'f', 'o', 'o', 'b', 'a', 'r', 'a', 'z'} 

result := runeremove(slice, 6) 

if len(result) != len(expectedSlice) { 
t.Errorf("expected slice: %v, got slice: %v", expectedSlice, result) 
} 

for i := range expectedSlice { 
if expectedSlice[i] != result[i] { 
t.Errorf("expected slice: %v, got slice: %v", expectedSlice, result) 
} 
} 
} 

func TestTrimatoi(t *testing.T) { 
tests := []struct { 
input    string 
expected int 
}{ 
{"123", 123}, 
{"456", 456}, 
{"789", 789}, 
} 

for _, test := range tests { 
result := trimatoi(test.input) 
if result != test.expected { 
t.Errorf("expected: %d, got: %d", test.expected, result) 
} 
} 
} 

func TestGetAlpha(t *testing.T) { 
tests := []struct { 
input    []rune 
expected string 
}{ 
{[]rune{'f', 'o', 'o'}, "foo"}, 
{[]rune{'B', 'A', 'R'}, "BAR"}, 
{[]rune{'B', 'a', 'Z'}, "BaZ"}, 
} 

for _, test := range tests { 
result := getAlpha(test.input) 
if result != test.expected { 
t.Errorf("expected: %s, got: %s", test.expected, result) 
} 
} 
} 

func TestHexaNumberToInteger(t *testing.T) { 
tests := []struct { 
input    string 
expected string 
}{ 
{"A", "10"}, 
{"B", "11"}, 
{"C", "12"}, 
} 

for _, test := range tests { 
result := hexaNumberToInteger(test.input) 
if result != test.expected { 
t.Errorf("expected: %s, got: %s", test.expected, result) 
} 
} 
} 

func TestBintodecimal(t *testing.T) { 
tests := []struct { 
input    string 
expected string 
}{ 
{"1010", "10"}, 
{"1101", "13"}, 
{"1110", "14"}, 
} 

for _, test := range tests { 
result := bintodecimal(test.input) 
if result != test.expected { 
t.Errorf("expected: %s, got: %s", test.expected, result) 
} 
} 
}