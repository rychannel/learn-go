# Go Learning Repository - Copilot Instructions

## Project Overview

This is a **GOPATH-based Go learning repository** containing independent example programs demonstrating Go language fundamentals. Each subdirectory in `src/` is a self-contained example focusing on a specific concept.

**Critical**: This project uses the **legacy GOPATH workspace structure**, NOT Go modules. The workspace root (`/workspaces/learn-go`) IS the GOPATH.

## Architecture & Structure

```
/workspaces/learn-go/          # GOPATH root
├── src/                       # All source code (GOPATH convention)
│   ├── hello/                 # Independent example programs
│   ├── funcs/                 # Each directory = one concept
│   ├── cube/                  # Reusable package (no main)
│   ├── verify/                # Reusable package (no main)
│   ├── prog/                  # Example importing local packages
│   └── ...                    # 30+ examples total
├── pkg/linux_amd64/           # Compiled package objects
└── bin/                       # Compiled binaries (when using GOPATH/bin)
```

### Package Organization Patterns

1. **Executable examples**: Directories with `main.go` containing `package main`
   - Examples: [src/hello/main.go](../src/hello/main.go), [src/funcs/main.go](../src/funcs/main.go), [src/method/main.go](../src/method/main.go)
   - Compiled binaries remain in their source directories (e.g., `src/hello/hello`)

2. **Reusable packages**: Directories with non-main packages
   - [src/cube/cube.go](../src/cube/cube.go) - demonstrates encapsulation with exported/unexported methods
   - [src/verify/verify.go](../src/verify/verify.go) - demonstrates error handling patterns
   - Import path is just the package name (e.g., `"cube"`, `"verify"`) due to GOPATH

## Critical Workflows

### Building & Running Programs

**GOPATH-style compilation** (not `go build` in module mode):

```bash
# Navigate to example directory first
cd src/hello
go build          # Creates binary in current directory (./hello)
./hello           # Run the compiled binary

# Or run directly without creating persistent binary
go run main.go
```

**For programs importing local packages** (e.g., [src/prog/main.go](../src/prog/main.go)):

```bash
cd src/prog
go run main.go    # GOPATH automatically resolves "verify" package
```

### Creating New Examples

Follow the established pattern:

1. Create new directory under `src/` with descriptive name (e.g., `src/maps/`)
2. Create `main.go` with `package main` for executables
3. Binaries compile to their source directory (gitignored by pattern in [.gitignore](../.gitignore))

## Code Patterns & Conventions

### Struct Methods & OOP Patterns

This codebase teaches Go OOP through composition:

- **Pointer receivers** for methods that modify state: `func (d *Dims) SetSize(w, l, h int)`
  - See [src/cube/cube.go](../src/cube/cube.go) for encapsulation patterns
  
- **Value receivers** for read-only methods: `func (c car) accelerate() string`
  - See [src/method/main.go](../src/method/main.go) for basic method syntax

- **Composition over inheritance**: Embedded structs for "inheritance"
  - See [src/inher/main.go](../src/inher/main.go) - article embeds member, gains `fullName()` method
  - See [src/embed/main.go](../src/embed/main.go) - circle embeds coords, accesses x/y directly

### Encapsulation Patterns

Go uses capitalization for visibility (GOPATH packages demonstrate this clearly):

- **Exported** (public): Start with uppercase - `GetVolume()`, `SetSize()`, `Dims`
- **Unexported** (private): Start with lowercase - `area()` (internal to package)
- Example in [src/cube/cube.go](../src/cube/cube.go):
  - `GetArea()` is public wrapper around private `area()` method
  - Note: [src/encap/main.go](../src/encap/main.go) shows this pattern (has compilation error accessing private members)

### Error Handling Pattern

See [src/verify/verify.go](../src/verify/verify.go) for idiomatic error returns:
```go
func IsPosInt(num int) (int, error) {
    if num < 1 {
        err := fmt.Errorf("%v not a positive number", num)
        return -1, err
    }
    return num, nil
}
```

Consumer code in [src/prog/main.go](../src/prog/main.go) checks `if err != nil` immediately after call.

## Environment & Tooling

- **GOPATH**: Set to `/workspaces/learn-go` (workspace root)
- **Git config**: Copy [.env.local.example](../.env.local.example) to `.env.local` for user details
- **Compiled artifacts**: Executables in `src/*/` directories are gitignored (see [.gitignore](../.gitignore) lines 19-23)

## What NOT to Do

- ❌ Don't create `go.mod` files - this is a GOPATH project, not modules
- ❌ Don't use `go install` expecting binaries in `/workspaces/learn-go/bin` - they build locally
- ❌ Don't import with full paths like `"github.com/user/learn-go/src/cube"` - use just `"cube"`
- ❌ Don't suggest modern Go module practices - preserve the learning context
