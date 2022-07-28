package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"time"
)

const SelectQuery = "SELECT temperature, time FROM temperature_sensors WHERE sensor_id = $1 ORDER BY time DESC LIMIT 1"

const InsertQuery = "INSERT INTO temperature_sensors (sensor_id, temperature, time) VALUES ($1, $2, now())"

type Sensor struct {
	Temperature float32
	Time        time.Time
}

type Database struct {
	*sql.DB
}

func (database Database) GetLatestTemperature(sensorId int64) (Sensor, error) {
	var res Sensor

	rows, err := database.Query(SelectQuery, sensorId)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	if rows.Next() {
		errScan := rows.Scan(&res.Temperature, &res.Time)
		if errScan != nil {
			return res, err
		}
	}
	return res, nil
}

func (database Database) InsertTemperature(sensorId int64, temperature float64) (int64, error) {
	result, err := database.Exec(InsertQuery, sensorId, temperature)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func Open(config string) Database {
	db, err := sql.Open("postgres", config)
	if err != nil {
		log.Fatal(err)
	}

	Ping(db)

	return Database{db}
}

func Ping(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Postgres: ping is ok")
	}
}
