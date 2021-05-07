package internal

import (
	"os"
	"path"
	"testing"
)


// 현재 working directory는 Project root가 아니기 때문에
// 그냥 상대경로를 이용하면 fail
func TestRelPath(t *testing.T) {
	wd, _ := os.Getwd()
	t.Log("현재 working directory:", wd)
	_, err := os.ReadFile("relpath_test.txt")
	if err == nil {
		t.Error("Working directory에 relpath_test.txt가 존재합니다.")
		t.Fail()
	}
}

func TestRelPathWithProjectRootEnv(t *testing.T) {
	root, exists := os.LookupEnv("EZ_PROJECT_ROOT")
	if !exists {
		t.Error("EZ_PROJECT_ROOT 환경변수를 정의해주세요.")
		t.Fail()
	}

	wd, _ := os.Getwd()
	t.Log("현재 working directory:", wd)
	_, err := os.ReadFile(path.Join(root, "relpath_test.txt"))

	if err != nil {
		t.Error("$EZ_PROJECT_ROOT에 relpath_test.txt가 존재하지 않습니다.")
		t.Fail()
	}
}