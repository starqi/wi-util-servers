package main

import (
    "net/http"
    "log"
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
    "strings"
)

func main() {
    router := gin.Default()
    router.GET("/chat", chat)
    router.Run()
}

var upgrader = websocket.Upgrader {
    ReadBufferSize: 1024,
    WriteBufferSize: 1024,
    CheckOrigin: checkOrigin,
}

// WS no CORS
func checkOrigin(r *http.Request) bool {
    origin := r.Header.Get("origin")
    log.Print("Origin = ", origin)
    return strings.Index(origin, "http://localhost") == 0 || strings.Index(origin, "localhost") == 0
}

func chat(c *gin.Context) {
    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        log.Print("Chat init failed! ", err)
        return
    }
    go func() {
        for {
            messageType, p, err := conn.ReadMessage()
            if err != nil {
                log.Print("Read message failed! ", err)
                return
            }
            if messageType == websocket.TextMessage {
                hi := string(p)
                log.Print("Message: ", hi)
            }
        }
    }()
    c.Status(http.StatusOK);
}
