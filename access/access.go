package access

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strconv"

	"github.com/dghubble/oauth1"

	"github.com/cd-mbc/twitter_api_cli_go/model"
	"github.com/cd-mbc/twitter_api_cli_go/setting"
)

func Base_access(url string) []byte {
	config := oauth1.NewConfig(setting.CK, setting.CS)
	token := oauth1.NewToken(setting.AT, setting.AS)
	httpClient := config.Client(oauth1.NoContext, token)

	res, err := httpClient.Get(url)
	if err != nil {
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil
	}

	return body
}

func Search_tweets(q string, count int) *model.Tweets {
	path := setting.Search_tweets_path
	params := url.Values{}
	params.Add("q", q)
	params.Add("count", strconv.Itoa(count))

	body := Base_access(path + params.Encode())
	if body == nil {
		return nil
	}

	var tweets model.Tweets
	if err := json.Unmarshal(body, &tweets); err != nil {
		return nil
	}
	fmt.Println(tweets.Search_metadata.Max_id)

	return &tweets
}

func Get_trends(id int) *[]model.Trends {
	path := setting.Get_trends_path
	params := url.Values{}
	params.Add("id", strconv.Itoa(id))

	body := Base_access(path + params.Encode())
	if body == nil {
		return nil
	}

	var trends []model.Trends
	if err := json.Unmarshal(body, &trends); err != nil {
		fmt.Println(err)
		return nil
	}

	return &trends
}
