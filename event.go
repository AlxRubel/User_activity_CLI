package main

type Event struct {
	ID         int     `json:"id"`
	Type       string  `json:"type"`
	Actor      Actor   `json:"actor"`
	Repo       Repo    `json:"repo"`
	Payload    Payload `json:"payload"`
	Public     bool    `json:"public"`
	Created_at string  `json:"created_at"`
	Org        Org     `json:"org"`
}

type Actor struct {
	ID            int    `json:"id"`
	Login         string `json:"login"`
	Display_login string `json:"display_login"`
	Gravatar_ID   string `json:"gravatar_id"`
	URL           string `json:"url"`
	Avatar_URL    string `json:"avatar_url"`
}

type Repo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Org struct {
	ID          int    `json:"id"`
	Login       string `json:"login"`
	Gravatar_ID string `json:"gravatar_id"`
	URL         string `json:"url"`
	Avatar_URL  string `json:"avatar_url"`
}

type Payload interface {
	humanReadable() string
}
