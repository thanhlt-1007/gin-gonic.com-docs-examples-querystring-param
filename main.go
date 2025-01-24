package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func getWelcomeHandler(context *gin.Context) {
    firstname := context.Query("firstname")
    lastname := context.DefaultQuery("lastname", "NA")
    context.JSON(
        http.StatusOK,
        gin.H {
            "firstname": firstname,
            "lastname": lastname,
        },
    )
}

func main() {
    engine := gin.Default()
    engine.GET("/welcome", getWelcomeHandler)
    engine.Run()
}
