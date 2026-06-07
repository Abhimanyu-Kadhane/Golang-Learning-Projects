Bookstore Management APIs


Details:
1. database : my sql
2. orm : GORM
3. Json marshall unmarshall
4. Project structure
5. Gorilla Mux

Project Structure:
    CMD 
        - main.go
    PKG
        - config
            - app.go
        -controllers
            - book-controller
        -models
            - book.go
        -routes
            - boookstore-routes
        -utils
            - utils.go



Routes :
    GETBOOKS -> /book/ -> GET
    CREATEBOOKS -> /book/ -> POST
    GETBOOKBYID -> /book/{bookid} -> GET
    UPDATEBOOK -> /book/{bookid} -> PUT
    DELETEBOOK -> /book/{bookid} -> DELETE
    

