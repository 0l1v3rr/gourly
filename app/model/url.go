package model

type URL struct {
	ID         int    `json:"id"`
	RedirectTo string `json:"redirect_to"`
	UniqueId   string `json:"unique_id"`
	Clicks     int    `json:"clicks"`
}

func (URL) GetAll() ([]URL, error) {
	urls := []URL{}

	return urls, nil
}

func (URL) GetByID(id int) (URL, error) {
	return URL{}, nil
}

func (URL) Create(url URL) (URL, error) {
	return url, nil
}

func (URL) Delete(id int) error {
	return nil
}

func (URL) IncrementClicks(id int) (URL, error) {
	return URL{}, nil
}
