package analysis

import (
	"fmt"
	"os"
	"strings"
)

// AnalyzeCommand handles the "analyze" command in the Nuru CLI
func AnalyzeCommand(args []string) int {
	if len(args) < 1 {
		fmt.Println("Samahani: Hakuna faili au kamusi kupanga.")
		fmt.Println("Uwakilishi: nuru --chambua <faili_au_kamusi> [options]")
		return 1
	}

	targetPath := args[0]
	options := DefaultAnalysisOptions()

	// Parse additional options
	for i := 1; i < len(args); i++ {
		arg := args[i]
		if strings.HasPrefix(arg, "--warning-level=") {
			level := strings.TrimPrefix(arg, "--warning-level=")
			switch level {
			case "kimya":
				options.WarningLevel = WarningSilent
			case "kawaida":
				options.WarningLevel = WarningNormal
			case "kali":
				options.WarningLevel = WarningStrict
			case "kosa":
				options.WarningLevel = WarningAsError
			default:
				fmt.Printf("Kiwango cha onyo si sahihi: %s. Kutumia 'kawaida'.\n", level)
			}
		} else if arg == "--bila-moduli" {
			options.AnalyzeModules = false
		} else if arg == "--bila-kamusi" {
			options.AnalyzeDataStructures = false
		} else if arg == "--bila-aina" {
			options.AnalyzeTypes = false
		} else if strings.HasPrefix(arg, "--tenga=") {
			excludeDirs := strings.TrimPrefix(arg, "--tenga=")
			options.ExcludeDirs = append(options.ExcludeDirs, strings.Split(excludeDirs, ",")...)
		}
	}

	analyzer := NewAnalyzer(options)
	var result *AnalysisResult
	var err error

	// Check if the target is a file or directory
	fileInfo, err := os.Stat(targetPath)
	if err != nil {
		fmt.Printf("Samahani: Haiwezi kufikia %s: %v\n", targetPath, err)
		return 1
	}

	if fileInfo.IsDir() {
		fmt.Printf("Inachambua kamusi: %s\n", targetPath)
		result, err = analyzer.AnalyzeDirectory(targetPath)
	} else {
		fmt.Printf("Inachambua faili: %s\n", targetPath)
		// Read the file
		content, err := os.ReadFile(targetPath)
		if err != nil {
			fmt.Printf("Samahani: faili haisomeki %s: %v\n", targetPath, err)
			return 1
		}

		warnings, err := analyzer.AnalyzeFile(targetPath, string(content))
		if err != nil {
			fmt.Printf("Samahani: Faili haichambuliki %s: %v\n", targetPath, err)
			return 1
		}

		// Create a result with just this file's warnings
		result = &AnalysisResult{
			Warnings:             warnings,
			TotalFilesAnalyzed:   1,
			SummaryByWarningType: make(map[WarningType]int),
		}

		// Count warnings by type
		for _, warning := range warnings {
			result.SummaryByWarningType[warning.Type]++
		}
	}

	if err != nil {
		fmt.Printf("Samahani: hitilafu imeingia wakati wa kuchambua %v\n", err)
		return 1
	}

	// Print the results
	fmt.Println(FormatAnalysisResult(result, true))

	// Generate DOT graph if requested
	if result.ModuleDependencyGraph != nil {
		for i := 1; i < len(args); i++ {
			if strings.HasPrefix(args[i], "--dot=") {
				dotFile := strings.TrimPrefix(args[i], "--dot=")
				dotGraph := result.GenerateDOTGraph()
				err := os.WriteFile(dotFile, []byte(dotGraph), 0644)
				if err != nil {
					fmt.Printf("Hitilafu katika kuandika grafu ya DOT kwa %s: %v\n", dotFile, err)
				} else {
					fmt.Printf("Grafu ya utegemezi wa moduli iliyoandikwa %s\n", dotFile)
				}
				break
			}
		}
	}

	// Return appropriate exit code based on warnings and warning level
	if len(result.Warnings) > 0 && options.WarningLevel == WarningAsError {
		return 1
	}

	return 0
}

// HandleVisualizeDepsCommand handles the "visualize-deps" command
func HandleVisualizeDepsCommand(args []string) int {
	if len(args) < 1 {
		fmt.Println("Samahani: Hakuna faili au kamusi kupanga.")
		fmt.Println("Uwakilishi: nuru --tengeneza <faili_au_kamusi> [--aina=dot|json|svg]")
		return 1
	}

	targetPath := args[0]
	format := "dot"

	// Parse format option
	for i := 1; i < len(args); i++ {
		if strings.HasPrefix(args[i], "--aina=") {
			format = strings.TrimPrefix(args[i], "--aina=")
		}
	}

	// Ensure the path exists
	_, err := os.Stat(targetPath)
	if err != nil {
		fmt.Printf("Samahani: Haiwezi kufikia %s: %v\n", targetPath, err)
		return 1
	}

	// Analyze the target
	options := DefaultAnalysisOptions()
	options.AnalyzeDataStructures = false
	options.AnalyzeTypes = false
	analyzer := NewAnalyzer(options)

	var result *AnalysisResult
	if isDir, _ := isDirectory(targetPath); isDir {
		result, err = analyzer.AnalyzeDirectory(targetPath)
	} else {
		content, err := os.ReadFile(targetPath)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", targetPath, err)
			return 1
		}

		// For a single file, we need to create a mini dependency graph
		files := map[string]string{
			targetPath: string(content),
		}

		moduleGraph, _, err := BuildDependencyGraph(files)
		if err != nil {
			fmt.Printf("Samahani: hitilafu imeingia wakati wa kuchambua %v\n", err)
			return 1
		}

		result = &AnalysisResult{
			ModuleDependencyGraph: moduleGraph,
			TotalFilesAnalyzed:    1,
		}
	}

	if err != nil {
		fmt.Printf("Samahani: hitilafu imeingia wakati wa kuchambua %v\n", err)
		return 1
	}

	// Generate the visualization
	switch format {
	case "dot":
		fmt.Println(result.GenerateDOTGraph())
	case "json":
		fmt.Println(generateJSONGraph(result.ModuleDependencyGraph))
	default:
		fmt.Printf("Samahani: Aina imepungua: %s. Aina zilizotumika: dot, json\n", format)
		return 1
	}

	return 0
}

// isDirectory checks if a path is a directory
func isDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), nil
}

// generateJSONGraph generates a JSON representation of the dependency graph
func generateJSONGraph(graph *ModuleDependencyGraph) string {
	if graph == nil {
		return "{}"
	}

	var json strings.Builder
	json.WriteString("{\n  \"nodes\": [\n")

	// Add nodes
	nodes := make([]string, 0, len(graph.Modules))
	for name := range graph.Modules {
		nodes = append(nodes, fmt.Sprintf("    {\"id\": \"%s\"}", name))
	}
	json.WriteString(strings.Join(nodes, ",\n"))

	json.WriteString("\n  ],\n  \"links\": [\n")

	// Add edges
	links := []string{}
	for name, node := range graph.Modules {
		for _, dep := range node.Dependencies {
			links = append(links, fmt.Sprintf("    {\"source\": \"%s\", \"target\": \"%s\"}", name, dep.Name))
		}
	}
	json.WriteString(strings.Join(links, ",\n"))

	json.WriteString("\n  ]\n}")
	return json.String()
}
