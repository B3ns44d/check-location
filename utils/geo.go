package utils

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func GetGeo() fiber.Handler {
	client := fasthttp.Client{}
	return func(c *fiber.Ctx) error {
		ip := c.Params("ip", c.IP())
		fields := c.Params("fields")
		req := fasthttp.AcquireRequest()
		res := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseRequest(req)
		defer fasthttp.ReleaseResponse(res)

		req.SetRequestURI("http://ip-api.com/json/" + ip + "?" + fields)

		if err := client.DoTimeout(req, res, 30*time.Second); err != nil {
			return err
		}

		if res.StatusCode() != fiber.StatusOK {
			return errors.New("invalid statuscode")
		}

		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

		return c.Send(res.Body())
	}
}
