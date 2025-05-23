<h1 align="left">
  Fakjs - A fast Go-based tool to uncover sensitive information in JavaScript
</h1>

<p align="left">
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/license-MIT-_red.svg"></a>
  <a href="https://github.com/thd3r/fakjs/releases"><img src="https://img.shields.io/github/release/thd3r/fakjs.svg"></a>
  <a href="https://x.com/thd3r"><img src="https://img.shields.io/twitter/follow/thd3r.svg?logo=twitter"></a>
  <a href="https://github.com/thd3r/fakjs/issues"><img src="https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat"></a>

</p>

```sh
███████╗ █████╗ ██╗  ██╗     ██╗███████╗
██╔════╝██╔══██╗██║ ██╔╝     ██║██╔════╝
█████╗  ███████║█████╔╝      ██║███████╗
██╔══╝  ██╔══██║██╔═██╗ ██   ██║╚════██║
██║     ██║  ██║██║  ██╗╚█████╔╝███████║
╚═╝     ╚═╝  ╚═╝╚═╝  ╚═╝ ╚════╝ ╚══════╝
                  v1.0.1																		
```

**Fakjs** is a fast, lightweight, and extensible tool written in Go, designed to extract potentially sensitive information from publicly accessible JavaScript files. It plays a crucial role in reconnaissance during security assessments, allowing you to discover information that might aid in understanding the inner workings of a web application or reveal unintended data exposures.

## 🔍 Why Fakjs?

During penetration testing, bug bounty hunting, or red teaming, analyzing JavaScript files can lead to critical findings. Manual inspection is time-consuming — that's where **Fakjs** comes in. It uses pattern matching, regular expressions, and content heuristics to locate data that may be of interest from a security perspective.

## 🚀 Key Features

- ⚡ **High Performance:** Written in Go for fast execution and low memory usage.
- 🔎 **Automated Detection:** Identifies potential sensitive content through customizable regex-based scanning.
- 🌐 **Remote & Local Support:** Analyze JavaScript from URLs or local file paths.
- 🧠 **Recon-Friendly:** Ideal for OSINT, bug bounty, pentesting, or passive reconnaissance.
- 🛠️ **Easily Integratable:** Can be used standalone or integrated into larger recon pipelines.

## 🧪 What It Looks For

- Hardcoded secrets (API keys, tokens, etc.)
- Internal or hidden endpoints
- Configuration data
- Potential indicators of exposed logic or backend connections

---
> [!NOTE]
> By default, **Fakjs** uses a set of regular expressions to detect common patterns. You can expand or customize these patterns for specific targets or use cases.
---

## Installation

```sh
go install -v github.com/thd3r/fakjs@latest
```

## Usage

### The file contains JavaScript URLs

```sh
cat jsUrls.txt | fakjs
```

###  Or

```sh
cat cunks.js | fakjs
```

---
> [!TIP]
> **Fakjs** automatically generates a report and saves it to a temporary folder.
---

## Acknowledments

Since this tool includes some contributions, I'd like to publicly thank the following users for their help and resources, which provided regex patterns and guidance during the development of this project

- [@DarkHacker420](https://github.com/DarkHacker420)
- [@profmoriarity](https://github.com/profmoriarity)
- [@dwisiswant0](https://github.com/dwisiswant0)
