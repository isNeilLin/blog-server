package models

import (
	"strings"
	"blog/utils"
)

// 获取标签列表以及对应的文章列表
func GetTags()([]map[string]interface{}, error)  {
	var (
		tags 	[]*Tag
		result 	[]map[string]interface{}
	)
	err := DB.Find(&tags).Error
	if err != nil {
		return result, err
	}
	for _,tag := range tags {
		posts, err := GetPostsByTag(tag.ID)
		if err != nil {
			return result, err
		} else {
			result = append(result,map[string]interface{}{
				"name": 	tag.Name,
				"color":	tag.Color,
				"posts":	posts,
			})
		}
	}
	return result, err
}

// 根据标签获取对应的文章列表
func GetPostsByTag(tagId uint) ([]*Post, error){
	var (
		posts 	[]*Post
	)
	err := DB.Joins("JOIN post_tag on post.id = post_tag.post_id").Where("post_tag.tag_id = ?",tagId).Find(&posts).Error
	return posts,err
}

// 添加标签
func (tag *Tag)Insert() error {
	return DB.Create(tag).Error
}

// 更新标签
func (tag *Tag)Update() error {
	return DB.Model(tag).Update(map[string]interface{}{
		"color": tag.Color,
		"name":  tag.Name,
	}).Error
}

// 删除标签
func (tag *Tag)Delete() error {
	return DB.Delete(tag).Error
}


// 为单个文章批量添加标签
func CreateTags(tags string, post *Post) error {
	if len(tags) == 0 {
		return nil
	}
	sliceTags := strings.Split(tags,",")
	for _,tag := range sliceTags {
		tagID,err := utils.StringToUint(tag)
		if err != nil {
			return err
		}
		pt := PostTag{
			PostId: post.ID,
			TagId: 	tagID,
		}
		err = pt.Insert()
		if err != nil {
			return err
		}
	}
	return nil
}

// 根据文章ID删除对应的标签映射
func DeleteTagsFromPostId(id uint) error {
	return DB.Delete(&PostTag{}, "post_id = ?", id).Error
}

// 添加文章和标签的关系映射
func (pt *PostTag)Insert() error  {
	return DB.Create(pt).Error
}
