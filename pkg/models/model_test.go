package models

import (
	"reflect"
	"testing"
)

func TestCalculatePacks(t *testing.T) {
	tests := []struct {
		name      string
		orderSize int
		expected  map[int]int
	}{
		{
			name:      "Order a single item",
			orderSize: 1,
			expected: map[int]int{
				250: 1,
			},
		},
		{
			name:      "Order size equals the smallest pack size",
			orderSize: 250,
			expected: map[int]int{
				250: 1,
			},
		},
		{
			name:      "Order size just higher than the smallest pack size",
			orderSize: 251,
			expected: map[int]int{
				500: 1,
			},
		},
		{
			name:      "Order size just higher than the second smallest pack size",
			orderSize: 501,
			expected: map[int]int{
				500: 1,
				250: 1,
			},
		},
		{
			name:      "Order a large amount",
			orderSize: 12001,
			expected: map[int]int{
				5000: 2,
				2000: 1,
				250:  1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item := NewItems("TestItem", 0)
			got := item.calculatePacks(tt.orderSize)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("calculatePacks(%d) = %v; want %v", tt.orderSize, got, tt.expected)
			}
		})
	}
}
