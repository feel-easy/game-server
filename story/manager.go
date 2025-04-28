package story

var storyData = map[string]StoryNode{
	"start": {
		ID:        "start",
		Text:      "你遇到了青梅竹马，要打招呼吗？",
		Character: "青梅",
		Options: []StoryOption{
			{Text: "打招呼", NextNodeID: "greet", LoveChange: 10},
			{Text: "假装没看到", NextNodeID: "ignore", LoveChange: -5},
		},
	},
	"greet": {
		ID:        "greet",
		Text:      "她开心地回应了你，好感度上升！",
		Character: "青梅",
		Options:   []StoryOption{},
	},
	"ignore": {
		ID:        "ignore",
		Text:      "她有点失望地走开了，好感度下降...",
		Character: "青梅",
		Options:   []StoryOption{},
	},
}

var currentNodeID = "start"

// func StartGame() StoryNode {
//     currentNodeID = "start"
//     return storyData[currentNodeID]
// }

// func ProgressGame(choiceText string) StoryNode {
//     node := storyData[currentNodeID]
//     for _, opt := range node.Options {
//         if opt.Text == choiceText {
//             currentNodeID = opt.NextNodeID
//             player.AddLove(opt.LoveChange) // 改变好感度
//             break
//         }
//     }
//     return storyData[currentNodeID]
// }
