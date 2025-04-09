package file

import (
	stderr "errors"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	NotFound     = stderr.New("file not found")
	MissingParam = stderr.New("missing searchDir or fileName parameter")
)

type Option func(searcher *Searcher)
type Searcher struct {
	// directory to search into
	searchDir string

	// file name to search for
	fileName string

	// ds ---- deepSearch: whether to search recursively, default is false
	ds bool

	// whether to match file name exactly, default is false
	exactMatch bool

	// whether to use regex pattern to search for file name, default is false
	regex bool

	// regex pattern to search for file name
	pattern string
}

// NewFileSearcher creates a new Searcher instance
func NewFileSearcher(options ...Option) *Searcher {
	fileSearcher := &Searcher{}

	for _, o := range options {
		o(fileSearcher)
	}
	return fileSearcher
}

func WithDeepSearch() Option {
	return func(searcher *Searcher) {
		searcher.ds = true
	}
}

func WithRegex(pattern string) Option {
	return func(searcher *Searcher) {
		searcher.regex = true
		searcher.pattern = pattern
	}
}

func WithExactMatch() Option {
	return func(searcher *Searcher) {
		searcher.exactMatch = true
	}
}

func (s *Searcher) SetSearchDir(searchDir string) *Searcher {
	s.searchDir = searchDir
	return s
}

func (s *Searcher) SetFileName(fileName string) *Searcher {
	s.fileName = fileName
	return s
}

func (s *Searcher) SetDeepSearch() *Searcher {
	s.ds = true
	return s
}

func (s *Searcher) SetRegex(pattern string) *Searcher {
	s.regex = true
	s.pattern = pattern
	return s
}

func (s *Searcher) SetExactMatch() *Searcher {
	s.exactMatch = true
	return s
}

// Find finds a file by the searchStr. It returns the file path if found, otherwise it returns an empty string, with an error.
func (s *Searcher) Find(searchDir, searchStr string) (string, error) {
	if searchDir != "" {
		s.searchDir = searchDir
	}
	if s.fileName == "" {
		s.fileName = searchStr
	}

	if s.searchDir == "" || s.fileName == "" {
		return "", MissingParam
	}
	switch {
	case s.ds:
		return s.deepSearch()
	case s.regex:
		return s.regexSearch()
	default:
		return s.lightSearch()
	}
}

// lightSearch searches for a file by the fileName in the searchDir, without searching recursively.
// It just searches for the file name in the current directory.
func (s *Searcher) lightSearch() (string, error) {
	var filePath string
	files, err := os.ReadDir(s.searchDir)
	if err != nil {
		return "", err
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if s.exactMatch && file.Name() == s.fileName {
			filePath = filepath.Join(s.searchDir, file.Name())
			break
		} else if !s.exactMatch && strings.Contains(file.Name(), s.fileName) {
			filePath = filepath.Join(s.searchDir, file.Name())
			break
		}
	}
	if filePath == "" {
		return "", NotFound
	}
	return filePath, nil
}

// deepSearch searches for a file by the fileName in the searchDir, searching recursively.
// It searches for the file name in all subdirectories of the searchDir.
// If the file is found, it returns the file path, otherwise it returns an empty string, with an error.
func (s *Searcher) deepSearch() (string, error) {
	var filePath string
	searchFunc := func(path string, info os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && !s.exactMatch && strings.Contains(info.Name(), s.fileName) {
			filePath = path
			return filepath.SkipAll
		} else if s.exactMatch && info.Name() == s.fileName {
			filePath = path
			return filepath.SkipAll
		}
		return nil
	}

	err := filepath.WalkDir(s.searchDir, searchFunc)
	if err != nil {
		return "", err
	}

	if filePath == "" {
		return "", NotFound
	}
	return filePath, nil
}

// regexSearch searches for a file by the regex pattern in the searchDir.
// Whether to search recursively is determined by the ds option of the Searcher.
func (s *Searcher) regexSearch() (string, error) {
	re, err := regexp.Compile(s.pattern)
	if err != nil {
		return "", err
	}

	if s.ds {
		return s.deepRegexSearch(re)
	} else {
		return s.lightRegexSearch(re)
	}
}

func (s *Searcher) lightRegexSearch(re *regexp.Regexp) (string, error) {
	var filePath string
	files, err := os.ReadDir(s.searchDir)
	if err != nil {
		return "", err
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if re.MatchString(file.Name()) {
			filePath = filepath.Join(s.searchDir, file.Name())
			break
		}
	}
	if filePath == "" {
		return "", NotFound
	}
	return filePath, nil
}

func (s *Searcher) deepRegexSearch(re *regexp.Regexp) (string, error) {
	var filePath string
	searchFunc := func(path string, info os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && re.MatchString(info.Name()) {
			filePath = path
			return filepath.SkipAll
		}
		return nil
	}

	err := filepath.WalkDir(s.searchDir, searchFunc)
	if err != nil {
		return "", err
	}

	if filePath == "" {
		return "", NotFound
	}
	return filePath, nil
}
