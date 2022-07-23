package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/packagefoundation/yap/constants"
)

func ReadFile(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("%s❌ :: %sfailed to read file '%s'%s\n",
			string(constants.ColorBlue),
			string(constants.ColorYellow),
			path,
			string(constants.ColorWhite))

		return data, err
	}

	return data, err
}

func ReadDir(path string) ([]os.FileInfo, error) {
	items, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("%s❌ :: %sfailed to read dir '%s'%s\n",
			string(constants.ColorBlue),
			string(constants.ColorYellow),
			path,
			string(constants.ColorWhite))

		return items, err
	}

	return items, err
}

func MkdirAll(path string) error {
	err := os.MkdirAll(path, 0o755)
	if err != nil {
		fmt.Printf("%s❌ :: %sfailed to mkdir '%s'%s\n",
			string(constants.ColorBlue),
			string(constants.ColorYellow),
			path,
			string(constants.ColorWhite))

		return err
	}

	return err
}

func Chmod(path string, perm os.FileMode) error {
	err := os.Chmod(path, perm)
	if err != nil {
		fmt.Printf("%s❌ :: %sfailed to chmod '%s'%s\n",
			string(constants.ColorBlue),
			string(constants.ColorYellow),
			path,
			string(constants.ColorWhite))

		return err
	}

	return err
}

func ChownR(path string, user, group string) error {
	err := Exec("",
		"chown",
		"-R",
		fmt.Sprintf("%s:%s", user, group),
		path,
	)

	if err != nil {
		return err
	}

	return err
}

func Remove(path string) error {
	err := os.Remove(path)
	if err != nil {
		fmt.Printf("%s❌ :: %sfailed to remove '%s'%s\n",
			string(constants.ColorBlue),
			string(constants.ColorYellow),
			path,
			string(constants.ColorWhite))

		return err
	}

	return err
}

func RemoveAll(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		fmt.Printf("%s❌ :: %sfailed to remove '%s'%s\n",
			string(constants.ColorBlue),
			string(constants.ColorYellow),
			path,
			string(constants.ColorWhite))

		return err
	}

	return err
}

func ExistsMakeDir(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			err = MkdirAll(path)
			if err != nil {
				return err
			}
		} else {
			fmt.Printf("%s❌ :: %sfailed to stat '%s'%s\n",
				string(constants.ColorBlue),
				string(constants.ColorYellow),
				path,
				string(constants.ColorWhite))

			return err
		}

		return err
	}

	return err
}

func Create(path string) (*os.File, error) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Printf("%s❌ :: %sfailed to create '%s'%s\n",
			string(constants.ColorBlue),
			string(constants.ColorYellow),
			path,
			string(constants.ColorWhite))
	}

	return file, err
}

func CreateWrite(path string, data string) error {
	file, err := Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		fmt.Printf("%s❌ :: %sfailed to write to file '%s'%s\n",
			string(constants.ColorBlue),
			string(constants.ColorYellow),
			path,
			string(constants.ColorWhite))

		return err
	}

	return err
}

func Open(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("%s❌ :: %sfailed to open file '%s'%s\n",
			string(constants.ColorBlue),
			string(constants.ColorYellow),
			path,
			string(constants.ColorWhite))
	}

	return file, err
}

func Move(source, dest string) error {
	err := Exec("", "mv", source, dest)

	if err != nil {
		return err
	}

	return err
}

func Copy(dir, source, dest string, presv bool) error {
	args := []string{"-r", "-T", "-f"}

	if presv {
		args = append(args, "-p")
	}

	args = append(args, source, dest)

	err := Exec(dir, "cp", args...)
	if err != nil {
		return err
	}

	return err
}

func CopyFile(dir, source, dest string, presv bool) error {
	args := []string{"-f"}

	if presv {
		args = append(args, "-p")
	}

	args = append(args, source, dest)

	err := Exec(dir, "cp", args...)
	if err != nil {
		return err
	}

	return err
}

func CopyFiles(source, dest string, presv bool) error {
	files, err := ioutil.ReadDir(source)
	if err != nil {
		fmt.Printf("%s❌ :: %sfailed to read dir '%s'%s\n",
			string(constants.ColorBlue),
			string(constants.ColorYellow),
			source,
			string(constants.ColorWhite))

		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		err = CopyFile("", filepath.Join(source, file.Name()), dest, presv)
		if err != nil {
			return err
		}
	}

	return err
}

func FindExt(path, ext string) ([]string, error) {
	var matches []string

	files, err := ioutil.ReadDir(path)

	if err != nil {
		fmt.Printf("%s❌ :: %sfailed to read dir '%s'%s\n",
			string(constants.ColorBlue),
			string(constants.ColorYellow),
			path,
			string(constants.ColorWhite))

		return matches, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if strings.HasSuffix(file.Name(), ext) {
			matches = append(matches, filepath.Join(path, file.Name()))
		}
	}

	return matches, err
}

func FindMatch(path, match string) ([]string, error) {
	var matches []string

	files, err := ioutil.ReadDir(path)

	if err != nil {
		fmt.Printf("%s❌ :: %sfailed to read dir '%s'%s\n",
			string(constants.ColorBlue),
			string(constants.ColorYellow),
			path,
			string(constants.ColorWhite))

		return matches, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if strings.Contains(file.Name(), match) {
			matches = append(matches, filepath.Join(path, file.Name()))
		}
	}

	return matches, err
}

func Filename(path string) string {
	n := strings.LastIndex(path, "/")
	if n == -1 {
		return path
	}

	return path[n+1:]
}

func GetDirSize(path string) (int, error) {
	output, err := ExecOutput("", "du", "-c", "-s", path)
	if err != nil {
		os.Exit(1)
	}

	split := strings.Fields(output)

	size, err := strconv.Atoi(split[len(split)-2])
	if err != nil {
		fmt.Printf("%s❌ :: %sfailed to get dir size '%s'%s\n",
			string(constants.ColorBlue),
			string(constants.ColorYellow),
			path,
			string(constants.ColorWhite))

		return size, err
	}

	return size, err
}

func Exists(path string) (bool, error) {
	exists := false
	_, err := os.Stat(path)

	if err != nil {
		if os.IsNotExist(err) {
			err = nil
		} else {
			fmt.Printf("utils: Exists check error for '%s'\n", path)
			log.Fatal(err)

			return exists, err
		}
	} else {
		exists = true
	}

	return exists, err
}
