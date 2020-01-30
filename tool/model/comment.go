package model

type Comment struct {
	CommentId        string `json:"comment_id"`
	MainId           string `json:"main_id"`
	Title            string `json:"title"`
	Content          string `json:"content"`
	Rating           string `json:"rating"`
	Version          string `json:"version"`
	PublishTime      string `json:"publish_time"`
	PublishTimeStamp int64  `json:"publish_timestamp"`
	CrawlTime        string `json:"crawl_time"`
	CrawlTimeStamp   int64  `json:"crawl_timestamp"`
}
