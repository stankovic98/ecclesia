package model

type Parish struct {
	UID       string
	Name      string
	Priest    string
	Info      string
	DioceseID string
	Aritcles  []Aritcle
}

type Diocese struct {
	UID      string
	Name     string
	Info     string
	Aritcles []Aritcle
}

type Aritcle struct {
	Author    string
	Title     string
	Content   string
	CreatedAt string
}
