<!-- https://github.com/Zeronetsec/Znskit -->

# Installation
`install.sh` optional option:
- `--backup`

Use `--backup` to create a backup of the existing Znskit installation before replacing it.

## Termux & Linux (root)
```bash
git clone https://github.com/Zeronetsec/Znskit
cd Znskit
chmod +x install.sh
./install.sh
```

## Linux (user)
```bash
git clone https://github.com/Zeronetsec/Znskit
cd Znskit
chmod +x install.sh
sudo ./install.sh
```

## Uninstallation
```bash
export prefix="${PREFIX:-/usr}"
rm -f "${prefix}/bin/znskit"
rm -rf "${prefix}/opt/znskit"
rm -rf ~/.config/znskit
```

<!-- Copyright (c) 2026 Zeronetsec -->