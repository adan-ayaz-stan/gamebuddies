package db

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Temporarily putting structs here
type User struct {
	ID                 string         `gorm:"type:uuid;primaryKey;"`
	CreatedAt          time.Time      `gorm:"autoCreateTime"`
	UpdatedAt          time.Time      `gorm:"autoUpdateTime"`
	DeletedAt          gorm.DeletedAt `gorm:"index"`
	InstanceId         string         `gorm:"not null"`
	Aud                string         `gorm:"not null"`
	Role               string         `gorm:"not null"`
	Email              string         `gorm:"not null"`
	EncryptedPassword  string         `gorm:"not null"`
	EmailConfirmedAt   *time.Time
	ConfirmationToken  *string
	ConfirmationSentAt *time.Time
	RecoveryToken      *string
	RecoverySentAt     *time.Time
	LastSignInAt       *time.Time
	BannedUntil        *time.Time
}

type Sessions struct {
	ID          int    `gorm:"primaryKey"`
	UserId      string `gorm:"type:uuid;not null"`
	NotAfter    *time.Time
	RefreshedAt *time.Time
	UserAgent   *string
	Ip          *string
	Tag         *string
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

type RefreshTokens struct {
	ID         uint      `gorm:"primaryKey"`
	UserId     string    `gorm:"type:uuid;not null"`
	InstanceId string    `gorm:"type:varchar(255);default:'000-000-000'"`
	Token      string    `gorm:"type:varchar(255);"`
	SessionID  uint      `gorm:"not null"`
	Revoked    bool      `gorm:"default:false"`
	Parent     *string   `gorm:"type:varchar(255);"`
	CreatedAt  time.Time `gorm:"not null;default:now()"`
	UpdatedAt  time.Time `gorm:"not null;default:now()"`
}

var DB *gorm.DB

// BeforeCreate will set a UUID rather than numeric ID.
func (b *User) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	return
}

func InitDB() {
	dsn := "root:root@tcp(localhost:3307)/gamebuddy?parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
