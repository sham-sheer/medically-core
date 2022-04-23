package main

// User is a model in the "users" table.
type User struct {
	ID   	int     `json:"id,omitempty"`
	Name 	*string `json:"name" gorm:"not null"`
	Email 	*string `json:"email" gorm:"not null"`
	Contact *string `json:"contact" gorm:"not null"`
}

// Med is a model in the "medications" table.
type Med struct {
	ID    int     `json:"id,omitempty"`
	Name  *string `json:"name"  gorm:"not null"`
	Desc  *string `json:"desc" 	gorm:"not null"`
}

// Disease is a model in the "diseases" table.
type Disease struct {
	ID    int     `json:"id,omitempty"`
	Name  *string `json:"name"  gorm:"not null"`
	Desc  *string `json:"desc" 	gorm:"not null"`
}

// Clinics is a model in the "clinics" table.
type Clinic struct {
	ID    int     `json:"id,omitempty"`
	Name  *string `json:"name"  gorm:"not null"`
	Desc  *string `json:"desc" 	gorm:"not null"`
}
