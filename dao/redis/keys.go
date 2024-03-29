package redis

const (
	KeyPrefix              = "myForum:"
	KeyPostTimeZSet        = "post:time"   //zset;帖子及发帖时间
	KeyPostScoreZSet       = "post:score"  //zset帖子及投票分数
	KeyPostVotedZSetPrefix = "post:voted:" //zset;记录用户及投票类型 参数是帖子--post_id
)

// 为key加上前缀
func getRedisKey(key string) string {
	return KeyPrefix + key
}
