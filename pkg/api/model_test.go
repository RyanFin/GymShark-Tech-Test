package api

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
			item := NewItem("TestItem", 0)
			got := item.calculatePacks(tt.orderSize)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("calculatePacks(%d) = %v; want %v", tt.orderSize, got, tt.expected)
			}
		})
	}
}

func TestRemovePackSize(t *testing.T) {
	tests := []struct {
		name         string
		initial      []int
		sizeToRemove int
		expected     []int
	}{
		{
			name:         "Remove existing pack size",
			initial:      []int{5000, 2000, 1000, 500, 250},
			sizeToRemove: 2000,
			expected:     []int{5000, 1000, 500, 250},
		},
		{
			name:         "Remove non-existing pack size",
			initial:      []int{5000, 2000, 1000, 500, 250},
			sizeToRemove: 50,
			expected:     []int{5000, 2000, 1000, 500, 250},
		},
		{
			name:         "Remove size that is at the beginning",
			initial:      []int{5000, 2000, 1000, 500, 250},
			sizeToRemove: 5000,
			expected:     []int{2000, 1000, 500, 250},
		},
		{
			name:         "Remove size that is at the end",
			initial:      []int{5000, 2000, 1000, 500, 250},
			sizeToRemove: 250,
			expected:     []int{5000, 2000, 1000, 500},
		},
		{
			name:         "Remove size from an empty slice",
			initial:      []int{},
			sizeToRemove: 10,
			expected:     []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item := Item{
				name:      "TestItem",
				packSizes: tt.initial,
				price:     10.0,
			}
			item.removePackSize(tt.sizeToRemove)

			for i, v := range item.packSizes {
				if v != tt.expected[i] {
					t.Errorf("For %s, expected %v but got %v", tt.name, tt.expected, item.packSizes)
					break
				}
			}
		})
	}
}
