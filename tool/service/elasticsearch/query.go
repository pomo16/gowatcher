package elasticsearch

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gowatcher/go_analyze/consts"
	"gowatcher/go_analyze/exceptions"
	"gowatcher/go_analyze/model"
)

//CrawlTimeFilter 返回爬取时间范围条件时间戳比对
func CrawlTimeFilter(boolQ *elastic.BoolQuery, bTime int64, eTime int64) *elastic.BoolQuery {
	return boolQ.Filter(elastic.NewRangeQuery("crawl_timestamp").Gte(bTime).Lte(eTime))
}

//BrokenTimeFilter 缺省时间过滤器
func BrokenTimeFilter(boolQ *elastic.BoolQuery) *elastic.BoolQuery {
	return boolQ.MustNot(elastic.NewExistsQuery("publish_timestamp"))
}

//MainIDFilter ID过滤器
func MainIDFilter(boolQ *elastic.BoolQuery, mainID string) *elastic.BoolQuery {
	return boolQ.Must(elastic.NewTermQuery("main_id", mainID))
}

//BrokenTitleFilter 缺省标题过滤器
func BrokenTitleFilter(boolQ *elastic.BoolQuery) *elastic.BoolQuery {
	return boolQ.MustNot(elastic.NewExistsQuery("title"))
}

//QueryCount 根据指定条件获取计数
func QueryCount(ctx context.Context, params *model.CountParams) (*model.Count, error) {
	boolQuery := elastic.NewBoolQuery()
	//boolQuery = CrawlTimeFilter(boolQuery, params.BeginTime, params.EndTime)
	boolQuery = BrokenTitleFilter(boolQuery)

	result, err := elasticClient.Count().
		Index(consts.ESTempIndex).
		Query(boolQuery).
		Do(ctx)

	if err != nil {
		logrus.Warnf("get count error: %v", err)
		return nil, err
	}

	cnt := &model.Count{
		BeginTime: params.BeginTime,
		EndTime:   params.EndTime,
		Num:       result,
	}

	return cnt, nil
}

//GetComments 获取评论
func GetComments(ctx context.Context, params *model.CommentParams) ([]*model.Comment, error) {
	boolQuery := elastic.NewBoolQuery()
	//boolQuery = CrawlTimeFilter(boolQuery, params.BeginTime, params.EndTime)
	//boolQuery = BrokenTimeFilter(boolQuery)
	boolQuery = BrokenTitleFilter(boolQuery)

	result, err := elasticClient.Search().
		Index(consts.ESTempIndex).
		Query(boolQuery).
		From(params.OffSet).
		Size(params.Limit).
		Do(ctx)

	if err != nil {
		logrus.Warnf("get comments error: %v", err)
		return nil, err
	}

	commentList, err := commentOutputter(result)
	if err != nil {
		return nil, err
	}

	return commentList, nil
}

//commentOutputter 评论输出
func commentOutputter(res *elastic.SearchResult) ([]*model.Comment, error) {
	commentList := []*model.Comment{}
	if res.Hits.TotalHits.Value > 0 {
		for _, hit := range res.Hits.Hits {
			comment := &model.Comment{}
			if err := json.Unmarshal(hit.Source, comment); err != nil {
				return nil, exceptions.ErrParseResult
			}
			commentList = append(commentList, comment)
		}
	}
	return commentList, nil
}

func QueryByID(ctx context.Context, mainID string) ([]*model.Comment, error) {
	boolQuery := elastic.NewBoolQuery()
	boolQuery = MainIDFilter(boolQuery, mainID)

	result, err := elasticClient.Search().
		Index(consts.ESTempIndex).
		Query(boolQuery).
		Do(ctx)

	if err != nil {
		logrus.Warnf("get comments error: %v", err)
		return nil, err
	}

	commentList, err := commentOutputter(result)
	if err != nil {
		return nil, err
	}

	return commentList, nil
}
