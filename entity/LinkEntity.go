package entity

type Link struct {
	Id				*string
	IsDeleted		*bool
	CurrentUsr		*string
}

func (Link) TableName() string {
	return "link"
}