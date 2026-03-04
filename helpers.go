package semconv

import "fmt"

// InputMessageAttribute returns the flattened attribute key for a field within
// an indexed input message, e.g. "llm.input_messages.0.message.role".
func InputMessageAttribute(index int, suffix string) string {
	return fmt.Sprintf("%s.%d.%s", LLMInputMessages, index, suffix)
}

// OutputMessageAttribute returns the flattened attribute key for a field within
// an indexed output message, e.g. "llm.output_messages.0.message.content".
func OutputMessageAttribute(index int, suffix string) string {
	return fmt.Sprintf("%s.%d.%s", LLMOutputMessages, index, suffix)
}

// InputMessageContentAttribute returns the flattened attribute key for a content
// item within an indexed input message, e.g.
// "llm.input_messages.0.message.contents.1.message_content.text".
func InputMessageContentAttribute(messageIndex, contentIndex int, suffix string) string {
	return fmt.Sprintf("%s.%d.message.contents.%d.message_content.%s",
		LLMInputMessages, messageIndex, contentIndex, suffix)
}

// OutputMessageToolCallAttribute returns the flattened attribute key for a tool
// call within an indexed output message, e.g.
// "llm.output_messages.0.message.tool_calls.0.tool_call.function.name".
func OutputMessageToolCallAttribute(messageIndex, toolCallIndex int, suffix string) string {
	return fmt.Sprintf("%s.%d.%s.%d.%s",
		LLMOutputMessages, messageIndex, MessageToolCalls, toolCallIndex, suffix)
}
