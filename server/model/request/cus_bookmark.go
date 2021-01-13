package request

// 2	cus_web_id	int	NULL	NULL	NO	0			web_id
// 3	cus_user_id	int	NULL	NULL	NO	0			user_id
// 4	path	varchar(255)	utf8mb4	utf8mb4_0900_ai_ci	NO				路径
// 5	query	varchar(255)	utf8mb4	utf8mb4_0900_ai_ci	NO				查询
// 6	title	varchar(255)	utf8mb4	utf8mb4_0900_ai_ci	NO				标题
// 7	description	varchar(255)	utf8mb4	utf8mb4_0900_ai_ci	NO				描述
// 8	cus_group_id	int	NULL	NULL	NO	0			分组ID
// 9	cus_tag_str	varchar(255)	utf8mb4	utf8mb4_0900_ai_ci	NO				标签
type NewBookmark struct {
	Link        string `json:"link"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
