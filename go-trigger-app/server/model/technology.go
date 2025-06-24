package model

type Technology struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Name    string `json:"name"`
	Details string `json:"details"`
}
