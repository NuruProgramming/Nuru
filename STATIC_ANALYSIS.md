# Static Analysis for Circular Dependency Detection in Nuru

This document outlines the design and implementation of static analysis tools for detecting circular dependencies in Nuru code before runtime execution.

## Overview

Static analysis allows us to identify potential circular references during the compilation/interpretation phase rather than at runtime. This helps developers catch issues early and prevent memory leaks or unexpected behavior.

## Types of Circular Dependencies to Detect

1. **Module Import Cycles** - When modules import each other directly or indirectly
2. **Type Definition Cycles** - When types reference each other in their definitions
3. **Data Structure Cycles** - When data structures (dictionaries, arrays) contain circular references
4. **Object Property Cycles** - When object properties create reference cycles

## Implementation Approach

### 1. Module Dependency Analyzer

Create a system to track and analyze module import relationships:

```go
// Represents a directed graph of module dependencies
type ModuleDependencyGraph struct {
    modules map[string]*ModuleNode
}

type ModuleNode struct {
    name         string
    dependencies []*ModuleNode
    visited      bool  // Used during cycle detection
    inStack      bool  // Used for detecting cycles in current DFS path
}

// Checks if there are cycles in the dependency graph
func (g *ModuleDependencyGraph) DetectCycles() [][]string {
    var cycles [][]string
    
    // For each module, perform DFS to detect cycles
    for _, module := range g.modules {
        path := []string{}
        g.detectCyclesHelper(module, path, &cycles)
    }
    
    return cycles
}
```

### 2. Type Definition Analyzer

Analyze type definitions for circular dependencies:

```go
// Track relationships between type definitions
type TypeDependencyTracker struct {
    typeDependencies map[string][]string  // Maps type name to the types it depends on
}

// Detect type definition cycles
func (t *TypeDependencyTracker) DetectCycles() [][]string {
    // Similar DFS approach as module dependency analyzer
}
```

### 3. Data Structure Analyzer

Analyze data structure initializations and assignments for potential circular references:

```go
// Analyze variable assignments to detect potential circular references
func AnalyzeAssignments(program *ast.Program) []CircularReferenceWarning {
    var warnings []CircularReferenceWarning
    
    // Track variable relationships
    varRelationships := map[string][]string{}
    
    // Analyze variable assignments and detect cycles
    // ...
    
    return warnings
}
```

## Integration with Parser

The static analysis will be integrated with the Nuru parser:

```go
func Parse(input string, filename string) (*ast.Program, []Error, []Warning) {
    lexer := lexer.New(input)
    parser := parser.New(lexer)
    
    program := parser.ParseProgram()
    
    // Run static analysis
    warnings := RunStaticAnalysis(program, filename)
    
    return program, parser.Errors(), warnings
}
```

## Warning System

Implement a warning system to report findings without blocking execution:

```go
type Warning struct {
    Message  string
    Line     int
    Column   int
    Filename string
    Type     WarningType
}

type WarningType string

const (
    CircularImportWarning       WarningType = "CIRCULAR_IMPORT"
    CircularTypeDefWarning      WarningType = "CIRCULAR_TYPE_DEF"
    PotentialCircularRefWarning WarningType = "POTENTIAL_CIRCULAR_REF"
)
```

## Command Line Interface

Add flags to control static analysis behavior:

```
nuru --analyze <file.nr>          # Run static analysis only without execution
nuru --warning-level=strict       # Treat warnings as errors
nuru --disable-warnings           # Disable static analysis warnings
nuru --report-format=json         # Output warnings in JSON format
```

## Example Usage

Below is an example of how warnings would be displayed for circular dependencies:

```
$ nuru --analyze program.nr

⚠️  WARNING: Circular module dependency detected
   → moduleA imports moduleB imports moduleA
   → File: moduleA.nr, Line: 3, Column: 1

⚠️  WARNING: Potential circular reference in data structure
   → obj1["refToObj2"] = obj2, obj2["refToObj1"] = obj1
   → Consider using weak references
   → File: program.nr, Line: 42, Column: 5
```

## Implementation Plan

1. **Phase 1**: Implement basic module import cycle detection
2. **Phase 2**: Add type definition cycle detection
3. **Phase 3**: Implement data structure and assignment analysis
4. **Phase 4**: Add warning system and CLI integration
5. **Phase 5**: Add visualization tools for dependency graphs

## Visualization Tools

For complex dependency structures, add visualization tools to help developers understand relationships:

```
$ nuru --visualize-deps program.nr --format=dot > deps.dot
$ dot -Tpng deps.dot -o deps.png
```

## Benefits

- Early detection of potential memory leaks
- Improved code quality and maintainability
- Better developer experience with clear warnings
- Visualization tools for understanding complex relationships

## Future Extensions

1. **IDE Integration**: Provide integration with popular IDEs for inline warnings
2. **Automated Fixes**: Suggest or automatically apply fixes for common circular dependency patterns
3. **Performance Analysis**: Extend static analysis to identify performance bottlenecks
4. **Custom Rules**: Allow developers to define custom static analysis rules 