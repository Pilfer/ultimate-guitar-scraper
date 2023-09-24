# Ultimate Guitar Scraper  

Ultimate-Guitar.com is the world's largest online database of guitar tablature. They also have a horrible UI that is filled with advertisements, interstitials, and other annoying stuff.  

This package allows you to programmatically fetch tabs and do pretty much whatever you want with them depending on the data structure of the response. 

#### Potential use-cases might include...

- CLI tab viewer/manager  
- A utility that calculates the most used chords or progressions in a specific set of songs  
- Automatic transposition service  
- A tab "player" - similar to the "GuitarPro" application  
- Save text-based tabs + associated meta  
- Generate and save HTML, PDF, etc tabs  
- Download tabs then upload to popular services like Google Drive, Dropbox, etc  

### Features  

- [X] Commandline interface (WIP)
- [X] Fetch a tab by id  
- [X] Fetch all your saved tabs.
- [X] Fetch tab by URL
- [X] Search for tabs  
- [X] Explore popular tabs  
- [X] Export tab as `.wav` (Thanks to [https://github.com/timiskhakov/music](https://github.com/timiskhakov/music))!!  
- [X] Fetch popular tabs (see: `ultimateguitar.Explore`)  
- [ ] Scrape all tabs by artist  
  -  Fun fact: on mobile, UG doesn't have a "list tabs by artist name/id" endpoint. They just load ~7 pages. Weird. The functionality for this is technically here already, I just didn't add a helper method. Go nuts.  

### Building  

1. `go build` (lol)  

### Using the CLI  

Run `./ultimate-guitar-scraper -h` if you're curious, buuuut...

- Fetch a tab: `./ultimate-guitar-scraper fetch -id 96835 -output wee.wav`  
- Fetch all your saved tabs: `./ultimate-guitar-scraper get_all --output ./out`  
- Fetch and export tab as HTML (using `cmd/data/template.tmpl`): `./ultimate-guitar-scraper export -id 96835`  
- Fetch a tab and export it as a .wav file: `./ultimate-guitar-scraper wav -id 113039 -output hallelujah.wav`  


#### ... But why?  

As much as I appreciate the work UG has done compiling the largest online guitar tabs database, I can't bring myself to use their website or mobile app (and definitely not their website on mobile!). I started working on this package (originally a node module) as a way for me to view tabs/chord charts without dealing with their display ads and interstitials.  


#### Technology Used

- Golang (duh)  
- Frida - [https://frida.re/](https://frida.re/)  
- JEB Decompiler - [https://www.pnfsoftware.com/jeb/android](https://www.pnfsoftware.com/jeb/android)  
- Charles / mitmproxy  
- This awesome experimental package: https://github.com/timiskhakov/music  


## Disclaimer / Legal  

This software's purpose is purely educational. I am not responsible for how you use this package. This repository and all others associated with it are not affiliated with, authorized, or endorsed by Ultimate-Guitar.com. 


