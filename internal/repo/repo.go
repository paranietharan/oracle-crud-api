package repo

import (
	"database/sql"
	"oracle-crud-api/internal/model"
)

type PersonRepository struct {
	db *sql.DB
}

func NewPersonRepository(db *sql.DB) *PersonRepository {
	return &PersonRepository{db: db}
}

func (r *PersonRepository) Create(p *model.Person) error {
	_, err := r.db.Exec(`INSERT INTO persons (id, name, phone, email) VALUES (:1, :2, :3, :4)`,
		p.Id, p.Name, p.Phone, p.Email)
	return err
}

func (r *PersonRepository) GetByID(id string) (*model.Person, error) {
	row := r.db.QueryRow(`SELECT id, name, phone, email FROM persons WHERE id = :1`, id)

	var p model.Person
	if err := row.Scan(&p.Id, &p.Name, &p.Phone, &p.Email); err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *PersonRepository) GetAll() ([]model.Person, error) {
	rows, err := r.db.Query(`SELECT id, name, phone, email FROM persons`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var people []model.Person
	for rows.Next() {
		var p model.Person
		if err := rows.Scan(&p.Id, &p.Name, &p.Phone, &p.Email); err != nil {
			return nil, err
		}
		people = append(people, p)
	}
	return people, nil
}

func (r *PersonRepository) Update(p *model.Person) error {
	_, err := r.db.Exec(`UPDATE persons SET name = :1, phone = :2, email = :3 WHERE id = :4`,
		p.Name, p.Phone, p.Email, p.Id)
	return err
}

func (r *PersonRepository) Delete(id string) error {
	_, err := r.db.Exec(`DELETE FROM persons WHERE id = :1`, id)
	return err
}
