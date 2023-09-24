package controller

import (
	"AKO/ch/teko/ako/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type rankJson struct {
	Rank string `json:"rank"`
}

func (h Handler) getAllRanks(context *gin.Context) {

	allRanks := make([]rankJson, len(model.AllRanks))
	for index, rank := range model.AllRanks {
		allRanks[index] = rankJson{Rank: rank}
	}
	context.JSON(http.StatusOK, allRanks)
}
