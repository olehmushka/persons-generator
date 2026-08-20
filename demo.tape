# VHS tape script -- generates examples/demo.gif. Run via
# .github/workflows/demo-gif.yml (charmbracelet/vhs-action), since this
# environment has no terminal renderer to run VHS end-to-end locally.
# See https://github.com/charmbracelet/vhs

Output examples/demo.gif

Set Theme "Dracula"
Set FontSize 16
Set Width 1200
Set Height 650
Set Padding 20

Type "go run main.go generate_religion"
Enter
Sleep 4s
