package yelp

import (
	"testing"
)

func TestClient_Business(t *testing.T) {
	c := NewClient("your token here")
	biz, err := c.Business("garaje-san-francisco")
	if err != nil {
		t.Log(err)
	}
	t.Logf("we have %+v", biz)
}
