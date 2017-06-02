package elasticsearch

import (
	"gopkg.in/olivere/elastic.v5"
)

type ESClient struct {
	es       *elastic.Client
	baseURL  string
	user     string
	password string
	index    string
}

func NewESClient(base, user, password, index string) (*ESClient, error) {

	// Create the ElasticSearch client.
	var (
		c   *elastic.Client
		err error
	)
	if user != "" && password != "" {
		c, err = elastic.NewSimpleClient(elastic.SetURL(base), elastic.SetBasicAuth(user, password))
	} else {
		c, err = elastic.NewSimpleClient(elastic.SetURL(base))
	}
	if err != nil {
		return nil, err
	}

	return &ESClient{
		es:       c,
		baseURL:  base,
		user:     user,
		password: password,
		index:    index,
	}, nil
}
