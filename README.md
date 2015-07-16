# configexp
Discovery tool for finding and parsing .Net config web.config and AppName.config
file which contains the <appSettings></appSettings>.

This tool will assume all files are in one root directory.
## Tasks
- Create a function which returns a slice of values from a directory containing
the filepaths to all of the .config files.
- Create a function which can parse one xml file and return the key values pairs
- Create a function which can take a string value and parse a group of xml files and return fuzzy matches
- Create a http net server to serve up JSON for the results
- Add tests
