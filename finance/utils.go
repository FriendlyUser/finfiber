package finance

// http://localhost:3000/api/v1/book?quotes=BB.TO,ACB.TO
import (
	"github.com/gofiber/fiber"
	"unsafe"
)

func queryMulti(ctx *fiber.Ctx, key string) (values []string) {
	valuesBytes := ctx.Fasthttp.QueryArgs().PeekMulti(key)
	values = make([]string, len(valuesBytes))
	for i, v := range valuesBytes {
		values[i] = getString(v)
	}
	return values
}

// #nosec G103
// getString converts byte slice to a string without memory allocation.
// See https://groups.google.com/forum/#!msg/Golang-Nuts/ENgbUzYvCuU/90yGx7GUAgAJ .
var getString = func(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}