package spider_lib

import (
	"github.com/PuerkitoBio/goquery"
	."splider/models"
	."splider/helper"
	."splider/spider_lib/landing_page"
	"fmt"
)


func ZhiHuBianJi(channel chan <- []*Crawler){
	doc, err := goquery.NewDocument("https://www.zhihu.com/explore/recommendations")

	if err != nil{
		panic(err.Error())
	}

	var urlList []string
	doc.Find("#zh-recommend-list-full .zh-general-list .zm-item h2 a").
		Each(func(i int, selection *goquery.Selection) {
		url, isExist := selection.Attr("href")

		fmt.Println(url)
		if isExist{
			urlList = append(urlList, url)
		}
	})

	var data []*Crawler

	for _, url := range FilterZhihuURLs(ChangeToAbspath(urlList, "https://www.zhihu.com")){
		crawlerData, err := PaserZhihuQuestion(url)
		if err == nil{
			data = append(data, crawlerData)
		}
	}

	channel <- data
}


