package semconv_test

import (
	"testing"

	semconv "github.com/ngoldack/semantic-conventions-go"
)

func TestInputOutputAttributes(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{"InputValue", semconv.InputValue, "input.value"},
		{"InputMimeType", semconv.InputMimeType, "input.mime_type"},
		{"OutputValue", semconv.OutputValue, "output.value"},
		{"OutputMimeType", semconv.OutputMimeType, "output.mime_type"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("got %q, want %q", tt.value, tt.want)
			}
		})
	}
}

func TestLLMAttributes(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{"LLMInputMessages", semconv.LLMInputMessages, "llm.input_messages"},
		{"LLMOutputMessages", semconv.LLMOutputMessages, "llm.output_messages"},
		{"LLMModelName", semconv.LLMModelName, "llm.model_name"},
		{"LLMPrompts", semconv.LLMPrompts, "llm.prompts"},
		{"LLMInvocationParameters", semconv.LLMInvocationParameters, "llm.invocation_parameters"},
		{"LLMProvider", semconv.LLMProvider, "llm.provider"},
		{"LLMSystem", semconv.LLMSystem, "llm.system"},
		{"LLMFunctionCall", semconv.LLMFunctionCall, "llm.function_call"},
		{"LLMTools", semconv.LLMTools, "llm.tools"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("got %q, want %q", tt.value, tt.want)
			}
		})
	}
}

func TestLLMTokenCountAttributes(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{"LLMTokenCountPrompt", semconv.LLMTokenCountPrompt, "llm.token_count.prompt"},
		{"LLMTokenCountCompletion", semconv.LLMTokenCountCompletion, "llm.token_count.completion"},
		{"LLMTokenCountTotal", semconv.LLMTokenCountTotal, "llm.token_count.total"},
		{"LLMTokenCountPromptDetails", semconv.LLMTokenCountPromptDetails, "llm.token_count.prompt_details"},
		{"LLMTokenCountPromptDetailsCacheWrite", semconv.LLMTokenCountPromptDetailsCacheWrite, "llm.token_count.prompt_details.cache_write"},
		{"LLMTokenCountPromptDetailsCacheRead", semconv.LLMTokenCountPromptDetailsCacheRead, "llm.token_count.prompt_details.cache_read"},
		{"LLMTokenCountPromptDetailsCacheInput", semconv.LLMTokenCountPromptDetailsCacheInput, "llm.token_count.prompt_details.cache_input"},
		{"LLMTokenCountPromptDetailsAudio", semconv.LLMTokenCountPromptDetailsAudio, "llm.token_count.prompt_details.audio"},
		{"LLMTokenCountCompletionDetails", semconv.LLMTokenCountCompletionDetails, "llm.token_count.completion_details"},
		{"LLMTokenCountCompletionDetailsReasoning", semconv.LLMTokenCountCompletionDetailsReasoning, "llm.token_count.completion_details.reasoning"},
		{"LLMTokenCountCompletionDetailsAudio", semconv.LLMTokenCountCompletionDetailsAudio, "llm.token_count.completion_details.audio"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("got %q, want %q", tt.value, tt.want)
			}
		})
	}
}

func TestLLMCostAttributes(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{"LLMCost", semconv.LLMCost, "llm.cost"},
		{"LLMCostPrompt", semconv.LLMCostPrompt, "llm.cost.prompt"},
		{"LLMCostCompletion", semconv.LLMCostCompletion, "llm.cost.completion"},
		{"LLMCostTotal", semconv.LLMCostTotal, "llm.cost.total"},
		{"LLMCostInput", semconv.LLMCostInput, "llm.cost.prompt_details.input"},
		{"LLMCostOutput", semconv.LLMCostOutput, "llm.cost.completion_details.output"},
		{"LLMCostCompletionDetailsReasoning", semconv.LLMCostCompletionDetailsReasoning, "llm.cost.completion_details.reasoning"},
		{"LLMCostCompletionDetailsAudio", semconv.LLMCostCompletionDetailsAudio, "llm.cost.completion_details.audio"},
		{"LLMCostPromptDetailsCacheWrite", semconv.LLMCostPromptDetailsCacheWrite, "llm.cost.prompt_details.cache_write"},
		{"LLMCostPromptDetailsCacheRead", semconv.LLMCostPromptDetailsCacheRead, "llm.cost.prompt_details.cache_read"},
		{"LLMCostPromptDetailsCacheInput", semconv.LLMCostPromptDetailsCacheInput, "llm.cost.prompt_details.cache_input"},
		{"LLMCostPromptDetailsAudio", semconv.LLMCostPromptDetailsAudio, "llm.cost.prompt_details.audio"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("got %q, want %q", tt.value, tt.want)
			}
		})
	}
}

func TestPromptTemplateAttributes(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{"PromptTemplateVariables", semconv.PromptTemplateVariables, "llm.prompt_template.variables"},
		{"PromptTemplateTemplate", semconv.PromptTemplateTemplate, "llm.prompt_template.template"},
		{"PromptTemplateVersion", semconv.PromptTemplateVersion, "llm.prompt_template.version"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("got %q, want %q", tt.value, tt.want)
			}
		})
	}
}

func TestMessageAttributes(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{"MessageRole", semconv.MessageRole, "message.role"},
		{"MessageContent", semconv.MessageContent, "message.content"},
		{"MessageContents", semconv.MessageContents, "message.contents"},
		{"MessageName", semconv.MessageName, "message.name"},
		{"MessageFunctionCallName", semconv.MessageFunctionCallName, "message.function_call_name"},
		{"MessageFunctionCallArgumentsJSON", semconv.MessageFunctionCallArgumentsJSON, "message.function_call_arguments_json"},
		{"MessageToolCalls", semconv.MessageToolCalls, "message.tool_calls"},
		{"MessageToolCallID", semconv.MessageToolCallID, "message.tool_call_id"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("got %q, want %q", tt.value, tt.want)
			}
		})
	}
}

func TestMessageContentAttributes(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{"MessageContentType", semconv.MessageContentType, "message_content.type"},
		{"MessageContentText", semconv.MessageContentText, "message_content.text"},
		{"MessageContentImage", semconv.MessageContentImage, "message_content.image"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("got %q, want %q", tt.value, tt.want)
			}
		})
	}
}

func TestToolCallAttributes(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{"ToolCallFunctionName", semconv.ToolCallFunctionName, "tool_call.function.name"},
		{"ToolCallFunctionArgumentsJSON", semconv.ToolCallFunctionArgumentsJSON, "tool_call.function.arguments"},
		{"ToolCallID", semconv.ToolCallID, "tool_call.id"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("got %q, want %q", tt.value, tt.want)
			}
		})
	}
}

func TestImageAudioAttributes(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{"ImageURL", semconv.ImageURL, "image.url"},
		{"AudioURL", semconv.AudioURL, "audio.url"},
		{"AudioMimeType", semconv.AudioMimeType, "audio.mime_type"},
		{"AudioTranscript", semconv.AudioTranscript, "audio.transcript"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("got %q, want %q", tt.value, tt.want)
			}
		})
	}
}

func TestDocumentAttributes(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{"DocumentID", semconv.DocumentID, "document.id"},
		{"DocumentContent", semconv.DocumentContent, "document.content"},
		{"DocumentScore", semconv.DocumentScore, "document.score"},
		{"DocumentMetadata", semconv.DocumentMetadata, "document.metadata"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("got %q, want %q", tt.value, tt.want)
			}
		})
	}
}

func TestEmbeddingAttributes(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{"EmbeddingEmbeddings", semconv.EmbeddingEmbeddings, "embedding.embeddings"},
		{"EmbeddingText", semconv.EmbeddingText, "embedding.text"},
		{"EmbeddingModelName", semconv.EmbeddingModelName, "embedding.model_name"},
		{"EmbeddingVector", semconv.EmbeddingVector, "embedding.vector"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("got %q, want %q", tt.value, tt.want)
			}
		})
	}
}

func TestRetrievalRerankerAttributes(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{"RetrievalDocuments", semconv.RetrievalDocuments, "retrieval.documents"},
		{"RerankerInputDocuments", semconv.RerankerInputDocuments, "reranker.input_documents"},
		{"RerankerOutputDocuments", semconv.RerankerOutputDocuments, "reranker.output_documents"},
		{"RerankerQuery", semconv.RerankerQuery, "reranker.query"},
		{"RerankerModelName", semconv.RerankerModelName, "reranker.model_name"},
		{"RerankerTopK", semconv.RerankerTopK, "reranker.top_k"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("got %q, want %q", tt.value, tt.want)
			}
		})
	}
}

func TestToolAttributes(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{"ToolName", semconv.ToolName, "tool.name"},
		{"ToolDescription", semconv.ToolDescription, "tool.description"},
		{"ToolParameters", semconv.ToolParameters, "tool.parameters"},
		{"ToolJSONSchema", semconv.ToolJSONSchema, "tool.json_schema"},
		{"ToolID", semconv.ToolID, "tool.id"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("got %q, want %q", tt.value, tt.want)
			}
		})
	}
}

func TestSessionUserAttributes(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{"SessionID", semconv.SessionID, "session.id"},
		{"UserID", semconv.UserID, "user.id"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("got %q, want %q", tt.value, tt.want)
			}
		})
	}
}

func TestPromptAttributes(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{"PromptVendor", semconv.PromptVendor, "prompt.vendor"},
		{"PromptID", semconv.PromptID, "prompt.id"},
		{"PromptURL", semconv.PromptURL, "prompt.url"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("got %q, want %q", tt.value, tt.want)
			}
		})
	}
}

func TestAgentGraphAttributes(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{"AgentName", semconv.AgentName, "agent.name"},
		{"GraphNodeID", semconv.GraphNodeID, "graph.node.id"},
		{"GraphNodeName", semconv.GraphNodeName, "graph.node.name"},
		{"GraphNodeParentID", semconv.GraphNodeParentID, "graph.node.parent_id"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("got %q, want %q", tt.value, tt.want)
			}
		})
	}
}

func TestGenericAttributes(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{"Metadata", semconv.Metadata, "metadata"},
		{"TagTags", semconv.TagTags, "tag.tags"},
		{"OpenInferenceSpanKindKey", semconv.OpenInferenceSpanKindKey, "openinference.span.kind"},
		{"ProjectName", semconv.ProjectName, "openinference.project.name"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("got %q, want %q", tt.value, tt.want)
			}
		})
	}
}

func TestOpenInferenceSpanKind(t *testing.T) {
	tests := []struct {
		name  string
		value semconv.OpenInferenceSpanKind
		want  string
	}{
		{"SpanKindLLM", semconv.SpanKindLLM, "LLM"},
		{"SpanKindChain", semconv.SpanKindChain, "CHAIN"},
		{"SpanKindTool", semconv.SpanKindTool, "TOOL"},
		{"SpanKindRetriever", semconv.SpanKindRetriever, "RETRIEVER"},
		{"SpanKindReranker", semconv.SpanKindReranker, "RERANKER"},
		{"SpanKindEmbedding", semconv.SpanKindEmbedding, "EMBEDDING"},
		{"SpanKindAgent", semconv.SpanKindAgent, "AGENT"},
		{"SpanKindGuardrail", semconv.SpanKindGuardrail, "GUARDRAIL"},
		{"SpanKindEvaluator", semconv.SpanKindEvaluator, "EVALUATOR"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if string(tt.value) != tt.want {
				t.Errorf("got %q, want %q", tt.value, tt.want)
			}
		})
	}
}

func TestMimeType(t *testing.T) {
	tests := []struct {
		name  string
		value semconv.MimeType
		want  string
	}{
		{"MimeTypeText", semconv.MimeTypeText, "text/plain"},
		{"MimeTypeJSON", semconv.MimeTypeJSON, "application/json"},
		{"MimeTypeAudioWAV", semconv.MimeTypeAudioWAV, "audio/wav"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if string(tt.value) != tt.want {
				t.Errorf("got %q, want %q", tt.value, tt.want)
			}
		})
	}
}

func TestLLMSystem(t *testing.T) {
	tests := []struct {
		name  string
		value semconv.LLMSystem
		want  string
	}{
		{"LLMSystemOpenAI", semconv.LLMSystemOpenAI, "openai"},
		{"LLMSystemAnthropic", semconv.LLMSystemAnthropic, "anthropic"},
		{"LLMSystemMistralAI", semconv.LLMSystemMistralAI, "mistralai"},
		{"LLMSystemCohere", semconv.LLMSystemCohere, "cohere"},
		{"LLMSystemVertexAI", semconv.LLMSystemVertexAI, "vertexai"},
		{"LLMSystemAI21", semconv.LLMSystemAI21, "ai21"},
		{"LLMSystemMeta", semconv.LLMSystemMeta, "meta"},
		{"LLMSystemAmazon", semconv.LLMSystemAmazon, "amazon"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if string(tt.value) != tt.want {
				t.Errorf("got %q, want %q", tt.value, tt.want)
			}
		})
	}
}

func TestLLMProvider(t *testing.T) {
	tests := []struct {
		name  string
		value semconv.LLMProvider
		want  string
	}{
		{"LLMProviderOpenAI", semconv.LLMProviderOpenAI, "openai"},
		{"LLMProviderAnthropic", semconv.LLMProviderAnthropic, "anthropic"},
		{"LLMProviderMistralAI", semconv.LLMProviderMistralAI, "mistralai"},
		{"LLMProviderCohere", semconv.LLMProviderCohere, "cohere"},
		{"LLMProviderGoogle", semconv.LLMProviderGoogle, "google"},
		{"LLMProviderAWS", semconv.LLMProviderAWS, "aws"},
		{"LLMProviderAzure", semconv.LLMProviderAzure, "azure"},
		{"LLMProviderXAI", semconv.LLMProviderXAI, "xai"},
		{"LLMProviderDeepSeek", semconv.LLMProviderDeepSeek, "deepseek"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if string(tt.value) != tt.want {
				t.Errorf("got %q, want %q", tt.value, tt.want)
			}
		})
	}
}
