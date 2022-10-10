package pkg

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
    "os"
)

var (
    DatabaseMysqlUser = os.Getenv("DATABASE_MYSQL_USER")
    DatabaseMysqlPassword = os.Getenv("DATABASE_MYSQL_PASSWORD")
    DatabaseMysqlHost = os.Getenv("DATABASE_MYSQL_HOST")
    DatabaseMysqlPort = os.Getenv("DATABASE_MYSQL_PORT")
    DatabaseMysqlSchema = os.Getenv("DATABASE_MYSQL_SCHEMA")
)

func GetMysqlConfig() string {
    return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DatabaseMysqlUser,DatabaseMysqlPassword,DatabaseMysqlHost,DatabaseMysqlPort,DatabaseMysqlSchema) + "?parseTime=true&loc=Asia%2FTokyo"
}

var DatabaseMysql *sqlx.DB

type ChatRepository struct {}

func (c *ChatRepository) Insert(body string) {
    _, err := DatabaseMysql.Exec(`INSERT INTO chat (body) VALUES (?) `, body)
    if err != nil {
        fmt.Println(err)
    }
}
