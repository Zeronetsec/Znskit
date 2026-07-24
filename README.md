<!-- https://github.com/Zeronetsec/Znskit -->

<div align="center">
    <img src="https://img.shields.io/badge/Znskit-Version%200.1-blue?style=square&logo=go&v=1" />
    <img src="https://img.shields.io/badge/Supported%20OS-Linux-blue?style=square&logo=linux&v=1" />
    <a href="LICENSE">
        <img src="https://img.shields.io/badge/License-Apache--2.0-blue?style=square&logo=github&v=1" />
    </a>
</div>

# Znskit
Znskit is a personal automation CLI tool for managing Zeronetsec repositories.

## Features
- Automate repository searching and listing.
- Streamline tool installation by executing native setup scripts.
- Track and list currently installed tools on the system.
- Quick reinstallation and clean uninstallation of tools.
- And more features.

## Disclaimer
Please read [.docs/disclaimer.md](.docs/disclaimer.md) before using this tool. </br>
Use this software at your own risk. </br>
The author is not responsible for any damage, data loss, or issues that may result from its use.

## Installation
Quick install:
```bash
git clone https://github.com/Zeronetsec/Znskit
bash Znskit/install.sh
```
For more detailed installation and uninstallation instructions, see [.docs/install_and_uninstall.md](.docs/install_and_uninstall.md).

## Usage Example
```bash
znskit --search cli
znskit --install chprompt --iflag --backup
znskit --uninstall chprompt --uflag --home=/home/zeronetsec
znskit --list details
znskit --info comet
```
And more commands.

<!-- Copyright (c) 2026 Zeronetsec -->