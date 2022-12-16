package types

type AddImageRegistryRequest struct {
	Name         string `json:"name" binding:"required"`
	Desc         string `json:"desc"`
	Host         string `json:"host" binding:"required"`
	Username     string `json:"username" binding:"required"`
	Password     string `json:"password" binding:"required"`
	HistoryCount int64  `json:"history_count" binding:"required"`
	Protocol     string `json:"protocol"  binding:"required"`
}

type UpdateImageRegistryRequest struct {
	ID           int64   `json:"id"  binding:"required"`
	Name         *string `json:"name"`
	Desc         *string `json:"desc"`
	Host         *string `json:"host"`
	HistoryCount *int64  `json:"history_count"`
	Username     *string `json:"username"`
	Password     *string `json:"password"`
	Protocol     *string `json:"protocol"`
}

type DeleteImageRegistryRequest struct {
	ID int64 `json:"id"  binding:"required"`
}

type ConnectImageRegistryRequest struct {
	ID int64 `json:"id"  binding:"required"`
}
