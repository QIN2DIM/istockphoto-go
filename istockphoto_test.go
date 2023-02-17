package main

import (
	"fmt"
	"github.com/QIN2DIM/getproxies"
	"github.com/QIN2DIM/istockphoto-go/downloader"
	"testing"
)

func TestStandardDownloader(t *testing.T) {
	downloader.NewDownloader("bug").Mining()
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
	d := downloader.NewDownloader("turtle")
	d.Pages = 3
	d.Orientations = downloader.Orientations.Undefined
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
	d := downloader.NewDownloader("turtle")
	d.Orientations = downloader.Orientations.Undefined
	d.Flag = "turtle-similar-content-1"
	d.MoreLikeThis(1136464536).Mining()
}

func TestMoreLikeThisColor(t *testing.T) {
	tag := "turtle"
	d := downloader.NewDownloader(tag)
	d.Orientations = downloader.Orientations.Undefined
	d.Flag = fmt.Sprintf("%s-similar-color", tag)
	d.Similar = downloader.Color
	d.MoreLikeThis(834997362).Mining()
}
