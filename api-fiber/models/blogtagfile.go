package models

type BlogTagFile struct {
	Id               int    `json:"id"`
	BlogtagId        int    `json:"blogtag_id"`
	BlogfileId       int    `json:"blogfile_id"`
	BlogtagName      string `json:"blogtag_name"`
	BlogfileFilename string `json:"blogfile_filename"`
}
