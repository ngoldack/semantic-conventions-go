# semantic-conventions-go

Go package providing [OpenInference semantic conventions](https://arize-ai.github.io/openinference/spec/semantic_conventions.html) for tracing LLM applications, agents, and related AI systems.

## Installation

```sh
go get github.com/ngoldack/semantic-conventions-go
```

Requires Go 1.24+.

## Usage

### Attribute key constants

Use predefined string constants for span attribute keys:

```go
import semconv "github.com/ngoldack/semantic-conventions-go"

span.SetAttributes(
    attribute.String(semconv.InputValue, "What is the capital of France?"),
    attribute.String(semconv.InputMimeType, semconv.MimeTypeText.String()),
    attribute.String(semconv.LLMModelName, "gpt-4o"),
    attribute.String(semconv.OpenInferenceSpanKindKey, semconv.SpanKindLLM.String()),
)
```

### Typed enums

The package provides typed string enums for well-known values:

```go
semconv.SpanKindLLM       // OpenInferenceSpanKind = "LLM"
semconv.SpanKindAgent     // OpenInferenceSpanKind = "AGENT"
semconv.MessageRoleUser   // MessageRole = "user"
semconv.MessageRoleSystem // MessageRole = "system"
semconv.MimeTypeJSON      // MimeType = "application/json"
semconv.LLMSystemOpenAI   // LLMSystem = "openai"
semconv.LLMProviderAzure  // LLMProvider = "azure"
```

### Helper functions for flattened attributes

OpenInference flattens nested lists into indexed attribute keys. Use the helper functions to build these keys:

```go
// "llm.input_messages.0.message.role"
semconv.InputMessageAttribute(0, semconv.MessageRoleKey)

// "llm.output_messages.0.message.content"
semconv.OutputMessageAttribute(0, semconv.MessageContent)

// "llm.output_messages.0.message.tool_calls.0.tool_call.function.name"
semconv.OutputMessageToolCallAttribute(0, 0, semconv.ToolCallFunctionName)

// "retrieval.documents.2.document.content"
semconv.RetrievalDocumentAttribute(2, semconv.DocumentContent)

// "embedding.embeddings.0.embedding.text"
semconv.EmbeddingAttribute(0, semconv.EmbeddingText)
```

### Configuration environment variables

The package exposes constants for OpenInference privacy/redaction environment variables:

```go
semconv.EnvHideInputs          // "OPENINFERENCE_HIDE_INPUTS"
semconv.EnvHideOutputs         // "OPENINFERENCE_HIDE_OUTPUTS"
semconv.EnvHideInputMessages   // "OPENINFERENCE_HIDE_INPUT_MESSAGES"
semconv.EnvHideOutputMessages  // "OPENINFERENCE_HIDE_OUTPUT_MESSAGES"
semconv.EnvHideInputImages     // "OPENINFERENCE_HIDE_INPUT_IMAGES"
semconv.EnvBase64ImageMaxLength // "OPENINFERENCE_BASE64_IMAGE_MAX_LENGTH"
```

When content is hidden, attributes are replaced with `semconv.RedactedValue` (`"__REDACTED__"`).

## Attribute Categories

| Category | Example Constants |
|---|---|
| Input/Output | `InputValue`, `OutputValue`, `InputMimeType`, `OutputMimeType` |
| LLM | `LLMModelName`, `LLMInvocationParameters`, `LLMTokenCountTotal`, `LLMCostTotal` |
| Messages | `MessageRoleKey`, `MessageContent`, `MessageToolCalls` |
| Message Content | `MessageContentTypeKey`, `MessageContentText`, `MessageContentImage` |
| Documents | `DocumentID`, `DocumentContent`, `DocumentScore`, `DocumentMetadata` |
| Embeddings | `EmbeddingText`, `EmbeddingModelName`, `EmbeddingVector` |
| Retrieval | `RetrievalDocuments` |
| Reranker | `RerankerQuery`, `RerankerModelName`, `RerankerTopK` |
| Tools | `ToolName`, `ToolDescription`, `ToolParameters`, `ToolJSONSchema` |
| Tool Calls | `ToolCallFunctionName`, `ToolCallFunctionArgumentsJSON`, `ToolCallID` |
| Session/User | `SessionID`, `UserID` |
| Agent | `AgentName` |
| Graph | `GraphNodeID`, `GraphNodeName`, `GraphNodeParentID` |
| Prompt Registry | `PromptVendor`, `PromptID`, `PromptURL` |
| Exceptions | `ExceptionType`, `ExceptionMessage`, `ExceptionStacktrace` |

## License

See [LICENSE](LICENSE).