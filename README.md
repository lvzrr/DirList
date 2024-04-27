# Directory List
I got tired of the default dir command in windows
## Usage
Just add the .exe folder to your path and type `ls`, feel free to change the source code and rebuild the code with: 
```bash
go build ls.go
```
It uses unicode symbols for the icons so it should be compatible with different machines and OS

Note: the screenshot uses the theme Solarized-osaka, so it might look a bit different for you

![Captura de pantalla 2024-04-27 234956](https://github.com/lvzrr/DirList/assets/161524890/0c308adc-e488-47dc-9e42-ae566eb5175e)
### Performance 
In terms of performance, at least in my machine, matches the windows `dir` command in windows 11. Sometimes even surpasses it in speed, with an average runtime of around 0.03 seconds 
