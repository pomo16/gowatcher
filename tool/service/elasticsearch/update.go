package elasticsearch

import (
	"context"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gowatcher/go_analyze/consts"
	"gowatcher/go_analyze/model"
	"gowatcher/go_analyze/utils"
)

//FixBrokenTime 修复缺省时间
func FixBrokenTime(ctx context.Context, comment *model.Comment) error {

	publishTimeStamp := utils.GetTimeStampByTimeStr(comment.PublishTime)
	crawlTimeStamp := utils.GetTimeStampByTimeStr(comment.CrawlTime)

	result, err := elasticClient.Update().
		Index(consts.ESTempIndex).
		Id(comment.MainId).
		Script(elastic.NewScript("ctx._source.publish_timestamp = params.publishTimeStamp; ctx._source.crawl_timestamp = params.crawlTimeStamp").
			Params(map[string]interface{}{
				"publishTimeStamp": publishTimeStamp,
				"crawlTimeStamp":   crawlTimeStamp,
			})).
		Do(ctx)

	if err != nil {
		logrus.Warnf("update comment error: %+v", err)
		return err
	}

	logrus.Infof("result: %+v", result)
	return nil
}

//FixBrokenTitle 修复缺省标题
func FixBrokenTitle(ctx context.Context, comment *model.Comment) error {

	title, content := utils.SplitTitleAndContent(comment.Content)

	_, err := elasticClient.Update().
		Index(consts.ESTempIndex).
		Id(comment.MainId).
		Script(elastic.NewScript("ctx._source.title = params.title; ctx._source.content = params.content").
			Params(map[string]interface{}{
				"title":   title,
				"content": content,
			})).
		Do(ctx)

	if err != nil {
		logrus.Warnf("update comment error: %+v", err)
		return err
	}

	return nil
}
