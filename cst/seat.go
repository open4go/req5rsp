package cst

// SeatStatus 座位状态
type SeatStatus int

const (
	// SeatAvailable 空闲
	SeatAvailable SeatStatus = iota
	// SeatOccupied 座位占用
	SeatOccupied
	// SeatBook 座位预定
	SeatBook
	// SeatDisabled 座位禁用
	SeatDisabled
)
