package api

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Item holds the pack sizes
type Item struct {
	name      string
	packSizes []int
	price     float64
}

// NewItem creates a new Item struct with initialized pack sizes
func NewItem(itemName string, price float64) *Item {
	return &Item{
		name:      itemName,
		packSizes: []int{5000, 2000, 1000, 500, 250},
		price:     price,
	}
}

// Function to ensure packSizes is unique and sorted in descending order
func (i *Item) preparePackSizes() {
	// Remove duplicates
	uniquePackSizes := make(map[int]struct{})
	for _, size := range i.packSizes {
		uniquePackSizes[size] = struct{}{}
	}

	// Convert map keys to a slice
	distinctSizes := make([]int, 0, len(uniquePackSizes))
	for size := range uniquePackSizes {
		distinctSizes = append(distinctSizes, size)
	}

	// Sort slice in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(distinctSizes)))

	// Assign the cleaned and sorted slice back to packSizes
	i.packSizes = distinctSizes
}

// Function to calculate the packs required
func (i *Item) calculatePacks(orderSize int) map[int]int {
	// Ensure packSizes is unique and sorted
	i.preparePackSizes()

	// fmt.Println("Pack sizes:", i.packSizes)

	// Initialize a map to store the count of each pack size
	packs := make(map[int]int)
	// Copy of the orderSize to keep track of the remaining item
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
		// Increment the smallest item in the map (250) by 1
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

// removePackSize removes the first occurrence of size from the packSizes slice.
func (i *Item) removePackSize(size int) {
	for index, packSize := range i.packSizes {
		if packSize == size {
			// Remove the element by slicing around it
			i.packSizes = append(i.packSizes[:index], i.packSizes[index+1:]...)
			return
		}
	}
}

// view packsizes for the item
func (i Item) viewPackSizes() []int {
	return i.packSizes
}

// Add a new pack size to the item
func (i *Item) addPackSize(size int) {

	// Ensure the pack size is positive
	if size <= 0 {
		return
	}

	// Add the pack size if it's not already present
	for _, packSize := range i.packSizes {
		if packSize == size {
			// Pack size already exists
			return
		}
	}
	i.packSizes = append(i.packSizes, size)
}

// Handler Functions

// @Summary Get pack sizes
// @Description Get the available pack sizes for the item
// @Tags packs
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /pack-sizes [get]
func (server *Server) getPackSizesHandler(ctx *gin.Context) {
	packSizes := server.item.viewPackSizes()
	response := gin.H{
		"packSizes": packSizes,
	}

	ctx.JSON(http.StatusOK, response)
}

type getOrderedPacksRequest struct {
	// add binding for at least one item requested
	OrderSize int64 `uri:"ordersize" binding:"required,min=1"`
}

// @Summary Calculate packs
// @Description Calculate the required packs for a given order size
// @Tags packs
// @Produce json
// @Param ordersize path int true "Order Size"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /calculate-packs/{ordersize} [get]
func (server *Server) calculatePacksHandler(ctx *gin.Context) {
	var req getOrderedPacksRequest

	// Get the order size from the URL parameter
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Calculate the packs
	packs := server.item.calculatePacks(int(req.OrderSize))

	// Create the response struct
	response := gin.H{
		"name":  server.item.name,
		"price": server.item.price,
		"packs": formatPacks(packs),
	}

	// Return the response
	ctx.JSON(http.StatusOK, response)
}

// formatPacks formats the packs map into a suitable structure for JSON response
func formatPacks(packs map[int]int) gin.H {
	formattedPacks := gin.H{}
	for k, v := range packs {
		formattedPacks[strconv.Itoa(k)] = v
	}
	return formattedPacks
}

// @Summary Remove pack size
// @Description Remove a pack size from the item
// @Tags packs
// @Produce json
// @Param packsize query int true "Pack Size"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /remove-pack-size [delete]
func (server *Server) removePackSizeHandler(ctx *gin.Context) {
	// Extract the pack size from the query parameter
	packSizeStr := ctx.DefaultQuery("packsize", "")
	if packSizeStr == "" {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("packsize parameter is required")))
		return
	}

	// Convert the pack size to an integer
	packSize, err := strconv.Atoi(packSizeStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Ensure the pack size is positive
	if packSize <= 0 {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("packsize must be a positive integer")))
		return
	}

	// Remove the pack size from the item
	server.item.removePackSize(packSize)

	// Return a success response
	ctx.JSON(http.StatusOK, gin.H{"message": "Pack size removed successfully"})
}

// @Summary Add pack size
// @Description Add a new pack size to the item
// @Tags packs
// @Produce json
// @Param packsize query int true "Pack Size"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /add-pack-size [post]
func (server *Server) addPackSizeHandler(ctx *gin.Context) {
	// Extract the pack size from the query parameter
	packSizeStr := ctx.DefaultQuery("packsize", "")
	if packSizeStr == "" {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("packsize parameter is required")))
		return
	}

	// Convert the pack size to an integer
	packSize, err := strconv.Atoi(packSizeStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Ensure the pack size is positive
	if packSize <= 0 {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("packsize must be a positive integer")))
		return
	}

	// Add the pack size to the item
	server.item.addPackSize(packSize)

	// Return a success response
	ctx.JSON(http.StatusOK, gin.H{"message": "Pack size added successfully"})
}
