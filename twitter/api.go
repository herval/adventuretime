package twitter

import (
	"time"

	"github.com/ChimeraCoder/anaconda"
	"image"
	"net/url"
	"strconv"
	"log"
	"github.com/herval/adventuretime/util"
)

type Api struct {
	api *anaconda.TwitterApi
}

func NewApi(consumerKey string, consumerSecret string, accessToken string, tokenSecret string) *Api {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)

	return &Api{
		api: anaconda.NewTwitterApi(accessToken, tokenSecret),
	}
}

func (self *Api) Post(message string) {
	self.api.PostTweet(message, nil)
}

func (self *Api) PostWithMedia(message string, img *image.RGBA) {
	data, err := util.ImageToBase64(img)
	if err != nil {
		log.Fatal(err)
	}

	mediaResponse, err := self.api.UploadMedia(data)
	if err != nil {
		log.Fatal(err)
	}

	v := url.Values{}
	v.Set("media_ids", strconv.FormatInt(mediaResponse.MediaID, 10))

	self.api.PostTweet(message, v)
}

func (self *Api) MentionsStream(since time.Time) <-chan anaconda.Tweet {
	out := make(chan anaconda.Tweet)

	go func() {
		for {
			mentions, err := self.api.GetMentionsTimeline(nil)
			if err == nil && len(mentions) > 0 {
				created, err := mentions[0].CreatedAtTime()
				if err == nil && created.After(since) {
					since = created
					out <- mentions[0]
				}
			}
			time.Sleep(time.Duration(1) * time.Minute) // make configurable?
		}
	}()

	return out
}
