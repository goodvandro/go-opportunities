package schemas

import "gorm.io/gorm"

type Openings struct {
	gorm.Model
	Role     string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   int64
}
