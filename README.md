# AnimeNewsGo

__News from animenewsnetwork.com__

<a href="https://pkg.go.dev/github.com/NksamaX/AnimeNewsGo"><img src="https://pkg.go.dev/badge/github.com/NksamaX/AnimeNewsGo.svg" alt="Go Reference"></a>


# How to install

```
go get github.com/NksamaX/AnimeNewsGo
```

# How to Use

```
package main

import (
	"github.com/NksamaX/AnimeNewsGo/news"
	"fmt"
)

func main() {
	x := news.Get_News()
	fmt.Println(x["title"])
	fmt.Println(x["news"])
	
}

```
