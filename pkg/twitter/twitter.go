package twitter
import (
	"github.com/gofiber/fiber/v2"
	twitter "github.com/FriendlyUser/finfiber/pkg/twitter/util"
)

// destructure the response to something meaningful
// allow the inclusion of query params, etc ...
func GetTwitterSimple(c *fiber.Ctx) error {

	body, err := twitter.MakeRequest()
	if err != nil {
		return err
	}
	return c.JSON(body)
}
