function install::installer() {
    if [[ "${__BACKUP__}" == true && -d "${opt}/znskit" ]]; then
        (
            cd "${opt}"
            install::getinstall \
                "
                    command zip -r \
                    znskit_${bkdate}.bak.zip \
                    znskit
                " \
                "Backup: ${GG}${opt}/znskit ${DG}-> ${GG}${opt}/znskit_${bkdate}.bak.zip${N}"
            cd
        )
    fi

    if [[ -d "${opt}/znskit" ]]; then
        install::getinstall \
            "command rm -rf ${opt}/znskit" \
            "Removing old source..."
    fi

    install::getinstall \
        "command mv ${root} ${opt}/znskit" \
        "Moving: ${GG}${root} ${DG}-> ${GG}${opt}/znskit${N}"

    (
        cd "${opt}/znskit"
        install::getinstall \
            "command go mod tidy" \
            "Retidy: ${GG}znskit${N}"

        install::getinstall \
            "command go build -v -o znskit" \
            "Building: ${GG}znskit${N}"
        cd
    )

    install::getinstall \
        "
            command ln -sf \
            ${opt}/znskit/znskit \
            ${bin}/znskit
        " \
        "Symlink: ${GG}${opt}/znskit/znskit ${DG}-> ${GG}${bin}/znskit${N}"
}