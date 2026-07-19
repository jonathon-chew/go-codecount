# CHANGELOG

	## v0.5.9, tag: v0.5.8 

		### UPDATES
		1. README clean up

	## v0.5.7 

		### UPDATES
		1. upgrading Aphrodite so only outputs ANSI if output is a terminal
		1. removing local md files for notes and the binary i create when conducting more manual test for behaviour

	## v0.5.6 

		### UPDATES
		1. moving allowed to check on each file rather than global
		1. replace the hard coded  path with a local one

	## v0.5.5 

		### NEW
		1. break when found the right file

	## v0.5.4 

		### MISC
		1. exclusive was blocking all path ways even when exclusive was empty

	## v0.5.3 

		### NEW
		1. now can search for only specific types of files exclusively

	## v0.5.2 

		### MISC
		1. include files added to the right list

	## v0.5.1 

		### UPDATES
		1. rename cli package from cmd to cli; new: adding common build names for the bianry to the .gitignore
		1. rename project and file structre moved over to match traditional go structures

	## v0.5.0 

		### NEW
		1. sort results in alphabetical order
		1. license

	## v0.4.2 

		### NEW
		1. option for including specific files
		### UPDATES
		1. README.md for formatting on github

	## v0.4.1 

		### UPDATES
		1. README.md to reflect current implimentation, output and features

	## v0.4.0 

		### UPDATES
		1. version number

	## v0.3.0 

		### NEW
		1. addToList function and countWords function for significant performance benefits - 66,000 files from 34s to 15s; refactor final print to be only a single print, map instead of slice for storing languages for O(1) instead of O(N)

	## v0.2.1 

		### UPDATES
		1. total treated as a normal line instead of a seperate line

	## v0.2.0 

		### UPDATES
		1. if i == 0 was meaning a file type was being missed out every time!

	## v0.1.1 

		### NEW
		1. Empty Lines count, ignore errors as an option; refactor: printing to the screen so header and rows are added to the same formatting, biggest number worked out in same loop instead of in it's own loop

	## v0.1.0 

		### UPDATES
		1. file path breaker for ignore so will work on windows

	## v0.0.7 

		### REFACTORS
		1. total line to be human readable instread of just an int; update: spacing on print out - still in development

	## v0.0.6 

		### NEW
		1. added a toal final row and a header title

	## v0.0.5 

		### NEW
		1. added the ability to ignore file names / extensions, this is currently for anything in the file name or extension by design

	## v0.0.4 

		### NEW
		1. flag to remove any folder

	## v0.0.3 

		### NEW
		1. added the ability to add unknown file types - but uses the extension and counts lines
		### UPDATES
		1. powershell files working on mac

	## v0.0.2 

		### UPDATES
		1. added different powershell file types