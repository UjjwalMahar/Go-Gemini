package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)


type Content struct{
	Parts []string `json:Parts`
	Role string `json:Role`
} 
type Candidates struct {
	Content *Content `json:Content`
}
type ContentResponse struct{
	Candidates *[]Candidates `json:Candidates`
}

func main() {

	//Load .env key
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env key")
	}

	api_key := os.Getenv("API_KEY")

	//initializing the gemini
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(api_key))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	model := client.GenerativeModel("gemini-pro-vision")


	img1, err := os.ReadFile("images/earth.jpeg")
	if err != nil {
		log.Fatal("Error Loading img1")
	}
	img2, err := os.ReadFile("images/modi.jpeg")
	if err != nil {
		log.Fatal("Error Loading img2")
	}
	img3, err := os.ReadFile("images/trump.jpeg")
	if err != nil {
		log.Fatal("Error Loading img3")
	}

	prompt := []genai.Part{
		genai.ImageData("jpeg", img1),
		genai.ImageData("jpeg", img2),
		genai.ImageData("jpeg", img3),
		genai.Text("First identify the pictures then, Write a story out of it "),
	}

	resp, err := model.GenerateContent(ctx, prompt...)
	if err !=nil{
		log.Fatal( err)
	}

	marshalResponse,_ := json.MarshalIndent(resp,"","  ")

	var generateResponse ContentResponse
	if err := json.Unmarshal(marshalResponse, &generateResponse); err !=nil{
		log.Fatal(err)
	}

	for _, cad := range *generateResponse.Candidates{
		if cad.Content !=nil{
			for _, part := range cad.Content.Parts{
				fmt.Print(part)
			}
		}
	}


}

