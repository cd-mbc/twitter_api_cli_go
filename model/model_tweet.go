
package model

type Tweets struct {
  Statuses []Tweet `json:statuses`
  Search_metadata Search_meta `json:search_metadata`
}

type Tweet struct {
  Created_at string `json:created_at`
  Id int `json:id`
  Text string `json:text`
  Lang string `json:lang`
}

type Search_meta struct {
  Completed_in float64 `json:completed_in`
  Max_id int `json:max_id`
  Max_id_str string `json:max_id_str`
  Next_results string `json:next_results`
  Query string `json:query`
  Refresh_url string `json:refresh_url`
  Count int `json:count`
  Since_id int `json:since_id`
  Since_id_str string `json:since_id_str`
}
