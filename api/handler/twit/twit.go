package twit

import (
	"fmt"
	"milliy/api/auth"
	"milliy/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateTwit godoc
// @Security ApiKeyAuth
// @Summary Create Twit
// @Description it will Create Twit
// @Tags TWIT API
// @Param info body model.CreateTwitRequestApi true "info"
// @Success 200 {object} string "twit_id"
// @Failure 400 {object} string "Invalid data"
// @Failure 401 {object} string "Invalid token"
// @Failure 500 {object} string "Server error"
// @Router /v1/twit [post]
func (h *newTwits) CreateTwit(c *gin.Context) {
	h.Log.Info("CreateTwit called")
	var req model.CreateTwitRequestApi
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Log.Error("Invalid request", "error", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	token := c.GetHeader("Authorization")
	userId, _, err := auth.GetUserInfoFromRefreshToken(token)
	if err != nil {
		h.Log.Error("Invalid token", "error", err)
		c.JSON(401, gin.H{"error": "Invalid token"})
		return
	}

	twitId, err := h.Twit.CreateTwit(&model.CreateTwitRequest{
		UserID:       userId,
		PublisherFIO: req.PublisherFIO,
		Type:         req.Type,
		Texts:        req.Texts,
		Title:        req.Title,
	})
	if err != nil {
		h.Log.Error("Error creating twit", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	h.Log.Info("Twit created successfully")
	c.JSON(200, gin.H{"twit_id": twitId})
}

// GetTwit godoc
// @Summary Get Twit
// @Description it will Get Twit By id
// @Tags TWIT API
// @Param id path string true "id"
// @Success 200 {object} model.TwitResponse
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /v1/twit/{id} [get]
func (h *newTwits) GetTwit(c *gin.Context) {
	h.Log.Info("GetTwit called")
	id := c.Param("id")
	if id == "" {
		h.Log.Error("Invalid id")
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	twit, err := h.Twit.GetTwit(id)
	if err != nil {
		h.Log.Error("Error getting twit", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	h.Log.Info("Twit retrieved successfully")
	c.JSON(200, twit)
}

// DeleteTwit godoc
// @Security ApiKeyAuth
// @Summary Delete Twit
// @Description it will Delete Twit
// @Tags TWIT API
// @Param id path string true "id"
// @Success 200 {object} string "message"
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /v1/twit/{id} [delete]
func (h *newTwits) DeleteTwit(c *gin.Context) {
	h.Log.Info("DeleteTwit called")
	id := c.Param("id")
	if id == "" {
		h.Log.Error("Invalid id")
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}
	err := h.Twit.DeleteTwit(id)
	if err != nil {
		h.Log.Error("Error deleting twit", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	h.Log.Info("Twit deleted successfully")
	c.JSON(200, gin.H{"message": "Twit deleted"})
}

// AddReadersCount godoc
// @Summary Add view to Twit
// @Description it will add views Twit
// @Tags TWIT API
// @Param id path string true "id"
// @Success 200 {object} string "message"
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /v1/twit/{id} [post]
func (h *newTwits) AddReadersCount(c *gin.Context) {
	h.Log.Info("AddReadersCount called")
	id := c.Param("id")
	if id == "" {
		h.Log.Error("Invalid id")
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}
	err := h.Twit.AddReadersCount(id)
	if err != nil {
		h.Log.Error("Error adding readers count", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	h.Log.Info("Readers count added successfully")
	c.JSON(200, gin.H{"message": "Readers count added"})
}

// GetAllTwits godoc
// @Summary get all Twits
// @Description it will get all Twits
// @Tags TWIT API
// @Success 200 {object} []string
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /v1/twit/all [get]
func (h *newTwits) GetAllTwits(c *gin.Context) {
	h.Log.Info("GetAllTwits called")
	twits, err := h.Twit.GetAllTwits()
	if err != nil {
		h.Log.Error("Error getting all twits", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	h.Log.Info("All twits retrieved successfully")
	c.JSON(200, twits)
}

// GetTwitsByType godoc
// @Summary get twits by type
// @Description it will get twits by type
// @Tags TWIT API
// @Param type path string true "type"
// @Success 200 {object} []string
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /v1/twit/type/{type} [get]
func (h *newTwits) GetTwitsByType(c *gin.Context) {
	h.Log.Info("GetTwitsByType called")
	typeStr := c.Param("type")
	if typeStr == "" {
		h.Log.Error("Invalid type")
		c.JSON(400, gin.H{"error": "Invalid type"})
		return
	}

	twits, err := h.Twit.GetTwitsByType(typeStr)
	if err != nil {
		h.Log.Error("Error getting twits by type", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	h.Log.Info("Twits retrieved successfully")
	c.JSON(200, twits)
}

// GetMostViewedTwit godoc
// @Summary get most view twits
// @Description it will get most view twits
// @Tags TWIT API
// @Param limit query string false "limit"
// @Success 200 {object} []string
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /v1/twit/most-viewed [get]
func (h *newTwits) GetMostViewedTwits(c *gin.Context) {
	h.Log.Info("GetMostViewedTwits called")
	limitStr := c.Query("limit")
	limit := 10
	if limitStr != "" {
		limit, _ = strconv.Atoi(limitStr)
	}

	twits, err := h.Twit.GetMostViewedTwit(limit)
	if err != nil {
		h.Log.Error("Error getting most viewed twits", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	h.Log.Info("Most viewed twits retrieved successfully")
	c.JSON(200, twits)
}

// GetLatestTwits godoc
// @Summary get latest twits
// @Description it will get latest twits
// @Tags TWIT API
// @Param limit query string false "limit"
// @Success 200 {object} []string
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /v1/twit/latest-uploaded [get]
func (h *newTwits) GetLatestTwits(c *gin.Context) {
	h.Log.Info("GetLatestTwits called")
	limitStr := c.Query("limit")
	limit := 10
	if limitStr != "" {
		limit, _ = strconv.Atoi(limitStr)
	}

	twits, err := h.Twit.GetLatestTwits(limit)
	if err != nil {
		h.Log.Error("Error getting latest twits", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	h.Log.Info("Latest twits retrieved successfully")
	c.JSON(200, twits)
}

// SearchTwit godoc
// @Summary search twit by keyword from twit text and twit title and publisher-name
// @Description it will search twit by keyword from twit text and twit title and publisher-name
// @Tags TWIT API
// @Param keyword query string false "keyword"
// @Success 200 {object} []string
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /v1/twit/search [get]
func (h *newTwits) SearchTwit(c *gin.Context) {
	h.Log.Info("SearchTwit called")
	keyword := c.Query("keyword")
	if keyword == "" {
		h.Log.Error("Invalid keyword")
		c.JSON(400, gin.H{"error": "Invalid keyword"})
		return
	}

	twits, err := h.Twit.SearchTwit(keyword)
	if err != nil {
		h.Log.Error("Error searching twits", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	h.Log.Info("Twits retrieved successfully")
	c.JSON(200, twits)
}

// CreateLocation godoc
// @Security ApiKeyAuth
// @Summary Create Twit's Location
// @Description it will Create Twit's Location
// @Tags TWIT API
// @Param info body model.CreateLocationRequest true "info"
// @Success 200 {object} string "location_id"
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /v1/twit/location [post]
func (h *newTwits) CreateLocation(c *gin.Context) {
	h.Log.Info("CreateLocation called")
	var info model.CreateLocationRequest
	if err := c.ShouldBindJSON(&info); err != nil {
		h.Log.Error("Invalid data", "error", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	locationID, err := h.Twit.CreateLocation(&info)
	if err != nil {
		h.Log.Error("Error creating location", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	h.Log.Info("Location created successfully")
	c.JSON(200, gin.H{"location_id": locationID})
}

// CreateUrl godoc
// @Security ApiKeyAuth
// @Summary Create Twit's urls like youtube url
// @Description it will Create Twit's urls like youtube url
// @Tags TWIT API
// @Param info body model.CreateURLRequest true "info"
// @Success 200 {object} string "url_id"
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /v1/twit/url [post]
func (h *newTwits) CreateUrl(c *gin.Context) {
	h.Log.Info("CreateUrl called")
	var info model.CreateURLRequest
	if err := c.ShouldBindJSON(&info); err != nil {
		h.Log.Error("Invalid data", "error", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	urlID, err := h.Twit.CreateUrl(&info)
	if err != nil {
		h.Log.Error("Error creating url", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	h.Log.Info("Url created successfully")
	c.JSON(200, gin.H{"url_id": urlID})
}

// @Summary CreatePhoto
// @Security ApiKeyAuth
// @Description Upload Twit Photo
// @Tags TWIT API
// @Param twit_id path string true "twit_id"
// @Param file formData file true "UploadMediaForm"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /v1/twit/photo/{twit_id} [post]
func (h *newTwits) CreatePhoto(c *gin.Context) {
	h.Log.Info("UploadProductPhoto called")

	Id := c.Param("twit_id")
	if len(Id) == 0 {
		h.Log.Error("twit_id is required")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Twit ID is required"})
		return
	}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		h.Log.Error("Error retrieving the file", "error", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error retrieving the file"})
		return
	}
	fmt.Println(file)
	defer file.Close()
	fmt.Println("minioga kirvoti")
	url, err := h.MINIO.UploadFile("photos", file, header)
	if err != nil {
		h.Log.Error("Error uploading the file to MinIO", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	photo_id, err := h.Twit.CreatePhoto(&model.CreatePhotoRequest{
		TwitID: Id,
		Photo:  url,
	})
	if err != nil {
		h.Log.Error("Error creating photo", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error creating photo"})
		return
	}
	h.Log.Info("Photo uploaded successfully")
	c.JSON(http.StatusOK, gin.H{"photo_id": photo_id, "url": url})
}

// @Summary CreateMusic
// @Security ApiKeyAuth
// @Description Upload Twit Music
// @Tags TWIT API
// @Accept multipart/form-data
// @Param twit_id path string true "twit_id"
// @Param file formData file true "UploadMediaForm"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /v1/twit/music/{twit_id} [post]
func (h *newTwits) CreateMusic(c *gin.Context) {
	h.Log.Info("UploadProductMusic called")

	Id := c.Param("twit_id")
	if len(Id) == 0 {
		h.Log.Error("twit_id is required")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Twit ID is required"})
		return
	}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		h.Log.Error("Error retrieving the file", "error", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error retrieving the file"})
		return
	}
	defer file.Close()
	url, err := h.MINIO.UploadFile("musics", file, header)
	if err != nil {
		h.Log.Error("Error uploading the file to MinIO", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error uploading the file to MinIO"})
		return
	}
	music_id, err := h.Twit.CreateMusic(&model.CreateMusicRequest{
		TwitID: Id,
		MP3:    url,
	})
	if err != nil {
		h.Log.Error("Error creating music", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error creating music"})
		return
	}
	h.Log.Info("Music uploaded successfully")
	c.JSON(http.StatusOK, gin.H{"music_id": music_id, "url": url})
}

// @Summary CreateVideo
// @Security ApiKeyAuth
// @Description Upload Twit Video
// @Tags TWIT API
// @Accept multipart/form-data
// @Param twit_id path string true "twit_id"
// @Param file formData file true "UploadMediaForm"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /v1/twit/video/{twit_id} [post]
func (h *newTwits) CreateVideo(c *gin.Context) {
	h.Log.Info("UploadProductVideo called")

	Id := c.Param("twit_id")
	if len(Id) == 0 {
		h.Log.Error("twit_id is required")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Twit ID is required"})
		return
	}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		h.Log.Error("Error retrieving the file", "error", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error retrieving the file"})
		return
	}
	defer file.Close()
	url, err := h.MINIO.UploadFile("videos", file, header)
	if err != nil {
		h.Log.Error("Error uploading the file to MinIO", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error uploading the file to MinIO"})
		return
	}
	video_id, err := h.Twit.CreateVideo(&model.CreateVideoRequest{
		TwitID: Id,
		Video:  url,
	})
	if err != nil {
		h.Log.Error("Error creating video", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error creating video"})
		return
	}
	h.Log.Info("Video uploaded successfully")
	c.JSON(http.StatusOK, gin.H{"video_id": video_id, "url": url})
}

// @Summary GetUniqueTypes
// @Description Get types list
// @Tags TWIT API
// @Success 200 {object} string
// @Failure 500 {object} string
// @Router /v1/twit/types [get]
func (h *newTwits) GetUniqueTypes(c *gin.Context) {
	h.Log.Info("GetUniqueTypes called")

	types, err := h.Twit.GetUniqueTypes()
	if err != nil {
		h.Log.Error("Error getting unique types", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error getting unique types"})
		return
	}
	h.Log.Info("Unique types retrieved successfully")
	c.JSON(http.StatusOK, types)
}

// AddMainTwit godoc
// @Security ApiKeyAuth
// @Summary Add a main twit
// @Description Adds a new main twit with start_time and end_time
// @Tags TWIT API
// @Param info body model.SavedRequestApi true "Twit details"
// @Success 200 {object} string "Twit saved successfully"
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /v1/twit/main [post]
func (h *newTwits) AddMainTwit(c *gin.Context) {
	h.Log.Info("AddMainTwit called")
	var req model.SavedRequestApi

	if err := c.ShouldBindJSON(&req); err != nil {
		h.Log.Error("Invalid data", "error", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := h.Twit.AddMainTwit(req.TwitId, req.StartTime, req.EndTime)
	if err != nil {
		h.Log.Error("Error adding main twit", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.Log.Info("Main twit added successfully")
	c.JSON(200, gin.H{"message": "Twit saved successfully"})
}

// GetMainTwit godoc
// @Summary Get active main twits
// @Description Retrieves main twits that are not deleted and match the current time range
// @Tags TWIT API
// @Success 200 {object} []string "List of twit IDs"
// @Failure 500 {object} string "Server error"
// @Router /v1/twit/main [get]
func (h *newTwits) GetMainTwit(c *gin.Context) {
	h.Log.Info("GetMainTwit called")

	twits, err := h.Twit.GetMainTwit()
	if err != nil {
		h.Log.Error("Error getting main twits", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.Log.Info("Retrieved main twits successfully")
	c.JSON(200, gin.H{"twits": twits})
}

// DeleteMainTwit godoc
// @Security ApiKeyAuth
// @Summary Delete a main twit
// @Description Soft deletes a main twit
// @Tags TWIT API
// @Param twit_id path string true "Twit ID"
// @Success 200 {object} string "Twit deleted successfully"
// @Failure 400 {object} string "Invalid twit ID"
// @Failure 500 {object} string "Server error"
// @Router /v1/twit/main/{twit_id} [delete]
func (h *newTwits) DeleteMainTwit(c *gin.Context) {
	h.Log.Info("DeleteMainTwit called")
	twitID := c.Param("twit_id")

	if twitID == "" {
		h.Log.Error("Invalid twit ID")
		c.JSON(400, gin.H{"error": "Invalid twit ID"})
		return
	}

	err := h.Twit.DeleteMainTwit(twitID)
	if err != nil {
		h.Log.Error("Error deleting main twit", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.Log.Info("Main twit deleted successfully")
	c.JSON(200, gin.H{"message": "Twit deleted successfully"})
}
