package models

import (
	"errors"
	"net/http"
	"time"
)

type GameMode struct {
	ID        uint       `gorm:"primary_key" json:"id"` // ID of game mode
	CreatedAt time.Time  `json:"created_at"`            // creation time of game mode
	UpdatedAt time.Time  `json:"updated_at"`            // updated time of game mode
	DeletedAt *time.Time `json:"deleted_at"`            // deletation time of game mode

	GameID uint `json:"game_id"` // ID of the game the mode belongs to

	Name      string     `json:"name"`                                                                                // the mode's name
	GameCards []GameCard `json:"cards" gorm:"many2many:gamecard_modes;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // the mode's cards
} // @name GameMode

type GameModePostPutPayload struct {
	Name *string `json:"name" validate:"required"`
} // @name GameModePostPutPayload

func (g *GameModePostPutPayload) Bind(r *http.Request) error {
	if g.Name == nil {
		return errors.New("missing required content property")
	}

	return nil
}
