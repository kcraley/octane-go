package trivia

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

const (
	triviaURL     = "https://jservice.io/api"
	cluesAPI      = "/clues"
	randomAPI     = "/random"
	categoriesAPI = "/categories"
	categoryAPI   = "/category"
	invalidAPI    = "/invalid"
)

// Questions represents a slice of trive questions
type Questions []Question

// Question represents a trivia question
type Question struct {
	ID           int64     `json:"id"`
	Answer       string    `json:"answer"`
	Question     string    `json:"question,omitempty"`
	Value        int32     `json:"value,omitempty"`
	Airdate      time.Time `json:"airdate"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	CategoryID   int32     `json:"category_id"`
	GameID       int32     `json:"game_id"`
	InvalidCount int8      `json:"invalid_count,omitempty"`
	Category     Category  `json:"category"`
}

// Category represents a trivia category
type Category struct {
	ID         int64     `json:"id"`
	Title      string    `json:"title"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	CluesCount int32     `json:"clues_count,omitempty"`
}

// GetRandomQuestions retreives a given number of random trivia questions
func GetRandomQuestions(count int) ([]Question, error) {
	// Create a slice to hold our questions to return
	questions := Questions{}

	// Create a new http client
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Create a new request
	req, err := http.NewRequest("GET", triviaURL+randomAPI, nil)
	if err != nil {
		return nil, err
	}

	// Build our query to pull multiple questions
	query := req.URL.Query()
	query.Add("count", strconv.Itoa(count))
	req.URL.RawQuery = query.Encode()

	// Submit the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Unmarshal our JSON response and return
	json.NewDecoder(resp.Body).Decode(&questions)
	return questions, nil
}
