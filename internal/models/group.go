package models

type Group struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Village   string `json:"village"`
	District  string `json:"district"`
	LeaderID  int    `json:"leader_id"`
	CreatedAt string `json:"created_at"`
}

type GroupWithLeader struct {
	Group
	LeaderName  string `json:"leader_name"`
	LeaderPhone string `json:"leader_phone"`
	MemberCount int    `json:"member_count"`
}
