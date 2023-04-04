package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func ReadFileToString(filePath string) (string, error) {
	result := ""
	// "./assets/template/html/register-company/1_response_register.html"

	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Printf("failed reading data from file: %s", err)
		return result, err
	}
	// fmt.Printf("\nLength: %d bytes", len(data))
	// fmt.Printf("\nData: %s", data)
	// fmt.Printf("\nError: %v", err)

	result = string(data)

	return result, nil
}

func GetFileListFromHeader(c echo.Context, fieldFile string) ([]*multipart.FileHeader, error) {
	form, err := c.MultipartForm()
	if err != nil {
		if err.Error() != "no multipart boundary param in Content-Type" {
			log.Println(err.Error())
			return nil, err
		}
		log.Println(err.Error())
		return nil, err
	}

	files := form.File[fieldFile]

	return files, nil
}

type FileTypeInSystem struct {
	FileTypeName string
}

type FileTypeInSystemList []FileTypeInSystem

func (fileTypeInSystemList *FileTypeInSystemList) SetFileTypeNameList() {
	var FileTypeDataList = [][]interface{}{
		//FileTypeName
		{"image"},
		{"pdf"},
		{"excel"},
	}

	for _, data := range FileTypeDataList {
		fileTypeInSystem := new(FileTypeInSystem)
		fileTypeInSystem.FileTypeName = data[0].(string)
		*fileTypeInSystemList = append(*fileTypeInSystemList, *fileTypeInSystem)
	}
}

func CheckFileTypeInSystem(fileType, fileTypeReq string) error {
	fileTypeInSystemList := new(FileTypeInSystemList)
	fileTypeInSystemList.SetFileTypeNameList()
	checkHave := false
	for _, data := range *fileTypeInSystemList {
		if strings.Contains(fileTypeReq, data.FileTypeName) {
			switch data.FileTypeName {
			case "image":
				if fileType == "image/png" || fileType == "image/jpeg" || fileType == "image/jpg" {
					checkHave = true
				}
			case "pdf":
				if fileType == "application/pdf" {
					checkHave = true
				}
			case "excel":
				if fileType == "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" {
					checkHave = true
				}
			default:
				message := "file type not support in system"
				log.Println(message)
				errMessage := errors.New(message)
				return errMessage
			}
		}
	}

	if !checkHave {
		message := fmt.Sprintf("file type %s (%s) not support in: ", fileType, fileTypeReq)
		log.Println(message)
		errNew := errors.New(message)
		return errNew
	}
	return nil
}

func FileUrlDownload(pathFile, urlDownload, contentType string) (*http.Response, error) {
	// อย่าลืมปิด .Close()
	// ที่ไม่ปิดตรงนี้เพราะฟังก์ชั่นอื่นจะมีการอ่านไฟล์อยู่
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			r.Close = true
			return nil
		},
	}

	resp, err := client.Get(urlDownload)
	if err != nil {
		log.Println(err.Error())
		return resp, err
	}

	return resp, nil

}
