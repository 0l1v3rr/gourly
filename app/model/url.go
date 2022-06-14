package model

import "github.com/0l1v3rr/gourly/urlutil"

type URL struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	RedirectTo string `json:"redirect_to" gorm:"not null"`
	UniqueId   string `json:"unique_id" gorm:"unique;not null"`
	Clicks     int    `json:"clicks"`
}

func (URL) GetAll() ([]URL, error) {
	var urls []URL

	tx := DB.Find(&urls)
	if tx.Error != nil {
		return []URL{}, tx.Error
	}

	return urls, nil
}

func (URL) GetByID(id int) (URL, error) {
	var url URL

	tx := DB.Where("id = ?", id).First(&url)
	if tx.Error != nil {
		return URL{}, tx.Error
	}

	return url, nil
}

func (URL) Create(url URL) (URL, error) {
	url.UniqueId = urlutil.GenerateRandomUrl(32)
	url.Clicks = 0

	tx := DB.Create(&url)
	return url, tx.Error
}

func (URL) Delete(id int) error {
	tx := DB.Unscoped().Delete(&URL{}, id)
	return tx.Error
}

func (URL) IncrementClicks(id int) (URL, error) {
	url, err := URL{}.GetByID(id)
	if err != nil {
		return URL{}, err
	}

	url.Clicks++

	tx := DB.Save(&url)
	if tx.Error != nil {
		return URL{}, tx.Error
	}

	return url, nil
}
