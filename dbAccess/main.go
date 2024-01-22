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
    host, port, os.Getenv("DBUSER"),os.Getenv("DBPASS"), dbname)

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
	album, err := albumByID(db,2)
	if err != nil {
		log.Fatal(err)
	}
	albID, err := addAlbum(db,Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)
	fmt.Printf("Album found: %v\n",album)
	fmt.Printf("ID of added album: %v\n", albID)
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

func albumByID(db *sql.DB,id int64) (Album, error) {
    var alb Album

    row := db.QueryRow("SELECT * FROM album WHERE id = $1", id)
    if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
        if err == sql.ErrNoRows {
            return alb, fmt.Errorf("albumsById %d: no such album", id)
        }
        return alb, fmt.Errorf("albumsById %d: %v", id, err)
    }
    return alb, nil
}

func addAlbum(db *sql.DB,alb Album) (int64, error) {
	var result int64
    err := db.QueryRow("INSERT INTO album (title, artist, price) VALUES ($1, $2, $3) RETURNING id", alb.Title, alb.Artist, alb.Price).Scan(&result)
	if err != nil {
        return 0, fmt.Errorf("addAlbum: %v", err)
    }
    return result, nil
}