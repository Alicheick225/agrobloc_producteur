package controllers

import (
	"agrobloc_producteur/config"
	"agrobloc_producteur/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAllParcelles(c *gin.Context) {
	var producteurs []models.Parcelle

	if err := config.DB.Find(&producteurs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, producteurs)
}

func CreateParcelle(c *gin.Context) {
	var parcelle models.Parcelle

	// Bind JSON to struct
	if err := c.ShouldBindJSON(&parcelle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	// Insert into DB
	if err := config.DB.Create(&parcelle).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error: " + err.Error()})
		return
	}

	// Return created parcelle (with auto-generated ID)
	c.JSON(http.StatusCreated, parcelle)
}

func DeleteParcelle(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}

	var parcelle models.Parcelle

	// Vérifie si la parcelle existe
	if err := config.DB.First(&parcelle, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Parcelle non trouvée"})
		return
	}

	// Supprime la parcelle
	if err := config.DB.Delete(&parcelle).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la suppression"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Parcelle supprimée avec succès"})
}

func UpdateParcelle(c *gin.Context) {
	// 1. Lire l'ID dans l'URL
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}

	// 2. Chercher la parcelle existante
	var parcelle models.Parcelle
	if err := config.DB.First(&parcelle, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Parcelle non trouvée"})
		return
	}

	// 3. Lire les nouvelles données du corps de la requête
	var updatedData models.Parcelle
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Données invalides : " + err.Error()})
		return
	}

	// 4. Mettre à jour les champs (tu peux personnaliser ici si besoin)
	parcelle.Libelle = updatedData.Libelle
	parcelle.Geolocalisation = updatedData.Geolocalisation
	parcelle.Surface = updatedData.Surface
	parcelle.UserID = updatedData.UserID

	// 5. Sauvegarder dans la base
	if err := config.DB.Save(&parcelle).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur de mise à jour"})
		return
	}

	// 6. Retourner la parcelle mise à jour
	c.JSON(http.StatusOK, parcelle)
}

func GetParcelleByID(c *gin.Context) {
	// 1. Récupérer l'ID dans l'URL
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}

	// 2. Chercher la parcelle par ID
	var parcelle models.Parcelle
	if err := config.DB.First(&parcelle, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Parcelle non trouvée"})
		return
	}

	// 3. Retourner la parcelle trouvée
	c.JSON(http.StatusOK, parcelle)
}

func GetParcellesByUser(c *gin.Context) {
	userIDParam := c.Param("user_id")

	// Parse l'UUID
	userID, err := uuid.Parse(userIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID utilisateur invalide"})
		return
	}

	var parcelles []models.Parcelle

	// Filtrer les parcelles par user_id
	if err := config.DB.Where("user_id = ?", userID).Find(&parcelles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, parcelles)
}
