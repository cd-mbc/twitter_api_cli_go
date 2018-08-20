package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/cd-mbc/twitter_api_cli_go/access"
	"github.com/cd-mbc/twitter_api_cli_go/database"
	"github.com/cd-mbc/twitter_api_cli_go/model"
	"github.com/cd-mbc/twitter_api_cli_go/setting"
)

func get_trends_from_db() []model.Trend {
	trend_seq_db := database.Select_trend_all()
	return trend_seq_db
}

func trends_db_update(trends model.Trends) []model.Trend {
	trend_seq := trends.Trends
	if len(trend_seq) != 0 {
		database.Delete_trend_all()
	}

	for i := 0; i < len(trend_seq); i++ {
		fmt.Println(trend_seq[i].Name)
		database.Insert_trend(&trend_seq[i])
	}

	return trend_seq
}

func get_tweet_of_trend(name string) []model.Tweet {
	tweets := access.Search_tweets(name, setting.Tweet_max)
	if tweets != nil {
		return tweets.Statuses
	} else {
		return []model.Tweet{}
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	woeid := setting.Woeid_japan

	trends := access.Get_trends(woeid)
	trend_seq := []model.Trend{}
	if trends != nil {
		trend_seq = trends_db_update((*trends)[0])
	} else {
		trend_seq = get_trends_from_db()
	}

	t, _ := template.ParseFiles("./template/index.html")
	t.Execute(w, trend_seq)
}

func trend_tweet(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	query := r.Form
	name := query["name"][0]
	tweet_seq := get_tweet_of_trend(name)

	t, _ := template.ParseFiles("./template/trend_tweet.html")
	t.Execute(w, tweet_seq)
}

func main() {

	database.Init_db()

	server := http.Server{
		Addr: setting.Host + ":" + setting.Port,
	}
	http.HandleFunc("/", index)
	http.HandleFunc("/trend_tweet", trend_tweet)
	server.ListenAndServe()

}
