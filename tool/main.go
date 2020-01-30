package main

import (
	"context"
	"fmt"
	"gowatcher/go_analyze/model"
	"gowatcher/go_analyze/service/elasticsearch"
)

func Init() {
	elasticsearch.InitElasticSearch()
}

func Task() {

	ctx := context.Background()

	cntParams := &model.CountParams{
		BeginTime: 0,
		EndTime:   0,
	}

	count, _ := elasticsearch.QueryCount(ctx, cntParams)
	fmt.Println(int(count.Num))

	//listParams := &model.CommentParams{
	//	OffSet: 0,
	//	Limit:  int(count.Num),
	//}
	//
	//list, _ := elasticsearch.GetComments(ctx, listParams)
	//for k, v := range list {
	//	//fmt.Printf("%v, %+v\n", k, v)
	//
	//	fmt.Println(k)
	//	elasticsearch.FixBrokenTitle(ctx, v)
	//}
}

func main() {
	Init()
	Task()
}
