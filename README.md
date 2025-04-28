# Game Server Demo (Fiber版)

## 运行步骤
```bash
go mod tidy
go run main.go
```


```bash
game-server/
├── go.mod
├── main.go          // Fiber服务器入口
├── story/
│   ├── story.go     // 剧情结构体定义
│   └── manager.go   // 剧情数据+推进逻辑
├── player/
│   └── player.go    // 玩家状态管理（简单版）
└── README.md        // 简单使用说明
```
