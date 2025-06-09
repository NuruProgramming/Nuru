package analysis

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/NuruProgramming/Nuru/ast"
	"github.com/NuruProgramming/Nuru/lexer"
	"github.com/NuruProgramming/Nuru/parser"
)

// ModuleDependencyGraph represents a directed graph of module dependencies
type ModuleDependencyGraph struct {
	Modules       map[string]*ModuleNode
	AnalyzedFiles map[string]bool // Tracks which files have been analyzed
}

// ModuleNode represents a module in the dependency graph
type ModuleNode struct {
	Name         string
	Dependencies []*ModuleNode
	Visited      bool // Used during cycle detection
	InStack      bool // Used for detecting cycles in current DFS path
	FilePath     string
	ImportLines  map[string]int // Maps imported module to line number
}

// Warning represents a static analysis warning
type Warning struct {
	Message  string
	Line     int
	Column   int
	Filename string
	Type     WarningType
}

type WarningType string

const (
	CircularImportWarning       WarningType = "UINGIZI_WA_MDUARA"
	CircularTypeDefWarning      WarningType = "UINGIZI_WA_KITAMBULISHO"
	PotentialCircularRefWarning WarningType = "UINGIZI_WA_KUMBUKUMBU"
)

// NewModuleDependencyGraph creates a new empty dependency graph
func NewModuleDependencyGraph() *ModuleDependencyGraph {
	return &ModuleDependencyGraph{
		Modules:       make(map[string]*ModuleNode),
		AnalyzedFiles: make(map[string]bool),
	}
}

// AddModule adds a module to the dependency graph if it doesn't exist
func (g *ModuleDependencyGraph) AddModule(name string, filePath string) *ModuleNode {
	if node, exists := g.Modules[name]; exists {
		return node
	}

	node := &ModuleNode{
		Name:         name,
		Dependencies: []*ModuleNode{},
		Visited:      false,
		InStack:      false,
		FilePath:     filePath,
		ImportLines:  make(map[string]int),
	}

	g.Modules[name] = node
	return node
}

// AddDependency adds a dependency between two modules
func (g *ModuleDependencyGraph) AddDependency(from, to string, line int) {
	fromNode := g.Modules[from]
	toNode := g.Modules[to]

	// Check if the dependency already exists
	for _, dep := range fromNode.Dependencies {
		if dep.Name == to {
			return // Dependency already exists
		}
	}

	// Add the dependency
	fromNode.Dependencies = append(fromNode.Dependencies, toNode)
	fromNode.ImportLines[to] = line
}

// DetectCycles detects cycles in the dependency graph
func (g *ModuleDependencyGraph) DetectCycles() []Warning {
	var warnings []Warning

	// Reset visited and inStack flags
	for _, node := range g.Modules {
		node.Visited = false
		node.InStack = false
	}

	// Check each module for cycles
	for _, node := range g.Modules {
		if !node.Visited {
			path := []string{}
			g.detectCyclesHelper(node, path, &warnings)
		}
	}

	return warnings
}

// detectCyclesHelper is a helper function for detecting cycles using DFS
func (g *ModuleDependencyGraph) detectCyclesHelper(node *ModuleNode, path []string, warnings *[]Warning) {
	node.Visited = true
	node.InStack = true

	path = append(path, node.Name)

	for _, dep := range node.Dependencies {
		if !dep.Visited {
			g.detectCyclesHelper(dep, path, warnings)
		} else if dep.InStack {
			// Found a cycle
			cycleStart := -1
			for i, moduleName := range path {
				if moduleName == dep.Name {
					cycleStart = i
					break
				}
			}

			if cycleStart != -1 {
				cycle := append(path[cycleStart:], dep.Name)
				// Create a warning
				*warnings = append(*warnings, Warning{
					Message:  fmt.Sprintf("Utegemezi wa moduli wa mduara umegunduliwa: %s", strings.Join(cycle, " → ")),
					Line:     node.ImportLines[dep.Name],
					Column:   1,
					Filename: node.FilePath,
					Type:     CircularImportWarning,
				})
			}
		}
	}

	node.InStack = false
}

// AnalyzeFile analyzes a file for module dependencies
func (g *ModuleDependencyGraph) AnalyzeFile(filePath string, fileContent string) error {
	// Check if we've already analyzed this file
	if g.AnalyzedFiles[filePath] {
		return nil
	}

	// Mark as analyzed
	g.AnalyzedFiles[filePath] = true

	// Get the module name from the file path
	moduleName := filepath.Base(filePath)
	if strings.HasSuffix(moduleName, ".nr") {
		moduleName = moduleName[:len(moduleName)-3]
	}

	// Add the module to the graph
	g.AddModule(moduleName, filePath)

	// Parse the file
	l := lexer.New(fileContent)
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		return fmt.Errorf("kuchanganua makosa ndani %s: %v", filePath, p.Errors())
	}

	// Find import statements
	for _, stmt := range program.Statements {
		// Import is typically an expression, so check for ExpressionStatement first
		if exprStmt, ok := stmt.(*ast.ExpressionStatement); ok {
			// Check if the expression inside is an Import
			if importExpr, ok := exprStmt.Expression.(*ast.Import); ok {
				for importAs, importIdent := range importExpr.Identifiers {
					importedModule := importIdent.Value
					importLine := importIdent.Token.Line

					// Add the imported module to the graph
					g.AddModule(importedModule, "") // We may not know the actual file path yet

					// Add the dependency
					g.AddDependency(moduleName, importedModule, importLine)

					// Also record the import under the "as" name if different
					if importAs != importedModule {
						g.AddModule(importAs, "")
						g.AddDependency(moduleName, importAs, importLine)
					}
				}
			}
		}
	}

	return nil
}

// BuildDependencyGraph builds a dependency graph from a set of files
func BuildDependencyGraph(files map[string]string) (*ModuleDependencyGraph, []Warning, error) {
	graph := NewModuleDependencyGraph()

	// First pass: analyze all files
	for filePath, content := range files {
		if err := graph.AnalyzeFile(filePath, content); err != nil {
			return nil, nil, err
		}
	}

	// Detect cycles
	warnings := graph.DetectCycles()

	return graph, warnings, nil
}

// GenerateDOTGraph generates a DOT format graph for visualization
func (g *ModuleDependencyGraph) GenerateDOTGraph() string {
	var dot strings.Builder

	dot.WriteString("digraph ModuleDependencies {\n")
	dot.WriteString("  node [shape=box];\n")

	// Add nodes
	for name := range g.Modules {
		dot.WriteString(fmt.Sprintf("  \"%s\";\n", name))
	}

	// Add edges
	for name, node := range g.Modules {
		for _, dep := range node.Dependencies {
			dot.WriteString(fmt.Sprintf("  \"%s\" -> \"%s\";\n", name, dep.Name))
		}
	}

	dot.WriteString("}\n")

	return dot.String()
}
