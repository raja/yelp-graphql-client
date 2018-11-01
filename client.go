package yelp

import (
	"context"
	"time"

	"github.com/shurcooL/graphql"
	"golang.org/x/oauth2"
)

// Client represents a yelp graphql client
type Client struct {
	client *graphql.Client
	token  string
}

//NewClient returns a client by providing a token
func NewClient(token string) *Client {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	httpClient.Timeout = 5 * time.Second
	return &Client{
		client: graphql.NewClient("https://api.yelp.com/v3/graphql", httpClient),
		token:  token,
	}
}

// Business loads information about a specific business
func (c *Client) Business(id string) (*Business, error) {
	var query struct {
		Business Business `graphql:"business(id: $id)"`
	}
	variables := map[string]interface{}{
		"id": graphql.String(id),
	}
	err := c.client.Query(context.Background(), &query, variables)
	if err != nil {
		return nil, err
	}
	return &query.Business, nil
}

// Search for a Business on Yelp with term and location
func (c *Client) Search(term, location string, limit int) ([]*Business, error) {
	var query struct {
		Search struct {
			Total    graphql.Int
			Business []Business
		} `graphql:"search(term: $term, location: $location, limit: $limit)"`
	}
	queryVars := map[string]interface{}{
		"term":     graphql.String(term),
		"location": graphql.String(location),
		"limit":    graphql.Int(limit),
	}
	err := c.client.Query(context.Background(), &query, queryVars)
	if err != nil {
		return nil, err
	}
	var results []*Business
	for _, biz := range query.Search.Business {
		result := biz
		results = append(results, &result)
	}
	return results, nil
}
