package connection

import (
	"api-address-golang/repository/configuration"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	configuration := configuration.GetDatabase()

	stringConnection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", configuration.Host, configuration.Port, configuration.User, configuration.Password, configuration.Database)

	connection, err := sql.Open("postgres", stringConnection)
	err = connection.Ping()

	return connection, err
}
