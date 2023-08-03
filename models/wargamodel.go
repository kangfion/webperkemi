package models

import (
	"database/sql"
	"log"

	"webperkemi/config"
	"webperkemi/entities"
)

type WargaModel struct {
	db *sql.DB
}

func NewWargaModel() *WargaModel {
	conn, err := config.DBConn()

	if err != nil {
		panic(err)
	}

	return &WargaModel{
		db: conn,
	}
}

func (p *WargaModel) FindAll() ([]entities.Warga, error) {

	rows, err := p.db.Query("select id,name,username,avatar_file_name from users")
	if err != nil {
		return []entities.Warga{}, err
	}
	defer rows.Close()

	var dataWarga []entities.Warga

	for rows.Next() {
		var warga entities.Warga
		rows.Scan(&warga.Id,
			&warga.Name,
			&warga.Username,
			&warga.AvatarFileName)

		dataWarga = append(dataWarga, warga)
	}

	return dataWarga, nil

}

func (p *WargaModel) Find(id int64, warga *entities.Warga) error {
	return p.db.QueryRow("select id,name,username,avatar_file_name from users where id = ?", id).Scan(
		&warga.Id,
		&warga.Name,
		&warga.Username,
		&warga.AvatarFileName,
	)
}

func (p *WargaModel) Update(id int64, warga *entities.Warga) error {

	log.Println("ID MODEL WARGA ", id)

	_, err := p.db.Exec("update tbl_warga set  nama=? where id = ?",
		&warga.Name, id)

	if err != nil {
		return err
	}

	return nil
}

func (p *WargaModel) Create(warga *entities.Warga) error {
	result, err := p.db.Exec("insert into users (name,username, alamat,created_at) values(?,?,?,now())",
		warga.Name, warga.Username, warga.Alamat)

	if err != nil {
		return err
	}

	lastInsertId, _ := result.LastInsertId()
	warga.Id = lastInsertId
	return nil
}

func (p *WargaModel) Delete(id int64) error {
	_, err := p.db.Exec("delete from users where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
