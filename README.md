# Basic CRUD project

DB used - Postgres

Install postgres --
`brew install postgresql`

Start postgres db -- `brew services start postgresql`
Stop postgres db -- `brew services stop postgresql`

You need Soda(Buffalo) binary to run the migrations which can be found [here](https://gobuffalo.io/en/docs/db/getting-started/)

run the project `go run cmd/api/main.go`

test commands <br>
`curl -X POST  http://localhost:4444/movies -d '{"name":"Lord of the rings","description":"Epic movie"}'` <br>
` curl -X DELETE http://localhost:4444/movies/1`<br>
`curl http://localhost:4444/movies/1`<br>
`curl -X PUT http://localhost:4444/movies/1 -d '{ "title":"The Lord of the rings","description":"epic movie"}'`<br>