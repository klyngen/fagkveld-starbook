package repository

type Star struct {
	ID          int
	Description string
	PersonID    int
	UserID      string
}

type Person struct {
	ID        int
	Name      string
	Image     string
	BelongsTo string
}
