package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var redOpen string = "\x1b[31m"
var redClose string = "\x1b[0m"
var blueOpen string = "\x1b[34m"
var blueClose string = "\x1b[0m"
var orangeOpen string = "\x1b[33m"
var orangeClose string = "\x1b[0m"
var yellowOpen string = "\x1b[33m"
var yellowClose string = "\x1b[0m"
var greyOpen string = "\x1b[90m"
var greyClose string = "\x1b[0m"
var greenOpen string = "\x1b[32m"
var greenClose string = "\x1b[0m"
var cyanOpen string = "\x1b[36m"
var cyanClose string = "\x1b[0m"
var magentaOpen string = "\x1b[35m"
var magentaClose string = "\x1b[0m"
const tealOpen = "\033[96m" // Teal color open tag
const tealClose = "\033[0m" // Reset color tag
const purpleOpen = "\033[95m" // Purple color open tag
const purpleClose = "\033[0m" // Reset color tag
var lightRedOpen string = "\033[91m" // Light red color open tag
var lightRedClose string = "\033[0m" // Reset color tag
var lightGreenOpen string = "\033[92m" // Light green color open tag
var lightGreenClose string = "\033[0m" // Reset color tag
var lightyellowOpen string = "\033[93m" // Light yellow color open tag
var lightyellowClose string = "\033[0m" // Reset color tag
var lightBlueOpen string = "\033[94m" // Light blue color open tag
var lightBlueClose string = "\033[0m" // Reset color tag
var lightMagentaOpen string = "\033[95m" // Light magenta color open tag
var lightMagentaClose string = "\033[0m" // Reset color tag
var lightCyanOpen string = "\033[96m" // Light cyan color open tag
var lightCyanClose string = "\033[0m" // Reset color tag
var lightOrangeOpen string = "\033[38;5;214m" // Light orange color open tag
var lightOrangeClose string = "\033[0m"      // Reset color tag
var whiteOpen string = "\033[97m"
var whiteClose string = "\033[0m"


func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		return
	}

	current, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("\nCurrent directory: %-s\n\n", current)
	fmt.Printf("%-16s %-23s %-16s %-9s %-9s %-26s\n\n", "MOD", "LASTMODTIME","SIZE", "TYPE","ICO", "NAME")

	for _, f := range files {
    var extension string = filepath.Ext(f.Name())
		var c, cc,ico  string

		if !f.IsDir() {
			c, cc , ico = checkcolor(extension)[0], checkcolor(extension)[1],checkcolor(extension)[2]
			fmt.Printf("%-16s %-23s %-16d %-14s %-9s%-26s\n", f.Mode(), f.ModTime().Format("02/01/06 15:04"), f.Size(), c+extension,ico, f.Name()+cc)
		} else {
			c = greyOpen
			cc = greyClose
      ico="📁"
			fmt.Printf("%-16s %-23s %-16d %-18s %-9s%-26s\n", f.Mode(), f.ModTime().Format("02/01/06 15:04"), f.Size(), c+"dir"+cc, ico,f.Name())
		}
	}
fmt.Println("\n")
}
func checkcolor(extention string) []string{
  var c,cc,ico string
    switch extention {
    case ".py":
        c = blueOpen
        cc = blueClose
        ico = "🐍" // 🐍 Python snake
    case ".go":
        c = blueOpen
        cc = blueClose
        ico = "🐧" // 🐧 Gopher
    case ".c", ".cpp":
        c = redOpen
        cc = redClose
        ico = "🔨" // 🔨 Hammer and wrench
    case ".exe", ".sh":
        c = lightGreenOpen
        cc = lightGreenClose
        ico = "💾" // 💻 Laptop
    case ".java":
        c = yellowOpen
        cc = yellowClose
        ico = "☕" // ☕ Coffee
    case ".rs":
        c = orangeOpen
        cc = orangeClose
        ico = "🦀" // 🦀 Rust crab
    case ".mp4", ".avi", ".mov":
        c = magentaOpen
        cc = magentaClose
        ico = "🎥" // 🎥 Movie camera
    case ".mp3", ".wav", ".flac", ".ogg", ".opus":
        c = cyanOpen
        cc = cyanClose
        ico = "🎵" // 🎵 Musical note
    case ".html", ".css" :
        c = purpleOpen
        cc = purpleClose
        ico = "🌐" // 🌐 Globe with meridians
    case ".txt", ".md":
        c = whiteOpen
        cc = whiteClose
        ico = "📄" // 📄 Page facing up
    case ".svg", ".png", ".jpg":
        c = tealOpen
        cc = tealClose
        ico = "🖼️" // 🖼️ Framed picture
    case ".pdf":
        c = lightRedOpen
        cc = lightRedClose
        ico = "📝" 
    case ".pptx":
        c = lightOrangeOpen
        cc = lightOrangeClose
        ico = "\U0001F4DA" // 🖼️ Framed picture
    case ".json", ".js":
        c = lightyellowOpen
        cc = lightyellowClose
        ico = "📜" // 📝 Page facing up with pen (light yellow icon)
    case ".zip":
        c = lightGreenOpen
        cc = lightGreenClose
        ico = "📦" // 📦 Package (light green icon)

    default:
        c = whiteOpen
        cc = whiteClose
        ico = "📰"
    }
return []string{c,cc,ico}
}
