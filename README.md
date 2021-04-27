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
if err != nil {
    log.Fatal(err)
    return
}


// Initialize the first connection to the database, to see if everything works correctly.
// Make sure to check the error.
err := db.Ping()
if err != nil {
    log.Fatal(err)
    return
}

```

### How to create table

SQL command query create table for example 
```
CREATE TABLE users (
    id INT AUTO_INCREMENT,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    created_at DATETIME,
    PRIMARY KEY (id)
);
```

How to excute SQL comannad query
```
query := `
    CREATE TABLE users (
        id INT AUTO_INCREMENT,
        username TEXT NOT NULL,
        password TEXT NOT NULL,
        created_at DATETIME,
        PRIMARY KEY (id)
    );`

// Executes the SQL query in our database. Check err to ensure there was no error.
_, err := db.Exec(query)
if err != nil {
    log.Fatal(err)
    return
}
```


