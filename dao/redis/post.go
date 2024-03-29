package redis

import (
	"forum/vo"
	"github.com/go-redis/redis"
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

// GetPostVoteData 根据ids查询每篇帖子的投赞成票的数据 (只统计赞成票，不包含反对)
func GetPostVoteData(ids []string) (data []int64, err error) {
	data = make([]int64, 0, len(ids))
	//for _, id := range ids {
	//	key := getRedisKey(KeyPostVotedZSetPrefix + id)
	//	//统计每篇帖子赞成票的数量
	//	v := rdb.ZCount(key, "1", "1").Val()
	//	data = append(data, v)
	//}
	//使用pipeline一次发送多条命令，减少RTT
	pipeline := client.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPrefix + id)
		pipeline.ZCount(key, "1", "1")
	}
	cmders, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return
}
