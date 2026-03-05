package semconv

// RedactedValue is the placeholder used when content is hidden due to privacy
// configuration settings. Consumers of trace data can use this value to
// distinguish intentionally hidden content from missing or empty values.
const RedactedValue = "__REDACTED__"

// OpenInference configuration environment variable names.
//
// Each variable controls the masking of a specific category of span attribute.
// When a value is hidden its attribute is replaced with [RedactedValue].
// Variables default to false / off unless noted otherwise.
const (
	// EnvHideLLMInvocationParameters hides llm.invocation_parameters (bool, default false).
	EnvHideLLMInvocationParameters = "OPENINFERENCE_HIDE_LLM_INVOCATION_PARAMETERS"

	// EnvHideInputs hides input.value and all input messages (bool, default false).
	// Input messages are also hidden when [EnvHideInputMessages] is set.
	EnvHideInputs = "OPENINFERENCE_HIDE_INPUTS"

	// EnvHideOutputs hides output.value and all output messages (bool, default false).
	// Output messages are also hidden when [EnvHideOutputMessages] is set.
	EnvHideOutputs = "OPENINFERENCE_HIDE_OUTPUTS"

	// EnvHideInputMessages hides all input messages, independent of [EnvHideInputs]
	// (bool, default false).
	EnvHideInputMessages = "OPENINFERENCE_HIDE_INPUT_MESSAGES"

	// EnvHideOutputMessages hides all output messages, independent of [EnvHideOutputs]
	// (bool, default false).
	EnvHideOutputMessages = "OPENINFERENCE_HIDE_OUTPUT_MESSAGES"

	// EnvHideInputImages hides images in input messages; only applies when input
	// messages are not already hidden (bool, default false).
	EnvHideInputImages = "OPENINFERENCE_HIDE_INPUT_IMAGES"

	// EnvHideInputText hides text in input messages; only applies when input
	// messages are not already hidden (bool, default false).
	EnvHideInputText = "OPENINFERENCE_HIDE_INPUT_TEXT"

	// EnvHideOutputText hides text in output messages; only applies when output
	// messages are not already hidden (bool, default false).
	EnvHideOutputText = "OPENINFERENCE_HIDE_OUTPUT_TEXT"

	// EnvHidePrompts hides llm.prompts (completions API legacy input)
	// (bool, default false).
	EnvHidePrompts = "OPENINFERENCE_HIDE_PROMPTS"

	// EnvHideChoices hides llm.choices (completions API legacy output)
	// (bool, default false).
	EnvHideChoices = "OPENINFERENCE_HIDE_CHOICES"

	// EnvHideEmbeddingsVectors replaces embedding vector values with [RedactedValue]
	// (bool, default false).
	EnvHideEmbeddingsVectors = "OPENINFERENCE_HIDE_EMBEDDINGS_VECTORS"

	// EnvHideEmbeddingVectors is deprecated: use [EnvHideEmbeddingsVectors].
	EnvHideEmbeddingVectors = "OPENINFERENCE_HIDE_EMBEDDING_VECTORS"

	// EnvHideEmbeddingsText replaces embedding text values with [RedactedValue]
	// (bool, default false).
	EnvHideEmbeddingsText = "OPENINFERENCE_HIDE_EMBEDDINGS_TEXT"

	// EnvBase64ImageMaxLength limits the character length of base64-encoded images
	// (int, default 32000).
	EnvBase64ImageMaxLength = "OPENINFERENCE_BASE64_IMAGE_MAX_LENGTH"
)

// DefaultBase64ImageMaxLength is the default value for [EnvBase64ImageMaxLength].
const DefaultBase64ImageMaxLength = 32_000
