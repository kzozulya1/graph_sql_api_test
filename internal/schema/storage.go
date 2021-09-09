package schema

import (
	"errors"

	"github.com/go-pg/pg/v10"
	"github.com/kzozulya1/graph_sql_api_test/internal/storage"
)

var errTooSmallObjID = errors.New("id expected to be > 0")

//getClient return client by filter
func getClient(db *pg.DB, id int) (interface{}, error) {
	var client = &storage.Client{}
	err := db.Model(client).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}
	return client, nil
}

//getClients return clients slice by filter
func getClients(db *pg.DB, clientName string, first, offset int) (interface{}, error) {
	var (
		clients []*storage.Client
		model   = db.Model(&clients)
	)

	if clientName != "" {
		model.Where("client_name ILIKE '%" + clientName + "%'")
	}
	if first != 0 {
		model.Limit(first)
	}
	if offset != 0 {
		model.Offset(offset)
	}

	if err := model.Select(); err != nil {
		return nil, err
	}
	return clients, nil
}

//getClientsCount counts record number in table
func getClientsCount(db *pg.DB) (interface{}, error) {
	return db.Model(&storage.Client{}).Count()
}

// createClient creates a Client
func createClient(db *pg.DB, clientName, urAdr string) (interface{}, error) {
	var model = &storage.Client{
		ClientName: clientName,
		UrAdr:      urAdr,
	}
	_, err := db.Model(model).Insert()
	return model, err
}

// updateClient updates data of the Client
func updateClient(db *pg.DB, id int, clientName, urAdr string) (interface{}, error) {

	if id > 0 {
		var model = &storage.Client{ID: id}

		err := db.Model(model).WherePK().Select()
		if err != nil {
			return nil, err
		}
		if clientName != "" {
			model.ClientName = clientName
		}
		if urAdr != "" {
			model.UrAdr = urAdr
		}

		_, err = db.Model(model).WherePK().Update()
		if err != nil {
			return nil, err
		}

		return model, nil
	}

	return nil, errTooSmallObjID
}

// deleteClient removes client from DB
func deleteClient(db *pg.DB, id int) (interface{}, error) {
	if id > 0 {
		var model = &storage.Client{ID: id}

		err := db.Model(model).WherePK().Select()
		if err != nil {
			return nil, err
		}

		_, err = db.Model(model).WherePK().Delete()
		if err != nil {
			return nil, err
		}
		return model, nil
	}
	return nil, errTooSmallObjID
}
