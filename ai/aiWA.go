package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var aiReady bool

func InitAi() {
	if os.Getenv("GEMINI_API_KEY") != "" {
		aiReady = true
		fmt.Println("Gemini AI Engine berhasil diinisialisasi.")
	} else {
		fmt.Println("Warning: GEMINI_API_KEY belum di-set!")
	}
}

type GeminiContent struct {
	Parts []GeminiPart `json:"parts"`
	Role  string       `json:"role,omitempty"`
}

type GeminiPart struct {
	Text string `json:"text"`
}

type GeminiRequest struct {
	Contents          []GeminiContent `json:"contents"`
	SystemInstruction *GeminiContent  `json:"systemInstruction,omitempty"`
}

type GeminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []GeminiPart `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

func TanyaAi(userID string, userInput string) string {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return "Mohon maaf, API Key AI belum diatur di server cloud."
	}

	url := "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.5-flash:generateContent?key=" + apiKey

	reqBody := GeminiRequest{
		SystemInstruction: &GeminiContent{
			Parts: []GeminiPart{{Text: "Anda adalah FikomBot, asisten virtual resmi Fakultas Ilmu Komputer UDB Surakarta."}},
		},
		Contents: []GeminiContent{
			{
				Role:  "user",
				Parts: []GeminiPart{{Text: userInput}},
			},
		},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "Error saat menyusun data AI."
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "Mohon maaf, terjadi gangguan saat terhubung ke Gemini AI."
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Sprintf("Mohon maaf, respon error dari Gemini AI (Status: %d).", resp.StatusCode)
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var geminiResp GeminiResponse
	err = json.Unmarshal(bodyBytes, &geminiResp)
	if err != nil {
		return "Gagal memproses jawaban dari AI."
	}

	if len(geminiResp.Candidates) > 0 && len(geminiResp.Candidates[0].Content.Parts) > 0 {
		return geminiResp.Candidates[0].Content.Parts[0].Text
	}

	return "Maaf, AI tidak memberikan respons."
}
