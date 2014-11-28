package async

import (
	"github.com/darkhelmet/twitterstream"
	"time"
)

type Client struct {
	*twitterstream.Client
}

func NewClient(consumerKey, consumerSecret, accessToken, accessSecret string) *Client {
	return &Client{
		twitterstream.NewClient(consumerKey, consumerSecret, accessToken, accessSecret),
	}
}

func (client *Client) TrackAndServe(channels string, callback func(tweet *twitterstream.Tweet)) {
	for {
		conn, err := client.Track(channels)
		if err != nil {
			time.Sleep(1 * time.Minute)
			continue
		}
		client.decode(callback, conn)
	}
}

func (client *Client) decode(callback func(tweet *twitterstream.Tweet), conn *twitterstream.Connection) (*twitterstream.Tweet, error) {
	for {
		if tweet, err := conn.Next(); err == nil {
			callback(tweet)
		} else {
			return nil, err
		}
	}
}
