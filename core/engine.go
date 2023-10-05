package core

type Role string

const (
	None   Role = "none"
	Reader Role = "reader"
	Writer Role = "writer"
	Admin  Role = "admin"
)
