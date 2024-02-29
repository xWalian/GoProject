package controller
import (
	"database/sql"
	"log"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)
db, err := sql.Open("postgres", "user=youruser password=yourpassword dbname=yourdb sslmode=disable")
if err != nil {
    log.Fatal(err)
}
defer db.Close()
