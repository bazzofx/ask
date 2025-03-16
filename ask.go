package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/go-deepseek/deepseek"
	"github.com/go-deepseek/deepseek/request"
)

func main() {
	// Get the DeepSeek API token from environment variables
	apiDeepSeek := os.Getenv("apiDeepSeek")
	if apiDeepSeek == "" {
		fmt.Println("API key is not set. Please set the 'apiDeepSeek' environment variable.")
		return
	}

	// Initialize the DeepSeek client with the API token
	client, err := deepseek.NewClient(apiDeepSeek)
	if err != nil {
		fmt.Println("Error initializing DeepSeek client:", err)
		return
	}

	// Define your role and base question format
	role := "You are a master in PowerShell, Azure CLI, and DevOps. You understand Windows OS and all its variations. You provide the best cmd-let codes to accomplish tasks, verified by Microsoft documentation. The answers must be either a DOS command or a PowerShell cmd-let, nothing else."
	baseQuestion := "Give the cmd-let only, only give the best answer, only the first best result cmd-let command only, nothing else, no marks, no extra information, only the cmd-let to give the user the answer to the question. Do not include ```powershell on the answer, give the cmd-let ONLY. QUESTION:"

	// Ask user for a question
	fmt.Print("»")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userInput := scanner.Text()

	// Form the complete query
	query := baseQuestion + userInput

	// Create a ChatCompletionsRequest for the DeepSeek API
	chatReq := &request.ChatCompletionsRequest{
		Model:  deepseek.DEEPSEEK_CHAT_MODEL, // Ensure the model name is correct
		Stream: false,
		Messages: []*request.Message{
			{
				Role:    "system",
				Content: role,
			},
			{
				Role:    "user",
				Content: query,
			},
		},
	}

	// Make the API call to get the response
	chatResp, err := client.CallChatCompletionsChat(context.Background(), chatReq)
	if err != nil {
		fmt.Println("Error in API request:", err)
		return
	}

	// Output the response from DeepSeek
	if len(chatResp.Choices) > 0 {
		fmt.Println(chatResp.Choices[0].Message.Content)
	} else {
		fmt.Println("No response received.")
	}
}
