package models

// Items holds the pack sizes
type Items struct {
	name      string
	packSizes []int
	price     float64
}

// NewItems creates a new Items struct with initialized pack sizes
func NewItems(itemName string, price float64) *Items {
	return &Items{
		name:      itemName,
		packSizes: []int{5000, 2000, 1000, 500, 250},
		price:     price,
	}
}

// Function to calculate the packs required
func (i *Items) calculatePacks(orderSize int) map[int]int {
	packs := make(map[int]int)
	for _, size := range i.packSizes {
		if orderSize == 0 {
			break
		}
		if orderSize >= size {
			packs[size] = orderSize / size
			orderSize = orderSize % size
		}
	}
	if orderSize > 0 {
		packs[250]++
	}
	return packs
}
