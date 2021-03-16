# Welcome Aktify!

Hi! Thanks for taking the time to look over this challenge! I decided to reimplement it
with Go and Postgres. Dockerizing the app idiomatically was my fun focus for this project. Routes, if you're looking for them, will be in the `aktify.go` file in the root. Also handlers is just an idiomatic
naming convention for what other languages refer to as controllers in MVC.

Here's the steps to check it out:

1. Clone the repo
2. Then run `docker-compose up -d`
3. To shut it down  run `docker-compose down`
4. To run the tests in docker use `docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit` 

> **Note:** Unfortunately I didn't have time to fully implement the unit tests in their fullest. However, there is a fair amount of database setup in place. See `tests/handlertests/handler_test.go` to view the test code. 

To view the database you can go to `http:localhost:5050` using pgAdmin in a browser as usual. 

For convenience here is the database information:
Host: fullstack-postgres
User: postgres
Password: postgres (gasp)
Database Name: aktify_db_1
Port: 5432

And here is the curl commands to run:

**Create Campaigns** 

```
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name":"Aktify Campaign 1","description":"Typical campaign description","isActive":true}' \
  http://localhost:8080/campaigns
```

**Read Campaigns** 
```
curl --header "Content-Type: application/json" \
  --request GET \
  http://localhost:8080/campaigns/1
```
  **Delete Campaigns** 
```
curl --header "Content-Type: application/json" \
  --request DELETE \
  http://localhost:8080/campaigns/1
  ``` 

Thanks again for taking the time to look it over!!

-Michael