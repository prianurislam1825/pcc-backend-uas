package ai

import (
	"fmt"
	"strings"
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
	input := strings.ToLower(userInput)

	if strings.Contains(input, "siapa") {
		return "Halo! Saya adalah FikomBot, asisten virtual resmi Fakultas Ilmu Komputer UDB Surakarta. Ada yang bisa saya bantu terkait informasi akademik?"
	} else if strings.Contains(input, "halo") || strings.Contains(input, "hai") {
		return "Halo juga! Selamat datang di layanan FikomBot. Silakan tanyakan apa saja."
	} else if strings.Contains(input, "bantu") {
		return "Tentu! Saya bisa membantu kamu menjawab pertanyaan seputar jadwal kuliah, KRS, atau informasi pendaftaran. Apa yang ingin kamu tanyakan?"
	} else if strings.Contains(input, "kuliah") {
		return "Untuk informasi perkuliahan Fakultas Ilmu Komputer UDB, kamu bisa mengecek jadwal terupdate di website resmi kampus atau menghubungi BAAK."
	} else if strings.Contains(input, "terima kasih") || strings.Contains(input, "makasih") {
		return "Sama-sama! Senang bisa membantu. Jangan ragu untuk bertanya lagi jika butuh bantuan."
	}

	// Default respon jika keyword tidak ada
	return "Maaf, saya tidak mengerti pertanyaanmu. Sebagai asisten demo FikomBot, saya hanya dapat merespons beberapa pertanyaan umum saat ini."
}
