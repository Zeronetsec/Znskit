<!-- https://github.com/Zeronetsec/Znskit -->

[![version](https://img.shields.io/badge/Znskit-Version%200.1-blue.svg)]()
[![os](https://img.shields.io/badge/Supported%20OS-Linux-blue.svg)]()
[![license](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

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
cd Znskit
chmod +x install.sh
./install.sh
```
For more detailed installation and uninstallation instructions, see [.docs/install_and_uninstall.md](.docs/install_and_uninstall.md).

## Usage Example
```bash
znskit --search cli
znskit --install woofind
znskit --uninstall woofind
znskit --list-installed
znskit --info woofind
```
And more commands.

## License
This project is licensed under the MIT License.

<!-- Copyright (c) 2026 Zeronetsec -->