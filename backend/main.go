package main

import (
	"game-server/player"
	"game-server/story"
	"io"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/google/uuid"
)

func main() {
	app := fiber.New()
	file, _ := os.OpenFile("logs/test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	iw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(iw)
	log.SetLevel(log.LevelInfo)

	app.Use(logger.New())
	// CORS，允许跨域（可选，因为静态托管后基本不会出现跨域了）
	app.Use(cors.New())

	// ✨ 静态资源托管：访问 / 时，加载 frontend 目录下的index.html
	app.Static("/", "../frontend/game")
	app.Static("/admin", "../frontend/game-admin/dist")

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

	// 获取所有节点
	app.Get("/api/nodes", func(c *fiber.Ctx) error {
		return c.JSON(nodes)
	})

	// 创建一个新节点
	app.Post("/api/nodes", func(c *fiber.Ctx) error {
		var node Node
		if err := c.BodyParser(&node); err != nil {
			return err
		}

		node.ID = uuid.New().String()
		nodes = append(nodes, node)
		return c.Status(fiber.StatusCreated).JSON(node)
	})

	// 获取所有边
	app.Get("/api/edges", func(c *fiber.Ctx) error {
		return c.JSON(edges)
	})

	// 创建一个新边
	app.Post("/api/edges", func(c *fiber.Ctx) error {
		var edge Edge
		if err := c.BodyParser(&edge); err != nil {
			return err
		}

		edges = append(edges, edge)
		return c.Status(fiber.StatusCreated).JSON(edge)
	})

	app.Listen(":3000")
}

type Node struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Edge struct {
	From string `json:"from"`
	To   string `json:"to"`
}

var nodes []Node
var edges []Edge
