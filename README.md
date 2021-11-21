![readme cover image](https://user-images.githubusercontent.com/1287098/142776854-83abf265-a1ba-485f-a8b6-995da7f7ef8b.png)


DustiLock is a tool to find which of your dependencies is susceptible to Dependency Confusion attack.

## What is Dependency Confusion?

A technique discovered by [@alex.birsan](https://medium.com/@alex.birsan/dependency-confusion-4a5d60fec610) to hijack a privately used package by registering its name (if available) on a public registry with a higher version number. This may cause artifact servers and build tools to "confuse" and use the attacker's package.


![Frame 237 (1)](https://user-images.githubusercontent.com/1287098/142776859-7c6c3ef6-6a15-4e34-99f6-b4bac029a036.png)


## CLI

### Build
```
go build
./dustilock
```

### CLI Arguments

- `-r` - Recursive scan (default=false)
- `-p <path>` - Custom path to scan (default=current working dir)
- `-a` - Audit only, will not fail for detections (default=false)

**Example**

```
./dustilock -p /tmp/code -r

DustiLock started
scanning directory "/tmp/code" (recursive=true) ...
error - npm package "private-org-infra" is available for public registration. /tmp/code/test-project/package.json
one or more packages are available for public registration
```

## Using in CI Workflows
This tool can be easily integrated into modern CI workflows to test new code contributions.  

### Using in GitHub Actions

Add the following to your workflow file as a step:
```
- name: DustiLock
  uses: checkmarx/dustilock@v1.0.0
```

Override defaults:
```
- name: DustiLock
  uses: checkmarx/dustilock@v1.0.0
  with:
    recursive: true
    path: my-nested-project
    # disable failing the build when having alerts (default = false)
    audit: true
```

### How to Deal With Vulnerable Packages?
To mitigate the risk, you need to register a dummy placeholder package with the same names as your internal packages (if any), to prevent such an attack. When registered, do use a low version number (e.g. 0.0.1), so it won't be used instead of your internal package. 


### Languages Support
At this moment, this tool can check:

- `requirements.txt` - Python
- `package.json` - npm 


### Credits

- Alex Birsan for sharing [his research](https://medium.com/@alex.birsan/dependency-confusion-4a5d60fec610) and helping to secure the open-source supply-chain.
- https://github.com/davidfischer/requirements-parser for examples on python requirements.txt dependencies


# License

Apache-2.0 License

Copyright (c) 2021 Checkmarx
