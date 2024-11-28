# CRUD
Simple project CRUD movies. Without database, only local data.

## Start project
- Crate go mod
```go
    go mod init <name>
```
- Install gorilla mux
```go
    go get "github.com/gorilla/mux"
```
- Create file main
```go
    touch main.go
```
## CRUD

- `getMovies` (GET)
    - return all movies
- `getMovie` (GET)
    - get movie by ID
    - need pass {id} in route
- `createMovie` (POSt)
    - create a new Movie
    - need pass struct Movie in the body request
- `updateMovie` (PUT)
    - update data movie by id
    - need pass  {id} in route
- `deleteMovie` (DELETE)
    - delete movie by id
    - need pass  {id} in route
