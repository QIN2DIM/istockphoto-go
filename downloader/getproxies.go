package downloader

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"regexp"
	"runtime"
	"strings"
)

const (
	Windows = "windows"
	Darwin  = "darwin"
	Linux   = "linux"
)

func GetProxies() map[string]string {
	var solutions = map[string]func() map[string]string{
		Windows: getProxiesOnWindows,
		Darwin:  getProxiesOnDarwin,
		Linux:   getProxiesOnLinux,
	}
	return solutions[runtime.GOOS]()
}

func getProxiesOnWindows() map[string]string {
	proxies := make(map[string]string)
	openPath := "Software\\Microsoft\\Windows\\CurrentVersion\\Internet Settings"

	k, _ := registry.OpenKey(registry.CURRENT_USER, openPath, registry.QUERY_VALUE)
	defer k.Close()

	s, _, _ := k.GetStringValue("ProxyServer")

	if strings.Contains(s, "=") {
		for _, p := range strings.Split(s, ";") {
			ps := strings.SplitN(p, "=", 1)
			protocol, address := ps[0], ps[1]
			res, _ := regexp.MatchString("(:[^/:+]://)", address)
			if !res {
				address = fmt.Sprintf("%s://%s", protocol, address)
			}
			proxies[protocol] = address
		}
	} else {
		if strings.HasPrefix(s, "http:") {
			proxies["http"] = s
		} else {
			proxies["http"] = fmt.Sprintf("http://%s", s)
			proxies["https"] = fmt.Sprintf("https://%s", s)
			proxies["ftp"] = fmt.Sprintf("ftp://%s", s)
		}
	}

	return proxies
}

func getProxiesOnDarwin() map[string]string {
	return nil
}

func getProxiesOnLinux() map[string]string {
	return nil
}
