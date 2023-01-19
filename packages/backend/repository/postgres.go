package repository

import "database/sql"

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		sql: db,
	}
}

type Repository struct {
	sql *sql.DB
}

func (r *Repository) CreatePerson(person *Person) error {
	var id int
	row := r.sql.QueryRow("INSERT INTO public.person (user_id, personName , image) VALUES($1, $2, $3) RETURNING id", person.BelongsTo, person.Name, person.Image)

	if row.Err() != nil {
		return row.Err()
	}

	err := row.Scan(&id)
	if err != nil {
		return err
	}

	person.ID = id

	return nil
}

func (r *Repository) GetPersonsByUserId(userId string) ([]Person, error) {
	rows, err := r.sql.Query("SELECT id, personName, image FROM public.person WHERE user_id = $1", userId)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	result := make([]Person, 0)

	var id int
	var name string
	var image string

	for rows.Next() {
		err = rows.Scan(&id, &name, &image)

		if err != nil {
			continue
		}

		result = append(result, Person{
			ID:    id,
			Name:  name,
			Image: image,
		})
	}

	return result, nil
}

func (r *Repository) CreateStar(star *Star) error {
	var id int
	row := r.sql.QueryRow("INSERT INTO public.star (person_id, description, user_id) VALUES($1, $2, $3) RETURNING id", star.PersonID, star.Description, star.UserID)

	if row.Err() != nil {
		return row.Err()
	}

	err := row.Scan(&id)

	if err != nil {
		return err
	}

	star.ID = id

	return nil

}

func (r *Repository) GetStarsByUser(userId string) ([]Star, error) {
	rows, err := r.sql.Query("SELECT id, person_id, description FROM public.star WHERE user_id = $1", userId)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	result := make([]Star, 0)

	var id int
	var personId int
	var description string

	for rows.Next() {
		err = rows.Scan(&id, &personId, &description)

		if err != nil {
			continue
		}

		result = append(result, Star{
			ID:          id,
			PersonID:    personId,
			Description: description,
		})
	}

	return result, nil
}
