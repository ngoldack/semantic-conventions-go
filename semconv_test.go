package semconv_test

import (
	"testing"

	semconv "github.com/ngoldack/semantic-conventions-go"
)

// --- Attribute key constants ---

func TestInputOutputAttributes(t *testing.T) {
	cases := []struct{ name, got, want string }{
		{"InputValue", semconv.InputValue, "input.value"},
		{"InputMimeType", semconv.InputMimeType, "input.mime_type"},
		{"OutputValue", semconv.OutputValue, "output.value"},
		{"OutputMimeType", semconv.OutputMimeType, "output.mime_type"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.got != c.want {
				t.Errorf("got %q, want %q", c.got, c.want)
			}
		})
	}
}

func TestLLMAttributes(t *testing.T) {
	cases := []struct{ name, got, want string }{
		{"LLMInputMessages", semconv.LLMInputMessages, "llm.input_messages"},
		{"LLMOutputMessages", semconv.LLMOutputMessages, "llm.output_messages"},
		{"LLMModelName", semconv.LLMModelName, "llm.model_name"},
		{"LLMPrompts", semconv.LLMPrompts, "llm.prompts"},
		{"LLMChoices", semconv.LLMChoices, "llm.choices"},
		{"LLMInvocationParameters", semconv.LLMInvocationParameters, "llm.invocation_parameters"},
		{"LLMProviderKey", semconv.LLMProviderKey, "llm.provider"},
		{"LLMSystemKey", semconv.LLMSystemKey, "llm.system"},
		{"LLMFunctionCall", semconv.LLMFunctionCall, "llm.function_call"},
		{"LLMTools", semconv.LLMTools, "llm.tools"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.got != c.want {
				t.Errorf("got %q, want %q", c.got, c.want)
			}
		})
	}
}

func TestLLMTokenCountAttributes(t *testing.T) {
	cases := []struct{ name, got, want string }{
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
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.got != c.want {
				t.Errorf("got %q, want %q", c.got, c.want)
			}
		})
	}
}

func TestLLMCostAttributes(t *testing.T) {
	cases := []struct{ name, got, want string }{
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
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.got != c.want {
				t.Errorf("got %q, want %q", c.got, c.want)
			}
		})
	}
}

func TestPromptTemplateAttributes(t *testing.T) {
	cases := []struct{ name, got, want string }{
		{"PromptTemplateVariables", semconv.PromptTemplateVariables, "llm.prompt_template.variables"},
		{"PromptTemplateTemplate", semconv.PromptTemplateTemplate, "llm.prompt_template.template"},
		{"PromptTemplateVersion", semconv.PromptTemplateVersion, "llm.prompt_template.version"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.got != c.want {
				t.Errorf("got %q, want %q", c.got, c.want)
			}
		})
	}
}

func TestMessageAttributes(t *testing.T) {
	cases := []struct{ name, got, want string }{
		{"MessageRole", semconv.MessageRole, "message.role"},
		{"MessageContent", semconv.MessageContent, "message.content"},
		{"MessageContents", semconv.MessageContents, "message.contents"},
		{"MessageName", semconv.MessageName, "message.name"},
		{"MessageFunctionCallName", semconv.MessageFunctionCallName, "message.function_call_name"},
		{"MessageFunctionCallArgumentsJSON", semconv.MessageFunctionCallArgumentsJSON, "message.function_call_arguments_json"},
		{"MessageToolCalls", semconv.MessageToolCalls, "message.tool_calls"},
		{"MessageToolCallID", semconv.MessageToolCallID, "message.tool_call_id"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.got != c.want {
				t.Errorf("got %q, want %q", c.got, c.want)
			}
		})
	}
}

func TestMessageContentAttributes(t *testing.T) {
	cases := []struct{ name, got, want string }{
		{"MessageContentType", semconv.MessageContentType, "message_content.type"},
		{"MessageContentText", semconv.MessageContentText, "message_content.text"},
		{"MessageContentImage", semconv.MessageContentImage, "message_content.image"},
		{"MessageContentAudio", semconv.MessageContentAudio, "message_content.audio"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.got != c.want {
				t.Errorf("got %q, want %q", c.got, c.want)
			}
		})
	}
}

func TestToolCallAttributes(t *testing.T) {
	cases := []struct{ name, got, want string }{
		{"ToolCallFunctionName", semconv.ToolCallFunctionName, "tool_call.function.name"},
		{"ToolCallFunctionArgumentsJSON", semconv.ToolCallFunctionArgumentsJSON, "tool_call.function.arguments"},
		{"ToolCallID", semconv.ToolCallID, "tool_call.id"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.got != c.want {
				t.Errorf("got %q, want %q", c.got, c.want)
			}
		})
	}
}

func TestImageAudioAttributes(t *testing.T) {
	cases := []struct{ name, got, want string }{
		{"ImageURL", semconv.ImageURL, "image.url"},
		{"AudioURL", semconv.AudioURL, "audio.url"},
		{"AudioMimeType", semconv.AudioMimeType, "audio.mime_type"},
		{"AudioTranscript", semconv.AudioTranscript, "audio.transcript"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.got != c.want {
				t.Errorf("got %q, want %q", c.got, c.want)
			}
		})
	}
}

func TestDocumentAttributes(t *testing.T) {
	cases := []struct{ name, got, want string }{
		{"DocumentID", semconv.DocumentID, "document.id"},
		{"DocumentContent", semconv.DocumentContent, "document.content"},
		{"DocumentScore", semconv.DocumentScore, "document.score"},
		{"DocumentMetadata", semconv.DocumentMetadata, "document.metadata"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.got != c.want {
				t.Errorf("got %q, want %q", c.got, c.want)
			}
		})
	}
}

func TestEmbeddingAttributes(t *testing.T) {
	cases := []struct{ name, got, want string }{
		{"EmbeddingEmbeddings", semconv.EmbeddingEmbeddings, "embedding.embeddings"},
		{"EmbeddingText", semconv.EmbeddingText, "embedding.text"},
		{"EmbeddingModelName", semconv.EmbeddingModelName, "embedding.model_name"},
		{"EmbeddingVector", semconv.EmbeddingVector, "embedding.vector"},
		{"EmbeddingInvocationParameters", semconv.EmbeddingInvocationParameters, "embedding.invocation_parameters"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.got != c.want {
				t.Errorf("got %q, want %q", c.got, c.want)
			}
		})
	}
}

func TestRetrievalRerankerAttributes(t *testing.T) {
	cases := []struct{ name, got, want string }{
		{"RetrievalDocuments", semconv.RetrievalDocuments, "retrieval.documents"},
		{"RerankerInputDocuments", semconv.RerankerInputDocuments, "reranker.input_documents"},
		{"RerankerOutputDocuments", semconv.RerankerOutputDocuments, "reranker.output_documents"},
		{"RerankerQuery", semconv.RerankerQuery, "reranker.query"},
		{"RerankerModelName", semconv.RerankerModelName, "reranker.model_name"},
		{"RerankerTopK", semconv.RerankerTopK, "reranker.top_k"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.got != c.want {
				t.Errorf("got %q, want %q", c.got, c.want)
			}
		})
	}
}

func TestToolAttributes(t *testing.T) {
	cases := []struct{ name, got, want string }{
		{"ToolName", semconv.ToolName, "tool.name"},
		{"ToolDescription", semconv.ToolDescription, "tool.description"},
		{"ToolParameters", semconv.ToolParameters, "tool.parameters"},
		{"ToolJSONSchema", semconv.ToolJSONSchema, "tool.json_schema"},
		{"ToolID", semconv.ToolID, "tool.id"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.got != c.want {
				t.Errorf("got %q, want %q", c.got, c.want)
			}
		})
	}
}

func TestSessionUserAttributes(t *testing.T) {
	cases := []struct{ name, got, want string }{
		{"SessionID", semconv.SessionID, "session.id"},
		{"UserID", semconv.UserID, "user.id"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.got != c.want {
				t.Errorf("got %q, want %q", c.got, c.want)
			}
		})
	}
}

func TestPromptAttributes(t *testing.T) {
	cases := []struct{ name, got, want string }{
		{"PromptVendor", semconv.PromptVendor, "prompt.vendor"},
		{"PromptID", semconv.PromptID, "prompt.id"},
		{"PromptURL", semconv.PromptURL, "prompt.url"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.got != c.want {
				t.Errorf("got %q, want %q", c.got, c.want)
			}
		})
	}
}

func TestAgentGraphAttributes(t *testing.T) {
	cases := []struct{ name, got, want string }{
		{"AgentName", semconv.AgentName, "agent.name"},
		{"GraphNodeID", semconv.GraphNodeID, "graph.node.id"},
		{"GraphNodeName", semconv.GraphNodeName, "graph.node.name"},
		{"GraphNodeParentID", semconv.GraphNodeParentID, "graph.node.parent_id"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.got != c.want {
				t.Errorf("got %q, want %q", c.got, c.want)
			}
		})
	}
}

func TestGenericAndOpenInferenceAttributes(t *testing.T) {
	cases := []struct{ name, got, want string }{
		{"Metadata", semconv.Metadata, "metadata"},
		{"TagTags", semconv.TagTags, "tag.tags"},
		{"OpenInferenceSpanKindKey", semconv.OpenInferenceSpanKindKey, "openinference.span.kind"},
		{"ProjectName", semconv.ProjectName, "openinference.project.name"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.got != c.want {
				t.Errorf("got %q, want %q", c.got, c.want)
			}
		})
	}
}

func TestExceptionAttributes(t *testing.T) {
	cases := []struct{ name, got, want string }{
		{"ExceptionEscaped", semconv.ExceptionEscaped, "exception.escaped"},
		{"ExceptionMessage", semconv.ExceptionMessage, "exception.message"},
		{"ExceptionStacktrace", semconv.ExceptionStacktrace, "exception.stacktrace"},
		{"ExceptionType", semconv.ExceptionType, "exception.type"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.got != c.want {
				t.Errorf("got %q, want %q", c.got, c.want)
			}
		})
	}
}

// --- Enum types ---

func TestOpenInferenceSpanKind(t *testing.T) {
	cases := []struct {
		name string
		got  semconv.OpenInferenceSpanKind
		want string
	}{
		{"SpanKindLLM", semconv.SpanKindLLM, "LLM"},
		{"SpanKindEmbedding", semconv.SpanKindEmbedding, "EMBEDDING"},
		{"SpanKindChain", semconv.SpanKindChain, "CHAIN"},
		{"SpanKindRetriever", semconv.SpanKindRetriever, "RETRIEVER"},
		{"SpanKindReranker", semconv.SpanKindReranker, "RERANKER"},
		{"SpanKindTool", semconv.SpanKindTool, "TOOL"},
		{"SpanKindAgent", semconv.SpanKindAgent, "AGENT"},
		{"SpanKindGuardrail", semconv.SpanKindGuardrail, "GUARDRAIL"},
		{"SpanKindEvaluator", semconv.SpanKindEvaluator, "EVALUATOR"},
		{"SpanKindPrompt", semconv.SpanKindPrompt, "PROMPT"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.got.String() != c.want {
				t.Errorf("got %q, want %q", c.got, c.want)
			}
		})
	}
}

func TestMimeType(t *testing.T) {
	cases := []struct {
		name string
		got  semconv.MimeType
		want string
	}{
		{"MimeTypeText", semconv.MimeTypeText, "text/plain"},
		{"MimeTypeJSON", semconv.MimeTypeJSON, "application/json"},
		{"MimeTypeAudioWAV", semconv.MimeTypeAudioWAV, "audio/wav"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.got.String() != c.want {
				t.Errorf("got %q, want %q", c.got, c.want)
			}
		})
	}
}

func TestLLMSystem(t *testing.T) {
	cases := []struct {
		name string
		got  semconv.LLMSystem
		want string
	}{
		{"LLMSystemOpenAI", semconv.LLMSystemOpenAI, "openai"},
		{"LLMSystemAnthropic", semconv.LLMSystemAnthropic, "anthropic"},
		{"LLMSystemVertexAI", semconv.LLMSystemVertexAI, "vertexai"},
		{"LLMSystemCohere", semconv.LLMSystemCohere, "cohere"},
		{"LLMSystemMistralAI", semconv.LLMSystemMistralAI, "mistralai"},
		{"LLMSystemXAI", semconv.LLMSystemXAI, "xai"},
		{"LLMSystemDeepSeek", semconv.LLMSystemDeepSeek, "deepseek"},
		{"LLMSystemAmazon", semconv.LLMSystemAmazon, "amazon"},
		{"LLMSystemMeta", semconv.LLMSystemMeta, "meta"},
		{"LLMSystemAI21", semconv.LLMSystemAI21, "ai21"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.got.String() != c.want {
				t.Errorf("got %q, want %q", c.got, c.want)
			}
		})
	}
}

func TestLLMProvider(t *testing.T) {
	cases := []struct {
		name string
		got  semconv.LLMProvider
		want string
	}{
		{"LLMProviderOpenAI", semconv.LLMProviderOpenAI, "openai"},
		{"LLMProviderAnthropic", semconv.LLMProviderAnthropic, "anthropic"},
		{"LLMProviderCohere", semconv.LLMProviderCohere, "cohere"},
		{"LLMProviderMistralAI", semconv.LLMProviderMistralAI, "mistralai"},
		{"LLMProviderAzure", semconv.LLMProviderAzure, "azure"},
		{"LLMProviderGoogle", semconv.LLMProviderGoogle, "google"},
		{"LLMProviderAWS", semconv.LLMProviderAWS, "aws"},
		{"LLMProviderXAI", semconv.LLMProviderXAI, "xai"},
		{"LLMProviderDeepSeek", semconv.LLMProviderDeepSeek, "deepseek"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.got.String() != c.want {
				t.Errorf("got %q, want %q", c.got, c.want)
			}
		})
	}
}

// --- Helper functions ---

func TestInputMessageAttribute(t *testing.T) {
	cases := []struct {
		name   string
		index  int
		suffix string
		want   string
	}{
		{"role at 0", 0, "message.role", "llm.input_messages.0.message.role"},
		{"content at 1", 1, "message.content", "llm.input_messages.1.message.content"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := semconv.InputMessageAttribute(c.index, c.suffix); got != c.want {
				t.Errorf("got %q, want %q", got, c.want)
			}
		})
	}
}

func TestOutputMessageAttribute(t *testing.T) {
	cases := []struct {
		name   string
		index  int
		suffix string
		want   string
	}{
		{"role at 0", 0, "message.role", "llm.output_messages.0.message.role"},
		{"content at 2", 2, "message.content", "llm.output_messages.2.message.content"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := semconv.OutputMessageAttribute(c.index, c.suffix); got != c.want {
				t.Errorf("got %q, want %q", got, c.want)
			}
		})
	}
}

func TestInputMessageContentAttribute(t *testing.T) {
	cases := []struct {
		name                       string
		messageIndex, contentIndex int
		suffix                     string
		want                       string
	}{
		{"text at msg0 content1", 0, 1, "text", "llm.input_messages.0.message.contents.1.message_content.text"},
		{"type at msg1 content0", 1, 0, "type", "llm.input_messages.1.message.contents.0.message_content.type"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := semconv.InputMessageContentAttribute(c.messageIndex, c.contentIndex, c.suffix)
			if got != c.want {
				t.Errorf("got %q, want %q", got, c.want)
			}
		})
	}
}

func TestOutputMessageContentAttribute(t *testing.T) {
	cases := []struct {
		name                       string
		messageIndex, contentIndex int
		suffix                     string
		want                       string
	}{
		{"text at msg0 content0", 0, 0, "text", "llm.output_messages.0.message.contents.0.message_content.text"},
		{"type at msg1 content2", 1, 2, "type", "llm.output_messages.1.message.contents.2.message_content.type"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := semconv.OutputMessageContentAttribute(c.messageIndex, c.contentIndex, c.suffix)
			if got != c.want {
				t.Errorf("got %q, want %q", got, c.want)
			}
		})
	}
}

func TestOutputMessageToolCallAttribute(t *testing.T) {
	cases := []struct {
		name                        string
		messageIndex, toolCallIndex int
		suffix                      string
		want                        string
	}{
		{"function.name at msg0 call0", 0, 0, "tool_call.function.name", "llm.output_messages.0.message.tool_calls.0.tool_call.function.name"},
		{"id at msg1 call2", 1, 2, "tool_call.id", "llm.output_messages.1.message.tool_calls.2.tool_call.id"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := semconv.OutputMessageToolCallAttribute(c.messageIndex, c.toolCallIndex, c.suffix)
			if got != c.want {
				t.Errorf("got %q, want %q", got, c.want)
			}
		})
	}
}

func TestLLMPromptAttribute(t *testing.T) {
	cases := []struct {
		name  string
		index int
		want  string
	}{
		{"prompt text at 0", 0, "llm.prompts.0.prompt.text"},
		{"prompt text at 1", 1, "llm.prompts.1.prompt.text"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := semconv.LLMPromptAttribute(c.index); got != c.want {
				t.Errorf("got %q, want %q", got, c.want)
			}
		})
	}
}

func TestLLMChoiceAttribute(t *testing.T) {
	cases := []struct {
		name  string
		index int
		want  string
	}{
		{"choice text at 0", 0, "llm.choices.0.completion.text"},
		{"choice text at 2", 2, "llm.choices.2.completion.text"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := semconv.LLMChoiceAttribute(c.index); got != c.want {
				t.Errorf("got %q, want %q", got, c.want)
			}
		})
	}
}

func TestLLMToolAttribute(t *testing.T) {
	cases := []struct {
		name   string
		index  int
		suffix string
		want   string
	}{
		{"json_schema at 0", 0, "tool.json_schema", "llm.tools.0.tool.json_schema"},
		{"json_schema at 1", 1, "tool.json_schema", "llm.tools.1.tool.json_schema"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := semconv.LLMToolAttribute(c.index, c.suffix); got != c.want {
				t.Errorf("got %q, want %q", got, c.want)
			}
		})
	}
}

func TestEmbeddingAttribute(t *testing.T) {
	cases := []struct {
		name   string
		index  int
		suffix string
		want   string
	}{
		{"text at 0", 0, "embedding.text", "embedding.embeddings.0.embedding.text"},
		{"vector at 1", 1, "embedding.vector", "embedding.embeddings.1.embedding.vector"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := semconv.EmbeddingAttribute(c.index, c.suffix); got != c.want {
				t.Errorf("got %q, want %q", got, c.want)
			}
		})
	}
}

func TestRetrievalDocumentAttribute(t *testing.T) {
	cases := []struct {
		name   string
		index  int
		suffix string
		want   string
	}{
		{"content at 0", 0, "document.content", "retrieval.documents.0.document.content"},
		{"score at 1", 1, "document.score", "retrieval.documents.1.document.score"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := semconv.RetrievalDocumentAttribute(c.index, c.suffix); got != c.want {
				t.Errorf("got %q, want %q", got, c.want)
			}
		})
	}
}

func TestRerankerInputDocumentAttribute(t *testing.T) {
	cases := []struct {
		name   string
		index  int
		suffix string
		want   string
	}{
		{"id at 0", 0, "document.id", "reranker.input_documents.0.document.id"},
		{"score at 2", 2, "document.score", "reranker.input_documents.2.document.score"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := semconv.RerankerInputDocumentAttribute(c.index, c.suffix); got != c.want {
				t.Errorf("got %q, want %q", got, c.want)
			}
		})
	}
}

func TestRerankerOutputDocumentAttribute(t *testing.T) {
	cases := []struct {
		name   string
		index  int
		suffix string
		want   string
	}{
		{"id at 0", 0, "document.id", "reranker.output_documents.0.document.id"},
		{"content at 1", 1, "document.content", "reranker.output_documents.1.document.content"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := semconv.RerankerOutputDocumentAttribute(c.index, c.suffix); got != c.want {
				t.Errorf("got %q, want %q", got, c.want)
			}
		})
	}
}

// --- Legacy completions sub-field constants ---

func TestLegacyCompletionsSubfields(t *testing.T) {
	cases := []struct{ name, got, want string }{
		{"PromptText", semconv.PromptText, "prompt.text"},
		{"CompletionText", semconv.CompletionText, "completion.text"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.got != c.want {
				t.Errorf("got %q, want %q", c.got, c.want)
			}
		})
	}
}

// --- Configuration constants ---

func TestConfigEnvVarNames(t *testing.T) {
	cases := []struct{ name, got, want string }{
		{"EnvHideLLMInvocationParameters", semconv.EnvHideLLMInvocationParameters, "OPENINFERENCE_HIDE_LLM_INVOCATION_PARAMETERS"},
		{"EnvHideInputs", semconv.EnvHideInputs, "OPENINFERENCE_HIDE_INPUTS"},
		{"EnvHideOutputs", semconv.EnvHideOutputs, "OPENINFERENCE_HIDE_OUTPUTS"},
		{"EnvHideInputMessages", semconv.EnvHideInputMessages, "OPENINFERENCE_HIDE_INPUT_MESSAGES"},
		{"EnvHideOutputMessages", semconv.EnvHideOutputMessages, "OPENINFERENCE_HIDE_OUTPUT_MESSAGES"},
		{"EnvHideInputImages", semconv.EnvHideInputImages, "OPENINFERENCE_HIDE_INPUT_IMAGES"},
		{"EnvHideInputText", semconv.EnvHideInputText, "OPENINFERENCE_HIDE_INPUT_TEXT"},
		{"EnvHideOutputText", semconv.EnvHideOutputText, "OPENINFERENCE_HIDE_OUTPUT_TEXT"},
		{"EnvHidePrompts", semconv.EnvHidePrompts, "OPENINFERENCE_HIDE_PROMPTS"},
		{"EnvHideChoices", semconv.EnvHideChoices, "OPENINFERENCE_HIDE_CHOICES"},
		{"EnvHideEmbeddingsVectors", semconv.EnvHideEmbeddingsVectors, "OPENINFERENCE_HIDE_EMBEDDINGS_VECTORS"},
		{"EnvHideEmbeddingVectors", semconv.EnvHideEmbeddingVectors, "OPENINFERENCE_HIDE_EMBEDDING_VECTORS"},
		{"EnvHideEmbeddingsText", semconv.EnvHideEmbeddingsText, "OPENINFERENCE_HIDE_EMBEDDINGS_TEXT"},
		{"EnvBase64ImageMaxLength", semconv.EnvBase64ImageMaxLength, "OPENINFERENCE_BASE64_IMAGE_MAX_LENGTH"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.got != c.want {
				t.Errorf("got %q, want %q", c.got, c.want)
			}
		})
	}
}

func TestRedactedValue(t *testing.T) {
	if semconv.RedactedValue != "__REDACTED__" {
		t.Errorf("got %q, want %q", semconv.RedactedValue, "__REDACTED__")
	}
}

func TestDefaultBase64ImageMaxLength(t *testing.T) {
	if semconv.DefaultBase64ImageMaxLength != 32_000 {
		t.Errorf("got %d, want 32000", semconv.DefaultBase64ImageMaxLength)
	}
}
