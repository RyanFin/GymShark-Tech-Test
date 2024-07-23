package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
	// Initialize a map to store the count of each pack size
	packs := make(map[int]int)
	// Copy of the orderSize to keep track of the remaining items
	remainingOrder := orderSize

	// Iterate over each pack size in descending order
	for _, size := range i.packSizes {
		// If there are no remaining items to pack, exit the loop
		if remainingOrder == 0 {
			break
		}
		// If the remaining order is larger than or equal to the current pack size
		if remainingOrder >= size {
			// Calculate how many packs of the current size are needed
			packs[size] = remainingOrder / size
			// Update the remaining order with the remainder
			remainingOrder = remainingOrder % size
		}
	}

	// If there are any remaining items, add the smallest available pack (250)
	if remainingOrder > 0 {
		// increment the smallest item in the map (250) by 1
		packs[250]++
	}

	// Check if it's beneficial to combine smaller packs into a larger one
	for size := range packs {
		// Skip pack sizes larger than 250 as we're focusing on combining smaller packs
		if size > 250 {
			continue
		}
		// Iterate over pack sizes from the smallest to largest (excluding the largest)
		for j := len(i.packSizes) - 1; j > 0; j-- {
			// If the current pack size matches and we have more than one of these packs
			if size == i.packSizes[j] && packs[size] > 1 {
				// Determine the next larger pack size
				largerPack := i.packSizes[j-1]
				// Calculate how many larger packs would be equivalent to the smaller packs
				largerPackCount := (packs[size] * size) / largerPack
				// If the smaller packs can be perfectly combined into larger packs
				if (packs[size]*size)%largerPack == 0 {
					// Update the map with the count of the larger packs
					packs[largerPack] = largerPackCount
					// Remove the smaller packs from the map
					delete(packs, size)
				}
			}
		}
	}

	// Return the map with the count of each pack size
	return packs
}

func (server *Server) getOrderedPacks(ctx *gin.Context) {
	// Get the order size from the URL parameter
	orderSizeStr := ctx.Param("orderSize")
	orderSize, err := strconv.Atoi(orderSizeStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Get name and price from query parameters
	itemName := ctx.Query("name")
	priceStr := ctx.Query("price")
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Create an Items object with the provided name and price
	items := NewItems(itemName, price)

	// Calculate the packs
	packs := items.calculatePacks(orderSize)

	// Create the response
	response := gin.H{}
	for k, v := range packs {
		response[strconv.Itoa(k)] = v
	}

	// Return the response
	ctx.JSON(http.StatusOK, response)
}
