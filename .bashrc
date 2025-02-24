eval "$(oh-my-posh init bash --config ~/.poshthemes/kali.omp.json)"

function vine() {
    if [[ "$1" == "-d" ]]; then
        case "$2" in
            e) cd /c/substE ;;
            d) cd /c/substD ;; # Add more if needed
            *) echo "Unknown drive: $(go run C:/substE/projects/GoPaly/vine exec inputtoecho)" ;;
        esac
    elif [[ "$1" == "-e" ]]; then
        case "$2" in
            vm) source ~/connectVM.sh;;
            *) echo "ne" ;;
        esac
    else
        echo "Usage: vine -d/-e <drive_letter/script>"
    fi
}

