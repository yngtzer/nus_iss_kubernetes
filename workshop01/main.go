package main

import (
  "net/http"

  "github.com/gin-gonic/gin"

  "math/rand"
)

var router *gin.Engine

func main() {

	list_of_texts := []string{
		"Logic will get you from A to B. Imagination will take you everywhere.",
		"There are 10 kinds of people. Those who know binary and those who don't.",
		"There are two ways of constructing a software design. One way is to make it so simple that there are obviously no deficiencies and the other is to make it so complicated that there are no obvious deficiencies.",
		"It's not that I'm so smart, it's just that I stay with problems longer.",
		"It is pitch dark. You are likely to be eaten by a grue.",
  }

  // Set the router as the default one provided by Gin
  router = gin.Default()

  // Process the templates at the start so that they don't have to be loaded
  // from the disk again. This makes serving HTML pages very fast.
  router.LoadHTMLGlob("templates/*")
  router.StaticFile("/golang-image.jpeg", "./assets/golang-image.jpeg")

  // Define the route for the index page and display the index.html template
  // To start with, we'll use an inline route handler. Later on, we'll create
  // standalone functions that will be used as route handlers.
  router.GET("/", func(c *gin.Context) {

    // Call the HTML method of the Context to render a template
    c.HTML(
      // Set the HTTP status to 200 (OK)
      http.StatusOK,
      // Use the index.html template
      "index.html",
      // Pass the data that the page uses (in this case, 'title')
      gin.H{
        "title": "Home Page",
		"randomize_text": list_of_texts[rand.Intn(len(list_of_texts))],
      },
    )

  })

  // Start serving the application
  router.Run()

}