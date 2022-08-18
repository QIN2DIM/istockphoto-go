![image-20220818084304218](https://user-images.githubusercontent.com/62018067/185268084-23e4db6a-7162-4297-ba41-bb401a1e9ec6.png)

# IstockPhoto Go

📸 Gracefully download dataset from [istockphoto](https://www.istockphoto.com/).

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

