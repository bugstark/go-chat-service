package cron

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"time"
	"ws/app/websocket"
)
var s *gocron.Scheduler

func clearChannel()  {
	websocket.AdminManager.ClearInactiveChannel()
	websocket.UserManager.ClearInactiveChannel()
}

func Run()  {
	fmt.Println("start cron")
	s = gocron.NewScheduler(time.UTC)
	s.Every(1).Minute().Do(closeSessions)
	s.Every(1).Minute().Do(clearChannel)
	s.StartAsync()
}
func Stop()  {
	if s != nil{
		s.Stop()
	}
}
