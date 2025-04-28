package player

var playerState = &PlayerState{}

func InitPlayer() {
	playerState = &PlayerState{
		Love:    0,
		XiaoAi:  0,
		XiaoXue: 0,
		XiaoMei: 0,
	}
}

// 玩家状态
type PlayerState struct {
	Love    int `json:"love"`    // 总好感度
	XiaoAi  int `json:"xiaoAi"`  // 小艾的好感度
	XiaoXue int `json:"xiaoXue"` // 小雪的好感度
	XiaoMei int `json:"xiaoMei"` // 小美的好感度
}

// 更新好感度
func UpdateLove(change int) {
	playerState.Love += change
}

// 获取所有好感度
func GetState() map[string]int {
	return map[string]int{
		"love":    playerState.Love,
		"xiaoAi":  playerState.XiaoAi,
		"xiaoXue": playerState.XiaoXue,
		"xiaoMei": playerState.XiaoMei,
	}
}

// 更新角色好感度
func UpdateCharacterLove(character string, change int) {
	switch character {
	case "小艾":
		playerState.XiaoAi += change
	case "小雪":
		playerState.XiaoXue += change
	case "小美":
		playerState.XiaoMei += change
	}
}

// 获取当前状态
func GetCurrentState() *PlayerState {
	return playerState
}
