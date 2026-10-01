package models

type Profile struct {
	Name       string       `json:"name"`
	Handle     string       `json:"handle"`
	Title      string       `json:"title"`
	Email      string       `json:"email"`
	Phone      string       `json:"phone"`
	Summary    string       `json:"summary"`
	Links      Links        `json:"links"`
	Experience []Experience `json:"experience"`
	Projects   []Project    `json:"projects"`
	Education  []Education  `json:"education"`
	Skills     Skills       `json:"skills"`
	Rules      []string     `json:"rules"`
}

type Links struct {
	Website  string `json:"website,omitempty"`
	LinkedIn string `json:"linkedin,omitempty"`
	GitHub   string `json:"github,omitempty"`
}

type Experience struct {
	Slug       string   `json:"slug"`
	Role       string   `json:"role"`
	Company    string   `json:"company"`
	Location   string   `json:"location"`
	Start      string   `json:"start"`
	End        string   `json:"end"`
	Highlights []string `json:"highlights"`
	Stack      []string `json:"stack"`
}

type Project struct {
	Slug         string   `json:"slug"`
	Name         string   `json:"name"`
	Subtitle     string   `json:"subtitle"`
	Live         bool     `json:"live"`
	Technologies []string `json:"technologies"`
	Highlights   []string `json:"highlights"`
}

type Education struct {
	Institution string `json:"institution"`
	Affiliation string `json:"affiliation"`
	Degree      string `json:"degree"`
	Start       string `json:"start"`
	End         string `json:"end"`
}

type Skills struct {
	Languages []string `json:"languages"`
	Databases []string `json:"databases"`
	Frontend  []string `json:"frontend"`
	Backend   []string `json:"backend"`
	Mobile    []string `json:"mobile"`
	DevOps    []string `json:"devops"`
	Testing   []string `json:"testing"`
	Tools     []string `json:"tools"`
}
