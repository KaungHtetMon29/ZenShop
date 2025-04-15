package models

type Category struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"unique;not null"`
	Products  []Product `gorm:"foreignKey:CategoryID"`
	CreatedAt string    `gorm:"not null"`
}

type Product struct {
	ID                uint     `gorm:"primaryKey;autoIncrement"`
	ProductPerOrderID uint     `gorm:"not null"` // Foreign key
	Name              string   `gorm:"unique;not null"`
	Price             int      `gorm:"not null"`
	Stock             int      `gorm:"not null"`
	CategoryID        uint     `gorm:"not null"` // Foreign key
	Category          Category `gorm:"foreignKey:CategoryID"`
	CreatedAt         string   `gorm:"not null"`
	UpdatedAt         string   `gorm:"not null"`
	UpdateBy          string   `gorm:"not null"`
}

type ProductUpdateHistory struct {
	ID        uint    `gorm:"primaryKey;autoIncrement"`
	ProductID uint    `gorm:"not null"` // Foreign key
	Product   Product `gorm:"foreignKey:ProductID"`
	AdminID   string  `gorm:"not null"` // Foreign key
	UpdatedAt string  `gorm:"not null"`
	Summary   string  `gorm:"not null"`
}

type RepairStatus struct {
	ID        uint     `gorm:"primaryKey;autoIncrement"`
	UpdatedBy string   `gorm:"not null"`
	UpdatedAt string   `gorm:"not null"`
	Status    string   `gorm:"not null"`
	Repair    []Repair `gorm:"foreignKey:RepairStatusID"`
}

type Repair struct {
	ID             uint         `gorm:"primaryKey;autoIncrement"`
	UserId         string       `gorm:"not null"` // Foreign key
	RepairStatusID uint         `gorm:"not null"` // Foreign key
	RepairStatus   RepairStatus `gorm:"foreignKey:RepairStatusID"`
	Product        string       `gorm:"not null"`
	Category       string       `gorm:"not null"`
	CreatedAt      string       `gorm:"not null"`
	UpdatedAt      string       `gorm:"not null"`
	Description    string       `gorm:"not null"`
	Status         string       `gorm:"not null"`
}

type Order struct {
	ID                uint            `gorm:"primaryKey;autoIncrement"`
	UserId            string          `gorm:"not null"` // Foreign key
	ProductPerOrderID uint            `gorm:"not null"` // Foreign key
	ProductPerOrder   ProductPerOrder `gorm:"foreignKey:ProductPerOrderID"`
	CreatedAt         string          `gorm:"not null"`
	Payment           Payment         `gorm:"foreignKey:OrderID"`
	Shipping          Shipping        `gorm:"foreignKey:OrderID"`
}

type ProductPerOrder struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Order     []Order   `gorm:"foreignKey:ProductPerOrderID"`
	Product   []Product `gorm:"foreignKey:ProductPerOrderID"`
	CreatedAt string    `gorm:"not null"`
}

type Payment struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Amount    int    `gorm:"not null"`
	CreatedAt string `gorm:"not null"`
	Type      string `gorm:"not null"`
	OrderID   uint   `gorm:"not null"` // Foreign key
}

type Shipping struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Address   string `gorm:"not null"`
	CreatedAt string `gorm:"not null"`
	OrderID   uint   `gorm:"not null"` // Foreign key
}
