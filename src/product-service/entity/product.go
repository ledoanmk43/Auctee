package entity

type Product struct {
	Id          string `json:"id" gorm:"primary_key;type:varchar(20);not null;unique"`
	Name        string `json:"name" gorm:"type:nvarchar(100);not null"`
	MinPrice    string `json:"minPrice" gorm:"type:nvarchar(100);not null"`
	Description string `json:"description" gorm:"type:nvarchar(500);not null"`
	Quantity    int    `json:"quantity" gorm:"type:nvarchar(500);not null"`
}
