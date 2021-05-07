package ezconfig

import (
	"os"
	"testing"
)

/**
 * PROJECT_ROOT 와 같은 환경변수를 설정하지 않을 경우 테스트 코드나 타 경로에서는
 * PROJECT_ROOT에 대한 상대 경로를 인식하지 못하는 오류가 발생할 수 있다.
 * 본 테스트는 프로젝트의 루트에서 수행되기 때문에 문제 없이 /relpath_test.txt를 찾을 수 있지만
 * internal/relpath_issue_test.go에서는 /relpath_test.txt를 찾을 수 없다.
 */
func TestRelPath(t *testing.T) {
	wd, _ := os.Getwd()
	t.Log("현재 working directory:", wd)
	_, err := os.ReadFile("relpath_test.txt")
	if err != nil {
		t.Error("Working directory에 relpath_test.txt가 존재하지 않습니다..")
		t.Fail()
	}
}