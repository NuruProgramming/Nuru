#!/usr/bin/env sh

# Hii ni skripti ya shell ili kusakinisha programu ya Nuru.
# Programu zinazohitajika:
#   - curl/wget: Kupakua faili za 'tar' kutoka 'Github'
#   - cp: Nakili faili kuenda mahali sahihi
#   - jq: Kupata uhusiano kwenye fomati ya 'JSON'
#   - tar: Kufungua faili za tar.gz

set -e

ARCH="$(uname -m)"
OSNAME="$(uname -s)"
PREFIX_PATH="/usr"
BIN=""
VERSION="latest"
RELEASE_URL="https://github.com/NuruProgramming/Nuru/releases"
TEMP=""

# Cleanup function to remove temp directory on exit
cleanup() {
    if [ -n "$TEMP" ] && [ -d "$TEMP" ]; then
        rm -rf "$TEMP"
    fi
}
trap cleanup EXIT

# Check if command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Print usage information
usage() {
    echo "Usage: $0 [OPTIONS]"
    echo ""
    echo "Options:"
    echo "  -p, --prefix  The base path to be used when installing (default: /usr)"
    echo "  -v, --version The version to be downloaded from GitHub (default: latest)"
    echo "  -h, --help    Show this help message"
    echo ""
}

# Normalize architecture names
arch_name() {
    case "$ARCH" in
        x86_64)
            ARCH="amd64"
            ;;
        i386|i686)
            ARCH="i386"
            ;;
        arm64|aarch64)
            ARCH="arm64"
            ;;
        *)
            echo "Unsupported architecture: $ARCH"
            exit 2
            ;;
    esac
}

# Validate OS name
os_name() {
    case "$OSNAME" in
        Darwin|Linux|Android)
            ;;
        *)
            echo "Unsupported Operating System: $OSNAME"
            exit 2
            ;;
    esac
}

# Parse command line arguments
parse_args() {
    while [ "$#" -gt 0 ]; do
        case "$1" in
            -h|--help)
                usage
                exit 0
                ;;
            -p|--prefix)
                shift
                if [ -z "$1" ]; then
                    echo "Error: Missing argument for --prefix"
                    exit 1
                fi
                PREFIX_PATH="$1"
                ;;
            -v|--version)
                shift
                if [ -z "$1" ]; then
                    echo "Error: Missing argument for --version"
                    exit 1
                fi
                VERSION="$1"
                ;;
            --)
                shift
                break
                ;;
            *)
                echo "Unknown argument: $1"
                usage
                exit 1
                ;;
        esac
        shift
    done
    BIN="$PREFIX_PATH/bin"
}

# Download file using curl or wget
download() {
    URL="$1"
    if command_exists curl; then
        curl -fSL "$URL"
    elif command_exists wget; then
        wget -qO- "$URL"
    else
        echo "Error: Neither curl nor wget is installed."
        exit 1
    fi
}

main() {
    os_name
    arch_name
    parse_args "$@"

    # Check required commands
    for cmd in jq tar cp; do
        if ! command_exists "$cmd"; then
            echo "Error: Required command '$cmd' not found."
            exit 1
        fi
    done

    if [ "$VERSION" = "latest" ]; then
        echo "Fetching latest version tag from GitHub..."
        VERSION="$(download "https://api.github.com/repos/NuruProgramming/Nuru/releases/latest" | jq -r .tag_name)"
        if [ -z "$VERSION" ] || [ "$VERSION" = "null" ]; then
            echo "Error: Unable to determine latest version."
            exit 1
        fi
    fi

    TAR_URL="$RELEASE_URL/download/$VERSION/nuru_${OSNAME}_${ARCH}.tar.gz"
    echo "Downloading Nuru version $VERSION for $OSNAME/$ARCH..."

    TEMP="$(mktemp -d)"
    if ! download "$TAR_URL" | tar -xz -C "$TEMP"; then
        echo "Error: Failed to download or extract archive."
        exit 1
    fi

    # Ensure bin directory exists
    if [ ! -d "$BIN" ]; then
        echo "Creating directory $BIN"
        mkdir -p "$BIN"
    fi

    echo "Installing Nuru to $BIN/nuru"
    if ! cp "$TEMP/nuru" "$BIN/"; then
        echo "Error: Failed to copy binary to $BIN"
        exit 1
    fi

    echo "Installation complete."
}

main "$@"
