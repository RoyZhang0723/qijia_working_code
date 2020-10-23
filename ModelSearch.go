//ModelSearch Interface
//code by Zhang Xinyu
//2020.9.11

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)
import "github.com/olivere/elastic/v7"


//实现类
type AllKeys struct {
	Model_id                        int    `json:"model_id"`
	Model_type                      string `json:"model_type"`
	Renderer                        string `json:"renderer"`
	Software                        string `json:"software"`
	Title                           string `json:"title"`
	Keyword                         string `json:"keyword"`
	Pointsnum                       int    `json:"pointsnum"`
	Planesnum                       int    `json:"planesnum"`
	Size                            string `json:"size"`
	Style                           string `json:"style"`
	Cost                            int    `json:"cost"`
	Downloadurl                     string `json:"download_url"`
	Status                          string `json:"status"`
	List_thumbnails                 string `json:"list_thumbnails"`
	Detail_images                   string `json:"detail_images"`
	Create_time                     string `json:"create_time"`
	Update_time                     string `json:"update_time"`
	Pano_url                        string `json:"pano_url"`
	Cloud_sdxy_model_download_count  int   `json:"cloud_sdxy_model_download_count"`
	Cloud_sdxy_model_collection_count int   `json:"cloud_sdxy_model_collection_count"`
	Cloud_sdxy_model_view_count     int    `json:"cloud_sdxy_model_view_count"`
}

type Result struct {
	Keys []resultKeys         `json:"results"`
	TotalNum int         `json:"total_num"`
	TotalPage int        `json:"total_page"`
}

type SearchKeys struct {
	Style string          `json:"style"`
	Renderer string       `json:"renderer"`
	Isfree string         `json:"isfree"`
	Order string          `json:"order"`
	PageSize int          `json:"pageSize"`
	Page int              `json:"page"`
}

//返回的keys值
type resultKeys struct {
	ModelId int           `json:"model_id"`
	ModelType string      `json:"model_type"`
	Title string          `json:"title"`
	Thumbnail string      `json:"list_thumbnails"`
}

func OrderByDownload(a []AllKeys, lo int, hi int)  {
	if lo > hi {
		return
	}
	p := Partition(a, lo, hi)
	OrderByDownload(a,lo,p-1)
	OrderByDownload(a,p+1,hi-1)
}

func OrderByViewCount(a []AllKeys, lo int, hi int)  {
	if lo > hi {
		return
	}
	p := Partition1(a, lo, hi)
	OrderByViewCount(a,lo,p-1)
	OrderByViewCount(a,p+1,hi-1)
}

func Partition(a []AllKeys, lo int, hi int) int {
	pivot := a[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if a[j].Cloud_sdxy_model_download_count > pivot.Cloud_sdxy_model_download_count {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[hi] = a[hi], a[i+1]
	return i + 1
}

func Partition1(a []AllKeys, lo int, hi int) int {
	pivot := a[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if a[j].Cloud_sdxy_model_view_count < pivot.Cloud_sdxy_model_view_count {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[hi] = a[hi], a[i+1]
	return i + 1
}

//func RegisterAPIRouter() {
//	gin.SetMode(gin.DebugMode) //调试模式
//	router := gin.Default()     //获得路由实例
//
//	routerDatasource := router.Group("/data/source")
//	// 监听/data/source/connect的get和post请求，对应方法：ConnTest
//	routerDatasource.GET("/connect", SearchByKeys)
//	routerDatasource.POST("/connect", SearchByKeys)
//
//	//监听端口
//	http.ListenAndServe(":9000", router)
//
//}


//通过post请求，得到用户需要的key的值，返回一个结构体出来
func GetAllSearchKeys(kstyle string, krenderer string, kisfree string, korder string, kpageSize int, kpage int) SearchKeys {
	if kstyle == "" {
		kstyle = "不限"
	}
	if krenderer == "" {
		krenderer = "全部"
	}
	if kisfree == "" {
		kisfree = "全部"
	}
	if korder == "" {
		korder = "综合"
	}
	keys := SearchKeys{kstyle,krenderer,kisfree,korder,kpageSize,kpage}
	return keys
}

func initElasticsearchClient(hostname string, portname string) *elastic.Client{
	var err error
	var esClient *elastic.Client
	esClient, err = elastic.NewClient(elastic.SetURL(fmt.Sprintf("http://%s:%s",hostname,portname)))
	if err != nil {
		panic(err)
	}
	return esClient
}

func GetSearchResult1(indexName string, esClient *elastic.Client, Keys SearchKeys) *elastic.SearchResult {
	if Keys.Style == "不限" {
		if Keys.Renderer == "全部" {
			q := elastic.NewBoolQuery().Must(
				elastic.NewRangeQuery("cost").Lte(0),
				elastic.NewMatchPhraseQuery("status","on"),
			)
			searchService := esClient.Search(indexName).Size(10000)
			res, _ := searchService.Query(q).Do(context.Background())
			return res
		} else {
			q := elastic.NewBoolQuery().Must(
				elastic.NewMatchPhraseQuery("renderer",Keys.Renderer),
				elastic.NewRangeQuery("cost").Lte(0),
				elastic.NewMatchPhraseQuery("status","on"),
			)
			searchService := esClient.Search(indexName).Size(100000)
			res, _ := searchService.Query(q).Do(context.Background())
			return res
		}
	} else {
		if Keys.Renderer == "全部"{
			q := elastic.NewBoolQuery().Must(
				elastic.NewMatchPhraseQuery("style",Keys.Style),
				elastic.NewRangeQuery("cost").Lte(0),
				elastic.NewMatchPhraseQuery("status","on"),
			)
			searchService := esClient.Search(indexName).Size(100000)
			res, _ := searchService.Query(q).Do(context.Background())
			return res
		} else {
			q := elastic.NewBoolQuery().Must(
				elastic.NewMatchPhraseQuery("style",Keys.Style),
				elastic.NewMatchPhraseQuery("renderer",Keys.Renderer),
				elastic.NewRangeQuery("cost").Lte(0),
				elastic.NewMatchPhraseQuery("status","on"),
			)
			searchService := esClient.Search(indexName).Size(100000)
			res, _ := searchService.Query(q).Do(context.Background())
			return res
		}
	}
}

func GetSearchResult2(indexName string, esClient *elastic.Client, Keys SearchKeys) *elastic.SearchResult {
	if Keys.Style == "不限" {
		if Keys.Renderer == "全部" {
			q := elastic.NewBoolQuery().Must(
				elastic.NewRangeQuery("cost").Gt(0),
				elastic.NewMatchPhraseQuery("status","on"),
			)
			searchService := esClient.Search(indexName).Size(100000)
			res, _ := searchService.Query(q).Do(context.Background())
			return res
		} else {
			q := elastic.NewBoolQuery().Must(
				elastic.NewMatchPhraseQuery("renderer",Keys.Renderer),
				elastic.NewRangeQuery("cost").Gt(0),
				elastic.NewMatchPhraseQuery("status","on"),
			)
			searchService := esClient.Search(indexName).Size(100000)
			res, _ := searchService.Query(q).Do(context.Background())
			return res
		}
	} else {
		if Keys.Renderer == "全部" {
			q := elastic.NewBoolQuery().Must(
				elastic.NewMatchPhraseQuery("style",Keys.Style),
				elastic.NewRangeQuery("cost").Gt(0),
				elastic.NewMatchPhraseQuery("status","on"),
			)
			searchService := esClient.Search(indexName).Size(100000)
			res, _ := searchService.Query(q).Do(context.Background())
			return res
		} else {
			q := elastic.NewBoolQuery().Must(
				elastic.NewMatchPhraseQuery("style",Keys.Style),
				elastic.NewMatchPhraseQuery("renderer",Keys.Renderer),
				elastic.NewRangeQuery("cost").Gt(0),
				elastic.NewMatchPhraseQuery("status","on"),
			)
			searchService := esClient.Search(indexName).Size(100000)
			res, _ := searchService.Query(q).Do(context.Background())
			return res
		}
	}
}

func GetSearchResult3(indexName string, esClient *elastic.Client, Keys SearchKeys) *elastic.SearchResult {
	if Keys.Style == "不限" {
		if Keys.Renderer == "全部" {
			q := elastic.NewBoolQuery().Must(
				elastic.NewMatchPhraseQuery("status","on"),
				)
			searchService := esClient.Search(indexName).Size(100000)
			res, _ := searchService.Query(q).Do(context.Background())
			return res
		} else {
			q := elastic.NewBoolQuery().Must(
				elastic.NewMatchPhraseQuery("renderer",Keys.Renderer),
				elastic.NewMatchPhraseQuery("status","on"),
				elastic.NewMatchPhraseQuery("status","on"),
			)
			searchService := esClient.Search(indexName).Size(100000)
			res, _ := searchService.Query(q).Do(context.Background())
			return res
		}
	} else {
		if Keys.Renderer == "全部" {
			q := elastic.NewBoolQuery().Must(
				elastic.NewMatchPhraseQuery("style",Keys.Style),
				elastic.NewMatchPhraseQuery("status","on"),
			)
			searchService := esClient.Search(indexName).Size(100000)
			res, _ := searchService.Query(q).Do(context.Background())
			return res
		} else {
			q := elastic.NewBoolQuery().Must(
				elastic.NewMatchPhraseQuery("style",Keys.Style),
				elastic.NewMatchPhraseQuery("renderer",Keys.Renderer),
				elastic.NewMatchPhraseQuery("status","on"),
			)
			searchService := esClient.Search(indexName).Size(100000)
			res, _ := searchService.Query(q).Do(context.Background())
			return res
		}
	}
}

func query(indexName string,  esClient *elastic.Client, Keys SearchKeys) *elastic.SearchResult {
	if Keys.Isfree == "免费" {
		res := GetSearchResult1(indexName,esClient,Keys)
		return res
	} else if Keys.Isfree == "精品" {
		res := GetSearchResult2(indexName,esClient,Keys)
		return res
	} else {
		res := GetSearchResult3(indexName,esClient,Keys)
		return res
	}
}

//通过所有的keys，在ES中进行匹配，把匹配到的结果生成一个列表输出
func FindElementsInES(hostname string, portname string, Keys SearchKeys) []resultKeys  {
	var allkeys AllKeys
	var esClient *elastic.Client
	esClient = initElasticsearchClient(hostname,portname)
	result := query("3dxymodel",esClient,Keys)
	var resultALLKeysList []AllKeys
	if result != nil {
		if result.Hits != nil && result.Hits.Hits != nil {
			for _, hit := range result.Hits.Hits {
				err := json.Unmarshal(hit.Source, &allkeys)
				if err != nil {
					fmt.Println("Deserialization failed")
				}
				resultALLKeysList = append(resultALLKeysList, allkeys)
			}
			if len(resultALLKeysList) > 1 {
				if Keys.Order == "热门" {
					OrderByDownload(resultALLKeysList, 0, len(resultALLKeysList)-1)
				} else if Keys.Order == "最新" {
					OrderByViewCount(resultALLKeysList, 0, len(resultALLKeysList)-1)
				}
			}
			var resultList []resultKeys
			for _, ak := range resultALLKeysList {
				resultList = append(resultList, resultKeys{ak.Model_id, ak.Model_type, ak.Title, ak.List_thumbnails})
			}
			return resultList
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func main()  {
	router := gin.Default()
	router.POST("/data" , func(c *gin.Context) {
		SearchByKeys(c)
	})
	router.Run(":8100")
}

func SearchByKeys(c *gin.Context) {
	//var Rsearchkeys ReadSearchKeys
	var searchkeys SearchKeys
	err := c.BindJSON(&searchkeys)
	reqBody, err := ioutil.ReadAll(c.Request.Body)
	fmt.Println(reqBody)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 400, "description": "Post Data Err"})
		return
	}
	searchkeys.Style = searchkeys.Style
	searchkeys.Renderer = searchkeys.Renderer
	searchkeys.Order = searchkeys.Order
	searchkeys.Isfree = searchkeys.Isfree
	searchkeys.PageSize = searchkeys.PageSize
	searchkeys.Page = searchkeys.Page
	Keys := GetAllSearchKeys(searchkeys.Style, searchkeys.Renderer, searchkeys.Isfree, searchkeys.Order, searchkeys.PageSize, searchkeys.Page)
	//SK.SetUpKeys(Keys.word,Keys.style,Keys.renderer,Keys.isfree,Keys.order,Keys.pageSize,Keys.page)
	res := FindElementsInES("172.20.2.70","9200", Keys)//填写初始化es的url
	//x,err := json.Marshal(res)
	//if err != nil {
	//	panic(err)
	//}
	if res == nil {
		c.JSON(200,gin.H{})
	} else {
		var result Result
		var ans []resultKeys
		var start int
		var end int
		var i int
		result.TotalNum = len(res)
		result.TotalPage = result.TotalNum / searchkeys.PageSize
		a := result.TotalNum - searchkeys.PageSize * result.TotalPage
		if a > 0 {
			result.TotalPage = result.TotalPage + 1
		}
		start = (searchkeys.Page - 1) * searchkeys.PageSize
		end = start + searchkeys.PageSize - 1
		if end > len(res) {
			end = len(res) - 1
		}
		for i = start; i <= end; i++ {
			ans = append(ans,res[i])
		}
		result.Keys = ans
		c.JSON(200, result)
	}
}