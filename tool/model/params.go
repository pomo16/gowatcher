package model

type CountParams struct {
	BeginTime int64 //开始时间
	EndTime   int64 //结束时间
}

type CommentParams struct {
	OffSet int
	Limit  int
}
