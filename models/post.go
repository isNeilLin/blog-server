package models


// 获取全部文章列表
func GetAllPost() ([]*Post, error) {
	var posts []*Post
	err := DB.Find(&posts).Error
	return posts, err
}

// 获取已发布文章列表
func GetPublishPost() ([]*Post, error) {
	return getPostList(true)
}

// 获取私密文章列表
func GetPrivatePost() ([]*Post, error)  {
	return getPostList(false)
}

// 创建文章
func (post *Post)Insert() error {
	return DB.Create(post).Error
}

// 更新文章
func (post *Post)Update() error {
	return DB.Model(post).Updates(map[string]interface{}{
		"title":	post.Title,
		"content":	post.Content,
		"summary":	post.Summary,
		"publish":	post.Publish,
	}).Error
}

// 删除文章
func (post *Post)Delete() error {
	return DB.Delete(post).Error
}

// 根据ID获取文章详情
func (p *Post)GetPostById() (*Post, error){
	var post Post
	err := DB.First(&post, "id = ?", p.ID).Error
	return &post, err
}


func getPostList(publish bool) ([]*Post, error) {
	var posts []*Post
	err := DB.Where("publish = ?", publish).Find(&posts).Error
	return posts, err
}