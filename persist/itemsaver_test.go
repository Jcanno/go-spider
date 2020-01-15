package persist

import (
	"context"
	"encoding/json"
	"spider/model"
	"testing"

	"github.com/olivere/elastic"
)

func TestSave(t *testing.T) {
	// "basicInfo": ["未婚", "26岁", "魔羯座(12.22-01.19)", "165cm", "50kg", "工作地:苏州相城区", "月收入:5-8千", "职业技术教师", "高中及以下"],

	profile := model.Profile{
		Name:       "林YY",
		Marriage:   "未婚",
		Age:        "26岁",
		Xingzuo:    "魔羯座(12.22-01.19)",
		Height:     "165cm",
		Weight:     "50kg",
		Income:     "月收入:5-8千",
		Occupation: "职业技术教师",
		Education:  "高中及以下",
	}
	id, err := save(profile)
	if err != nil {
		panic(err)
	}
	//从ElasticSearch中获取，根据id
	client, err := elastic.NewClient(elastic.SetSniff(false))
	resp, err := client.Get().
		Index("datint_profile").
		Type("zhenai").
		Id(id).Do(context.Background())

	if err != nil {
		panic(err)
	}

	t.Logf("%s", resp.Source) //打印

	//反序列化
	var actual model.Profile
	err = json.Unmarshal([]byte(*resp.Source), &actual)

	if err != nil {
		panic(err)
	}

	//断言
	if actual != profile {
		t.Errorf("got %v; expected %v", actual, profile)
	}
}
