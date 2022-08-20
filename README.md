<div align="center">
    <p align="center"><img src="https://user-images.githubusercontent.com/62018067/168443374-8343a72a-dc70-41b3-85ed-4776b3595bd3.png" width="30%"></p>
    <h1>IstockPhoto Go</h1>
    <p><i>📸 Gracefully download dataset from istockphoto. </i></p>
    <img src="https://img.shields.io/static/v1?message=awesome&style=for-the-badge&logo=Awesome Lists&label=&logoColor=black&labelColor=c5a2bd&color=4e486a">
    <img src="https://img.shields.io/github/license/QIN2DIM/istockphoto-go?style=for-the-badge">
    <a href="https://github.com/QIN2DIM/istockphoto-go/releases"><img src="https://img.shields.io/github/downloads/qin2dim/istockphoto-go/total?style=for-the-badge"></a>
	<br>
    <a href="https://github.com/QIN2DIM/istockphoto-go/"><img src="https://img.shields.io/github/stars/QIN2DIM/istockphoto-go?style=social"></a>
	<a href = "https://t.me/+tJrSQ0_0ujkwZmZh"><img src="https://img.shields.io/static/v1?style=social&logo=telegram&label=chat&message=studio" ></a>
	<br>
	<br>
</div>


## Installation

```bash
go get -u github.com/QIN2DIM/istockphoto-go
```

## Example

See [wiki](https://github.com/QIN2DIM/istockphoto-go/wiki) for more detailed examples. 

[ForCN. 这是一个被墙的网站，需要开启系统代理]

```go
package main

import "github.com/QIN2DIM/istockphoto-go/downloader"

func main() {
	downloader.NewDownloader("cyberpunk").Mining()
}
```

![image-20220818084304218](https://user-images.githubusercontent.com/62018067/185268084-23e4db6a-7162-4297-ba41-bb401a1e9ec6.png)
