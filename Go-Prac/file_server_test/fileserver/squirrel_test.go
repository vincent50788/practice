package fileserver

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

var testPath = "./test.txt"

func testCfg() Config {
	//conf.Start()
	//logger.Start()
	return Config{
		Host:     "http://squirrel-dev.paradise-soft.com.tw:12345",
		Username: "arch",
		Password: "1qaz!QAZ",
	}
}

func copyZqbYml() {
	os.Remove("./conf.d")
	input, err := ioutil.ReadFile("../../../../conf.d/zqb.yml")
	if err != nil {
		log.Println(err)
		return
	}

	err = os.Mkdir("./conf.d", 0777)
	if err != nil {
		log.Println(err)
		return
	}

	err = ioutil.WriteFile("./conf.d/zqb.yml", input, 0777)
	if err != nil {
		log.Println(err)
		return
	}
}

func createTestFile() (file *os.File, err error) {
	_ = os.Remove(testPath)

	file, err = os.Create(testPath)
	if err != nil {
		return
	}

	_, err = file.WriteString("hello\n")
	if err != nil {
		return
	}
	return
}

func TestNew(t *testing.T) {
	cfg := testCfg()
	fileServer, err := New(cfg)
	if err != nil {
		t.Errorf("New failed=%v", err)
	}
	t.Logf("fileserver=%+v", fileServer)
}

func TestSquirrelClient_Login(t *testing.T) {
	cfg := testCfg()
	fileServer, err := New(cfg)
	if err != nil {
		t.Errorf("New failed=%v", err)
	}
	err = fileServer.Login()
	if err != nil {
		t.Errorf("Login failed=%v", err)
	}
}

func TestSquirrelClient_Upload(t *testing.T) {
	cfg := testCfg()
	fileServer, err := New(cfg)
	if err != nil {
		t.Errorf("New failed=%v", err)
	}
	err = fileServer.Login()
	if err != nil {
		t.Errorf("Login failed=%v", err)
	}

	file, err := createTestFile()
	if err != nil {
		t.Errorf("createTestFile failed=%v", err)
	}
	defer file.Close()

	err = fileServer.Upload("/zqb", "test.txt", file)
	if err != nil {
		t.Errorf("Upload failed=%v", err)
	}
	_ = os.Remove(testPath)
}

func TestSquirrelCli_Delete(t *testing.T) {
	cfg := testCfg()
	fileServer, err := New(cfg)
	if err != nil {
		t.Errorf("New failed=%v", err)
	}
	err = fileServer.Login()
	if err != nil {
		t.Errorf("Login failed=%v", err)
	}

	file, err := createTestFile()
	if err != nil {
		t.Errorf("createTestFile failed=%v", err)
	}
	defer file.Close()

	err = fileServer.Upload("/zqb", "test.txt", file)
	if err != nil {
		t.Errorf("Upload failed=%v", err)
	}

	ok, err := fileServer.Delete("/zqb/test.txt")
	if !ok {
		t.Errorf("Delete failed=%v", err)
	}
	_ = os.Remove(testPath)
}

func TestSquirrelCli_GetFileList(t *testing.T) {
	cfg := testCfg()
	fileServer, err := New(cfg)
	if err != nil {
		t.Errorf("New failed=%v", err)
	}
	err = fileServer.Login()
	if err != nil {
		t.Errorf("Login failed=%v", err)
	}

	file, err := createTestFile()
	if err != nil {
		t.Errorf("createTestFile failed=%v", err)
	}
	defer file.Close()

	err = fileServer.Upload("/zqb", "test.txt", file)
	if err != nil {
		t.Errorf("Upload failed=%v", err)
	}

	info, err := fileServer.GetFileList("/zqb/test.txt")
	if err != nil {
		t.Errorf("GetFileList failed=%v", err)
	}
	t.Logf("info=%+v", info)
	_ = os.Remove(testPath)
}
