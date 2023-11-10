# 🎙️ Go-Script-Tokenizer

![Build](https://github.com/kmesiab/go-token-sdk/actions/workflows/go.yml/badge.svg)

![Golang](https://img.shields.io/badge/Go-00add8.svg?labelColor=171e21&style=for-the-badge&logo=go)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Terraform](https://img.shields.io/badge/terraform-%235835CC.svg?style=for-the-badge&logo=terraform&logoColor=white)
![AWS](https://img.shields.io/badge/AWS-%23FF9900.svg?style=for-the-badge&logo=amazon-aws&logoColor=white)

![Go-Script-Tokenizer Mascot](https://github.com/kmesiab/go-script-tokenizer/blob/main/assets/go-script-tokenizer-mascot-png.png)


Go-Script-Tokenizer is a Go-based utility that converts AWS Transcribe JSON output into a structured, tokenized format, ready for ingestion by Large Language Models (LLMs).

## Features

- Parses AWS Transcribe JSON files.
- Extracts and tokenizes spoken words.
- Identifies and tags speakers.
- Preserves timestamps for contextual relevance.
- Outputs in a format compatible with NLP models and LLMs.

## Getting Started

### Prerequisites

- Golang installed on your machine.
- Access to AWS Transcribe output files.

### Installation

```shell
go get github.com/kevin-mesiab/go-script-tokenizer
```

### Usage

```shell
go run main.go -input="path/to/transcribe.json"
```

Replace `path/to/transcribe.json\` with the actual path to your AWS Transcribe JSON file.

## Contributing

Contributions are what make the open-source community such a fantastic place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request
