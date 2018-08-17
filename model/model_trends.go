
package model

type Trends struct {
  Trends []Trend `json:trends`
}

type Trend struct {
  Name string `json:name  db:name`
  Url string `json:url  db:url`
  Tweet_volume int `json:tweet_volume  db:tweet_volume`
}
