package models

import (
	"time"

	_ "github.com/lib/pq"
)

type Campaign struct {
	ID          int       `json:"id"` // TODO change to UUID
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedOn   time.Time `json:"createdOn"`
	UpdatedOn   time.Time `json:"updatedOn"`
	IsActive    bool      `json:"isActive"`
}

func (c *Campaign) Create() error {
	insertQuery := `INSERT INTO campaigns (id, name, description, created_on, updated_on, is_active)
                    VALUES (nextval('campaigns_id_sequence'), $1, $2, NOW(), NOW(), $3) RETURNING id;`

	err := DB.QueryRow(insertQuery, c.Name, c.Description, c.IsActive).Scan(&c.ID)
	if err != nil {
		return err
	}

	return c.read(c.ID)
}

func (c *Campaign) Read(id int) error {
	return c.read(id)
}

func (c *Campaign) read(id int) error {
	query := `SELECT id,
                    name,
                    description,
                    created_on,
                    updated_on,
                    is_active
                FROM campaigns WHERE id = $1;`

	err := DB.QueryRow(query, id).Scan(&c.ID, &c.Name, &c.Description, &c.CreatedOn, &c.UpdatedOn, &c.IsActive)
	if err != nil {
		return err
	}

	return nil
}

func (c *Campaign) Update() {
	//TODO: implement this soon
}

func (c *Campaign) Delete(id int) error {
	insertQuery := `DELETE FROM campaigns WHERE id = $1;`

	_, err := DB.Exec(insertQuery, id)
	if err != nil {
		return err
	}
	return nil
}

func (c *Campaign) List() {

	// company table unimplemented
	/* query := `SELECT id,
		    name,
		    description,
		    created_on,
		    updated_on,
	        is_active,
	        company_id -- foreign_key we need!
		FROM campaigns WHERE company_id = $1;` */
}
