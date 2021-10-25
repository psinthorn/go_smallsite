package drivers

// if use docker conotainer as environtment development then host = database service name
// if use local development then use localhost as usaul
var PgDsn = "host=db_postgres port=5432 dbname=f2xhotel user=postgres password=postgres"

//var PgDsn = "host=f2.co.th port=5432 dbname=f2coth user=f2coth password=f2coth"

// working with docker see mysql server details docker-compose file
var MySqlDsn = "root:root@tcp(db_mysql:3306)/specialist_ambassador?parseTime=True"
