# mtf
This script accepts a plain-text file containing a list of movie titles (one per line), it will search every title starting with lowercase letter on [TMDB](themoviedb.org) and produce an output file containig corrected movie titles. It automatically applies the correct title only if there is a case mismatch, if the input title is incorrect the script will ask the user to choose betweeen the first 5 results. Also lines with an * will be ignored because i said so.

## Usage

```bash
go run main.go textfile_with_titles
```

or just compile it.