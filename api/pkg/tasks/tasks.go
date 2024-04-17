package tasks

import "time"

type Task struct {
    Title       string
    Description string
    CreatedAt   time.Time
    CompletedAt time.Time
}

type TaskList struct {
    Tasks []Task
}
