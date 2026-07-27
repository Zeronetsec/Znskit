function install::installer() {
    if [[ ! -d "${HOME}/.${targetins}" ]]; then
        install::getinstall \
            "command mkdir -p ${HOME}/.${targetins}" \
            "Create directory: ${GG}${HOME}/.${targetins}${N}"
    fi

    if [[ ! -f "${HOME}/.${targetins}/packages.lst" ]]; then
        install::getinstall \
            "
                command cat \
                    ${opt}/${targetins}/config/packages.lst \
                    > ${HOME}/.${targetins}/packages.lst
            " \
            "Create file: ${GG}${HOME}/.${targetins}/packages.lst${N}"
    fi

    (
        cd "${opt}/${targetins}"
        install::getinstall \
            "command go mod tidy" \
            "Retidy: ${GG}${targetins}${N}"

        install::getinstall \
            "command go build -o ${targetins}" \
            "Compiling: ${GG}${targetins}${N}"
        cd
    )
}; readonly -f install::installer