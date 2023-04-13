package api

import (
	"net/http"

	db "github.com/alim7007/go_backend_bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type ListEntriesParams struct {
	AccountID int64 `uri:"id" binding:"required,min=1"`
	PageID    int32 `form:"page_size"`
	PageSize  int32 `form:"page_skip"`
}

func (server *Server) GetEntry(ctx *gin.Context) {
	var req ListEntriesParams
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := db.ListEntriesParams{
		AccountID: req.AccountID,
		Limit:     req.PageID,
		Offset:    req.PageSize,
	}
	entries, err := server.store.ListEntries(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, entries)
}
