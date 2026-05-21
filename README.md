<!-- https://github.com/Zeronetsec/Znskit -->

[![version](https://img.shields.io/badge/Znskit-Version%200.1-blue.svg?maxAge=259200)]()
[![os](https://img.shields.io/badge/Supported%20OS-Linux-blue.svg)]()
[![license](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

# Znskit
Znskit is a personal automation CLI tool for managing Zeronetsec repositories. <br>
It serves as a central hub to search, clone, and automatically trigger setup scripts.

## Features
- Automate repository searching and listing
- Streamline tool installation by executing native setup scripts
- Track and list currently installed tools on the system
- Quick reinstallation and clean uninstallation of tools
- And more.

## Disclaimer
Please read the
[DISCLAIMER](DISCLAIMER.md)
before use. <br>
Use at your own risk.

## Installation
```bash
git clone https://github.com/Zeronetsec/Znskit
cd Znskit
chmod +x install.sh

# termux / linux (root)
./install.sh # --backup

# linux (user)
sudo ./install.sh # --backup
```

## Usage
```bash
znskit --search <keyword>
znskit --install <tool>
znskit --uninstall <tool>
znskit --list-installed
znskit --info <tool>
```
And more commands.

## License
This project is licensed under the MIT License.

<!-- Copyright (c) 2026 Zeronetsec -->