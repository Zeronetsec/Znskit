<!-- https://github.com/Zeronetsec/Znskit -->

# Installation
`install.sh` optional option:
- `--backup`
- └── create a backup of the existing Znskit installation before replacing it.

### Usage
```bash
git clone https://github.com/Zeronetsec/Znskit
bash Znskit/install.sh <option>
```

# Uninstallation
`uninstall.sh` optional option:
- `--remove-backup`
- └── remove all backup found.

### Usage
```bash
export prefix="${PREFIX:-/usr}"
bash $prefix/opt/znskit/uninstall.sh <option>
```

<!-- Copyright (c) 2026 Zeronetsec -->