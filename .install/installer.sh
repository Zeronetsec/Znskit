function install::installer() {
    if [[ ! -d "${HOME}/.${targetins}" ]]; then
        install::getinstall \
            "command mkdir -p ${HOME}/.${targetins}" \
            "Create directory: ${color_GG}${HOME}/.${targetins}${color_N}"
    fi

    if [[ ! -f "${HOME}/.${targetins}/packages.lst" ]]; then
        install::getinstall \
            "
                command cat \
                    ${opt}/${targetins}/config/packages.lst \
                    > ${HOME}/.${targetins}/packages.lst
            " \
            "Create file: ${color_GG}${HOME}/.${targetins}/packages.lst${color_N}"
    fi

    (
        cd "${opt}/${targetins}"
        install::getinstall \
            "
                command go mod tidy
                command go build -o ${targetins}
            " \
            "Compiling: ${color_GG}${targetins}${color_N}"
    )
}; readonly -f install::installer