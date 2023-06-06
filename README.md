# books
This project is a simple book catalogue with the rest API.

For data storing usage postgres DB

## Running
To start the server you need to fill the constant databseDSN in file /cmd/server/main.go and run the command
~~~~
go run ./cmd/server/main.go
~~~~
After it, the server will be ready to handle requests  by port 8080

## API
The server provides rest API with the following methods which allow storing in DB information about books (title, author and publish year). 

Data is sent and received in JSON format
~~~~
{
    "title":"book title",
    "author":"book author",
    "year": 2010,
    "id":3
}
~~~~

### Add new book
To add a new book into the collection you need to send a POST or PUT request to the  endpoint /book
~~~~
 curl -v -X PUT -H 'Content-Type:application/json' -d '{"title":"book title", "author":"book author", "year": 2010}' http://localhost:8080/book
~~~~
In case of success body of a response will contain the JSON object of the new book

### Update book
To update a book in the collection you need to send a POST or PUT request to the endpoint /book/:id where :id is a unique identification in DB
 ~~~~
 curl -v -X PUT -H 'Content-Type:application/json' -d '{"title":"book title", "author":"book author", "year": 2010}' http://localhost:8080/book/4
~~~~

### Delete book
To delete a book from the collection you need to send a POST or PUT request to the endpoint /book/:id where :id is a unique identification in DB
 ~~~~
 curl -v -X DELETE http://localhost:8080/book/4
~~~~

### List of books
To get a list of  books you need to send a GET request to the endpoint /books 
 ~~~~
 curl -v -X GET http://localhost:8080/books
~~~~