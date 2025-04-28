package story

type StoryNode struct {
    ID        string       `json:"id"`
    Text      string       `json:"text"`
    Character string       `json:"character"`
    Options   []StoryOption `json:"options"`
}

type StoryOption struct {
    Text        string `json:"text"`
    NextNodeID  string `json:"next_node_id"`
    LoveChange  int    `json:"love_change"`
}
