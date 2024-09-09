package models

type BlogTag struct {
	Id               int    `json:"id"`
	BlogcategoryId   int    `json:"blogcategory_id"`
	Name             string `json:"name"`
	BlogcategoryName string `json:"blogcategory_name"`
}
