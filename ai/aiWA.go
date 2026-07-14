package ai

import (
	"fmt"
	"sync"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

var client openai.Client

var aiReady bool

var userHistories = make(map[string][]openai.ChatCompletionMessageParamUnion)

var mu sync.Mutex

func InitAi() {
	client = openai.NewClient(
		option.WithBaseURL("http://localhost:8080/v1/"),
		option.WithAPIKey("saksake-karena-gak-butuh-api-key"),
	)
	aiReady = true

	fmt.Println("AI Engine berhasil diinisialisasi.")
}

func TanyaAi(userID string, userInput string) string {
	// Karena di Railway tidak ada server AI lokal yang jalan di port 8080,
	// kita beri respons dummy (pura-pura) agar tidak error saat didemo.
	return "Halo! Saya adalah FikomBot (Bot AI Simulasi). \nSaat ini koneksi ke engine AI pusat sedang dinonaktifkan untuk demo server cloud.\n\nNamun, pastikan kamu mencoba fitur integrasi *Google Drive* kami yang berjalan 100% sempurna di cloud!"
}
