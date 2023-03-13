package transaction

import "time"

type CampaignTrasanctionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAT time.Time `json:"created_at"`
}

func FormatCampaignTransaction(transaction Transaction) CampaignTrasanctionFormatter {
	//menampung data dari struct CampaignTrasactionFormatter
	//Objek
	formatter := CampaignTrasanctionFormatter{}
	formatter.ID = transaction.ID
	formatter.Name = transaction.User.Name
	formatter.Amount = transaction.Amount
	formatter.CreatedAT = transaction.CreatedAt

	return formatter
}

//mengubah dari list atau slice transaction menjadi slice list campaign transaction

func FormatCampaignTransactions(transactions []Transaction) []CampaignTrasanctionFormatter {
	if len(transactions) == 0 {
		return []CampaignTrasanctionFormatter{}
	}

	var transactionFormatter []CampaignTrasanctionFormatter

	for _, transaction := range transactions {
		formatter := FormatCampaignTransaction(transaction)
		transactionFormatter = append(transactionFormatter, formatter)
	}

	return transactionFormatter
}
