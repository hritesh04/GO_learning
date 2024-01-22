package main

import (
	"database/sql"
	"log"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	dbname   = "recordings"
  )

  type Album struct {
    ID     int64
    Title  string
    Artist string
    Price  float32
}

func main(){
	connection := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, os.Getenv("DBUSER"), os.Getenv("DBPASS"), dbname)

	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("DB Connected")

	albums, err := albumsByArtist(db,"John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

}

func albumsByArtist(db *sql.DB, name string) ([]Album, error) {
    var albums []Album

    rows, err := db.Query("SELECT * FROM album WHERE artist = $1", name)
    if err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }
    defer rows.Close()

    for rows.Next() {
        var alb Album
        if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
            return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
        }
        albums = append(albums, alb)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }
    return albums, nil
}