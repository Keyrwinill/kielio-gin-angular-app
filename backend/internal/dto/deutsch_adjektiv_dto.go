package dto

type DeutschAdjektivResponse struct {
	Oid             string   `json:"oid"`
	Type            string   `json:"type"`
	Gender          string   `json:"gender"`
	Case            string   `json:"case"`
	ArticleEnding   string   `json:"article_ending"`
	AdjectiveEnding string   `json:"adjective_ending"`
	ArticleTypeList []string `json:"article_type_list"`
}
