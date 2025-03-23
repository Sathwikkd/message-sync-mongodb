package http

import (
    "net/http"
    "chat-sync-service/internal/app"
    "github.com/labstack/echo/v4"
)

type Handler struct {
    syncUC *app.SyncUsecase
}

func NewHandler(e *echo.Echo, syncUC *app.SyncUsecase) {
    h := &Handler{syncUC: syncUC}
    e.GET("/sync", h.SyncMessages)
}

func (h *Handler) SyncMessages(c echo.Context) error {
    userID := c.QueryParam("user_id")
    currentMsgID := c.QueryParam("messageid")

    if userID == "" || currentMsgID == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Missing user_id or current_msg_id"})
    }

    messages, err := h.syncUC.SyncMessages(userID, currentMsgID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }

    return c.JSON(http.StatusOK, messages) 
}
