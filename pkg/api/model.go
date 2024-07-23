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

	// Combine smaller packs into larger ones if beneficial
	// We iterate through pack sizes from smallest to largest
	for j := len(i.packSizes) - 1; j > 0; j-- {
		// Determine the current and next larger pack size
		currentPack := i.packSizes[j]
		nextPack := i.packSizes[j-1]

		// Check if we have multiple smaller packs that can be combined into a larger pack
		if count, exists := packs[currentPack]; exists && count > 1 {
			// Calculate how many larger packs can be formed from the smaller packs
			largerPackCount := (count * currentPack) / nextPack
			if (count*currentPack)%nextPack == 0 {
				// Update the map with the count of the larger packs
				packs[nextPack] = largerPackCount
				// Remove the smaller packs from the map
				delete(packs, currentPack)
			}
		}
	}

	// Return the map with the count of each pack size
	return packs
}

type getOrderedPacksRequest struct {
	// add binding for at least one item requested
	OrderSize int64 `uri:"ordersize" binding:"required,min=1"`
}

func (server *Server) getOrderedPacks(ctx *gin.Context) {
	var req getOrderedPacksRequest
	// Get the order size from the URL parameter
	err := ctx.ShouldBindUri(&req)
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
	packs := items.calculatePacks(int(req.OrderSize))

	// Create the response
	response := gin.H{}
	for k, v := range packs {
		response[strconv.Itoa(k)] = v
	}

	// Return the response
	ctx.JSON(http.StatusOK, response)
}
