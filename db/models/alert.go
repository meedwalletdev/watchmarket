package models

import "time"

type Interval string

type Alert struct {
	UpdatedAt time.Time
	AssetID   string
	Interval
	Difference float64
	Display    bool
}

const (
	Hour Interval = "1h"
	Day  Interval = "1d"
	Week Interval = "1w"
)
