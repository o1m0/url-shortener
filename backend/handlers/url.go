package handlers

import (
	"math/rand"
	"net/http"
	"time"

	"url-shortener/db"
	"url-shortener/models"

	"github.com/gin-gonic/gin"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateShortCode() string {
	code := make([]byte, 6)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}
	return string(code)
}

type CreateURLInput struct {
	OriginalURL string `json:"original_url" binding:"required"`
	ExpiresIn   *int   `json:"expires_in"` // 有効期限（時間単位）
}

func CreateURL(c *gin.Context) {
	var input CreateURLInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URLを入力してください"})
		return
	}

	userID, _ := c.Get("userID")

	url := models.URL{
		OriginalURL: input.OriginalURL,
		ShortCode:   generateShortCode(),
		UserID:      func() *uint { id := userID.(uint); return &id }(),
	}

	if input.ExpiresIn != nil {
		expiresAt := time.Now().Add(time.Duration(*input.ExpiresIn) * time.Hour)
		url.ExpiresAt = &expiresAt
	}

	if err := db.DB.Create(&url).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, url)
}

func Redirect(c *gin.Context) {
	shortCode := c.Param("shortCode")

	// Redisキャッシュを確認
	cachedURL, err := db.RDB.Get(db.Ctx, shortCode).Result()
	if err == nil {
		db.RDB.Incr(db.Ctx, "count:"+shortCode)
		c.Redirect(http.StatusMovedPermanently, cachedURL)
		return
	}

	// キャッシュミス → DBから検索
	var url models.URL
	if err := db.DB.Where("short_code = ?", shortCode).First(&url).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URLが見つかりません"})
		return
	}

	// 有効期限チェック
	if url.ExpiresAt != nil && time.Now().After(*url.ExpiresAt) {
		// Redisのキャッシュも削除
		db.RDB.Del(db.Ctx, shortCode)
		c.JSON(http.StatusGone, gin.H{"error": "このURLは期限切れです"})
		return
	}

	// Redisにキャッシュ保存
	db.RDB.Set(db.Ctx, shortCode, url.OriginalURL, time.Hour)

	// アクセスカウント更新
	db.DB.Model(&url).Update("access_count", url.AccessCount+1)

	c.Redirect(http.StatusMovedPermanently, url.OriginalURL)
}
