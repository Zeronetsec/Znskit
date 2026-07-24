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

    if [[ ! -d "${HOME}/.znskit" ]]; then
        install::getinstall \
            "command mkdir -p ${HOME}/.znskit" \
            "Create directory: ${GG}${HOME}/.znskit${N}"
    fi

    if [[ ! -f "${HOME}/.znskit/packages.lst" ]]; then
        install::getinstall \
            "
                command cat \
                    ${opt}/znskit/config/packages.lst \
                    > ${HOME}/.znskit/packages.lst
            " \
            "Create file: ${GG}${HOME}/.znskit/packages.lst${N}"
    fi

    (
        cd "${opt}/znskit"
        install::getinstall \
            "command go mod tidy" \
            "Retidy: ${GG}znskit${N}"

        install::getinstall \
            "command go build -o znskit" \
            "Compiling: ${GG}znskit${N}"
        cd
    )

    install::getinstall \
        "
            command ln -sf \
                ${opt}/znskit/znskit \
                ${bin}/znskit
        " \
        "Symlink: ${GG}${opt}/znskit/znskit ${DG}-> ${GG}${bin}/znskit${N}"
}; readonly -f install::installer