package analysis

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/NuruProgramming/Nuru/lexer"
	"github.com/NuruProgramming/Nuru/parser"
)

// AnalysisOptions configures which analyses to perform
type AnalysisOptions struct {
	AnalyzeModules        bool
	AnalyzeDataStructures bool
	AnalyzeTypes          bool
	WarningLevel          WarningLevel
	ExcludeDirs           []string
}

type WarningLevel string

const (
	WarningSilent  WarningLevel = "kimya"
	WarningNormal  WarningLevel = "kawaida"
	WarningStrict  WarningLevel = "kali"
	WarningAsError WarningLevel = "kosa"
)

// DefaultAnalysisOptions returns the default analysis options
func DefaultAnalysisOptions() AnalysisOptions {
	return AnalysisOptions{
		AnalyzeModules:        true,
		AnalyzeDataStructures: true,
		AnalyzeTypes:          true,
		WarningLevel:          WarningNormal,
		ExcludeDirs:           []string{".git", "node_modules"},
	}
}

// AnalysisResult contains the results of a static analysis
type AnalysisResult struct {
	Warnings              []Warning
	ModuleDependencyGraph *ModuleDependencyGraph
	TotalFilesAnalyzed    int
	SummaryByWarningType  map[WarningType]int
}

// Analyzer is the main static analysis engine
type Analyzer struct {
	Options      AnalysisOptions
	FileContents map[string]string
}

// NewAnalyzer creates a new static analyzer with the given options
func NewAnalyzer(options AnalysisOptions) *Analyzer {
	return &Analyzer{
		Options:      options,
		FileContents: make(map[string]string),
	}
}

// AnalyzeFile analyzes a single file
func (a *Analyzer) AnalyzeFile(filename string, content string) ([]Warning, error) {
	var warnings []Warning

	// Parse the file
	l := lexer.New(content)
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		return nil, fmt.Errorf("parsing errors in %s: %v", filename, p.Errors())
	}

	// Store the file content for later use
	a.FileContents[filename] = content

	// Run data structure analysis if enabled
	if a.Options.AnalyzeDataStructures {
		dsAnalyzer := NewDataStructureAnalyzer()
		dsAnalyzer.AnalyzeFile(filename, program)
		warnings = append(warnings, dsAnalyzer.Warnings...)
	}

	return warnings, nil
}

// AnalyzeDirectory analyzes all .nr files in a directory and its subdirectories
func (a *Analyzer) AnalyzeDirectory(rootDir string) (*AnalysisResult, error) {
	result := &AnalysisResult{
		Warnings:             []Warning{},
		SummaryByWarningType: make(map[WarningType]int),
	}

	// First pass: collect all files and their contents
	files, err := a.collectFiles(rootDir)
	if err != nil {
		return nil, err
	}

	// Store the number of files analyzed
	result.TotalFilesAnalyzed = len(files)

	// Second pass: analyze individual files
	for filename, content := range files {
		fileWarnings, err := a.AnalyzeFile(filename, content)
		if err != nil {
			return nil, err
		}

		result.Warnings = append(result.Warnings, fileWarnings...)

		// Count warnings by type
		for _, warning := range fileWarnings {
			result.SummaryByWarningType[warning.Type]++
		}
	}

	// Third pass: analyze module dependencies if enabled
	if a.Options.AnalyzeModules {
		moduleGraph, moduleWarnings, err := BuildDependencyGraph(files)
		if err != nil {
			return nil, err
		}

		result.ModuleDependencyGraph = moduleGraph
		result.Warnings = append(result.Warnings, moduleWarnings...)

		// Count module warnings by type
		for _, warning := range moduleWarnings {
			result.SummaryByWarningType[warning.Type]++
		}
	}

	return result, nil
}

// collectFiles recursively collects all .nr files in a directory
func (a *Analyzer) collectFiles(rootDir string) (map[string]string, error) {
	files := make(map[string]string)

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip excluded directories
		if info.IsDir() {
			dirName := filepath.Base(path)
			for _, excludeDir := range a.Options.ExcludeDirs {
				if dirName == excludeDir {
					return filepath.SkipDir
				}
			}
			return nil
		}

		// Only process .nr files
		if !strings.HasSuffix(path, ".nr") {
			return nil
		}

		// Read the file
		content, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		// Store the file content
		files[path] = string(content)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

// FormatWarning formats a warning for display
func FormatWarning(warning Warning, colorize bool) string {
	var prefix string

	if colorize {
		prefix = "\033[33m⚠️  WARNING:\033[0m"
	} else {
		prefix = "WARNING:"
	}

	location := fmt.Sprintf("\n   → File: %s, Line: %d", warning.Filename, warning.Line)

	return fmt.Sprintf("%s %s%s", prefix, warning.Message, location)
}

// FormatAnalysisResult formats the analysis result for display
func FormatAnalysisResult(result *AnalysisResult, colorize bool) string {
	var output strings.Builder

	// Format summary
	output.WriteString(fmt.Sprintf("Kupanga kwa kutafsiri:\n"))
	output.WriteString(fmt.Sprintf("Nambari ya faili zilizotafsiriwa: %d\n", result.TotalFilesAnalyzed))

	// Summary by warning type
	output.WriteString("Maonyo kwa aina:\n")
	for warningType, count := range result.SummaryByWarningType {
		output.WriteString(fmt.Sprintf("  %s: %d\n", warningType, count))
	}

	// Format individual warnings
	if len(result.Warnings) > 0 {
		output.WriteString("\nMaonyo Kwa kina:\n")
		for _, warning := range result.Warnings {
			output.WriteString(FormatWarning(warning, colorize))
			output.WriteString("\n")
		}
	} else {
		output.WriteString("\nHakuna maonyo zilizotafsiriwa.\n")
	}

	return output.String()
}

// GenerateDOTGraph generates a DOT format graph for visualization
func (result *AnalysisResult) GenerateDOTGraph() string {
	if result.ModuleDependencyGraph != nil {
		return result.ModuleDependencyGraph.GenerateDOTGraph()
	}
	return ""
}
