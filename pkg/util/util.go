package util

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"

	"k8s.io/apimachinery/pkg/types"
)

func MergeArrays(left []string, right []string) []string {
	for _, item := range right {
		left = append(left, item)
	}
	return left
}

func Contains(target string, array []string) bool {
	sort.Strings(array)
	index := sort.SearchStrings(array, target)
	if index < len(array) && array[index] == target {
		return true
	}
	return false
}

func String2NamespacedName(s string) types.NamespacedName {

	split := strings.Split(s, "/")
	if len(split) != 2 {
		return types.NamespacedName{
			Namespace: "",
			Name:      s,
		}
	}
	return types.NamespacedName{
		Namespace: split[0],
		Name:      split[1],
	}
}

func WriteToFile(fileName string, content string) error {

	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("file create failed. err: " + err.Error())
	} else {
		n, _ := f.Seek(0, os.SEEK_END)
		_, err = f.WriteAt([]byte(content), n)
		defer f.Close()
	}
	return err
}

func ExecuteCommand(cmdStr string, arg string) error {
	commandStr := fmt.Sprintf(cmdStr, arg)
	out, err := exec.Command("/bin/sh", "-c", commandStr).CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to execute cmd: %v, error: %v", string(out), err)
	}

	return nil
}
