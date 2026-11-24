// A small support assistant: drafts replies to tickets with OpenAI, and
// summarises them for the team with Anthropic.
package com.example.support;

import com.anthropic.client.AnthropicClient;
import com.anthropic.client.okhttp.AnthropicOkHttpClient;
import com.anthropic.models.messages.Message;
import com.anthropic.models.messages.MessageCreateParams;
import com.openai.client.OpenAIClient;
import com.openai.client.okhttp.OpenAIOkHttpClient;
import com.openai.models.chat.completions.ChatCompletion;
import com.openai.models.chat.completions.ChatCompletionCreateParams;

public class SupportAssistant {
    private final OpenAIClient openai = OpenAIOkHttpClient.fromEnv();
    private final AnthropicClient anthropic = AnthropicOkHttpClient.fromEnv();

    public ChatCompletion draftReply(String ticket) {
        return openai.chat().completions().create(ChatCompletionCreateParams.builder()
                .model("gpt-5.5")
                .addSystemMessage("You write short, friendly replies to support tickets.")
                .addUserMessage(ticket)
                .build());
    }

    public Message summarise(String ticket) {
        return anthropic.messages().create(MessageCreateParams.builder()
                .model("claude-sonnet-4-6")
                .maxTokens(512)
                .addUserMessage("Summarise this support ticket in one sentence:\n\n" + ticket)
                .build());
    }

    public static void main(String[] args) {
        SupportAssistant assistant = new SupportAssistant();
        String ticket = "My invoice for last month shows two charges for the same plan.";
        System.out.println(assistant.summarise(ticket));
        System.out.println(assistant.draftReply(ticket));
    }
}
