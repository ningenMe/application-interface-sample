package pkg

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
    interface_grpc "ningenme/application-interface-sample/server-grpc-go/interface-grpc"
    "os"
    "time"
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

type ChatDto struct {
    Body      string    `db:"body"`
    CreatedAt time.Time `db:"created_at"`
}


func (c *ChatRepository) Select(startAt time.Time, endAt time.Time) []*interface_grpc.Chat {
    rows, err := DatabaseMysql.Queryx(`SELECT body, created_at FROM chat WHERE ? <= created_at AND created_at < ? ORDER BY created_at ASC`,
        startAt.Format(time.RFC3339), endAt.Format(time.RFC3339))
    if err != nil {
        fmt.Println(err)
    }
    defer rows.Close()
    var list []*interface_grpc.Chat

    for rows.Next() {
        dto := &ChatDto{}
        rows.StructScan(dto)
        list = append(list, &interface_grpc.Chat{
            Body:      dto.Body,
            CreatedAt: dto.CreatedAt.Format(time.RFC3339),
        })
    }
    return list
}
