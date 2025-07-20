### Usage

Currently supports 2 modes: directory fuzzing and subdomain enumeration
They can be accessed with commands `dir` and `dns`

Available flags:
```
    -t, --threads int       Amount of threads used for fuzzing (default 10)
    -u, --url string        Url to fuzz
    -w, --wordlist string   Path to wordlist
```

Example:
```
	gofuzz dns -u google.com -w ./wordlist.txt -t 20
```

For directory fuzzing request timeout can be configured via flag `--timeout`
