

# 001 Better folder structure

Status Date: 2025-02-13 12:38
Driver: @aHolbreich
Contributors: ...
Status: Accepted

** This ADR is more of demonstration rather than documentation. Don't take it too serious** 

# Context 
The current project structure seem not to support me well for me.
The code seem to not support good extensibility and seem not to adhere to best golang practices. 

# Decision
I'm going to choose option 1 understanding that refactoring would cost some time. However this is great enable for future improvements

### Consequences
Better adherence to the Go dev practices. Better testability, better cohesion inside of packages via better "layering"
No significant downsizes.
## Options considered
### 1. Internal package for all the logic
```
.
├── cmd/                    // CLI Commands
│   ├── root.go
│   └── init.go
├── internal/
│   ├── config/             // Configuration and path resolution
│   │   └── path.go
│   │   └── manager.go
│   ├── template/           // Template handling logic
│   │   └── manager.go
│   └── model/
│   │   └── model.go        // Model
│   └── adr/                // ADR business logic
│       ├── manager.go
│       ├── sorter.go
│       └── model.go

└── templates/              // Static template files
    └── default.md

```

### 2. Not touching it
    
Since the future development is not clear. We can have it untouched.

## Advices

* ChatGPT find likes the idea internal/config: Handles all configuration-related tasks, including path resolution and config loading.
    internal/template: Centralizes template loading, parsing, and rendering logic.
    internal/adr: Manages ADR-specific business logic, like creating, listing, and updating ADRs.
    This structure keeps cmd/ focused on CLI commands while separating the business logic into reusable components.
