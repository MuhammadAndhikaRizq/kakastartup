package transaction

import (
	"kakastartup/campaign"
	"kakastartup/user"
	"time"
)

// Reperesentasi dari tabel transaction di database
type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	Code       string
	User       user.User
	Campaign   campaign.Campaign
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
