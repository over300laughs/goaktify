package handlertests

import (
	"database/sql"
	"fmt"
	"os"
	"projects/goaktify/handlers"
	"testing"

	_ "github.com/lib/pq"
)

// Create an exported global variable to hold the database connection pool.
var DB *sql.DB
var ch handlers.CampaignHandler

func TestMain(m *testing.M) {
	// setup DB connection

	connStr := "host=%s port=%s user=%s sslmode=disable password=%s dbname=%s"
	DBURL := fmt.Sprintf(connStr,
		os.Getenv("TEST_DB_HOST"),
		os.Getenv("TEST_DB_PORT"),
		os.Getenv("TEST_DB_USER"),
		os.Getenv("TEST_DB_PASSWORD"),
		os.Getenv("TEST_DB_NAME"),
	)

	var err error

	DB, err = sql.Open("postgres", DBURL)
	if err != nil {
		fmt.Print("Cannot connect to test database")
	}

	os.Exit(m.Run())
}

func TestRead(t *testing.T) {
	fmt.Println("TestRead unimplented due to time.")
	fmt.Println("See sudo code in TestRead function")

	/*
				            The tests required the router middleware to be abstracted further per idiomatic standards.
				            Didn't have enough time to implement. Below is where I got to

							err := refreshDB()
							if err != nil {
								t.Errorf("Cannot refresh to test database: %v", err)
							}

							id, err := seedCampaign()
							if err != nil {
								t.Errorf("Cannot seed campaign to database: %v", err)
							}

							req, err := http.NewRequest("GET", "/campaigns/"+strconv.Itoa(id), nil)
							if err != nil {
								t.Fatal(err)
							}

							// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
							rr := httptest.NewRecorder()
							handler := http.HandlerFunc(ch.Read)

							// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
							// directly and pass in our Request and ResponseRecorder.
		                    handler.ServeHTTP(rr, req)

		                    AssertTrue(Read(id = 1) returns 200  // pass
		                    AssertFalse(Read(id = 0) returns 400 // pass
		                    AssertFalse(Read(id = 10000) return 400 or 404 for resource not found // pass

	*/

}

func TestCreate(t *testing.T) {
	fmt.Println("TestCreate unimplented due to time.")
	fmt.Println("See sudo code in TestRead function")

}

func TestUpdate(t *testing.T) {
	fmt.Println("TestUpdate unimplented due to time.")
	fmt.Println("See sudo code in TestRead function")

}

func TestDelete(t *testing.T) {
	fmt.Println("TestDelete unimplented due to time.")
	fmt.Println("See sudo code in TestRead function")

}

func TestList(t *testing.T) {
	fmt.Println("TestList unimplented due to time.")
	fmt.Println("See sudo code in TestRead function")

}

func refreshDB() error {

	// create campaigns table
	_, err := DB.Exec(`DROP TABLE IF EXISTS campaigns`)
	if err != nil {
		return err
	}

	_, err = DB.Exec(`
        CREATE TABLE campaigns ( id int PRIMARY KEY, 
                                name varchar(255),
                                description varchar,
                                created_on timestamp default current_timestamp,
                                updated_on timestamp default current_timestamp,
                                is_active BOOL NOT NULL );`)
	if err != nil {
		return err
	}

	_, err = DB.Exec(`DROP SEQUENCE IF EXISTS campaigns_id_sequence;`)
	if err != nil {
		return err
	}

	_, err = DB.Exec(`CREATE SEQUENCE campaigns_id_sequence START 1 INCREMENT 1;`)
	if err != nil {
		return err
	}

	return nil

}

func seedCampaign() (int, error) {
	id := 0

	insertQuery := `INSERT INTO campaigns (id, name, description, created_on, updated_on, is_active)
                    VALUES (nextval('campaigns_id_sequence'), 'Test Campaign', 'Basic description', NOW(), NOW(), TRUE) RETURNING id;`

	err := DB.QueryRow(insertQuery).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, err
}

func seedCampaigns() error {

	insertQuery := `INSERT INTO campaigns (id, name, description, created_on, updated_on, is_active)
                    VALUES
                    (nextval('campaigns_id_sequence'), 'Test 1 Campaign', "Basic description 1", NOW(), NOW(), TRUE),
                    (nextval('campaigns_id_sequence'), 'Test 2 Campaign', "Basic description 2", NOW(), NOW(), TRUE),
                    (nextval('campaigns_id_sequence'), 'Test 3 Campaign', "Basic description 3", NOW(), NOW(), TRUE),
                    (nextval('campaigns_id_sequence'), 'Test 4 Campaign', "Basic description 4", NOW(), NOW(), TRUE),
                    (nextval('campaigns_id_sequence'), 'Test 5 Campaign', "Basic description 5", NOW(), NOW(), TRUE);`

	_, err := DB.Query(insertQuery)
	return err

}
