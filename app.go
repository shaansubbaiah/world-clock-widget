package main

import (
	"context"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

type Clock struct {
	Timezone string `json:"timezone"`
	Name string `json:"name"`
	TimeFormat string `json:"timeFormat"`
	DateFormat string `json:"dateFormat"`
	Color string `json:"color"`
}



func (a *App) GetClocks() []Clock {
	clocks := []Clock{
		{Timezone: "US/Eastern", Name: "EST", TimeFormat: "", DateFormat:  "",Color: "blue"},
		{Timezone: "Asia/Kolkata",Name:  "IST",TimeFormat:  "", DateFormat: "", Color:  "orange"},
	}
	return clocks
}