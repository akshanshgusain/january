package middleware

import (
	"github.com/akshanshgusain/january"
	"januaryApp/data"
)

type Middleware struct {
	App    *january.January
	Models data.Models
}
