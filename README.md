## DustiLock

A scanner for project's package dependencies, checks if one or more is available for public registration.

This is to mitigate **vulnerable package names** that can be used in open-source supply-chain attacks as demonstrated in [this research](https://medium.com/@alex.birsan/dependency-confusion-4a5d60fec610) by Alex Birsan.

### Usage

```
go build
./dusti-lock
```

Arguments

- `-r` - recursive scan (default is false)
- `-p <path>` - custom path to scan (default is current working dir)
- `-a` - audit only, will not fail for detections (default is false)

**Example**

```
./dusti-lock -p /tmp/code -r

DustiLock started
scanning directory "/tmp/code" (recursive=true) ...
error - npm package "private-org-infra" is available for public registration. /tmp/code/test-project/package.json
one or more packages are available for public registration
```

### Using in GitHub Actions

Add the following to your workflow file as a step:
```
- name: DustiLock
  uses: dustico/dusti-lock@v1.0.0
```

If needed, you can customize it like so:
```
- name: DustiLock
  uses: dustico/dusti-lock@v1.0.0
  with:
    recursive: true
    path: my-nested-project
    # disable failing the build when having alerts (default = false)
    audit: true
```

### How to Deal With Vulnerable Packages?
To mitigate the risk, you need to register a dummy placeholder package with the same names as your internal packages (if any), to prevent such an attack. When registered, do use a low version number (e.g. 0.0.1), so it won't be used instead of your internal package. 

If you need any assistance, you're welcome to contact us at - research@dusti.co

### Languages Support
At this moment, this tool can check:

- `requirements.txt` - Python
- `package.json` - npm 

Stay tuned and if possible please submit a PR to add more languages support

### Credits

- Alex Birsan for sharing [his research](https://medium.com/@alex.birsan/dependency-confusion-4a5d60fec610) and helping to secure the open-source supply-chain.
- https://github.com/davidfischer/requirements-parser for examples on python requirements.txt dependencies
