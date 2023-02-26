package campaign

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Campaign, error)
	FindByUserID(userID int) ([]Campaign, error)
	FindByID(ID int) (Campaign, error)
	Save(campaign Campaign) (Campaign, error)
	Update(campaign Campaign) (Campaign, error)
	CreateImage(campaignImage CampaignImage) (CampaignImage, error)
	MarksAllImagesAsNonPrimary(campaignID int) (bool, error)
}

type repository struct {
	//menghubungkan ke database
	db *gorm.DB
}

//instance repository
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign

	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindByUserID(userID int) ([]Campaign, error) {
	var campaigns []Campaign

	//preload campaignimages befungsi untuk memfilter primary yg diiginkan
	err := r.db.Where("user_id = ?", userID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil

}

func (r *repository) FindByID(ID int) (Campaign, error) {
	var campaign Campaign

	err := r.db.Preload("User").Preload("CampaignImages").Where("id = ?", ID).Find(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil

}

func (r *repository) Save(campaign Campaign) (Campaign, error) {
	err := r.db.Create(&campaign).Error //crate pada gorm berfungsi untuk menyimpan data

	if err != nil {
		return campaign, err

	}

	return campaign, nil
}

func (r *repository) Update(campaign Campaign) (Campaign, error) {
	err := r.db.Save(&campaign).Error //Save pada gorm berfungsi mengupdate data yang sudah ada di database

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *repository) CreateImage(campaignImage CampaignImage) (CampaignImage, error) {
	err := r.db.Create(&campaignImage).Error //crate pada gorm berfungsi untuk menyimpan data

	if err != nil {
		return campaignImage, err

	}

	return campaignImage, nil
}

func (r *repository) MarksAllImagesAsNonPrimary(campaignID int) (bool, error) {
	//Query
	//UPDATE campaign SET is_primary = false WHERE campaign_id = 1

	//code dalam bentuk gorm
	//Model berfungsi untuk mencari struct yang dituju
	//Model merupakan representasi struktur tabel pada database.
	err := r.db.Model(&CampaignImage{}).Where("campaign_id = ? ", campaignID).Update("is_primary", false).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
