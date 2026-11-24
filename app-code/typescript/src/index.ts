// A small support assistant: drafts replies to tickets with OpenAI, and
// summarises them for the team with Anthropic.
import Anthropic from "@anthropic-ai/sdk";
import OpenAI from "openai";

const openai = new OpenAI();
const anthropic = new Anthropic();

export async function draftReply(ticket: string): Promise<string> {
  const response = await openai.chat.completions.create({
    model: "gpt-5.5",
    messages: [
      { role: "system", content: "You write short, friendly replies to support tickets." },
      { role: "user", content: ticket },
    ],
  });
  return response.choices[0].message.content ?? "";
}

export async function summarise(ticket: string): Promise<string> {
  const message = await anthropic.messages.create({
    model: "claude-haiku-4-5",
    max_tokens: 512,
    messages: [{ role: "user", content: `Summarise this support ticket in one sentence:\n\n${ticket}` }],
  });
  const block = message.content[0];
  return block.type === "text" ? block.text : "";
}

const ticket = "My invoice for last month shows two charges for the same plan.";
console.log(await summarise(ticket));
console.log(await draftReply(ticket));
