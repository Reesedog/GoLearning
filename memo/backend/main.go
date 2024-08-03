package main

import (
    "memo/database"
    "memo/routers"
)

func main() {
    db := database.InitDB()

    router := routers.SetupRouter(db)
    router.Run(":8080")
}
