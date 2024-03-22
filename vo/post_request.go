package vo

const (
	OrderTime  = "time"
	OrderScore = "score"
)

// ParamPostList 获取帖子列表参数
type ParamPostList struct {
	Page  int64  `json:"page" form:"page"`
	Size  int64  `json:"size" form:"size"`
	Order string `json:"order" form:"order"`
}
