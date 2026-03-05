package semconv

// OpenInferenceSpanKind identifies the type of operation being traced.
// The [OpenInferenceSpanKindKey] attribute is required for all OpenInference spans.
type OpenInferenceSpanKind string

// String returns the string value of the span kind.
func (k OpenInferenceSpanKind) String() string { return string(k) }

const (
	// SpanKindLLM identifies a span that calls a large language model.
	SpanKindLLM OpenInferenceSpanKind = "LLM"

	// SpanKindEmbedding identifies a span that creates embeddings.
	SpanKindEmbedding OpenInferenceSpanKind = "EMBEDDING"

	// SpanKindChain identifies a starting point or link between LLM application steps.
	SpanKindChain OpenInferenceSpanKind = "CHAIN"

	// SpanKindRetriever identifies a span that retrieves documents.
	SpanKindRetriever OpenInferenceSpanKind = "RETRIEVER"

	// SpanKindReranker identifies a span that reranks a set of documents.
	SpanKindReranker OpenInferenceSpanKind = "RERANKER"

	// SpanKindTool identifies a span that calls an external tool.
	SpanKindTool OpenInferenceSpanKind = "TOOL"

	// SpanKindAgent identifies a span that encompasses calls to LLMs and tools.
	SpanKindAgent OpenInferenceSpanKind = "AGENT"

	// SpanKindGuardrail identifies a span that applies a guardrail to LLM output.
	SpanKindGuardrail OpenInferenceSpanKind = "GUARDRAIL"

	// SpanKindEvaluator identifies a span that evaluates LLM outputs.
	SpanKindEvaluator OpenInferenceSpanKind = "EVALUATOR"

	// SpanKindPrompt identifies a span that renders a prompt template.
	SpanKindPrompt OpenInferenceSpanKind = "PROMPT"
)

// MimeType represents common MIME types used in OpenInference span attributes.
type MimeType string

// String returns the string value of the MIME type.
func (m MimeType) String() string { return string(m) }

const (
	// MimeTypeText is the plain text MIME type.
	MimeTypeText MimeType = "text/plain"

	// MimeTypeJSON is the JSON MIME type.
	MimeTypeJSON MimeType = "application/json"

	// MimeTypeAudioWAV is the WAV audio MIME type.
	MimeTypeAudioWAV MimeType = "audio/wav"
)

// LLMSystem represents the AI product as identified by the client or server instrumentation.
// Used as the value for the [LLMSystemKey] attribute.
type LLMSystem string

// String returns the string value of the LLM system.
func (s LLMSystem) String() string { return string(s) }

const (
	LLMSystemOpenAI    LLMSystem = "openai"
	LLMSystemAnthropic LLMSystem = "anthropic"
	LLMSystemVertexAI  LLMSystem = "vertexai"
	LLMSystemCohere    LLMSystem = "cohere"
	LLMSystemMistralAI LLMSystem = "mistralai"
	LLMSystemXAI       LLMSystem = "xai"
	LLMSystemDeepSeek  LLMSystem = "deepseek"
	LLMSystemAmazon    LLMSystem = "amazon"
	LLMSystemMeta      LLMSystem = "meta"
	LLMSystemAI21      LLMSystem = "ai21"
)

// LLMProvider represents the hosting provider of the LLM.
// Used as the value for the [LLMProviderKey] attribute.
type LLMProvider string

// String returns the string value of the LLM provider.
func (p LLMProvider) String() string { return string(p) }

const (
	LLMProviderOpenAI    LLMProvider = "openai"
	LLMProviderAnthropic LLMProvider = "anthropic"
	LLMProviderCohere    LLMProvider = "cohere"
	LLMProviderMistralAI LLMProvider = "mistralai"
	LLMProviderAzure     LLMProvider = "azure"
	LLMProviderGoogle    LLMProvider = "google"
	LLMProviderAWS       LLMProvider = "aws"
	LLMProviderXAI       LLMProvider = "xai"
	LLMProviderDeepSeek  LLMProvider = "deepseek"
)

// MessageRole represents the role of an entity in a conversation message.
// Used as the value for the [MessageRoleKey] attribute.
type MessageRole string

// String returns the string value of the message role.
func (r MessageRole) String() string { return string(r) }

const (
	MessageRoleUser      MessageRole = "user"
	MessageRoleAssistant MessageRole = "assistant"
	MessageRoleSystem    MessageRole = "system"
	MessageRoleTool      MessageRole = "tool"
	MessageRoleFunction  MessageRole = "function"
)

// MessageContentType represents the type of content in a multimodal message.
// Used as the value for the [MessageContentTypeKey] attribute.
type MessageContentType string

// String returns the string value of the message content type.
func (c MessageContentType) String() string { return string(c) }

const (
	MessageContentTypeText  MessageContentType = "text"
	MessageContentTypeImage MessageContentType = "image"
	MessageContentTypeAudio MessageContentType = "audio"
)
