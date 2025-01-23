# 🚀 Yai 💬 - AI powered terminal assistant

[![build](https://github.com/mounta11n/yai/actions/workflows/build.yml/badge.svg)](https://github.com/mounta11n/yai/actions/workflows/build.yml)
[![release](https://github.com/mounta11n/yai/actions/workflows/release.yml/badge.svg)](https://github.com/mounta11n/yai/actions/workflows/release.yml)
[![doc](https://github.com/mounta11n/yai/actions/workflows/doc.yml/badge.svg)](https://github.com/mounta11n/yai/actions/workflows/doc.yml)

> Unleash the power of artificial intelligence to streamline your command line experience.

![Intro](docs/_assets/intro.gif)

## What is Yai ?

`Yai` (your AI) is an assistant for your terminal, using OpenAI compatible APIs to build and run commands for you. You just need to describe them in your everyday language, it will take care or the rest.

You have any questions on random topics in mind? You can also ask `Yai`, and get the power of AI without leaving `/home`.

It is already aware of your:

- operating system & distribution
- username, shell & home directory
- preferred editor

And you can also give any supplementary preferences to fine tune your experience.

## Documentation

A complete documentation is available at [https://mounta11n.github.io/yai/](https://mounta11n.github.io/yai/).

## Build

To build `Yai`, simply run:

```shell
make build
```

To install the local build, simply run:

```shell
make install
```

To uninstall the local build and remove the configuration, simply run:

```shell
make uninstall
```

## Configuration

Create `~/.config/yai.json`, with the following structure:

```json
{
  "openai_base_url": "https://api.openai.com/v1", // OpenAI compatible API URL
  "openai_key": "sk-xxxxxxxxx", // API key (mandatory)
  "openai_model": "gpt-3.5-turbo", // AI model (default gpt-3.5-turbo)
  "openai_proxy": "", // OpenAI API proxy (default disabled)
  "openai_temperature": 0.2, // OpenAI API temperature (defaut 0.2)
  "openai_max_tokens": 1000, // OpenAI API max tokens (default 1000)
  "user_default_prompt_mode": "exec", // user prefered prompt mode: "exec" (default) or "chat"
  "user_preferences": "" // user preferences, expressed in natural language (default none)
}
```

## Thanks

Thanks to [@K-arch27](https://github.com/K-arch27) for the `Yai` name suggestion.
