package main

import ( 
  "github.com/gofiber/fiber"
  "github.com/FriendlyUser/finfiber/discord"
  "github.com/FriendlyUser/finfiber/finance"
  "github.com/FriendlyUser/finfiber/nlp"
  "github.com/FriendlyUser/finfiber/twitter"
  "os"
  "strconv"
)

func getenv(key, fallback string) string {
  value := os.Getenv(key)
  if len(value) == 0 {
      return fallback
  }
  return value
}

// add function to check env vars before running
func main() {
  app := fiber.New()
  port, _ := strconv.Atoi(getenv("PORT", "8080"))
  app.Get("/api/v1/tickers", finance.GetTickersPandas)
  app.Get("/api/v1/sentiment", nlp.GetSentiment)
  app.Get("/api/v1/twitter", twitter.GetTwitterSimple)
  app.Post("/api/v1/discord", discord.SendDiscordHook)
  app.Listen(port)
}
