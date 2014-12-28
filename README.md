# Twitter Stream

```go
package main

import (
    "github.com/darkhelmet/twitterstream"
    "github.com/wdalmut/twitterstream/async"
    "log"
)

func main() {
	client := async.NewClient(
        "consumerKey",
        "consumerSecret",
        "accessToken",
        "accessSecret"
    )

    client.TrackAndServe("love", func(tweet *twitterstream.Tweet) {
        log.Println(tweet.Text)
    })
}
```
