coverage=0
total=0
while IFS='' read -r line || [[ -n "$line" ]]; do
    IFS=' ' read -r -a array <<< "$line"
    total=$(($total+${array[1]}))
    if [[ "${array[2]}" -gt 0 ]]; then
        covered=$(($covered+${array[1]}))
    fi
done < "$1"
awk "BEGIN {printf \"coverage: %.1f%s of statements\", 100*${covered}/${total}, \"%\"}"