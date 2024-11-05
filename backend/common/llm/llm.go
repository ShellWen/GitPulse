package llm

type SparkModelData struct {
	MaxTokens   int64                `json:"max_tokens"`
	TopK        int64                `json:"top_k"`
	Temperature float64              `json:"temperature"`
	Messages    [2]SparkModelMessage `json:"messages"`
	Model       string               `json:"model"`
}

type SparkModelMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type SparkModelResp struct {
	Code    int64              `json:"code"`
	Message string             `json:"message"`
	Sid     string             `json:"sid"`
	Choices []SparkModelChoice `json:"choices"`
	Usage   SparkModelUsage    `json:"usage"`
}

type SparkModelChoice struct {
	Message SparkModelMessage `json:"message"`
	Index   int64             `json:"index"`
}

type SparkModelUsage struct {
	PromptTokens     int64 `json:"prompt_tokens"`
	CompletionTokens int64 `json:"completion_tokens"`
	TotalTokens      int64 `json:"total_tokens"`
}
