eval "$(oh-my-posh init bash --config ~/.poshthemes/kali.omp.json)"



# Vine stuff

vine_dir="$HOME/.vine/"
vine_executable="$HOME/.vine/vine_aly"

function vine() {
    if [[ "$1" == "-d" ]]; then
        cd "$vine_dir" || return 1
        output=$("$vine_executable" exec "$2")
        cd - > /dev/null
        eval "$output"
    elif [[ "$1" == "-e" ]]; then
        cd "$vine_dir" || return 1
        output=$(vine_aly exec "$2")
        cd - > /dev/null
        cd "$vine_dir"
        "$output"
        cd - > /dev/null
    else
        echo "Usage: vine -d/-e <drive_letter/script>"
    fi
}

# ffuts eniV
