package istockphoto

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"
	"strings"
)

const (
	IstockQuerySample = "https://www.istockphoto.com/search/2/image?phrase=dragon"
	IstockSearchAPI   = "https://www.istockphoto.com/search/2/image"
	IstockMediaPrefix = "https://media.istockphoto.com/"
	MaxPages          = 20
	GlobalProxyStr    = "http://127.0.0.1:10809"
)

var Headers = map[string]string{
	"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko)" +
		" Chrome/103.0.5060.134 Safari/537.36 Edg/103.0.1264.77",
}

type Downloader struct {
	Phrase         string
	Mediatype      string
	NumberOfPeople string
	Orientations   string
	Pages          int
	Flag           bool
	Backend        string

	rawPointerCount     int
	handledPointerCount int

	pendingImageCount    int // pendingImageCount
	downloadedImageCount int

	dirLocal string
	client   *Client
}
type Client struct {
	Session *http.Client
}
type Miner interface {
}

func init() {
	log.SetPrefix("DEBUG: ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// NewDownloader 新建下载器对象
func NewDownloader() (d *Downloader) { return }

// newClient 创建客户端会话(make session)
func newClient() *Client {
	// 自动获取本地代理
	proxyURL, _ := url.Parse(GlobalProxyStr)

	client := &Client{
		Session: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyURL),
			},
		},
	}

	return client
}

func (d *Downloader) preload() {

}

func (d *Downloader) offload() {

}

func (c *Client) handleHtml(url string) (body string, err error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Printf("Get url error | url=%s err=%s", url, err.Error())
		return
	}
	for key, value := range Headers {
		req.Header.Set(key, value)
	}

	resp, err := c.Session.Do(req)
	if err != nil {
		log.Printf("Session do error | url=%s err=%s", url, err.Error())
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)

	data, err := io.ReadAll(resp.Body)
	return string(data), err
}

// parseHtmlAndGetImageUrls Parse web content and get image download links
func (c *Client) parseHtmlAndGetImageUrls(responseStr string) []string {
	var containerImgUrls []string
	root := soup.HTMLParse(responseStr)
	imgTags := root.Find("div", "data-testid", "gallery-items-container").FindAll("img")
	for _, tag := range imgTags {
		if link := tag.Attrs()["src"]; link != "" {
			containerImgUrls = append(containerImgUrls, link)
		}
	}
	return containerImgUrls
}

// adaptor 接口适配器
func (d *Downloader) adaptor() {
	// 从工作队列中获取网页链接
	url_ := ""

	//　匹配链接特征选择不同的执行模式
	switch true {
	case strings.HasPrefix(url_, IstockMediaPrefix):
		fmt.Println("下载图片 ...")
	case strings.HasPrefix(url_, IstockSearchAPI):
		fmt.Println("搜集图片下载链接 ...")
	case d.downloadedImageCount%60 == 0:
		fmt.Println("缓存卸载 ...")
	}
}

// GetImageUrls Add the image download link to the Goroutine Queue
func (d *Downloader) GetImageUrls(url string) bool {
	// testUrl = IstockQuerySample
	responseStr, err := d.client.handleHtml(url)
	if err != nil {
		log.Printf("HandleHtmlExption | url=%s err=%s", url, err.Error())
		return false
	}

	urls := d.client.parseHtmlAndGetImageUrls(responseStr)
	for url_ := range urls {
		log.Println(url_)
		d.pendingImageCount += 1
	}

	d.handledPointerCount += 1
	log.Printf("Get image url - progress=[%d/%d] src=%s", d.handledPointerCount, d.rawPointerCount, url)

	return true
}

// DownloadImage 下载图片
func (d *Downloader) DownloadImage(url_ string) bool {
	d.downloadedImageCount += 1
	query, _ := url.ParseQuery(url_)
	istockID := query["m"][0]
	fp := path.Join(d.dirLocal, fmt.Sprintf("%s.jpg", istockID))
	fmt.Println(fp)
	return true
}

func (d *Downloader) Mining(concurrentPower int) {

}

func parseFileCount() {

}

func (d *Downloader) MoreLikeThis(istockId int, similar string) {

}