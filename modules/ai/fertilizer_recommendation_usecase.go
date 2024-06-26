package ai

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/sashabaranov/go-openai"
)

type AIFertilizerRecommendationService interface {
	GetFertilizerRecommendation(plantName string) ([]FertilizerRecommendation, error)
	GetPlantingRecommendation(plantName string) ([]PlantingRecommendation, error)
}

type aiFertilizerRecommendationService struct {
	aiClient *openai.Client
}

func NewPlantService(apiKey string) AIFertilizerRecommendationService {
	client := openai.NewClient(apiKey)
	return &aiFertilizerRecommendationService{aiClient: client}
}

func (s *aiFertilizerRecommendationService) GetFertilizerRecommendation(plantName string) ([]FertilizerRecommendation, error) {
	prompt := fmt.Sprintf("Provide detailed fertilizer recommendations, including the type of fertilizer, dosage, and application instructions for %s. Please provide the information in a structured format: Fertilizer: <fertilizer>, Dosage: <dosage>, Application: <application>.", plantName)

	response, err := s.aiClient.CreateChatCompletion(context.TODO(), openai.ChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "You are an expert in agriculture, especially in fertilizers and plant care.",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
		MaxTokens: 200,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get recommendation: %w", err)
	}

	// Logging the response for debugging
	log.Println("Response from OpenAI:", response.Choices[0].Message.Content)

	// Extracting text from the response
	text := response.Choices[0].Message.Content
	lines := strings.Split(text, "\n")

	var recommendations []FertilizerRecommendation
	var currentRecommendation strings.Builder

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Fertilizer:") {
			if currentRecommendation.Len() > 0 {
				// Push previous recommendation
				recommendations = append(recommendations, FertilizerRecommendation{
					PlantName:      plantName,
					Recommendation: currentRecommendation.String(),
				})
				// Clear current recommendation
				currentRecommendation.Reset()
			}
			currentRecommendation.WriteString(strings.TrimSpace(strings.TrimPrefix(line, "Fertilizer:")))
			currentRecommendation.WriteString("\n")
		} else if strings.HasPrefix(line, "Dosage:") {
			currentRecommendation.WriteString(strings.TrimSpace(strings.TrimPrefix(line, "Dosage:")))
			currentRecommendation.WriteString("\n")
		} else if strings.HasPrefix(line, "Application:") {
			currentRecommendation.WriteString(strings.TrimSpace(strings.TrimPrefix(line, "Application:")))
			currentRecommendation.WriteString("\n")
		}
	}

	// Append the last recommendation
	if currentRecommendation.Len() > 0 {
		recommendations = append(recommendations, FertilizerRecommendation{
			PlantName:      plantName,
			Recommendation: currentRecommendation.String(),
		})
	}

	if len(recommendations) == 0 {
		return nil, fmt.Errorf("no valid recommendations found in the response")
	}

	return recommendations, nil
}

func (s *aiFertilizerRecommendationService) GetPlantingRecommendation(plantName string) ([]PlantingRecommendation, error) {
	prompt := fmt.Sprintf("Provide detailed planting recommendations, including weather conditions, tools, soil type, and other relevant information for planting %s. Please provide the information in a structured format: Weather: <weather>, Tools: <tools>, Soil: <soil>, Instructions: <instructions>.", plantName)

	response, err := s.aiClient.CreateChatCompletion(context.TODO(), openai.ChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "You are an expert in agriculture, especially in planting and plant care.",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
		MaxTokens: 200,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get recommendation: %w", err)
	}

	// Logging the response for debugging
	log.Println("Response from OpenAI:", response.Choices[0].Message.Content)

	// Extracting text from the response
	text := response.Choices[0].Message.Content
	lines := strings.Split(text, "\n")

	var recommendations []PlantingRecommendation
	var currentRecommendation strings.Builder

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Weather:") {
			if currentRecommendation.Len() > 0 {
				// Push previous recommendation
				recommendations = append(recommendations, PlantingRecommendation{
					PlantName:      plantName,
					Recommendation: currentRecommendation.String(),
				})
				// Clear current recommendation
				currentRecommendation.Reset()
			}
			currentRecommendation.WriteString(strings.TrimSpace(strings.TrimPrefix(line, "Weather:")))
			currentRecommendation.WriteString("\n")
		} else if strings.HasPrefix(line, "Tools:") {
			currentRecommendation.WriteString(strings.TrimSpace(strings.TrimPrefix(line, "Tools:")))
			currentRecommendation.WriteString("\n")
		} else if strings.HasPrefix(line, "Soil:") {
			currentRecommendation.WriteString(strings.TrimSpace(strings.TrimPrefix(line, "Soil:")))
			currentRecommendation.WriteString("\n")
		} else if strings.HasPrefix(line, "Instructions:") {
			currentRecommendation.WriteString(strings.TrimSpace(strings.TrimPrefix(line, "Instructions:")))
			currentRecommendation.WriteString("\n")
		}
	}

	// Append the last recommendation
	if currentRecommendation.Len() > 0 {
		recommendations = append(recommendations, PlantingRecommendation{
			PlantName:      plantName,
			Recommendation: currentRecommendation.String(),
		})
	}

	if len(recommendations) == 0 {
		return nil, fmt.Errorf("no valid recommendations found in the response")
	}

	return recommendations, nil
}
