//go:build integration

package expense_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/aRaimaiRu/assessment/expense"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	body := bytes.NewBufferString(
		`{
			"title": "strawberry smoothie",
			"amount": 79,
			"note": "night market promotion discount 10 bath", 
			"tags": ["food", "beverage"]
		}`,
	)
	var u expense.Expense

	res := request(http.MethodPost, uri("expenses"), body)
	err := res.Decode(&u)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)
	assert.NotEqual(t, 0, u.Id)
	assert.Equal(t, "strawberry smoothie", u.Title)
	assert.Equal(t, float32(79.0), u.Amount)
}

func TestGetExpense(t *testing.T) {
	u := seedExpense(t)

	var lastest expense.Expense
	res := request(http.MethodGet, uri("expenses", strconv.Itoa(u.Id)), nil)
	err := res.Decode(&lastest)
	assert.Nil(t, err)
	assert.Equal(t, lastest.Title, u.Title)
	assert.Equal(t, lastest.Amount, u.Amount)
	assert.Equal(t, lastest.Note, u.Note)
}

func seedExpense(t *testing.T) expense.Expense {
	var c expense.Expense
	body := bytes.NewBufferString(
		`{
			"title": "smoothie",
			"amount": 80,
			"note": "night market promotion discount 20 bath", 
			"tags": ["food", "beverage"]
		}`,
	)
	err := request(http.MethodPost, uri("expenses"), body).Decode(&c)
	if err != nil {
		t.Fatal("can't create uomer:", err)
	}
	return c
}

func uri(paths ...string) string {
	host := "http://localhost:2565"
	if paths == nil {
		return host
	}

	url := append([]string{host}, paths...)
	return strings.Join(url, "/")
}

func request(method, url string, body io.Reader) *Response {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Add("Authorization", os.Getenv("AUTH_TOKEN"))
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	return &Response{res, err}
}

type Response struct {
	*http.Response
	err error
}

func (r *Response) Decode(v interface{}) error {
	if r.err != nil {
		return r.err
	}

	return json.NewDecoder(r.Body).Decode(v)
}
