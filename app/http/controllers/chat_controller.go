package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/henkgo/chatgpt/common/config"
	"github.com/henkgo/chatgpt/common/logger"
	gogpt "github.com/sashabaranov/go-gpt3"
)

// ChatController ...
type ChatController struct {
	BaseController
}

// NewChatController ...
func NewChatController() *ChatController {
	return &ChatController{}
}

// Index  : page for index
func (c *ChatController) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main website",
	})
}

// Completion ...
func (c *ChatController) Completion(ctx *gin.Context) {
	var request gogpt.ChatCompletionRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		c.ResponseJSON(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	logger.Info(request)
	if len(request.Messages) == 0 {
		c.ResponseJSON(ctx, http.StatusBadRequest, "request messages required", nil)
		return
	}

	cnf := config.LoadConfig()
	client := gogpt.NewClient(cnf.APIKey)
	if request.Messages[0].Role != "system" {
		newMessage := append([]gogpt.ChatCompletionMessage{
			{Role: "system", Content: cnf.BotDesc},
		}, request.Messages...)
		request.Messages = newMessage
		logger.Info(request.Messages)
	}

	if cnf.Model == gogpt.GPT3Dot5Turbo0301 || cnf.Model == gogpt.GPT3Dot5Turbo {
		request.Model = cnf.Model
		resp, err := client.CreateChatCompletion(ctx, request)
		if err != nil {
			c.ResponseJSON(ctx, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		c.ResponseJSON(ctx, http.StatusOK, "", gin.H{
			"reply":    resp.Choices[0].Message.Content,
			"messages": append(request.Messages, resp.Choices[0].Message),
		})
	} else {
		prompt := ""
		for _, item := range request.Messages {
			prompt += item.Content + "/n"
		}
		prompt = strings.Trim(prompt, "/n")

		logger.Info("request prompt is %s", prompt)
		req := gogpt.CompletionRequest{
			Model:            cnf.Model,
			MaxTokens:        cnf.MaxTokens,
			TopP:             cnf.TopP,
			FrequencyPenalty: cnf.FrequencyPenalty,
			PresencePenalty:  cnf.PresencePenalty,
			Prompt:           prompt,
		}

		resp, err := client.CreateCompletion(ctx, req)
		if err != nil {
			c.ResponseJSON(ctx, http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.ResponseJSON(ctx, http.StatusOK, "", gin.H{
			"reply": resp.Choices[0].Text,
			"messages": append(request.Messages, gogpt.ChatCompletionMessage{
				Role:    "assistant",
				Content: resp.Choices[0].Text,
			}),
		})
	}

}
