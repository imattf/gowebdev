package models

import "time"

//Session is a ...
type Session struct {
	UserName     string
	LastActivity time.Time
}
