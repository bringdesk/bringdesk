package news

import (
	"encoding/xml"
	"fmt"
	"github.com/bringdesk/bringdesk/ctx"
	"github.com/bringdesk/bringdesk/evt"
	"github.com/bringdesk/bringdesk/widgets"
	"log"
	"time"
)

type RSSItem struct {
	//<title>Фигуристы Мишина и Галлямов удовлетворены прокатом короткой программы на чемпионате России</title>
	Title string `xml:"title"`
	//<link><![CDATA[https://tass.ru/sport/13287229]]></link>
	Link string `xml:"link"`
	//<guid><![CDATA[https://tass.ru/sport/13287229]]></guid>
	Guid string `xml:"guid"`
	//<pubDate>Thu, 23 Dec 2021 20:37:15 +0300</pubDate>
	PubDate string `xml:"pubDate"`
	//<description>За свое выступление спортсмены получили 83,74 балла</description>
	Description string `xml:"description"`
	//<enclosure url="https://cdn3.tass.ru/fit/400x300_b2b00b17/tass/m2/uploads/i/20211223/6529409.jpg" type="image/jpeg" length="2815904"></enclosure>
	Enclosure struct {
		URL string
	} `xml:"enclosure"`

	//<category>Спорт</category>
	//<category>Фигурное катание</category>
	//<category>Чемпионат России по фигурному катанию</category>
	Category []string `xml:"category"`
}

type RSSChannel struct {
	//<title>ТАСС</title>
	Title string `xml:"title"`
	//<description>ИНФОРМАЦИОННОЕ АГЕНТСТВО РОССИИ ТАСС</description>
	Description string `xml:"description"`
	//<language>ru-ru</language>
	Language string `xml:"language"`
	//<link><![CDATA[https://tass.ru]]></link>
	Link string `xml:"link"`
	//<copyright>ТАСС</copyright>
	Copyright string `xml:"copyright"`
	//<image>
	Image struct {
		//<url><![CDATA[https://tass.ru/i/rss/logo.png]]></url>
		URL string `xml:"url"`
		//<title>ТАСС</title>
		Title string `xml:"title"`
		//<link><![CDATA[https://tass.ru]]></link>
		Link string `xml:"link"`
	} `xml:"image"`
	//</image>
	//<atom:link href="https://tass.ru/rss/v2.xml" rel="self" type="application/rss+xml"></atom:link>
	//<item>
	//</item>
	Items []RSSItem `xml:"item"`
}

type RSSAtom struct {
	//<rss version="2.0" xmlns:atom="https://www.w3.org/2005/Atom">
	XMLName xml.Name `xml:"rss"`
	//<channel>
	Channel []RSSChannel `xml:"channel"`
	//</channel>
	//</rss>
}

type NewsItem struct {
	pubDate time.Time /* News stamp         */
	summary string    /* News summary       */
}

type NewsWidget struct {
	widgets.BaseWidget
	news []*NewsItem
}

type ByPubDate []RSSItem

func (a ByPubDate) Len() int { return len(a) }
func (a ByPubDate) Less(i, j int) bool {
	parserPattern := "Thu, 23 Dec 2021 21:04:12 +0300"
	pubDate1 := a[i].PubDate
	pubDate2 := a[j].PubDate
	pubDate1Time, err1 := time.Parse(parserPattern, pubDate1)
	pubDate2Time, err2 := time.Parse(parserPattern, pubDate2)
	if err1 != nil || err2 != nil {
		return pubDate1 < pubDate2
	}
	log.Printf("Sort comparision error: err1 = %#v err2 = %#v", err1, err2)
	return pubDate1Time.After(pubDate2Time)
}
func (a ByPubDate) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func NewNewsWidget() *NewsWidget {
	newNewsWidget := new(NewsWidget)
	go func() {
		for {
			newNewsWidget.updateData()
			time.Sleep(10 * time.Minute)
		}
	}()
	return newNewsWidget
}

func (self *NewsWidget) updateData() {
	log.Printf("NewsWidget: Update RSS news...")
	/* Step 1. Get RSS news */
	mainNetworkManager := ctx.GetNetworkManager()
	req, err1 := mainNetworkManager.MakeRequest("NewsWidget", "GET", "http://tass.ru/rss/v2.xml", 15)
	if err1 != nil {
		log.Printf("err = %#v", err1)
		return
	}

	resp, err2 := mainNetworkManager.Perform(req)
	if err2 != nil {
		log.Printf("err = %#v", err2)
		return
	}

	/* Step 2. Parse RSS news */
	out := resp.Bytes()
	var rssAtom RSSAtom
	err3 := xml.Unmarshal(out, &rssAtom)
	if err3 != nil {
		log.Printf("err = %#v", err3)
		return
	}

	/* Step 2. Populate news */
	self.populateNews(&rssAtom)

}

func (self *NewsWidget) ProcessEvent(e *evt.Event) {
}

func (self *NewsWidget) Render() {
	self.BaseWidget.Render()

	/* Show task */
	for idx, n := range self.news {
		if idx > 15 {
			break
		}
		newText := widgets.NewTextWidget("", 16)
		newText.SetRect(self.X, self.Y+20*idx, self.Width, self.Height)
		newText.SetText(n.summary)
		newText.Render()
		newText.Destroy()
	}

}

func (self *NewsWidget) populateNews(rssAtom *RSSAtom) {
	self.news = nil
	for _, c := range rssAtom.Channel {
		for _, i := range c.Items {
			newNews := new(NewsItem)
			newNews.pubDate = time.Now()
			newNews.summary = fmt.Sprintf("%s - %s", i.PubDate, i.Title)
			self.news = append(self.news, newNews)
			//
			if len(self.news) > 15 {
				break
			}
		}
	}
}
