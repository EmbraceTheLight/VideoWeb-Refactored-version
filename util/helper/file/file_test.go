package file_test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"
	"util/helper/file"
)

func TestSearchFile(t *testing.T) {
	// 1. Without DeepSearch, search will be failed
	defaultSearcher := file.NewFileSearcher()
	path, err := defaultSearcher.Find("../../../resources", "avatar")
	require.EqualError(t, err, file.NotFound.Error())

	// 2. With DeepSearch, search will be successful, and can open the file
	deepSearcher := file.NewFileSearcher(file.WithDeepSearch())
	path, err = deepSearcher.Find("../../../resources", "avatar")
	fmt.Println(path)
	basePath := filepath.Dir(path)
	fileName := filepath.Base(path)
	fmt.Println(basePath, fileName)
	require.NoError(t, err)
	f, err := os.Open(path)
	require.NoError(t, err)
	require.NoError(t, f.Close())

	// 3. With DeepSearch and ExactMatch, search will be successful, and can open the file
	exactSearcher := file.NewFileSearcher(file.WithDeepSearch())
	path, err = exactSearcher.Find("../../../resources", "avatar.jpg")
	require.NoError(t, err)
	f, err = os.Open(path)
	require.NoError(t, err)
	require.NoError(t, f.Close())

	// 4. With DeepSearch and ExactMatch, but wrong file name, search will be successful, and can open the file
	exactSearcher = file.NewFileSearcher(file.WithDeepSearch())
	path, err = exactSearcher.Find("../../../resources", "avatar.png")
	require.EqualError(t, err, file.NotFound.Error())

	// 5. Chained searchers
	chainedSearcher := file.NewFileSearcher()
	path, err = chainedSearcher.SetDeepSearch().SetExactMatch().Find("../../../resources", "avatar.jpg")
	require.NoError(t, err)
	f, err = os.Open(path)
	require.NoError(t, err)
	require.NoError(t, f.Close())
}

func Test1(t *testing.T) {
	_, err := os.Create("dash/test.txt")
	require.NoError(t, err)
}
