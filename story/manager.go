package story

import "game-server/player"

var storyData = map[string]Scene{
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
			{Text: "问她最近怎么样", NextSceneID: "how_are_you", LoveChange: 5},
		},
	},
	"milk_tea": {
		ID:        "milk_tea",
		Character: "小艾",
		Text:      "奶茶店里，你们聊得很开心。她似乎对你的某些话题特别感兴趣。",
		Options: []Option{
			{Text: "谈论共同的童年回忆", NextSceneID: "childhood_memory", LoveChange: 15},
			{Text: "问她未来的打算", NextSceneID: "future_plan_ai", LoveChange: 10},
			{Text: "安静地喝奶茶", NextSceneID: "quiet_drink", LoveChange: -5},
		},
	},
	"movie": {
		ID:        "movie",
		Character: "小艾",
		Text:      "电影院里灯光昏暗，气氛有些暧昧。她看起来很专注。",
		Options: []Option{
			{Text: "偷偷看她", NextSceneID: "peek_at_her", LoveChange: 5},
			{Text: "和她讨论剧情", NextSceneID: "discuss_plot", LoveChange: 10},
			{Text: "不小心碰到她的手", NextSceneID: "touch_hand", LoveChange: 20},
		},
	},
	"how_are_you": {
		ID:        "how_are_you",
		Character: "小艾",
		Text:      "她笑着说最近还不错，并反问你。",
		Options: []Option{
			{Text: "分享你最近的趣事", NextSceneID: "share_stories", LoveChange: 10},
			{Text: "抱怨一下最近的烦恼", NextSceneID: "complain", LoveChange: -5},
			{Text: "邀请她稍后一起活动", NextSceneID: "invite_later", LoveChange: 15},
		},
	},
	"ignore": {
		ID:        "ignore",
		Character: "小艾",
		Text:      "小艾有点失落地低下了头，默默走开了。",
		Options: []Option{
			{Text: "回头喊她", NextSceneID: "greet", LoveChange: 5},
			{Text: "继续装作不认识，去图书馆", NextSceneID: "go_library", LoveChange: -10},
			{Text: "感到一丝后悔", NextSceneID: "regret", LoveChange: 0},
		},
	},
	"go_library": {
		ID:   "go_library",
		Text: "你来到图书馆，试图专注于学习，但刚才的场景挥之不去。",
		Options: []Option{
			{Text: "巧遇学姐小雪", NextSceneID: "snow_greet", LoveChange: 0},
			{Text: "继续看书", NextSceneID: "study_alone", LoveChange: 0},
		},
	},
	"regret": {
		ID:   "regret",
		Text: "你心里有点不是滋味，但生活还要继续。",
		Options: []Option{
			{Text: "去别处转转", NextSceneID: "walk_around", LoveChange: 0},
		},
	},
	"walk_around": {
		ID:   "walk_around",
		Text: "你在校园里漫步，看到公告栏上贴着一张社团招新海报。",
		Options: []Option{
			{Text: "去看看是什么社团", NextSceneID: "check_club", LoveChange: 0},
			{Text: "没兴趣，离开", NextSceneID: "leave_poster", LoveChange: 0},
		},
	},
	"check_club": {
		ID:   "check_club",
		Text: "原来是摄影社在招新，你看到学姐小雪也在那里帮忙。",
		Options: []Option{
			{Text: "上前和小雪打招呼", NextSceneID: "snow_greet_club", LoveChange: 5},
			{Text: "默默观察一下", NextSceneID: "observe_club", LoveChange: 0},
		},
	},
	// 新增角色和剧情
	"snow_greet": {
		ID:        "snow_greet",
		Character: "小雪",
		Text:      "在图书馆安静的角落，你遇到了安静看书的学姐——小雪。",
		Options: []Option{
			{Text: "主动搭话，问她在看什么书", NextSceneID: "snow_talk_book", LoveChange: 10},
			{Text: "保持沉默，在她附近坐下", NextSceneID: "snow_sit_near", LoveChange: 5},
			{Text: "悄悄离开", NextSceneID: "snow_leave_quietly", LoveChange: -5},
		},
	},
	"snow_talk_book": {
		ID:        "snow_talk_book",
		Character: "小雪",
		Text:      "小雪抬起头，轻轻告诉你书名，并简单介绍了一下内容。她的声音很柔和。",
		Options: []Option{
			{Text: "和她讨论这本书", NextSceneID: "snow_discuss_book", LoveChange: 15},
			{Text: "问她推荐类似的书籍", NextSceneID: "snow_ask_recommendation", LoveChange: 10},
			{Text: "表示感谢后离开", NextSceneID: "snow_thank_leave", LoveChange: 0},
		},
	},
	"snow_sit_near": {
		ID:        "snow_sit_near",
		Character: "小雪",
		Text:      "你安静地坐在她附近，图书馆里只有翻书的沙沙声。她似乎没有注意到你，又或者注意到了但没有表示。",
		Options: []Option{
			{Text: "鼓起勇气再次搭话", NextSceneID: "snow_talk_again", LoveChange: 5},
			{Text: "就这样安静地待着", NextSceneID: "snow_stay_quiet", LoveChange: 0},
			{Text: "感觉无趣，离开", NextSceneID: "snow_leave_bored", LoveChange: -5},
		},
	},
	"snow_discuss_book": {
		ID:        "snow_discuss_book",
		Character: "小雪",
		Text:      "你们就书中的内容展开了讨论，发现彼此有一些共同的见解。她的脸上露出了难得的微笑。",
		Options: []Option{
			{Text: "邀请她下次一起去书店", NextSceneID: "snow_invite_bookstore", LoveChange: 20},
			{Text: "问她关于作者的其他作品", NextSceneID: "snow_ask_other_works", LoveChange: 10},
		},
	},
	"snow_greet_club": {
		ID:        "snow_greet_club",
		Character: "小雪",
		Text:      "你上前和小雪打招呼，她有些惊讶，但还是礼貌地回应了你。",
		Options: []Option{
			{Text: "询问关于摄影社的事情", NextSceneID: "ask_club_details", LoveChange: 10},
			{Text: "称赞她的摄影作品（如果看到的话）", NextSceneID: "praise_photo", LoveChange: 15},
			{Text: "简单寒暄几句", NextSceneID: "simple_chat_club", LoveChange: 5},
		},
	},
	"ask_club_details": {
		ID:        "ask_club_details",
		Character: "小雪",
		Text:      "小雪耐心地向你介绍了摄影社的活动和要求。",
		Options: []Option{
			{Text: "表示有兴趣加入", NextSceneID: "join_club", LoveChange: 10},
			{Text: "再考虑一下", NextSceneID: "consider_club", LoveChange: 0},
		},
	},
	// ... (此处可以继续添加更多的小艾和小雪的后续剧情)

	// 结局场景示例
	"childhood_memory": {
		ID:        "childhood_memory",
		Character: "小艾",
		Text:      "聊起童年的趣事，你们都笑了。感觉彼此的距离更近了。也许，有些感觉从未改变。",
		Options:   []Option{},
	},
	"touch_hand": {
		ID:        "touch_hand",
		Character: "小艾",
		Text:      "黑暗中，你的手不小心碰到了她的手。她似乎颤抖了一下，但没有移开。空气中弥漫着一丝甜蜜。",
		Options:   []Option{},
	},
	"snow_invite_bookstore": {
		ID:        "snow_invite_bookstore",
		Character: "小雪",
		Text:      "她犹豫了一下，然后轻轻点了点头，答应了你的邀请。你感到一阵欣喜。",
		Options:   []Option{},
	},
	"study_alone": {
		ID:      "study_alone",
		Text:    "你决定专注于学习，度过了一个平静的下午。",
		Options: []Option{},
	},
	"leave_poster": {
		ID:      "leave_poster",
		Text:    "你对社团活动不感兴趣，继续在校园里闲逛。",
		Options: []Option{},
	},
	"observe_club": {
		ID:      "observe_club",
		Text:    "你在一旁观察了一会儿摄影社的活动，然后离开了。",
		Options: []Option{},
	},
	"random_event": {
		ID:        "random_event",
		Character: "事件",
		Text:      "突然一阵电话铃声响起，是某个熟悉的名字。",
		Options: []Option{
			{Text: "接电话", NextSceneID: "call_received", LoveChange: 0},
			{Text: "忽略电话", NextSceneID: "call_ignored", LoveChange: 0},
		},
	},
	"call_received": {
		ID:        "call_received",
		Character: "系统",
		Text:      "电话那头是你的好友，约你晚上一起打游戏。",
		Options: []Option{
			{Text: "答应", NextSceneID: "game_night", LoveChange: 0},
			{Text: "拒绝", NextSceneID: "study_alone", LoveChange: 0},
		},
	},
	"call_ignored": {
		ID:        "call_ignored",
		Character: "系统",
		Text:      "你挂断了电话，世界清静了。",
		Options: []Option{
			{Text: "继续当前活动", NextSceneID: "start", LoveChange: 0},
		},
	},
}

var currentNodeID = "start"

func StartGame() Scene {
	currentNodeID = "start"
	return storyData[currentNodeID]
}

func ProgressGame(choiceText string) Scene {
	node := storyData[currentNodeID]
	for _, opt := range node.Options {
		if opt.Text == choiceText {
			currentNodeID = opt.NextSceneID
			player.UpdateLove(opt.LoveChange) // 改变好感度
			break
		}
	}
	return storyData[currentNodeID]
}

// 获取随机事件
func GetRandomEvent() Scene {
	// 简单起见，这里直接返回一个固定的随机事件
	// 在实际应用中，你可以从多个随机事件中选择一个
	return Scene{
		ID:        "random_event",
		Character: "意外事件",
		Text:      "突然一阵电话铃声响起，是你的好友打来的。",
		Options: []Option{
			{Text: "接电话", NextSceneID: "start", LoveChange: 5},
			{Text: "不接", NextSceneID: "start", LoveChange: -5},
		},
	}
}
