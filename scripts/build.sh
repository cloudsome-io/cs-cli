#!/usr/bin/env bash

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
ROOT_DIR="${SCRIPT_DIR%/scripts}"
DIST_DIR="${ROOT_DIR}/dist"

BINARY_NAME="${BINARY_NAME:-cs}"
VERSION="${VERSION:-0.0.0}"

TARGETS=(
    "linux:amd64"
    "linux:arm64"
    "darwin:amd64"
    "darwin:arm64"
    "windows:amd64"
)

usage() {
    cat <<EOF
Usage: $(basename "$0") [all|<os>:<arch> ...]

Examples:
  $(basename "$0")
  $(basename "$0") all
  $(basename "$0") linux:amd64 windows:amd64
EOF
}

ensure_dist_dir() {
    mkdir -p "$DIST_DIR"
}

build_target() {
    local os="$1"
    local arch="$2"
    local suffix=""

    if [[ "$os" == "windows" ]]; then
        suffix=".exe"
    fi

    # Create a temporary directory for building
    local temp_dir=$(mktemp -d)
    local binary_name="${BINARY_NAME}${suffix}"
    local output="${temp_dir}/${binary_name}"

    echo "==> Building ${os}/${arch} -> ${output}"
    CGO_ENABLED=0 GOOS="$os" GOARCH="$arch" go build -o "$output" "$ROOT_DIR"

    # Create tgz archive with naming convention: cs-cli_<VERSION>_<OS>_<ARCH>.tgz
    local archive_name="${BINARY_NAME}-cli_${VERSION}_${os}_${arch}.tgz"
    local archive_path="${DIST_DIR}/${archive_name}"

    echo "==> Creating archive ${archive_path}"
    cd "$temp_dir"
    tar -czf "$archive_path" "$binary_name"
    cd - > /dev/null

    # Cleanup temp directory
    rm -rf "$temp_dir"

    echo "==> Created ${archive_path}"
}

main() {
    if [[ ${1:-} == "-h" || ${1:-} == "--help" ]]; then
        usage
        exit 0
    fi

    ensure_dist_dir

    local targets=("${TARGETS[@]}")

    if [[ $# -gt 0 && $1 != "all" ]]; then
        targets=("$@")
    fi

    for target in "${targets[@]}"; do
        if [[ "$target" == "all" ]]; then
            continue
        fi

        IFS=":" read -r os arch <<< "$target"
        if [[ -z "$os" || -z "$arch" ]]; then
            echo "Invalid target format: $target" >&2
            usage
            exit 1
        fi

        build_target "$os" "$arch"
    done
}

main "$@"

