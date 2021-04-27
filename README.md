# go_smallsite
Small website development in golang

### Markdown Sheat-Sheet
https://www.markdownguide.org/cheat-sheet/


### How to install MySQL and Basic example
visit(https://gowebexamples.com/mysql-database/) for more infomation.
To install the MySQL database driver, go to your terminal of choice and run:
```
go get -u github.com/go-sql-driver/mysql
```

To check if we can connect to our database, import the database/sql and the go-sql-driver/mysql package and open up a connection:
```
import "database/sql"
import _ "go-sql-driver/mysql"


// Configure the database connection (always check errors)
db, err := sql.Open("mysql", "username:password@(127.0.0.1:3306)/dbname?parseTime=true")



// Initialize the first connection to the database, to see if everything works correctly.
// Make sure to check the error.
err := db.Ping()

```
