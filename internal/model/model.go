package model

type OpenRequest struct {
	UserID     string
	ShareType  int32
	ShareCount int32
}
type CloseRequest struct {
	PositionID string
	ShareType  int32
	UserID     string
}
type Share struct {
	Name      int32
	Bid       float32
	Ask       float32
	UpdatedAt string
}
