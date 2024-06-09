package repl

import (
	"fmt"
	"strings"

	"github.com/NuruProgramming/Nuru/evaluator"
	"github.com/NuruProgramming/Nuru/lexer"
	"github.com/NuruProgramming/Nuru/object"
	"github.com/NuruProgramming/Nuru/parser"
	"github.com/NuruProgramming/Nuru/styles"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

var (
	buttonStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder(), true, false).
			Padding(0, 3).
			MarginTop(1)

	activeButtonStyle = buttonStyle.Copy().
				Foreground(lipgloss.Color("#FFF7DB")).
				Background(lipgloss.Color("#aa6f5a")).
				Margin(0, 2).
				Underline(true)

	tableOfContentStyle = lipgloss.NewStyle().Margin(1, 2).BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#aa6f5a")).
				Foreground(lipgloss.Color("#aa6f5a")).
				Padding(2)
)

type item struct {
	title, desc, filename string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type languages struct {
	title, desc, dir string
}

func (l languages) Title() string       { return l.title }
func (l languages) Description() string { return l.desc }
func (l languages) FilterValue() string { return l.title }

type playground struct {
	id             string
	output         viewport.Model
	code           string
	editor         textarea.Model
	docs           viewport.Model
	ready          bool
	filename       string
	content        []byte
	mybutton       string
	fileSelected   bool
	toc            list.Model
	windowWidth    int
	windowHeight   int
	docRenderer    *glamour.TermRenderer
	language       string
	languageCursor list.Model
}

func (pg playground) Init() tea.Cmd {
	return textarea.Blink
}

func (pg playground) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		edCmd  tea.Cmd
		opCmd  tea.Cmd
		docCmd tea.Cmd
		tocCmd tea.Cmd
	)

	pg.editor, edCmd = pg.editor.Update(msg)
	pg.output, opCmd = pg.output.Update(msg)
	pg.languageCursor, _ = pg.languageCursor.Update(msg)
	if !pg.fileSelected {
		pg.toc, tocCmd = pg.toc.Update(msg)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			fmt.Println(pg.editor.Value())
			return pg, tea.Quit
		case tea.KeyEnter:
			if pg.language == "" {
				i, ok := pg.languageCursor.SelectedItem().(languages)
				if ok {
					pg.language = i.dir
					if pg.language == "en" {
						pg.toc = list.New(englishItems, list.NewDefaultDelegate(), pg.windowWidth/2-4, pg.windowHeight-8)
						pg.toc.Title = "Table of Contents"
					} else {
						pg.toc = list.New(kiswahiliItems, list.NewDefaultDelegate(), pg.windowWidth/2-4, pg.windowHeight-8)
						pg.toc.Title = "Yaliyomo"
					}
					return pg, tea.EnterAltScreen
				}
			}
			i, ok := pg.toc.SelectedItem().(item)
			if ok {
				pg.filename = i.filename
				content, err := res.ReadFile("docs/" + pg.language + "/" + pg.filename)
				if err != nil {
					panic(err)
				}
				pg.content = content
				str, err := pg.docRenderer.Render(string(pg.content))
				if err != nil {
					panic(err)
				}

				pg.docs.SetContent(str + "\n\n\n\n\n\n")

				if err != nil {
					panic(err)
				}
				pg.fileSelected = true
				pg.editor.Focus()
			}
		case tea.KeyCtrlR:
			if strings.Contains(pg.editor.Value(), "jaza") {
				pg.output.SetContent(styles.HelpStyle.Italic(false).Render("Samahani, huwezi kutumia `jaza()` kwa sasa."))
			} else {
				// this is just for the output will find a better solution
				code := strings.ReplaceAll(pg.editor.Value(), "andika", "_andika")
				pg.code = code
				env := object.NewEnvironment()
				l := lexer.New(pg.code)
				p := parser.New(l)
				program := p.ParseProgram()
				if len(p.Errors()) != 0 {
					pg.output.Style = styles.ErrorStyle.PaddingLeft(3)
					pg.output.SetContent(strings.Join(p.Errors(), "\n"))
				} else {
					evaluated := evaluator.Eval(program, env)
					if evaluated != nil {
						if evaluated.Type() != object.NULL_OBJ {
							pg.output.Style = styles.ReplStyle.PaddingLeft(3)
							content := evaluated.Inspect()
							l := strings.Split(content, "\n")
							if len(l) > 15 {
								content = strings.Join(l[len(l)-16:], "\n")
							}
							pg.output.SetContent(content)
						}
					}
				}
			}
		case tea.KeyEsc:
			if pg.fileSelected {
				pg.fileSelected = false
				pg.editor.Blur()
			}
		}

	case tea.MouseMsg:
		if zone.Get(pg.id + "docs").InBounds(msg) {
			pg.docs, docCmd = pg.docs.Update(msg)

		}
		switch msg.Type {
		case tea.MouseLeft:
			if zone.Get(pg.id + "run").InBounds(msg) {
				if strings.Contains(pg.editor.Value(), "jaza") {
					pg.output.SetContent(styles.HelpStyle.Italic(false).Render("Samahani, huwezi kutumia `jaza()` kwa sasa."))
				} else {
					// this is just for the output will find a better solution
					code := strings.ReplaceAll(pg.editor.Value(), "andika", "_andika")
					pg.code = code
					env := object.NewEnvironment()
					l := lexer.New(pg.code)
					p := parser.New(l)
					program := p.ParseProgram()
					if len(p.Errors()) != 0 {
						pg.output.Style = styles.ErrorStyle.PaddingLeft(3)
						pg.output.SetContent(strings.Join(p.Errors(), "\n"))
					} else {
						evaluated := evaluator.Eval(program, env)
						if evaluated != nil {
							if evaluated.Type() != object.NULL_OBJ {
								pg.output.Style = styles.ReplStyle.PaddingLeft(3)
								content := evaluated.Inspect()
								l := strings.Split(content, "\n")
								if len(l) > 15 {
									content = strings.Join(l[len(l)-16:], "\n")
								}
								pg.output.SetContent(content)
							}
						}
					}
				}

			}
		}
	case tea.WindowSizeMsg:
		if !pg.ready {
			// editor code
			pg.editor = textarea.New()
			if pg.language == "en" {
				pg.editor.Placeholder = "Write Nuru code here..."
			} else {
				pg.editor.Placeholder = "Andika code yako hapa..."
			}

			pg.editor.Prompt = "┃ "
			pg.editor.SetWidth(msg.Width / 2)
			pg.editor.SetHeight((2 * msg.Height / 3) - 4)

			pg.editor.CharLimit = 0
			pg.editor.FocusedStyle.CursorLine = lipgloss.NewStyle()
			pg.editor.FocusedStyle.Base = lipgloss.NewStyle().PaddingTop(2).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("238"))

			pg.editor.ShowLineNumbers = true

			// output of editor
			pg.output = viewport.New(msg.Width/2, msg.Height/3-4)
			pg.output.Style = lipgloss.NewStyle().PaddingLeft(3)
			var output string
			if pg.language == "en" {

				output = "Your code output will be displayed here..." + strings.Repeat(" ", msg.Width-6)
			} else {
				output = "Matokeo hapa..." + strings.Repeat(" ", msg.Width-6)
			}
			pg.output.SetContent(output)

			// documentation
			pg.docs = viewport.New(msg.Width/2, msg.Height)
			pg.docs.KeyMap = viewport.KeyMap{
				Up: key.NewBinding(
					key.WithKeys("up"),
				),
				Down: key.NewBinding(
					key.WithKeys("down"),
				),
			}
			pg.docs.Style = lipgloss.NewStyle().
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("62")).
				Padding(2)

			renderer, err := glamour.NewTermRenderer(
				glamour.WithAutoStyle(),
				glamour.WithWordWrap(msg.Width/2-4),
			)
			if err != nil {
				panic(err)
			}

			pg.docRenderer = renderer

			pg.toc.SetSize(msg.Width, msg.Height-8)
			pg.windowWidth = msg.Width
			pg.windowHeight = msg.Height

			if pg.language == "en" {
				pg.mybutton = activeButtonStyle.Width(msg.Width / 2).Height(1).Align(lipgloss.Center).Render("Run (CTRL + R)")
			} else {

				pg.mybutton = activeButtonStyle.Width(msg.Width / 2).Height(1).Align(lipgloss.Center).Render("Run (CTRL + R)")
			}
			pg.ready = true

		} else {
			pg.editor.SetHeight((2 * msg.Height / 3) - 4)
			pg.editor.SetWidth(msg.Width / 2)
			pg.output.Height = msg.Height/3 - 4
			pg.output.Width = msg.Width / 2

			renderer, err := glamour.NewTermRenderer(
				glamour.WithAutoStyle(),
				glamour.WithWordWrap(msg.Width/2-4),
			)
			if err != nil {
				panic(err)
			}

			pg.docRenderer = renderer
			str, err := pg.docRenderer.Render(string(pg.content))
			if err != nil {
				panic(err)
			}
			pg.docs.Height = msg.Height
			pg.docs.Width = msg.Width / 2

			pg.docs.SetContent(str + "\n\n\n\n\n\n")
			if pg.language == "en" {
				pg.mybutton = activeButtonStyle.Width(msg.Width / 2).Height(1).Align(lipgloss.Center).Render("Run (CTRL + R)")
			} else {
				pg.mybutton = activeButtonStyle.Width(msg.Width / 2).Height(1).Align(lipgloss.Center).Render("Run (CTRL + R)")
			}
			pg.toc.SetSize(msg.Width, msg.Height-8)
			pg.windowWidth = msg.Width
			pg.windowHeight = msg.Height
		}
	}

	return pg, tea.Batch(edCmd, opCmd, docCmd, tocCmd)
}

func (pg playground) View() string {
	if pg.language == "" {
		return lipgloss.NewStyle().PaddingTop(1).Render(pg.languageCursor.View())
	}
	if !pg.ready {
		return "\n Tunakuandalia....."
	}
	var docs string
	if !pg.fileSelected {
		docs = zone.Mark(pg.id+"toc", tableOfContentStyle.Width(pg.windowWidth/2-4).Height(pg.windowHeight-8).Render(pg.toc.View()))
	} else {
		docs = zone.Mark(pg.id+"docs", pg.docs.View())
	}
	button := zone.Mark(pg.id+"run", pg.mybutton)
	return zone.Scan(lipgloss.JoinHorizontal(lipgloss.Center, docs, lipgloss.JoinVertical(lipgloss.Left, pg.editor.View(), button, pg.output.View())))
}
