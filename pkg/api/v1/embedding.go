package v1

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quantumwake/alethic-ism-core-go/pkg/repository/embedding"
)

// EmbeddingHandler holds the embedding storage backend.
type EmbeddingHandler struct {
	store embedding.Storage
}

// NewEmbeddingHandler creates a handler backed by pgvector storage.
func NewEmbeddingHandler(dsn string) *EmbeddingHandler {
	pgStore := embedding.NewPgVectorStorage(dsn)
	_ = embedding.EnsureTable(pgStore.DB)
	return &EmbeddingHandler{store: pgStore}
}

// SearchRequest is the API request body for similarity search.
type SearchRequest struct {
	Embedding []float32              `json:"embedding" binding:"required"`
	Filter    SearchFilterRequest    `json:"filter" binding:"required"`
	Options   SearchOptionsRequest   `json:"options"`
}

// SearchFilterRequest is the JSON-friendly search filter.
type SearchFilterRequest struct {
	UserID    string              `json:"user_id" binding:"required"`
	ProjectID *string             `json:"project_id,omitempty"`
	SessionID *string             `json:"session_id,omitempty"`
	ScopeType *embedding.ScopeType `json:"scope_type,omitempty"`
	ScopeID   *string             `json:"scope_id,omitempty"`
}

// SearchOptionsRequest is the JSON-friendly search options.
type SearchOptionsRequest struct {
	Limit         int      `json:"limit,omitempty"`
	MinSimilarity *float64 `json:"min_similarity,omitempty"`
}

// HandleUpsert
// @Summary Upsert an embedding document
// @Description Insert or update a single embedding document
// @Tags Embedding
// @Accept  json
// @Produce  json
// @Param document body embedding.Document true "Embedding document"
// @Success 200 {object} embedding.Document
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /v1/nlp/embeddings [post]
func (h *EmbeddingHandler) HandleUpsert(c *gin.Context) {
	var doc embedding.Document
	if err := c.ShouldBindJSON(&doc); err != nil {
		log.Println("embedding upsert: invalid input:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	if err := h.store.Upsert(&doc); err != nil {
		log.Println("embedding upsert failed:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.JSON(http.StatusOK, doc)
}

// HandleUpsertBatch
// @Summary Batch upsert embedding documents
// @Description Insert or update multiple embedding documents in a single transaction
// @Tags Embedding
// @Accept  json
// @Produce  json
// @Param documents body []embedding.Document true "List of embedding documents"
// @Success 200 {object} map[string]int
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /v1/nlp/embeddings/batch [post]
func (h *EmbeddingHandler) HandleUpsertBatch(c *gin.Context) {
	var docs []*embedding.Document
	if err := c.ShouldBindJSON(&docs); err != nil {
		log.Println("embedding upsert batch: invalid input:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	if err := h.store.UpsertBatch(docs); err != nil {
		log.Println("embedding upsert batch failed:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"count": len(docs)})
}

// HandleFindByID
// @Summary Find embedding document by ID
// @Description Retrieve a single embedding document by its ID
// @Tags Embedding
// @Produce  json
// @Param id path string true "Document ID"
// @Success 200 {object} embedding.Document
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /v1/nlp/embeddings/{id} [get]
func (h *EmbeddingHandler) HandleFindByID(c *gin.Context) {
	id := c.Param("id")

	doc, err := h.store.FindByID(id)
	if err != nil {
		log.Println("embedding find by id failed:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, doc)
}

// HandleFindByParentID
// @Summary Find embedding documents by parent ID
// @Description Retrieve all embedding documents with the given parent ID
// @Tags Embedding
// @Produce  json
// @Param id path string true "Parent ID"
// @Success 200 {array} embedding.Document
// @Failure 500 {object} model.ErrorResponse
// @Router /v1/nlp/embeddings/parent/{id} [get]
func (h *EmbeddingHandler) HandleFindByParentID(c *gin.Context) {
	parentID := c.Param("id")

	docs, err := h.store.FindByParentID(parentID)
	if err != nil {
		log.Println("embedding find by parent id failed:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.JSON(http.StatusOK, docs)
}

// HandleDelete
// @Summary Delete embedding document by ID
// @Description Delete a single embedding document by its ID
// @Tags Embedding
// @Produce  json
// @Param id path string true "Document ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} model.ErrorResponse
// @Router /v1/nlp/embeddings/{id} [delete]
func (h *EmbeddingHandler) HandleDelete(c *gin.Context) {
	id := c.Param("id")

	if err := h.store.Delete(id); err != nil {
		log.Println("embedding delete failed:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

// HandleDeleteByParentID
// @Summary Delete embedding documents by parent ID
// @Description Delete all embedding documents with the given parent ID
// @Tags Embedding
// @Produce  json
// @Param id path string true "Parent ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} model.ErrorResponse
// @Router /v1/nlp/embeddings/parent/{id} [delete]
func (h *EmbeddingHandler) HandleDeleteByParentID(c *gin.Context) {
	parentID := c.Param("id")

	if err := h.store.DeleteByParentID(parentID); err != nil {
		log.Println("embedding delete by parent id failed:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

// HandleSearch
// @Summary Search embedding documents by similarity
// @Description Search for similar documents using cosine similarity
// @Tags Embedding
// @Accept  json
// @Produce  json
// @Param search body SearchRequest true "Search parameters"
// @Success 200 {array} embedding.SearchResult
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /v1/nlp/embeddings/search [post]
func (h *EmbeddingHandler) HandleSearch(c *gin.Context) {
	var req SearchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("embedding search: invalid input:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	filter := embedding.SearchFilter{
		UserID:    req.Filter.UserID,
		ProjectID: req.Filter.ProjectID,
		SessionID: req.Filter.SessionID,
		ScopeType: req.Filter.ScopeType,
		ScopeID:   req.Filter.ScopeID,
	}

	opts := embedding.SearchOptions{
		Limit:         req.Options.Limit,
		MinSimilarity: req.Options.MinSimilarity,
	}

	results, err := h.store.Search(req.Embedding, filter, opts)
	if err != nil {
		log.Println("embedding search failed:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.JSON(http.StatusOK, results)
}

// HandleMigrate
// @Summary Run embedding table migration
// @Description Create the embedding_document table and indexes if they do not exist
// @Tags Embedding
// @Produce  json
// @Success 200 {object} map[string]string
// @Failure 500 {object} model.ErrorResponse
// @Router /v1/nlp/embeddings/migrate [post]
func (h *EmbeddingHandler) HandleMigrate(c *gin.Context) {
	pgStore, ok := h.store.(*embedding.PgVectorStorage)
	if !ok {
		// If using cached storage, try to get the underlying pgvector store
		c.JSON(http.StatusInternalServerError, gin.H{"error": "migration not supported for this storage backend"})
		return
	}

	if err := embedding.EnsureTable(pgStore.DB); err != nil {
		log.Println("embedding migration failed:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "migration failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "migrated"})
}
