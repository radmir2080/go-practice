package domain

import "time"

type User struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

type Student struct {
	ID       uint   `json:"id"`
	UserID   uint   `json:"user_id"`
	FullName string `json:"full_name"`
	Group    string `json:"group"`
	Skills   string `json:"skills"`
}

type Company struct {
	ID          uint   `json:"id"`
	UserID      uint   `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Internship struct {
	ID          uint   `json:"id"`
	CompanyID   uint   `json:"company_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PlacesCount int    `json:"places_count"`
}
