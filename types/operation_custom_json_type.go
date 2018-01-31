package types

type FollowOperation struct {
	Follower  string   `json:"follower"`
	Following string   `json:"following"`
	What      []string `json:"what"`
	Account   string   `json:"account"`
	Author    string   `json:"author"`
	Permlink  string   `json:"permlink"`
}

/*type ReblogOperation struct {
	Account  string `json:"account"`
	Author   string `json:"author"`
	Permlink string `json:"permlink"`
}*/
