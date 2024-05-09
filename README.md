# GoLang Translation Server

## Overview

This GoLang project functions as a server designed to handle JSON requests containing transcriptions, translate them using Google's Gemini AI, and return the translated subtitles to the client. It serves as a backend component for applications that require real-time subtitle translation.

## Features

- **JSON Request Handling:** Receives JSON requests with the field `transcription: data`.
- **Google Gemini AI Integration:** Utilizes Google's advanced Gemini AI for accurate and context-aware translations.
- **Translation Response:** Returns the translated text in the JSON format with the field `subtitle: data`.

## Requirements

- GoLang environment set up on the server.
- Access to Google's Gemini API with the required authentication.

## Installation

1. Clone or download this repository to your server.
2. Ensure that GoLang is installed on your system. If not, install it from [the official GoLang website](https://golang.org/dl/).
3. Navigate to the project directory.
4. Rename the `.env.example` file to `.env` and add values for `API_KEY` and `PORT`:
    - `API_KEY` should contain the authentication key for accessing Google's Gemini API.
    - `PORT` should contain the port on which the server will run.
5. Run `go mod tidy` to fetch the required packages and clean up the module, ensuring all dependencies are correct and up-to-date.
6. To start the server, use the command `go run cmd/main.go` in the root directory of the project.

## Usage

- **Send Requests:** Use any HTTP client to send JSON-formatted requests containing a `transcription` field to the server.
- **Receive Translated Subtitles:** The server processes the request using Google's Gemini AI and returns the translated subtitles in JSON format with the `subtitle` field.

## API Endpoints

- **/api/transcribe** (POST): Accepts JSON requests with a `transcription` field. Returns a JSON response containing the `subtitle` field with the translated data.

## Example Request

```json
{
  "transcription": "Hello, world!"
}
```

## Example Response

```json
{
  "subtitle": "Olá, mundo!"
}
```

## Development

- **Setting Up Google's Gemini API:**
  - Make sure you have valid credentials and access tokens for Google's Gemini API.
  - Configure these credentials in the `.env` file to authenticate API requests.

## Contribution

Contributions are welcome! Please fork the repository and submit pull requests with your improvements. Ensure you follow coding standards and include unit tests for new features.

