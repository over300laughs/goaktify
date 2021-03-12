package models

import (
    "database/sql"
)

// Create an exported global variable to hold the database connection pool.
var DB *sql.DB

type Campaigns struct {
    ID   int // TODO change to UUID
    Description  string
    IsActive bool
}

func (c *Campaigns) Create() {

}

func (c *Campaigns) Read() {

}

func (c *Campaigns) Update() {

}

func (c *Campaigns) Delete() {

}

func (c *Campaigns) List() {

}

