package model

type User struct {
	ID        int64  `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`       // 用户名称
	AvatarUrl string `json:"avatar_url,omitempty"` // 头像地址
	SiteAdmin bool   `json:"site_admin,omitempty"` // 是否管理
	Email     string `json:"email,omitempty"`      // 用户邮箱
}

type Project struct {
	ID        int64  `json:"id"`
	FullName  string `json:"full_name"`
	Desc      string `json:"desc"`
	Owner     string `json:"owner"`
	Repo      string `json:"repo"`
	Name      string `json:"name"`
	IsPrivate bool   `json:"is_private"`
	Url       string `json:"url"`
	Ssh       string `json:"ssh"`
}

type Tag struct {
	Name       string `json:"name"`
	CommitID   string `json:"commit_id"`
	CommitTime string `json:"commit_time"`
	CommitName string `json:"commit_name"`
}

type Branch struct {
	Name       string `json:"name"`
	CommitID   string `json:"commit_id"`
	CommitTime string `json:"commit_time"`
	CommitName string `json:"commit_name"`
}
