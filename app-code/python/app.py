"""A small support assistant: drafts replies to tickets with OpenAI, and
summarises them for the team with Anthropic."""

from anthropic import Anthropic
from openai import OpenAI

openai_client = OpenAI()
anthropic_client = Anthropic()


def draft_reply(ticket: str) -> str:
    response = openai_client.chat.completions.create(
        model="gpt-4o",
        messages=[
            {"role": "system", "content": "You write short, friendly replies to support tickets."},
            {"role": "user", "content": ticket},
        ],
    )
    return response.choices[0].message.content


def summarise(ticket: str) -> str:
    message = anthropic_client.messages.create(
        model="claude-sonnet-4-6",
        max_tokens=512,
        messages=[{"role": "user", "content": f"Summarise this support ticket in one sentence:\n\n{ticket}"}],
    )
    return message.content[0].text


if __name__ == "__main__":
    ticket = "My invoice for last month shows two charges for the same plan."
    print(summarise(ticket))
    print(draft_reply(ticket))
