eval "$(oh-my-posh init bash --config ~/.poshthemes/kali.omp.json)"



# Vine stuff

vine_path="C:/substE/projects/GoPaly/vine"
vine_dir="$HOME/.vine/"

function vine() {
    if [[ "$1" == "-d" ]]; then
        cd "$vine_path" || return 1
        output=$(go run main.go exec "$2")
        cd - > /dev/null
        eval "$output"
    elif [[ "$1" == "-e" ]]; then
        cd "$vine_path" || return 1
        output=$(go run main.go exec "$2")
        cd - > /dev/null
        cd "$vine_dir"
        "$output"
        cd - > /dev/null
    else
        echo "Usage: vine -d/-e <drive_letter/script>"
    fi
}

# ffuts eniV
