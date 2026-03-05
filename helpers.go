package semconv

import "fmt"

// ── LLM input messages ──────────────────────────────────────────────────────

// InputMessageAttribute returns the flattened attribute key for a field within
// an indexed input message, e.g.
//
//	InputMessageAttribute(0, MessageRoleKey) → "llm.input_messages.0.message.role"
func InputMessageAttribute(index int, suffix string) string {
	return fmt.Sprintf("%s.%d.%s", LLMInputMessages, index, suffix)
}

// InputMessageContentAttribute returns the flattened attribute key for a
// content item within an indexed input message, e.g.
//
//	InputMessageContentAttribute(0, 1, "text") → "llm.input_messages.0.message.contents.1.message_content.text"
func InputMessageContentAttribute(messageIndex, contentIndex int, suffix string) string {
	return fmt.Sprintf("%s.%d.message.contents.%d.message_content.%s",
		LLMInputMessages, messageIndex, contentIndex, suffix)
}

// ── LLM output messages ─────────────────────────────────────────────────────

// OutputMessageAttribute returns the flattened attribute key for a field within
// an indexed output message, e.g.
//
//	OutputMessageAttribute(0, MessageContent) → "llm.output_messages.0.message.content"
func OutputMessageAttribute(index int, suffix string) string {
	return fmt.Sprintf("%s.%d.%s", LLMOutputMessages, index, suffix)
}

// OutputMessageContentAttribute returns the flattened attribute key for a
// content item within an indexed output message, e.g.
//
//	OutputMessageContentAttribute(0, 0, "text") → "llm.output_messages.0.message.contents.0.message_content.text"
func OutputMessageContentAttribute(messageIndex, contentIndex int, suffix string) string {
	return fmt.Sprintf("%s.%d.message.contents.%d.message_content.%s",
		LLMOutputMessages, messageIndex, contentIndex, suffix)
}

// OutputMessageToolCallAttribute returns the flattened attribute key for a tool
// call within an indexed output message, e.g.
//
//	OutputMessageToolCallAttribute(0, 0, ToolCallFunctionName) → "llm.output_messages.0.message.tool_calls.0.tool_call.function.name"
func OutputMessageToolCallAttribute(messageIndex, toolCallIndex int, suffix string) string {
	return fmt.Sprintf("%s.%d.%s.%d.%s",
		LLMOutputMessages, messageIndex, MessageToolCalls, toolCallIndex, suffix)
}

// ── Legacy completions API ──────────────────────────────────────────────────

// LLMPromptAttribute returns the flattened attribute key for the text of an
// indexed legacy prompt, e.g.
//
//	LLMPromptAttribute(0) → "llm.prompts.0.prompt.text"
func LLMPromptAttribute(index int) string {
	return fmt.Sprintf("%s.%d.%s", LLMPrompts, index, PromptText)
}

// LLMChoiceAttribute returns the flattened attribute key for the text of an
// indexed legacy completion choice, e.g.
//
//	LLMChoiceAttribute(0) → "llm.choices.0.completion.text"
func LLMChoiceAttribute(index int) string {
	return fmt.Sprintf("%s.%d.%s", LLMChoices, index, CompletionText)
}

// ── LLM tools ───────────────────────────────────────────────────────────────

// LLMToolAttribute returns the flattened attribute key for a field within an
// indexed tool definition, e.g.
//
//	LLMToolAttribute(0, ToolJSONSchema) → "llm.tools.0.tool.json_schema"
func LLMToolAttribute(index int, suffix string) string {
	return fmt.Sprintf("%s.%d.%s", LLMTools, index, suffix)
}

// ── Embeddings ──────────────────────────────────────────────────────────────

// EmbeddingAttribute returns the flattened attribute key for a field within an
// indexed embedding object, e.g.
//
//	EmbeddingAttribute(0, EmbeddingText) → "embedding.embeddings.0.embedding.text"
func EmbeddingAttribute(index int, suffix string) string {
	return fmt.Sprintf("%s.%d.%s", EmbeddingEmbeddings, index, suffix)
}

// ── Retrieval ────────────────────────────────────────────────────────────────

// RetrievalDocumentAttribute returns the flattened attribute key for a field
// within an indexed retrieved document, e.g.
//
//	RetrievalDocumentAttribute(0, DocumentContent) → "retrieval.documents.0.document.content"
func RetrievalDocumentAttribute(index int, suffix string) string {
	return fmt.Sprintf("%s.%d.%s", RetrievalDocuments, index, suffix)
}

// ── Reranker ─────────────────────────────────────────────────────────────────

// RerankerInputDocumentAttribute returns the flattened attribute key for a
// field within an indexed reranker input document, e.g.
//
//	RerankerInputDocumentAttribute(0, DocumentScore) → "reranker.input_documents.0.document.score"
func RerankerInputDocumentAttribute(index int, suffix string) string {
	return fmt.Sprintf("%s.%d.%s", RerankerInputDocuments, index, suffix)
}

// RerankerOutputDocumentAttribute returns the flattened attribute key for a
// field within an indexed reranker output document, e.g.
//
//	RerankerOutputDocumentAttribute(0, DocumentID) → "reranker.output_documents.0.document.id"
func RerankerOutputDocumentAttribute(index int, suffix string) string {
	return fmt.Sprintf("%s.%d.%s", RerankerOutputDocuments, index, suffix)
}
