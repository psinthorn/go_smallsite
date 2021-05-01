package category

var Category category

type category struct {
	ID          string `json: "id"`
	Title       string `json: "title"`
	Desc        string `json: "desc"`
	Section     string `json: "section"`
	Parent      string `json: "parent"`
	AccessLevel int    `json: "access_level"`
	StatusID    int    `json: "status_id"`
	CreatedAt   string `json: "created_at"`
	UpdatedAt   string `json: "updated_at"`
}

type CategoryStatus struct {
	ID        int    `json: "id"`
	Title     string `json: "title"`
	Status    string `json: "status"`
	CreatedAt string `json: "created_at"`
	UpdatedAt string `json: "updated_at"`
}
