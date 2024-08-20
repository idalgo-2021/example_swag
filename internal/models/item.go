package models

import "time"

// @Description Item creation request
type ItemRequest struct {
	Name string `json:"name"`
}

// @Description Item response
type ItemResponse struct {
	ID string `json:"id"`
}

// @Description Item's description
type Item struct {
	ID         string    `json:"id" example:"1"`                             // @example field tag используется для примера значений
	Name       string    `json:"name" example:"TV-set"`                      // @example field tag используется для примера значений
	CreateTime time.Time `json:"create_time" example:"2024-08-21T14:00:00Z"` // @example field tag используется для примера значений
}
