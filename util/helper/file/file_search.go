package file

import (
	stderr "errors"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	FileNotFound = stderr.New("file not found")
)

type Option func(searcher *FileSearcher)
type FileSearcher struct {
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

// NewFileSearcher creates a new FileSearcher instance
func NewFileSearcher(options ...Option) *FileSearcher {
	fileSearcher := &FileSearcher{}

	for _, o := range options {
		o(fileSearcher)
	}
	return fileSearcher
}

func WithDeepSearch() Option {
	return func(searcher *FileSearcher) {
		searcher.ds = true
	}
}

func WithRegex(pattern string) Option {
	return func(searcher *FileSearcher) {
		searcher.regex = true
		searcher.pattern = pattern
	}
}

func WithExactMatch() Option {
	return func(searcher *FileSearcher) {
		searcher.exactMatch = true
	}
}

func (s *FileSearcher) SetSearchDir(searchDir string) *FileSearcher {
	s.searchDir = searchDir
	return s
}

func (s *FileSearcher) SetFileName(fileName string) *FileSearcher {
	s.fileName = fileName
	return s
}

func (s *FileSearcher) SetDeepSearch() *FileSearcher {
	s.ds = true
	return s
}

func (s *FileSearcher) SetRegex(pattern string) *FileSearcher {
	s.regex = true
	s.pattern = pattern
	return s
}

func (s *FileSearcher) SetExactMatch() *FileSearcher {
	s.exactMatch = true
	return s
}

// Find finds a file by the searchStr. It returns the file path if found, otherwise it returns an empty string, with an error.
func (s *FileSearcher) Find(searchDir, searchStr string) (string, error) {
	s.searchDir = searchDir
	s.fileName = searchStr
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
func (s *FileSearcher) lightSearch() (string, error) {
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
		return "", FileNotFound
	}
	return filePath, nil
}

// deepSearch searches for a file by the fileName in the searchDir, searching recursively.
// It searches for the file name in all subdirectories of the searchDir.
// If the file is found, it returns the file path, otherwise it returns an empty string, with an error.
func (s *FileSearcher) deepSearch() (string, error) {
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
		return "", FileNotFound
	}
	return filePath, nil
}

// regexSearch searches for a file by the regex pattern in the searchDir.
// Whether to search recursively is determined by the ds option of the FileSearcher.
func (s *FileSearcher) regexSearch() (string, error) {
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

func (s *FileSearcher) lightRegexSearch(re *regexp.Regexp) (string, error) {
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
		return "", FileNotFound
	}
	return filePath, nil
}

func (s *FileSearcher) deepRegexSearch(re *regexp.Regexp) (string, error) {
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
		return "", FileNotFound
	}
	return filePath, nil
}
