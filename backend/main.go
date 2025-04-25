package main

import (
  "encoding/json"
  "fmt"
  "math/rand"
  "net/http"
  "net/url"
  "os" // Tambahkan ini untuk mengakses environment variable
  "time"
  "github.com/gorilla/mux"
)

type SuggestionResponse struct {
  Suggestions []string `json:"suggestions"`
}

type TranslationResponse struct {
  Translation string `json:"translation"`
}

type RandomWordResponse struct {
  Word string `json:"word"`
}

func suggestWords(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query().Get("prefix")
  if query == "" {
    http.Error(w, "Missing prefix parameter", http.StatusBadRequest)
    return
  }

  apiURL := fmt.Sprintf("https://api.datamuse.com/words?sp=%s*", url.QueryEscape(query))
  resp, err := http.Get(apiURL)
  if err != nil {
    http.Error(w, "Error fetching suggestions", http.StatusInternalServerError)
    return
  }
  defer resp.Body.Close()

  var datamuseResp []struct {
    Word string `json:"word"`
  }
  if err := json.NewDecoder(resp.Body).Decode(&datamuseResp); err != nil {
    http.Error(w, "Error decoding suggestions", http.StatusInternalServerError)
    return
  }

  var suggestions []string
  for _, item := range datamuseResp {
    suggestions = append(suggestions, item.Word)
  }

  response := SuggestionResponse{Suggestions: suggestions}
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(response)
}

func translateWord(w http.ResponseWriter, r *http.Request) {
  word := r.URL.Query().Get("word")
  if word == "" {
    http.Error(w, "Missing word parameter", http.StatusBadRequest)
    return
  }

  apiURL := fmt.Sprintf("https://api.mymemory.translated.net/get?q=%s&langpair=en|id", url.QueryEscape(word))
  resp, err := http.Get(apiURL)
  if err != nil {
    http.Error(w, "Error fetching translation", http.StatusInternalServerError)
    return
  }
  defer resp.Body.Close()

  var myMemoryResp struct {
    ResponseData struct {
      TranslatedText string `json:"translatedText"`
    } `json:"responseData"`
  }
  if err := json.NewDecoder(resp.Body).Decode(&myMemoryResp); err != nil {
    http.Error(w, "Error decoding translation", http.StatusInternalServerError)
    return
  }

  translation := myMemoryResp.ResponseData.TranslatedText

  response := TranslationResponse{Translation: translation}
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(response)
}

func getRandomWord(w http.ResponseWriter, r *http.Request) {
  topics := []string{"animal", "food", "color", "object", "action", "emotion", "place", "person"}
  rand.Seed(time.Now().UnixNano())
  topic := topics[rand.Intn(len(topics))]

  apiURL := fmt.Sprintf("https://api.datamuse.com/words?ml=%s&max=1", url.QueryEscape(topic))
  resp, err := http.Get(apiURL)
  if err != nil {
    http.Error(w, "Error fetching random word", http.StatusInternalServerError)
    return
  }
  defer resp.Body.Close()

  var datamuseResp []struct {
    Word string `json:"word"`
  }
  if err := json.NewDecoder(resp.Body).Decode(&datamuseResp); err != nil {
    http.Error(w, "Error decoding random word", http.StatusInternalServerError)
    return
  }

  if len(datamuseResp) == 0 {
    http.Error(w, "No random word found", http.StatusInternalServerError)
    return
  }

  response := RandomWordResponse{Word: datamuseResp[0].Word}
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(response)
}

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/suggest", suggestWords).Methods("GET")
  r.HandleFunc("/translate", translateWord).Methods("GET")
  r.HandleFunc("/random", getRandomWord).Methods("GET")

  r.Use(func(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Received request: %s %s\n", r.Method, r.URL.Path)
      w.Header().Set("Access-Control-Allow-Origin", "*")
      w.Header().Set("Access-Control-Allow-Methods", "GET")
      w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			fmt.Println("Handling OPTIONS request")
      }
      next.ServeHTTP(w, r)
    })
  })

  // Ambil port dari environment variable PORT, fallback ke 8080 jika tidak ada
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }

  fmt.Printf("Server running on port %s\n", port)
  if err := http.ListenAndServe(":"+port, r); err != nil {
    fmt.Println("Error starting server:", err)
  }
}