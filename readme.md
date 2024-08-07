[![Build Status](https://github.com/MrAndiw/instagram-go-scraper/actions/workflows/go.yml/badge.svg)](https://github.com/MrAndiw/instagram-go-scraper/actions)
<img src="https://camo.githubusercontent.com/33cfae3047a121e7811c7a54e7b6ef4029c9db941f3d180a176069220c878954/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f676f2d2532333030414444382e7376673f7374796c653d666f722d7468652d6261646765266c6f676f3d676f266c6f676f436f6c6f723d7768697465" alt="Go" data-canonical-src="https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&amp;logo=go&amp;logoColor=white" style="max-width: 100%;">
<img src="https://camo.githubusercontent.com/7f5701ed50f919cf2352cd028b5b2dc974b5e643fe4d78ad826eb9e74551157f/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f496e7374616772616d2d2532334534343035462e7376673f7374796c653d666f722d7468652d6261646765266c6f676f3d496e7374616772616d266c6f676f436f6c6f723d7768697465" alt="Instagram" data-canonical-src="https://img.shields.io/badge/Instagram-%23E4405F.svg?style=for-the-badge&amp;logo=Instagram&amp;logoColor=white" style="max-width: 100%;">
<img src="https://camo.githubusercontent.com/c99cf08227dfb9aaac41b41c1fb50a371d608a5ba85a1df02a1b31e1fdfd8deb/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f47656d696e692d3845373542323f7374796c653d666f722d7468652d6261646765266c6f676f3d676f6f676c6562617264266c6f676f436f6c6f723d666666" alt="gemini" data-canonical-src="https://img.shields.io/badge/Gemini-8E75B2?style=for-the-badge&amp;logo=googlebard&amp;logoColor=fff" style="max-width: 100%;">
<img src="https://camo.githubusercontent.com/c8af3418f8fe508aed1c66f474b50f9e9d8f64db648d1bd947527b35b6243a99/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f636861744750542d3734616139633f7374796c653d666f722d7468652d6261646765266c6f676f3d6f70656e6169266c6f676f436f6c6f723d7768697465" alt="ChatGPT" data-canonical-src="https://img.shields.io/badge/chatGPT-74aa9c?style=for-the-badge&amp;logo=openai&amp;logoColor=white" style="max-width: 100%;">

# Instagram Roasting

Instagram Roasting is a terminal-based application built with Go and the `tcell` library. It allows users to input text, which is then formatted and displayed in an output box.

## Features

- Terminal-based interface using the `tcell` library.
- Input text is displayed in a designated output box.
- Automatic word wrapping for long text.

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
