# Ultimate Guitar Scraper  

Ultimate-Guitar.com is the world's largest online database of guitar tablature. They also have a horrible UI that is filled with advertisements, interstitials, and other annoying stuff.  

This package allows you to programmatically fetch tabs and do pretty much whatever you want with them depending on the data structure of the response. 

#### Potential use-cases might include...

- CLI tab viewer/manager  
- A utility that calculates the most used chords or progressions in a specific set of songs  
- Automatic transposition service  
- A tab "player" - similar to the "GuitarPro" application  

#### ... But why?  

As much as I appreciate the work UG has done compiling the largest online guitar tabs database, I can't bring myself to use their website or mobile app (and definitely not their website on mobile!). I started working on this package (originally a node module) as a way for me to view tabs/chord charts without dealing with their display ads and interstitials.  


#### Technology Used

- Golang (duh)  
- Frida - [https://frida.re/](https://frida.re/)    
- JEB Decompiler - [https://www.pnfsoftware.com/jeb/android](https://www.pnfsoftware.com/jeb/android))  
- Charles / mitmproxy


### Features  

- [x] Commandline interface (WIP)
- [x] Fetch a tab by id  
- [x] Fetch tab by URL
- [ ] Search for tabs - Will add later  
- [ ] Fetch popular tabs - Will add later   


## Disclaimer / Legal  

This software's purpose is purely educational. I am not responsible for how you use this package. This repository and all others associated with it are not affiliated with, authorized, affiliated, or endorsed by Ultimate-Guitar.com. 