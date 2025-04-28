package story

import "game-server/player"

type Scene struct {
	ID        string
	Character string
	Text      string
	Options   []Option
}

type Option struct {
	Text        string
	NextSceneID string
	LoveChange  int
}

var scenes map[string]Scene
var currentSceneID string

// 每个女主有独立的好感度
var characters = map[string]int{
	"小艾": 0,
	"小雪": 0,
	"小美": 0,
}

func init() {
	scenes = map[string]Scene{
		"start": {
			ID:        "start",
			Character: "小艾",
			Text:      "你在公园散步，偶遇了你的青梅竹马——小艾。",
			Options: []Option{
				{Text: "打招呼", NextSceneID: "greet", LoveChange: 10},
				{Text: "装作没看见", NextSceneID: "ignore", LoveChange: -5},
			},
		},
		"greet": {
			ID:        "greet",
			Character: "小艾",
			Text:      "小艾开心地回应了你。",
			Options: []Option{
				{Text: "一起去喝奶茶", NextSceneID: "milk_tea", LoveChange: 10},
				{Text: "约她晚上看电影", NextSceneID: "movie", LoveChange: 15},
			},
		},
		"ignore": {
			ID:        "ignore",
			Character: "小艾",
			Text:      "小艾有点失落地低下了头。",
			Options: []Option{
				{Text: "回头喊她", NextSceneID: "greet", LoveChange: 5},
				{Text: "继续装作不认识", NextSceneID: "alone", LoveChange: -10},
			},
		},
		// 新增角色和剧情
		"snow_greet": {
			ID:        "snow_greet",
			Character: "小雪",
			Text:      "你遇到了安静的学姐——小雪，她似乎有些孤单。",
			Options: []Option{
				{Text: "主动搭话", NextSceneID: "snow_talk", LoveChange: 10},
				{Text: "保持沉默", NextSceneID: "snow_ignore", LoveChange: -5},
			},
		},
		"snow_talk": {
			ID:        "snow_talk",
			Character: "小雪",
			Text:      "小雪抬起头，轻轻地回应你，眼神里似乎有一丝温暖。",
			Options: []Option{
				{Text: "邀请她一起去看电影", NextSceneID: "snow_movie", LoveChange: 15},
				{Text: "只是聊聊生活", NextSceneID: "snow_chat", LoveChange: 10},
			},
		},
		"snow_ignore": {
			ID:        "snow_ignore",
			Character: "小雪",
			Text:      "小雪有些失落，默默走开。",
			Options: []Option{
				{Text: "重新找机会搭话", NextSceneID: "snow_talk", LoveChange: 5},
				{Text: "不再理会她", NextSceneID: "snow_alone", LoveChange: -10},
			},
		},
		// 增加随机事件
		"random_event": {
			ID:        "random_event",
			Character: "事件",
			Text:      "突然一阵电话铃声响起，是某个熟悉的名字。",
			Options: []Option{
				{Text: "接电话", NextSceneID: "call_received", LoveChange: 0},
				{Text: "忽略电话", NextSceneID: "call_ignored", LoveChange: 0},
			},
		},
	}
}

func StartGame() Scene {
	currentSceneID = "start"
	return scenes[currentSceneID]
}

func ProgressGame(choice string) Scene {
	current := scenes[currentSceneID]
	var nextID string
	for _, opt := range current.Options {
		if opt.Text == choice {
			nextID = opt.NextSceneID
			player.UpdateLove(opt.LoveChange)
			break
		}
	}
	if nextID == "" {
		// 没找到选项，停在当前
		return current
	}
	currentSceneID = nextID
	return scenes[currentSceneID]
}

func TriggerRandomEvent() Scene {
	// 随机选择一个事件
	events := []string{"random_event"}
	nextEventID := events[0]
	return scenes[nextEventID]
}
