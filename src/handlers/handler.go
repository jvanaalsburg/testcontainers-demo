package handlers

type Handler struct {
	connStr string
}

func NewHandler(connStr string) *Handler {
	return &Handler{connStr: connStr}
}
