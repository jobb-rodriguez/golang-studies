# Go Modules
- The module path
- The version of the Go language your project requires
- Optionally, any external package dependencies your project has

```
module github.com/bootdotdev/exampleproject

go 1.20

require github.com/google/examplepackage v1.3.0
```

**Creating Go Modules**

```bash
go mod init {REMOTE}/{USERNAME}/{module_name}

# {REMOTE} (i.e. github.com)
# {USERNAME} (Git username)
```

Use remote to simplify downloading of packages.

By convention, name the package the same as the directory.

# Package Best Practices
[Article](https://blog.boot.dev/golang/how-to-separate-library-packages-in-go/)