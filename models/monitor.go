package models

import "time"

// Monitor - base structure for saving information
// about all monitors in the system
type Monitor struct {
	Model

	Name         string    `json:"name" gorm:"size:255" form:"name"`
	Description  string    `json:"description" gorm:"size:2000" form:"description"`
	Address      string    `json:"address" gorm:"size:255" form:"address"`
	User         User      `json:"user"`
	Interval     uint16    `json:"interval" form:"interval"`
	IsActive     bool      `json:"is_active" form:"is_active"`
	LastActivity time.Time `json:"last_active"`
}

// WebMonitor - monitor for http/https checking
// can use login and password for HTTP Auth
// can use login via login form (execute set of actions for login)
type WebMonitor struct {
	Monitor
	Login    string `json:"login" gorm:"size:1000" form:"login"`
	Password string `json:"password" gorm:"size:1000" form:"password"`
}

// PingMonitor - monitor which uses ping command for checking uptime
type PingMonitor struct {
	Monitor
}

// KeywordsMonitor - monitor for checking if some keyword presents or not
// presents on the page
type KeywordsMonitor struct {
	Monitor
	Keywords []string `json:"keywords" form:"keywords"`
}
