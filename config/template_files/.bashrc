# Vine stuff
vine_dir="$HOME/.vine"
vine_auth_token="$vine_dir/.auth/.vine_auth_token"
vine_secret_key="$vine_dir/.auth/.vine_secret_key"
vine_executable="$HOME/.vine/vine_aly"

function vine() {
    _generate_token

    trap 'rm "$vine_auth_token"' RETURN

    if [[ "$1" == "-d" ]]; then
        output=$("$vine_executable" exec "$2")
        _cd_back
        eval "$output"
    elif [[ "$1" == "-e" ]]; then
        output=$("$vine_executable" exec "$2")
        trap '_cd_back' INT
        cd "$vine_dir"
        "$output"
        _cd_back
        trap - INT
    else
        echo "Usage: vine -d/-e <drive_letter/script>"
    fi
}

function _cd_back() {
    if [[ -n "$OLDPWD" ]]; then
        cd - > /dev/null
    fi
}

function _generate_token() {
    SECRET_KEY=$(< "$vine_secret_key")
    TIME_STAMP=$(date -u +"%H:%M")
    TOKEN=$(printf "%s" "$TIME_STAMP" | openssl dgst -sha256 -hmac "$SECRET_KEY" | awk '{print $2}')
    echo "$TOKEN" > "$vine_auth_token"
}
# ffuts eniV
