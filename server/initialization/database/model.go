package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// table User for user data
type User struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Email      string    `gorm:"type:varchar(255);uniqueIndex;not null"`
	Username   *string   `gorm:"type:varchar(50);unique"`
	Password   *string   `gorm:"type:text"`
	Provider   string    `gorm:"type:varchar(50);default:'local';not null;index"` // e.g. local, google, facebook, line
	ProviderID *string   `gorm:"type:varchar(255);index"`
	Role       string    `gorm:"type:varchar(30);default:'customer';not null;index"` // e.g. admin, customer, employee, owner_restaurant

	Activate    Activate      `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Profile     Profile       `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	SocialMedia []SocialMedia `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Restaurant  []Restaurant  `gorm:"foreignKey:OwnerID;references:ID;constraint:OnDelete:CASCADE"`

	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// table Activate for checking user with activate function
type Activate struct {
	ID              uint      `gorm:"primaryKey;autoIncrement"`
	UserID          uuid.UUID `gorm:"type:uuid;unique"`
	ActivatedEmail  bool      `gorm:"type:boolean;default:false"`
	ActivatedPhone  bool      `gorm:"type:boolean;default:false"`
	ActivatedWallet bool      `gorm:"type:boolean;default:false"`
}

type Profile struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	UserID    uuid.UUID `gorm:"type:uuid;unique"`
	AvatarURL string    `gorm:"type:varchar(255);default:''"`

	FirstName string `gorm:"type:varchar(100);default:''"`
	LastName  string `gorm:"type:varchar(100);default:''"`
	Phone     string `gorm:"type:varchar(20);default:''"`

	Address string `gorm:"type:varchar(255);default:''"`
	City    string `gorm:"type:varchar(100);default:''"`
	State   string `gorm:"type:varchar(100);default:''"`
	Country string `gorm:"type:varchar(100);default:''"`
	ZipCode string `gorm:"type:varchar(20);default:''"`
}

type Restaurant struct {
	ID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	OwnerID uuid.UUID `gorm:"type:uuid;not null"`
	Name    string    `gorm:"type:varchar(255);not null"`

	Address string `gorm:"type:varchar(255);default:''"`
	City    string `gorm:"type:varchar(100);default:''"`
	State   string `gorm:"type:varchar(100);default:''"`
	Country string `gorm:"type:varchar(100);default:''"`
	ZipCode string `gorm:"type:varchar(20);default:''"`
	Phone1  string `gorm:"type:varchar(20);default:''"`
	Phone2  string `gorm:"type:varchar(20);default:''"`

	SocialMedia []SocialMedia `gorm:"foreignKey:RestaurantID;references:ID;constraint:OnDelete:CASCADE"`

	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type SocialMedia struct {
	ID           uint       `gorm:"primaryKey;autoIncrement"`
	UserID       *uuid.UUID `gorm:"type:uuid"`
	RestaurantID *uuid.UUID `gorm:"type:uuid"`

	FaceBookURL  string `gorm:"type:varchar(255);default:''"`
	InstagramURL string `gorm:"type:varchar(255);default:''"`
	TwitterURL   string `gorm:"type:varchar(255);default:''"`
	YoutubeURL   string `gorm:"type:varchar(255);default:''"`
	WebsiteURL   string `gorm:"type:varchar(255);default:''"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	active := Activate{
		UserID: u.ID,
	}
	if err = tx.Create(&active).Error; err != nil {
		return err
	}

	profile := Profile{
		UserID: u.ID,
	}
	if err = tx.Create(&profile).Error; err != nil {
		return err
	}
	return nil
}
