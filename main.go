package main

import (
    "exprorstate-api/db"
    "exprorstate-api/server"
)

func main() {
    db.Init()
    server.Init()
    db.Close()
}
