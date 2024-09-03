[![Build Status](https://github.com/MrAndiw/instagram-roasting/actions/workflows/go.yml/badge.svg)](https://github.com/MrAndiw/instagram-roasting/actions)
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
[![Instagram](https://img.shields.io/badge/Instagram-%23E4405F.svg?logo=Instagram&logoColor=white)](#)
[![Google Gemini](https://img.shields.io/badge/Google%20Gemini-886FBF?logo=googlegemini&logoColor=fff)](#)
[![ChatGPT](https://img.shields.io/badge/ChatGPT-74aa9c?logo=openai&logoColor=white)](#)


# Instagram Roasting

Instagram Roasting is a terminal-based application built with Go and the `tcell` library. It allows users to input text, which is then formatted and displayed in an output box.

## Features

- Terminal-based interface using the `tcell` library.
- Input text is displayed in a designated output box.
- Automatic word wrapping for long text.
- Support Gemini AI

## Installation

To run the Instagram Roasting application, you need to have Go installed on your system. If you don't have Go installed, you can download and install it from [here](https://golang.org/dl/).

1. Clone the repository:

```sh
git clone https://github.com/MrAndiw/instagram-roasting.git
cd instagram-roasting
```

2. Create file Config.json:

```sh
{
    "API_KEY_GEMINI" : "xxxxxxxxxxxx",
    "TEXT_ROAST" : "Tolong Ejek dengan sopan Instagram Akun ini"
}
```

3. you can create your own API_KEY_GEMINI from this website page : https://aistudio.google.com/app/apikey

4. Run the App directly:

```sh
go run main.go
```

5. Build & Run from bulder:

```sh
BUILD :
Mac & Ubuntu : go build -o main .
Windows : go build -o main.exe .

RUN :
Mac & Ubuntu : ./main
Windows : ./main.exe
```

6. Demo App:

https://github.com/user-attachments/assets/967f4fda-62ee-4428-a40d-2fe9f8feb7a3
