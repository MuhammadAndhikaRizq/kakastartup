package campaign

import "strings"

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{}
	campaignFormatter.ID = campaign.ID
	campaignFormatter.UserID = campaign.UserID
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.Slug = campaign.Slug
	campaignFormatter.ImageURL = ""

	//Pengecekan panjang data CampaignImages
	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	return campaignFormatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	//{} nilai default
	campaignsFormatter := []CampaignFormatter{}

	//mengubah ke struct formatter
	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)

		//kumpulan campaign formatter
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}

type CampaignDetailFormatter struct {
	ID               int                      `json:"id"`
	Name             string                   `json:"name"`
	ShortDescription string                   `json:"short_description"`
	Description      string                   `json:"description"`
	ImageUrl         string                   `json:"image_url"`
	GoalAmount       int                      `json:"goal_amount"`
	CurrentAmount    int                      `json:"current_amount"`
	BackerCount      int                      `json:backer_count"`
	UserID           int                      `json:"user_id`
	Slug             string                   `json:"slug"`
	Perks            []string                 `json:"perks`
	User             CampaignUserFormatter    `json:"user"`
	Images           []CampaignImageFormatter `json:"images"`
}

//struct ini dipakai di dalam struct CampaignDetailFormatter
//User
type CampaignUserFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

//Image
type CampaignImageFormatter struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	campaignDetailFormatter := CampaignDetailFormatter{}
	campaignDetailFormatter.ID = campaign.ID
	campaignDetailFormatter.Name = campaign.Name
	campaignDetailFormatter.ShortDescription = campaign.ShortDescription
	campaignDetailFormatter.Description = campaign.Description
	campaignDetailFormatter.GoalAmount = campaign.GoalAmount
	campaignDetailFormatter.CurrentAmount = campaign.CurrentAmount
	campaignDetailFormatter.BackerCount = campaign.BackerCount
	campaignDetailFormatter.UserID = campaign.UserID
	campaignDetailFormatter.Slug = campaign.Slug
	campaignDetailFormatter.ImageUrl = ""

	if len(campaign.CampaignImages) > 0 {
		campaignDetailFormatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	//pemecah string berdasarkan ,
	var perks []string
	for _, perk := range strings.Split(campaign.Perks, ",") {
		//fungsi append bertujuan untuk menambahkan data/elemen pada slice perks
		perks = append(perks, strings.TrimSpace(perk))
	}

	campaignDetailFormatter.Perks = perks

	//membuat objek user
	user := campaign.User

	//objek dari struct CampaignUserFormatter diatas
	//Untuk mengisi field user dari struct CampaignDetailFormatter
	CampaignUserFormatter := CampaignUserFormatter{}
	CampaignUserFormatter.Name = user.Name
	CampaignUserFormatter.ImageURL = user.AvatarFileName

	campaignDetailFormatter.User = CampaignUserFormatter

	//objek bertipe slice CampaignImageFormatter
	images := []CampaignImageFormatter{}

	//image mewakili entitiy struct CampaignImages
	for _, image := range campaign.CampaignImages {
		//objek darii CampaignImageFormatter
		campaignImageFormatter := CampaignImageFormatter{}
		campaignImageFormatter.ImageURL = image.FileName

		//pengecekan true or false
		isPrimary := false

		if image.IsPrimary == 1 {
			isPrimary = true
		}
		campaignImageFormatter.IsPrimary = isPrimary

		//Menambahkan data hasil diatas ke dalam images yg bertipe slice
		images = append(images, campaignImageFormatter)
	}

	//set field images dari campaignDetailFormatter / mengisi field image nya
	campaignDetailFormatter.Images = images

	return campaignDetailFormatter

}
