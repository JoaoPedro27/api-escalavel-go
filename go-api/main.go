package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    // Cria um router com middlewares padrão (logger e recovery)
    r := gin.Default()

    // Rota de health check
    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "status": "OK",
            "message": "API está funcionando!",
        })
    })

    // Rota simples de exemplo
    r.GET("/hello", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Olá, mundo!",
        })
    })

    // Inicia o servidor na porta 8080
    r.Run(":8080") // "localhost:8080"
}