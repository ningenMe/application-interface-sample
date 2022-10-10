package pkg

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
)

const (
    DatabaseMysqlUser = "root"
    DatabaseMysqlPassword = "password"
    DatabaseMysqlHost = "host.docker.internal"
    DatabaseMysqlPort = "9190"
    DatabaseMysqlSchema = "sample"
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
