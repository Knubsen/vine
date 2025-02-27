# Vine stuff

vine_dir="$HOME/.vine/"
vine_executable="$HOME/.vine/vine_aly"

function vine() {
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

# ffuts eniV
