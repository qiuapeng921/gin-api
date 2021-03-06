package db

import (
	"context"
	"encoding/json"
	"fmt"
	"gin-api/app/utility/app"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
	"strconv"
)

type ElasticClient struct {
	EsCon *elastic.Client
}

var EsClient ElasticClient

func InitElastic() {
	host := os.Getenv("ELASTIC_HOST")
	var err error
	EsClient.EsCon, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(host))

	app.Error("elastic初始化失败", err)

	ctx := context.Background()
	_, _, err = EsClient.EsCon.Ping(host).Do(ctx)

	app.Error("elastic连接检测", err)

	fmt.Println("elastic连接成功")
}

//创建
func (es *ElasticClient) CreateIndex(index string) bool {
	_, err := es.EsCon.CreateIndex(index).Do(context.Background())

	app.Error("创建索引失败", err)

	return true
}

func (es *ElasticClient) DeleteIndex(index string) bool {
	_, err := es.EsCon.DeleteIndex(index).Do(context.Background())

	app.Error("删除索引失败", err)

	return true
}

func (es *ElasticClient) IsExistsIndex(index string) bool {
	exists, err := es.EsCon.IndexExists(index).Do(context.Background())
	if err != nil {
		return false
	}
	if !exists {
		return false
	}
	return true
}

// 往索引添加数据
// json字符串导入 json = `{"id":"1", "name":"admin"}`
// struct结构体导入 Task{id: "1", name: "admin"}
func (es *ElasticClient) Insert(index string, bodyJSON interface{}) bool {
	if !es.IsExistsIndex(index) {
		if !es.CreateIndex(index) {
			return false
		}
	}
	_, err := es.EsCon.Index().Index(index).Type(index).BodyJson(bodyJSON).Do(context.Background())

	app.Error("往索引添加数据失败", err)

	return true
}

//批量插入
func (es *ElasticClient) BatchInsert(index string, type_ string, datas ...interface{}) int {
	if !es.IsExistsIndex(index) {
		if !es.CreateIndex(index) {
			return 0
		}
	}
	bulkRequest := es.EsCon.Bulk()
	for i, data := range datas {
		doc := elastic.NewBulkIndexRequest().Index(index).Type(type_).Id(strconv.Itoa(i)).Doc(data)
		bulkRequest = bulkRequest.Add(doc)
	}

	response, err := bulkRequest.Do(context.TODO())

	app.Panic(err)

	failed := response.Failed()
	iter := len(failed)
	return iter
}

// 删除一条数据
func (es *ElasticClient) DeleteData(index, id string) bool {
	_, err := es.EsCon.Delete().Index(index).Type(index).Id(id).Do(context.Background())

	app.Error("删除索引数据失败", err)

	return true
}

// 更新数据
func (es *ElasticClient) UpdateData(index, id string, updateMap map[string]interface{}) bool {
	res, err := es.EsCon.Update().
		Index(index).
		Type(index).
		Id(id).
		Doc(updateMap).
		FetchSource(true).
		Do(context.Background())
	if err != nil {
		return false
	}
	if res == nil {
		return false
	}
	if res.GetResult == nil {
		return false
	}
	data, _ := json.Marshal(res.GetResult.Source)
	log.Printf("<Update> update success. data:%s", data)
	return true
}

//查找
func (es *ElasticClient) Gets(index, id string) *elastic.GetResult {
	result, err := es.EsCon.Get().Index(index).Type(index).Id(id).Do(context.Background())

	app.Error("查找索引数据失败", err)

	return result
}

// 搜索
func (es *ElasticClient) Query(index string, query ...string) []*elastic.SearchHit {
	var result *elastic.SearchResult
	var err error
	if len(query) > 0 {
		q := elastic.NewQueryStringQuery(query[0])
		result, err = es.EsCon.Search(index).Type(index).Query(q).Do(context.Background())
	} else {
		// 字段相等
		result, err = es.EsCon.Search(index).Type(index).Do(context.Background())
	}

	app.Error("搜索索引数据失败", err)

	return result.Hits.Hits
}

// 分页 List("test",map[string]string{"page":1"size":10,"query":"test","sort_type":"_source","sort_type":"desc"})
func (es *ElasticClient) List(index string, params map[string]string) []*elastic.SearchHit {
	page := 1
	size := 10
	if pg, ok := params["page"]; ok {
		paramsPage, _ := strconv.Atoi(pg)
		page = paramsPage
	}
	if pg, ok := params["size"]; ok {
		paramsSize, _ := strconv.Atoi(pg)
		size = paramsSize
	}

	search := es.EsCon.Search(index).Type(index)
	if q, ok := params["query"]; ok {
		query := elastic.NewQueryStringQuery(q)
		search.Query(query)
	}

	sortType := true
	if s, ok := params["sort_type"]; ok {
		if s == "desc" {
			sortType = false
		}
	}

	//排序类型 desc asc es 中只使用 bool 值  true or false
	if s, ok := params["sort"]; ok {
		search.Sort(s, sortType)
	}
	searchResult, err := search.Size(size).From((page - 1) * size).Do(context.Background())

	app.Error("func list error", err)

	return searchResult.Hits.Hits
}
