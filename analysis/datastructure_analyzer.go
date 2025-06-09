package analysis

import (
	"fmt"

	"github.com/NuruProgramming/Nuru/ast"
)

// VariableAssignment tracks variable assignments for potential circular references
type VariableAssignment struct {
	Name       string
	Value      ast.Expression
	Line       int
	Column     int
	Filename   string
	References []string // Variables this assignment directly references
}

// DataStructureAnalyzer analyzes data structures for potential circular references
type DataStructureAnalyzer struct {
	Variables       map[string]*VariableAssignment
	AssignmentOrder []string // Keeps track of assignment order
	Warnings        []Warning
	CurrentFilename string
}

// NewDataStructureAnalyzer creates a new instance of DataStructureAnalyzer
func NewDataStructureAnalyzer() *DataStructureAnalyzer {
	return &DataStructureAnalyzer{
		Variables:       make(map[string]*VariableAssignment),
		AssignmentOrder: []string{},
		Warnings:        []Warning{},
	}
}

// AnalyzeFile analyzes a file for potential circular references in data structures
func (a *DataStructureAnalyzer) AnalyzeFile(filename string, program *ast.Program) {
	a.CurrentFilename = filename

	// First pass: Collect all variable assignments
	for _, stmt := range program.Statements {
		a.analyzeStatement(stmt)
	}

	// Second pass: Check for circular references
	a.detectCircularReferences()
}

// analyzeStatement analyzes a statement for variable assignments
func (a *DataStructureAnalyzer) analyzeStatement(stmt ast.Statement) {
	switch s := stmt.(type) {
	case *ast.LetStatement:
		a.analyzeLetStatement(s)
	case *ast.ExpressionStatement:
		a.analyzeExpressionStatement(s)
	case *ast.BlockStatement:
		for _, blockStmt := range s.Statements {
			a.analyzeStatement(blockStmt)
		}
	}
}

// analyzeLetStatement analyzes a let statement for variable assignments
func (a *DataStructureAnalyzer) analyzeLetStatement(stmt *ast.LetStatement) {
	varName := stmt.Name.Value
	a.trackVariableAssignment(varName, stmt.Value, stmt.Token.Line, 1)
}

// analyzeExpressionStatement analyzes an expression statement for assignment expressions
func (a *DataStructureAnalyzer) analyzeExpressionStatement(stmt *ast.ExpressionStatement) {
	if assignExpr, ok := stmt.Expression.(*ast.AssignmentExpression); ok {
		if ident, ok := assignExpr.Left.(*ast.Identifier); ok {
			a.trackVariableAssignment(ident.Value, assignExpr.Value, assignExpr.Token.Line, 1)
		} else if indexExpr, ok := assignExpr.Left.(*ast.IndexExpression); ok {
			// Handle assignments to array or dictionary elements
			a.analyzeIndexAssignment(indexExpr, assignExpr.Value, assignExpr.Token.Line)
		}
	}
}

// analyzeIndexAssignment analyzes assignments to array or dictionary elements
func (a *DataStructureAnalyzer) analyzeIndexAssignment(indexExpr *ast.IndexExpression, value ast.Expression, line int) {
	// If we're assigning to an index of an object, check if we might be creating a circular reference
	if ident, ok := indexExpr.Left.(*ast.Identifier); ok {
		// We have something like obj[key] = value
		// Check if value refers to obj itself or creates a potential cycle
		a.checkPotentialCircularReference(ident.Value, value, line)
	}
}

// trackVariableAssignment tracks a variable assignment
func (a *DataStructureAnalyzer) trackVariableAssignment(name string, value ast.Expression, line, column int) {
	references := a.extractReferences(value)

	assignment := &VariableAssignment{
		Name:       name,
		Value:      value,
		Line:       line,
		Column:     column,
		Filename:   a.CurrentFilename,
		References: references,
	}

	a.Variables[name] = assignment
	a.AssignmentOrder = append(a.AssignmentOrder, name)
}

// extractReferences extracts variables referenced in an expression
func (a *DataStructureAnalyzer) extractReferences(expr ast.Expression) []string {
	var references []string

	switch e := expr.(type) {
	case *ast.Identifier:
		// Direct reference to another variable
		references = append(references, e.Value)

	case *ast.ArrayLiteral:
		// Check each element in the array
		for _, element := range e.Elements {
			references = append(references, a.extractReferences(element)...)
		}

	case *ast.DictLiteral:
		// Check keys and values in the dictionary
		for key, value := range e.Pairs {
			references = append(references, a.extractReferences(key)...)
			references = append(references, a.extractReferences(value)...)
		}

	case *ast.IndexExpression:
		// Reference to an indexed element
		if left, ok := e.Left.(*ast.Identifier); ok {
			references = append(references, left.Value)
		}
		references = append(references, a.extractReferences(e.Index)...)
	}

	return references
}

// checkPotentialCircularReference checks if an assignment might create a circular reference
func (a *DataStructureAnalyzer) checkPotentialCircularReference(objName string, value ast.Expression, line int) {
	// If we're assigning a variable to itself or something that references it
	if ident, ok := value.(*ast.Identifier); ok && ident.Value == objName {
		a.addCircularWarning(objName, objName, line)
		return
	}

	// Check if value references other variables that might create a cycle
	references := a.extractReferences(value)
	for _, ref := range references {
		if ref == objName {
			a.addCircularWarning(objName, ref, line)
		}
	}
}

// addCircularWarning adds a warning about potential circular references
func (a *DataStructureAnalyzer) addCircularWarning(objName, refName string, line int) {
	warning := Warning{
		Message:  fmt.Sprintf("Kumbukumbu imegunduliwa: %s ni kumbukumbu ya %s", objName, refName),
		Line:     line,
		Column:   1,
		Filename: a.CurrentFilename,
		Type:     PotentialCircularRefWarning,
	}

	a.Warnings = append(a.Warnings, warning)
}

// detectCircularReferences performs a deeper analysis to find potential circular references
func (a *DataStructureAnalyzer) detectCircularReferences() {
	// Create a dependency graph
	graph := make(map[string][]string)

	// Build the graph from variable references
	for name, assignment := range a.Variables {
		graph[name] = assignment.References
	}

	// For each variable, check if it can reach itself through references
	for name := range a.Variables {
		visited := make(map[string]bool)
		path := []string{}

		a.dfsCheckCycle(name, name, graph, visited, path)
	}
}

// dfsCheckCycle performs DFS to check for cycles in the reference graph
func (a *DataStructureAnalyzer) dfsCheckCycle(start, current string, graph map[string][]string, visited map[string]bool, path []string) {
	// If we've visited this node before, skip it
	if visited[current] {
		return
	}

	// Mark as visited and add to path
	visited[current] = true
	path = append(path, current)

	// Check all references from this variable
	for _, ref := range graph[current] {
		// If we've found a reference to our starting variable, we have a cycle
		if ref == start && len(path) > 1 {
			// Found a cycle
			a.addComplexCircularWarning(path, ref)
			return
		}

		// Continue DFS
		a.dfsCheckCycle(start, ref, graph, visited, path)
	}
}

// addComplexCircularWarning adds a warning for a complex circular reference
func (a *DataStructureAnalyzer) addComplexCircularWarning(path []string, closing string) {
	// Only add if we haven't already warned about this pattern
	// (avoid duplicate warnings for the same cycle)

	cyclePath := append(path, closing)
	cycleStr := ""
	for i, name := range cyclePath {
		if i > 0 {
			cycleStr += " → "
		}
		cycleStr += name
	}

	// Find the assignment for the first variable in the cycle
	if assignment, ok := a.Variables[path[0]]; ok {
		warning := Warning{
			Message:  fmt.Sprintf("Rejeleo changamano la duara limegunduliwa: %s", cycleStr),
			Line:     assignment.Line,
			Column:   assignment.Column,
			Filename: assignment.Filename,
			Type:     PotentialCircularRefWarning,
		}

		// Check if this exact warning already exists
		exists := false
		for _, w := range a.Warnings {
			if w.Message == warning.Message && w.Line == warning.Line {
				exists = true
				break
			}
		}

		if !exists {
			a.Warnings = append(a.Warnings, warning)
		}
	}
}
