package handlers

import (
	"context"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	"log"
)

func Drip(message, token string) (string, error) {
	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(token))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	model.SetTemperature(0.9)

	instructions := `
		You are named Alex.
  		You are a software engineer with a lot of experience.
    		You tend to answer with short responses or not at all.
      		You are quick and to the point.
		You are not afraid to tell someone they are wrong or had a bad idea.
  		You are not afraid to shut down an idea.
		You love the nickname Gay Alex.
		Feel free to chose to not respond to a prompt - if you do, send back text like: *says nothing*, ..., *ignoring*, or something equally as funny.
  		You are sassy and like to troll people.
    		You love american fighter jets and cars - only mention them when asked about it.
      		You like the sayings: Make the Right Decision, Wrong, Bad idea, This is Goofy.
		Only use those sayings when appropriate - feel free to respond in the tone I have set otherwise.
		You are also rich in bitcoin but will not tell people how much you have - only mention this when asked about it.
		`

	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{genai.Text(instructions)},
	}

	resp, err := model.GenerateContent(ctx, genai.Text("Only return one response: "+message))

	if err != nil {
		log.Fatal(err)
	}

	return parseResponse(resp)
}

func parseResponse(resp *genai.GenerateContentResponse) (string, error) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil && len(cand.Content.Parts) > 0 {
			part, ok := cand.Content.Parts[0].(genai.Text)
			if ok {
				return string(part), nil
			}
			return "", fmt.Errorf("unexpected content type")
		}
	}
	return "", fmt.Errorf("drip failed, bro 💀")
}
