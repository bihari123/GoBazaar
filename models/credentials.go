package models

type UserCred struct {
	ID   int    `json:"Id"`
	Pass string `json:"Pass"`
}

type MerchCred struct {
	ID   int    `json:"Id"`
	Pass string `json:"Pass"`
}
