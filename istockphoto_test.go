package main

import (
	"fmt"
	"github.com/QIN2DIM/getproxies"
	"github.com/QIN2DIM/istockphoto-go/downloader"
	"testing"
)

func TestStandardDownloader(t *testing.T) {
	downloader.NewDownloader("cyberpunk").Mining()
}

func TestDownloaderWithProxyURL(t *testing.T) {
	// (Only valid for Windows)
	// This project has a built-in GetProxies method
	// to automatically obtain local system proxies by registry
	d := downloader.NewDownloader("cyberpunk")
	d.ProxyURL = getproxies.GetProxies()["http"]
	d.Mining()
}

func TestDownloaderNotQuery(t *testing.T) {
	d := downloader.NewDownloader("dog")
	d.CloseFilter()
	d.Mining()
}

func TestDownloaderWithPages(t *testing.T) {
	d := downloader.NewDownloader("cat")
	d.Pages = 4
	d.Mining()
}

func TestDownloaderWithFlag(t *testing.T) {
	// Images with different label will be centrally stored in the same directory
	flag := "lion"
	phrases := []string{"lion closed eyes", "lion open mouth"}

	for _, phrase := range phrases {
		d := downloader.NewDownloader(phrase)
		d.Orientations = downloader.Orientations.Undefined
		d.Flag = flag
		d.Mining()
	}
}

func TestMoreLikeThisContent(t *testing.T) {
	d := downloader.NewDownloader("gun")
	d.Orientations = downloader.Orientations.Undefined
	d.Flag = "gun-similar-content"
	d.MoreLikeThis(529989264).Mining()
}

func TestMoreLikeThisColor(t *testing.T) {
	tag := "cyberpunk"
	d := downloader.NewDownloader(tag)
	d.Orientations = downloader.Orientations.Undefined
	d.Flag = fmt.Sprintf("%s-similar-color", tag)
	d.Similar = downloader.Color
	d.MoreLikeThis(1266931346).Mining()
}
