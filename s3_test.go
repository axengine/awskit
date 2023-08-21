package awskit

import (
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
	"testing"
)

func TestAWSKit_UploadBuf(t *testing.T) {
	cli, err := New("default", "ap-northeast-1")
	if err != nil {
		t.Fatal(err)
	}
	f, _ := os.Open("1.jfif")
	defer f.Close()
	bz, err := io.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}
	contentType := http.DetectContentType(bz)
	objectKey := "avatar/" + uuid.New().String()
	fmt.Println(objectKey)
	fmt.Println(contentType)
	if err := cli.UploadBuf("EXAMPLE-BUCKET-NAME", objectKey, bz, int64(len(bz)), contentType); err != nil {
		t.Fatal(err)
	}
}

func TestAWSKit_DeleteOBJ(t *testing.T) {
	cli, err := New("default", "ap-northeast-1")
	if err != nil {
		t.Fatal(err)
	}
	if err := cli.DeleteOBJ("EXAMPLE-BUCKET-NAME", "avatar/5e0cbd19-5f29-4d0d-9d15-6294044af8d0"); err != nil {
		t.Fatal(err)
	}
}
