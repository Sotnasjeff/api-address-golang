package service

import (
	"api-address-golang/entities_core"
	"api-address-golang/repository/connection"
)

const (
	queryInsert        = `INSERT INTO address (public_area, street, number, neighborhood, city, state, country) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	queryGetById       = `SELECT * FROM address WHERE id= $1`
	queryGetAllAddress = `SELECT * FROM address`
	queryUpdateAddress = `UPDATE address SET public_area=$2, street=$3, number=$4, neighborhood=$5, city=$6, state=$7, country=$8 WHERE id=$1`
	queryDeleteAddress = `DELETE FROM address WHERE id=$1`
)

func InsertIntoDB(address entities_core.Address) (id int64, err error) {
	conn, err := connection.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	err = conn.QueryRow(queryInsert, address.Public_area, address.Street, address.Number, address.Neighborhood, address.City, address.State, address.Country).Scan(&id)

	return

}

func GetAddressById(id int64) (address entities_core.Address, err error) {
	conn, err := connection.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	err = conn.QueryRow(queryGetById, id).Scan(&address.Id, &address.Public_area, &address.Street, &address.Number, &address.Neighborhood, &address.City, &address.State, &address.Country)

	return
}

func GetAllAddress() (address []entities_core.Address, err error) {
	conn, err := connection.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(queryGetAllAddress)

	for rows.Next() {
		var address_help entities_core.Address
		err = rows.Scan(&address_help.Id, &address_help.Public_area, &address_help.Street, &address_help.Number, &address_help.Neighborhood, &address_help.City, &address_help.State, &address_help.Country)
		if err != nil {
			continue
		}

		address = append(address, address_help)
	}

	return

}

func UpdateAddress(id int64, address entities_core.Address) (int64, error) {
	conn, err := connection.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	result, err := conn.Exec(queryUpdateAddress, id, address.Public_area, address.Street, address.Number, address.Neighborhood, address.City, address.State, address.Country)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func DeleteAddress(id int64) (int64, error) {
	conn, err := connection.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	result, err := conn.Exec(queryDeleteAddress, id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()

}
