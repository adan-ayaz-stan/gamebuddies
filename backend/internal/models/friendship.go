package models

import (
	"time"

	"gorm.io/gorm"
)

// FriendshipStatus reflects the state of a friend request
type FriendshipStatus string

const (
	FriendshipPending  FriendshipStatus = "pending"
	FriendshipAccepted FriendshipStatus = "accepted"
)

// Friendship represents a directional friend relationship.
// RequesterID → AddresseeID.  Accepted when status = "accepted".
type Friendship struct {
	gorm.Model

	RequesterID uint             `gorm:"not null;index:idx_friendship_pair,unique,priority:1"`
	AddresseeID uint             `gorm:"not null;index:idx_friendship_pair,unique,priority:2"`
	Status      FriendshipStatus `gorm:"type:varchar(20);not null;default:'pending'"`

	Requester User `gorm:"foreignKey:RequesterID;constraint:OnDelete:CASCADE"`
	Addressee User `gorm:"foreignKey:AddresseeID;constraint:OnDelete:CASCADE"`
}

// UserSession tracks platform presence for time-on-platform stats.
type UserSession struct {
	gorm.Model

	UserID    uint       `gorm:"not null;index"`
	StartedAt time.Time  `gorm:"not null"`
	EndedAt   *time.Time // nil = currently active

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

// DurationSeconds returns the duration of this session in seconds.
// If still active, measures from start until now.
func (s *UserSession) DurationSeconds() int64 {
	end := time.Now()
	if s.EndedAt != nil {
		end = *s.EndedAt
	}
	return int64(end.Sub(s.StartedAt).Seconds())
}
