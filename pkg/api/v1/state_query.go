package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/quantumwake/alethic-ism-core-go/pkg/repository/query"
	"github.com/quantumwake/alethic-ism-core-go/pkg/repository/query/dsl"
	"log"
	"net/http"
	"os"
)

var (
	DSN = os.Getenv("DSN")
)

// HandleQueryState
// @Summary Query state data
// @Description Query state data with filters
// @Tags State
// @Accept  json
// @Produce  json
// @Param id path string true "State ID"
// @Param dsl.StateQuery body dsl.StateQuery true "the query groups and filters"
// @Success 200 {array} dsl.StateQueryResult
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /state/query/{id} [post]
func HandleQueryState(c *gin.Context) {
	stateID := c.Param("id")
	log.Println("querying state data for stateID: ", stateID)

	// Parse JSON body for filter criteria into StateQuery
	var dql dsl.StateQuery
	//dql.UserID = stateID // temporary hardcoding to TODO: get from JWT
	if err := c.ShouldBindJSON(&dql); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	log.Println("querying state data for stateID: ", stateID, " with query: ", dql)
	dataAccess := query.NewBackend(DSN)

	// Execute the query
	results, err := dataAccess.Query(stateID, dql)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	log.Println("querying state data for stateID: ", stateID, " with query: ", dql, " returned results: ", results)

	// Return the results
	c.JSON(http.StatusOK, results)
}
