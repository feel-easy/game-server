package main

import (
	"game-server/player"
	"game-server/story"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	// CORS，允许跨域（可选，因为静态托管后基本不会出现跨域了）
	app.Use(cors.New())

	// ✨ 静态资源托管：访问 / 时，加载 frontend 目录下的index.html
	app.Static("/", "./frontend")

	// API接口
	app.Get("/start", func(c *fiber.Ctx) error {
		player.InitPlayer()
		scene := story.StartGame()
		return c.JSON(scene)
	})

	app.Post("/progress", func(c *fiber.Ctx) error {
		var req struct {
			Choice string `json:"choice"`
		}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).SendString("Invalid input")
		}
		nextScene := story.ProgressGame(req.Choice)
		return c.JSON(nextScene)
	})

	app.Get("/state", func(c *fiber.Ctx) error {
		state := player.GetCurrentState()
		return c.JSON(state)
	})

	app.Get("/random_event", func(c *fiber.Ctx) error {
		// 这里简单返回一个随机事件场景
		nextScene := story.GetRandomEvent()
		return c.JSON(nextScene)
	})

	app.Listen(":3000")
}
