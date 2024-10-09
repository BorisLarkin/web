package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) PredictionById(ctx *gin.Context) {
	id := ctx.Param("id")
	Prediction, err := h.Repository.GetPredictionByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.HTML(http.StatusOK, "prediction.tmpl", gin.H{
		"Prediction": Prediction,
	})
}

func (h *Handler) DeletePrediction(ctx *gin.Context) {
	id := ctx.Param("id")
	h.Repository.DeletePrediction(id)
	ctx.Redirect(http.StatusFound, "/")
}