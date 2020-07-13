package model

type Parish struct {
	UID       string
	Name      string
	Priest    string
	DioceseID string
}

type Diocese struct {
	UID  string
	Name string
}

// type Aritcle struct {
// 	Author string
// 	Title string
// 	Content string
// 	CreatedAt string

// }
