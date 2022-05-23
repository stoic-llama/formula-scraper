package interfaces

type ICompany interface {
	GetName() string
	SetName(name string)
}

type DefaultCompany struct {
	Name     string
	Products map[string]ProductItem
	BaseURL  string
}

type ProductItem struct {
	Price  float32 `json:"price,omitempty"`
	Count  int     `json:"count,omitempty"`
	URL    string  `json:"url,omitempty"`
	Stores []Store `json:"stores,omitempty"`
}

type Store struct {
	Street    string  `json:"street,omitempty"`
	City      string  `json:"city,omitempty"`
	State     string  `json:"state,omitempty"`
	Country   string  `json:"country,omitempty"`
	Longitude float32 `json:"longitude,omitempty"`
	Latitude  float32 `json:"latitude,omitempty"`
}

func NewCompany() *DefaultCompany {
	return &DefaultCompany{}
}

func (a *DefaultCompany) GetName() string {
	return a.Name
}

func (a *DefaultCompany) SetName(name string) {
	a.Name = name
}
