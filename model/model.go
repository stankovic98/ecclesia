package model

type Parish struct {
	UID       string    `json:"uid"`
	Name      string    `json:"name"`
	Priest    string    `json:"priest"`
	Info      string    `json:"info"`
	DioceseID string    `json:"dioceseID`
	Aritcles  []Aritcle `json:"articles"`
}

type Diocese struct {
	UID      string    `json:"uid"`
	Name     string    `json:"name"`
	Info     string    `json:"info"`
	Aritcles []Aritcle `json:"articles"`
}

type Aritcle struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
}
