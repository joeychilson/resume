package resume

type Resume struct {
	Basics       Basics        `json:"basics"`
	Work         []Work        `json:"work"`
	Volunteer    []Volunteer   `json:"volunteer"`
	Education    []Education   `json:"education"`
	Awards       []Award       `json:"awards"`
	Publications []Publication `json:"publications"`
	Skills       []Skill       `json:"skills"`
	Languages    []Language    `json:"languages"`
	Interests    []Interest    `json:"interests"`
	References   []Reference   `json:"references"`
	Projects     []Project     `json:"projects"`
}

type Basics struct {
	Name     string    `json:"name"`
	Label    string    `json:"label"`
	Image    string    `json:"image"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	URL      string    `json:"url"`
	Summary  string    `json:"summary"`
	Location Location  `json:"location"`
	Profiles []Profile `json:"profiles"`
}

type Location struct {
	Address     string `json:"address"`
	City        string `json:"city"`
	Region      string `json:"region"`
	CountryCode string `json:"countryCode"`
}

type Profile struct {
	Network  string `json:"network"`
	Username string `json:"username"`
	URL      string `json:"url"`
}

type Work struct {
	Name        string   `json:"name"`
	Location    string   `json:"location"`
	Description string   `json:"description"`
	Position    string   `json:"position"`
	URL         string   `json:"url"`
	StartDate   string   `json:"startDate"`
	EndDate     string   `json:"endDate"`
	Summary     string   `json:"summary"`
	Highlights  []string `json:"highlights"`
}

type Volunteer struct {
	Organization string   `json:"organization"`
	Position     string   `json:"position"`
	URL          string   `json:"url"`
	StartDate    string   `json:"startDate"`
	EndDate      string   `json:"endDate"`
	Summary      string   `json:"summary"`
	Highlights   []string `json:"highlights"`
}

type Education struct {
	Institution string   `json:"institution"`
	URL         string   `json:"url"`
	Area        string   `json:"area"`
	StudyType   string   `json:"studyType"`
	StartDate   string   `json:"startDate"`
	EndDate     string   `json:"endDate"`
	Score       string   `json:"score"`
	Courses     []string `json:"courses"`
}

type Award struct {
	Title   string `json:"title"`
	Date    string `json:"date"`
	Awarder string `json:"awarder"`
	Summary string `json:"summary"`
}

type Publication struct {
	Name        string `json:"name"`
	Publisher   string `json:"publisher"`
	ReleaseDate string `json:"releaseDate"`
	URL         string `json:"url"`
	Summary     string `json:"summary"`
}

type Skill struct {
	Name string `json:"name"`
}

type Language struct {
	Language string `json:"language"`
}

type Interest struct {
	Name string `json:"name"`
}

type Reference struct {
	Name      string `json:"name"`
	Reference string `json:"reference"`
}

type Project struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	URL         string   `json:"url"`
	Keywords    []string `json:"keywords"`
}
