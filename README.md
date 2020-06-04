# WORDER

Small program that finds best match for a given word inside of a given dictionary file.
Dictionary file must be formatted as plaintext each word on a separate line.
(check [this awesome repo](https://github.com/dwyl/english-words) that contains english words)

Usage:
```
go build ./src/main.go
./main --dictionary="./words_alpha.txt" fieher
# Output be like:
# "Did you mean fisher?"
```