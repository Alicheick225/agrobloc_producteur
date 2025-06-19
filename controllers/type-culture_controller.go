package controllers

import (
	"agrobloc_producteur/config"
	"agrobloc_producteur/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateTypeCulture(c *gin.Context) {
	var tc models.TypeCulture

	if err := c.ShouldBindJSON(&tc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Données invalides : " + err.Error()})
		return
	}

	if err := config.DB.Create(&tc).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur DB : " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, tc)
}

func GetAllTypesCultures(c *gin.Context) {
	var types []models.TypeCulture

	if err := config.DB.Find(&types).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur DB : " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, types)
}

func GetTypeCultureByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}

	var tc models.TypeCulture
	if err := config.DB.First(&tc, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Type de culture non trouvé"})
		return
	}

	c.JSON(http.StatusOK, tc)
}

func UpdateTypeCulture(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}

	var existing models.TypeCulture
	if err := config.DB.First(&existing, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Type de culture non trouvé"})
		return
	}

	var updated models.TypeCulture
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Données invalides : " + err.Error()})
		return
	}

	existing.Libelle = updated.Libelle
	existing.PrixBordChamp = updated.PrixBordChamp

	if err := config.DB.Save(&existing).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur DB : " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, existing)
}

func DeleteTypeCulture(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}

	var tc models.TypeCulture
	if err := config.DB.First(&tc, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Type de culture non trouvé"})
		return
	}

	if err := config.DB.Delete(&tc).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur suppression : " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Type de culture supprimé avec succès"})
}
