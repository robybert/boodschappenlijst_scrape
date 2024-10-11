package models

type Config struct {
    Host     string
    Port     string
    User     string
    Password string
    DBName   string
    SSLMode  string
}

type Product struct {
	ID 				uint	`gorm:"primaryKey; auto_increment" json:"id"`
	ProductName 	string	`gorm:"not null; type:varchar(128)" json:"product_name"`
	Weight			int	`gorm:"not null" json:"weight"`
}

type Price struct {
	ID				uint	`gorm:"primaryKey; auto_increment"`
	PID				uint
	Lidl 			uint
	Aldi			uint
	Deka			uint
	Ah				uint
	Product 		Product	`gorm:"foreignKey:PID"`
}


