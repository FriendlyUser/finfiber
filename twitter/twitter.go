package twitter
import (
	"github.com/gofiber/fiber"
	twitter "github.com/FriendlyUser/finfiber/twitter/util"
)

// destructure the response to something meaningful
// allow the inclusion of query params, etc ...
func GetTwitterSimple(c *fiber.Ctx) {

	body, err := twitter.MakeRequest()
	if err != nil {
		return
	}
	c.JSON(body)
}
