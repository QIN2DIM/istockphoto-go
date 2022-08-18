# IstockPhoto Go

<img align="right" width="200px" src="https://user-images.githubusercontent.com/62018067/168443374-8343a72a-dc70-41b3-85ed-4776b3595bd3.png">

[![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome)

📸 Gracefully download dataset from [istockphoto](https://www.istockphoto.com/). 

- You can download nearly a thousand images in a few tens of seconds with great efficiency thanks to a modern crawler infrastructure.
- You can customize keywords and fine-tune filters for the purpose of controlling the search scope of resources.

## Installation

```bash
go get -u github.com/QIN2DIM/istockphoto-go
```

## Example

See [wiki](https://github.com/QIN2DIM/istockphoto-go/wiki) for more detailed examples. 

[ForCN. 这是一个被墙的网站，需要开启系统代理]

```go
package main

import "github.com/QIN2DIM/istockphoto-go"

func main() {
	istockphoto.NewDownloader("cyberpunk").Mining()
}
```

![image-20220818084304218](https://user-images.githubusercontent.com/62018067/185268084-23e4db6a-7162-4297-ba41-bb401a1e9ec6.png)
