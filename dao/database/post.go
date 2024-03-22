package database

import (
	"database/sql"
	"forum/model"
	"go.uber.org/zap"
)

func CreatePost(post *model.Post) (err error) {
	err = db.Create(post).Error
	return
}

func GetPostByID(postID string) (post *model.Post, err error) {
	post = new(model.Post)
	err = db.Where("post_id = ?", postID).Find(post).Error
	return
}

func GetPostList(page, size int) ([]model.Post, error) {
	var postList []model.Post
	err := db.Order("create_time DESC").Offset((page - 1) * size).Limit(size).Find(&postList).Error

	if err == sql.ErrNoRows {
		zap.L().Warn("There is no post record in DB.")
	}
	return postList, nil
}

// GetPostListByIDs 根据给定的id列表查询帖子数据
func GetPostListByIDs(ids []string) (postList []model.Post, err error) {

	// 将 []string 转换为 []interface{}
	var interfaceIDs []interface{}
	for _, id := range ids {
		interfaceIDs = append(interfaceIDs, id)
	}

	// 查询
	err = db.Where("id IN (?)", interfaceIDs).Find(&postList).Error

	return

	//sqlStr := `select post_id,title,content,author_id,community_id,create_time
	//from post
	//where post_id in (?)
	//order by FIND_IN_SET(post_id,?)`
	//query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	//if err != nil {
	//	return nil, err
	//}
	//fmt.Println(query)
	//fmt.Println(args)
	//query = db.Rebind(query)
	//err = db.Select(&postList, query, args...)
	//return
}
