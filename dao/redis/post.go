package redis

import (
	"forum/vo"
)

// GetPostIDsInOrder 从redis查询id记录
func GetPostIDsInOrder(p *vo.ParamPostList) ([]string, error) {
	//从redis获取id
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == vo.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	start := (p.Page - 1) * p.Size
	stop := start + p.Size + 1
	//按照分数或者时间查询指定数量的id记录，由大到小排序
	return client.ZRevRange(key, start, stop).Result()
}
