package kitsu

import (
	"time"
)

// SkippedReactionEnum is the skipped reaction
type SkippedReactionEnum string

// skipped reactions
var (
	ReactionUnskipped SkippedReactionEnum = "unskipped"
	ReactionSkipped   SkippedReactionEnum = "skipped"
	ReactionIgnored   SkippedReactionEnum = "ignored"
)

// StatusEnum is a watch status
type StatusEnum string

// watch status
var (
	StatusCompleted StatusEnum = "completed"
	StatusCurrent   StatusEnum = "current"
	StatusDropped   StatusEnum = "dropped"
	StatusOnHold    StatusEnum = "on_hold"
	StatusPlanned   StatusEnum = "planned"
)

type kitsuTime time.Time

func (kt *kitsuTime) UnmarshalJSON(data []byte) error {
	return kt.UnmarshalJSON(data)
}

func (kt *kitsuTime) MarshalJSON() ([]byte, error) {
	return kt.MarshalJSON()
}

// LibraryEntry is an anime or manga in a user's library
type LibraryEntry struct {
	ID              string    `jsonapi:"primary,libraryEntries"`
	Status          string    `jsonapi:"attr,status"`
	Progress        int64     `jsonapi:"attr,progress"`
	VolumesOwned    int64     `jsonapi:"attr,volumesOwned"`
	Reconsuming     bool      `jsonapi:"attr,reconsuming"`
	ReconsumeCount  int64     `jsonapi:"attr,reconsumeCount"`
	Notes           string    `jsonapi:"attr,notes"`
	Private         bool      `jsonapi:"attr,private"`
	ReactionSkipped string    `jsonapi:"attr,reactionSkipped"`
	ProgressedAt    kitsuTime `jsonapi:"attr,progressedAt"`
	StartedAt       kitsuTime `jsonapi:"attr,startedAt"`
	FinishedAt      kitsuTime `jsonapi:"attr,finishedAt"`
	Rating          string    `jsonapi:"attr,rating"`
	RatingTwenty    float64   `jsonapi:"attr,ratingTwenty"`
	CreatedAt       kitsuTime `jsonapi:"attr,createdAt"`
	UpdatedAt       kitsuTime `jsonapi:"attr,updatedAt"`
}
