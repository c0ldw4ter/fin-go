package handler

import (
	"net/http"
	"strconv"

	"fin-api/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
    srv *service.Service
}

func NewHandler(srv *service.Service) *Handler {
    return &Handler{srv: srv}
}

func (h *Handler) TopUpBalance(c *gin.Context) {
    userID, err := strconv.Atoi(c.Param("user_id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
        return
    }

    var req struct {
        Amount float64 `json:"amount"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    ctx := c.Request.Context()
    if err := h.srv.TopUpBalance(ctx, userID, req.Amount); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "balance topped up successfully"})
}

func (h *Handler) TransferMoney(c *gin.Context) {
    fromUserID, err := strconv.Atoi(c.Param("from_user_id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid from user ID"})
        return
    }

    toUserID, err := strconv.Atoi(c.Param("to_user_id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid to user ID"})
        return
    }

    var req struct {
        Amount float64 `json:"amount"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    ctx := c.Request.Context()
    if err := h.srv.TransferMoney(ctx, fromUserID, toUserID, req.Amount); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "money transferred successfully"})
}

func (h *Handler) GetLastTransactions(c *gin.Context) {
    userID, err := strconv.Atoi(c.Param("user_id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
        return
    }

    ctx := c.Request.Context()
    transactions, err := h.srv.GetLastTransactions(ctx, userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, transactions)
}