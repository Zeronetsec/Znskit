#!/usr/bin/env bash

set -o errexit

src="${BASH_SOURCE[0]}"
while [[ -h "${src}" ]]; do
    dir="$(
        cd -P "$(
            command dirname "${src}"
        )" > /dev/null 2>&1 && pwd
    )"
    src="$(command readlink "${src}")"
    [[ "${src}" != /* ]] && src="${dir}/${src}"
done

dir="$(
    cd -P "$(
        command dirname "${src}"
    )" > /dev/null 2>&1 && pwd
)"

export root="${dir}"; readonly root
source "${root}/.install/include.sh"

include : '(
    .install/color
    .install/variable
    .install/error
    .install/getinstall
)'

__RMBK__=false
__NRMCFG__=false

while [[ ${#} -gt 0 ]]; do
    case "${1}" in
        "--remove-backup") export __RMBK__=true ;;
        "--no-remove-config") export __NRMCFG__=true ;;
    esac
    shift
done

if [[ "${__RMBK__}" == true ]]; then
    install::getinstall \
        "command rm -f ${opt}/znskit_*.zip.bak" \
        "Removing all backup..."
fi

install::getinstall \
    "command rm -rf ${opt}/znskit" \
    "Removing: ${GG}${opt}/znskit${N}"

install::getinstall \
    "command rm -f ${bin}/znskit" \
    "Removing: ${GG}${bin}/znskit${N}"

if [[ "${__NRMCFG__}" == false ]]; then
    install::getinstall \
        "command rm -rf ${HOME}/.znskit" \
        "Removing: ${GG}${HOME}/.znskit${N}"
fi

echo -e "${GG}[+] ${N}Znskit removed"

trap - EXIT
exit ${?}