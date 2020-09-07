// pulls data from yahoo finance and returns data in json format
package discord

// http://localhost:3000/api/v1/book?quotes=BB.TO,ACB.TO
import (
	"log"
	"os"
	"github.com/gofiber/fiber"
	"encoding/json"
	"net/http"
	"bytes"
)


type DiscordEmbed struct {
	Title string `json:"title"`
	Url string `json:"url"`
	Description string `json:"description"`
}

type DiscordWebhook struct {
	Content string `json:"content" xml:"content" form:"content" query:"content"`
	Embeds []DiscordEmbed `json:"embeds" xml:"embeds" form:"embeds" query:"embeds"`
}

// Run tests with the following curl commands

// curl -X POST -H "Content-Type: application/json" --data "{\"content\":\"john\",\"pass\":\"doe\"}" localhost:3000

// curl -X POST -H "Content-Type: application/xml" --data "<discord><content>john</content><pass>doe</pass></discord>" localhost:3000

// curl -X POST -H "Content-Type: application/x-www-form-urlencoded" --data "content=john&pass=doe" localhost:3000

// curl -X POST -F content=john -F pass=doe http://localhost:3000

// curl -X POST "http://localhost:3000/?content=john&pass=doe"

func SendDiscordHook(c *fiber.Ctx) {

	discordWebhook := new(DiscordWebhook)

	if err := c.BodyParser(discordWebhook); err != nil {
		log.Fatal(err)
	}

	webhookData, err := json.Marshal(discordWebhook)
	if err != nil {
		log.Fatal(err)
	}

	discordUrl := os.Getenv("DISCORD_WEBHOOK")
	resp, err := http.Post(discordUrl, "application/json", bytes.NewBuffer(webhookData))
	log.Println(resp)
	log.Println(discordWebhook.Content) // john
	log.Println(discordWebhook.Embeds) // doe
	c.JSON(discordWebhook)
}
