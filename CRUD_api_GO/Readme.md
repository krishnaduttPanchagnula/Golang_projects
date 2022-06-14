# CRUD application with GO

The present project is developed as a inital prototype to understand api building ing Go. Initially we have build CRUD on a Movies . The structs contain ID, Moviename, Budge, Hero, Heroine. A SLICE is used as data storage mechanism which will be updated to a database soon. 


## Dependencies

| Libraries  | Utility   |
|:------:|:--------:|
|[encoding/json](https://pkg.go.dev/encoding/json) | To encode the data into json when sending a response |
|[fmt](https://golangdocs.com/fmt-package-golang)| to print out and scan things|
|[log](https://golangdocs.com/log-package-golang) | to create logs for api|
|[math/rand](https://pkg.go.dev/math/rand)| To create ID to the input data|
|[net/http](https://pkg.go.dev/net/http)  |To create a web server|
|[strconv](https://pkg.go.dev/strconv) |To convert the data into strings|
|[MUX](https://github.com/gorilla/mux) |Creates request router|

## API structure

| Method  | Route   | Utility   |
|:------:|:--------:|:--------:|
|GET |/movies|   Returns all the movies in the slice|
|GET|/movies/{id}| Retuens movie with given id in the slice|
|POST|/movies| Create a movie in the slice|
|PUT|/movies/{id}| Updates the movie with given id in the slice|
|DELETE|/movie/{id}| Deletes movie with given id in the slice|
	
### Object structure
```GO
    type Movie struct {
	ID          string `json:id`
	Genre       string `json:Genre`
	Budget      int    `json:budget`
	Title       string `json:title`
	Technicians `json:Technicians`
}

type Technicians struct {
	Hero    string `json:hero`
	Heroine string `json:Heroine`
}
```

