package ipfs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

type IPFSAddResponse struct {
	Name string `json:"Name"`
	Hash string `json:"Hash"`
	Size string `json:"Size"`
}

const IPFSBasePath string = "./ipfs"

func uploadFileToIPFS(filePath string) (string, error) {
	err := os.Chdir("./uploads")
	if err != nil {
		fmt.Printf("Error changing directory: %s\n", err)
		// handle error
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		return "", err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Error closing file: %s\n", err)
		}
	}(file)

	// 创建一个 buffer 来存储 multipart 表单数据
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// 创建一个文件字段
	part, err := writer.CreateFormFile("file", filePath)
	if err != nil {
		fmt.Printf("Error creating form file: %s\n", err)
		return "", err
	}

	// 将文件内容复制到表单字段中
	_, err = io.Copy(part, file)
	if err != nil {
		fmt.Printf("Error copying file content: %s\n", err)
		return "", err
	}

	// 关闭 multipart writer 以便写入结尾字段
	err = writer.Close()
	if err != nil {
		fmt.Printf("Error closing multipart writer: %s\n", err)
		return "", err
	}

	// 创建一个 HTTP 请求
	req, err := http.NewRequest("POST", "http://localhost:5001/api/v0/add", &requestBody)
	if err != nil {
		fmt.Printf("Error creating request: %s\n", err)
		return "", err
	}

	// 设置表单内容类型
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %s\n", err)
		return "", err
	}
	defer resp.Body.Close()

	// 读取响应体
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		return "", err
	}
	var ipfsAddResponse IPFSAddResponse
	err = json.Unmarshal(respBody, &ipfsAddResponse)
	if err != nil {
		fmt.Printf("Error unmarshaling response body: %s\n", err)
		return "", err
	}
	fmt.Println(ipfsAddResponse.Hash)
	return ipfsAddResponse.Hash, nil
}

func catFileFromToIPFS(cid string) error {
	// 创建一个 HTTP 请求

	// 基础URL
	baseURL := "http://localhost:5001/api/v0/cat"

	// 使用net/url包构造完整的URL
	u, err := url.Parse(baseURL)
	if err != nil {
		fmt.Printf("Error parsing base URL: %s\n", err)
		return err
	}

	// 设置查询参数
	q := u.Query()
	q.Set("arg", cid)
	u.RawQuery = q.Encode()
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		fmt.Printf("Error creating request: %s\n", err)
		return err
	}

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %s\n", err)
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("Error closing response body: %s\n", err)
		}
	}(resp.Body)

	// 读取响应体
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		return err
	}

	err = os.WriteFile(IPFSBasePath+"/"+cid, respBody, 0644)
	if err != nil {
		fmt.Printf("Error writing file: %s\n", err)
		return err
	}
	return nil
}

// Function arguments don't work,deprecated!
func getFileFromToIPFS(cid string) error {
	// 基础URL
	baseURL := "http://localhost:5001/api/v0/get"

	// 使用net/url包构造完整的URL
	u, err := url.Parse(baseURL)
	if err != nil {
		fmt.Printf("Error parsing base URL: %s\n", err)
		return err
	}

	// 设置查询参数
	q := u.Query()
	q.Set("arg", cid)
	u.RawQuery = q.Encode()
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		fmt.Printf("Error creating request: %s\n", err)
		return err
	}

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %s\n", err)
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("Error closing response body: %s\n", err)
		}
	}(resp.Body)

	// 读取响应体
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		return err
	}

	err = os.WriteFile(IPFSBasePath+"/"+cid, respBody, 0644)
	if err != nil {
		fmt.Printf("Error writing file: %s\n", err)
		return err
	}
	return nil
}

func UploadFileHandler(c *gin.Context) {
	// 从表单中获取文件
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", header.Filename)
	fmt.Printf("File Size: %+v\n", header.Size)
	fmt.Printf("MIME Header: %+v\n", header.Header)

	// 保存文件到指定路径
	dst := filepath.Join("./uploads", header.Filename)
	if err := c.SaveUploadedFile(header, dst); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	fmt.Println("File Upload Success!")

	_, _ = uploadFileToIPFS(header.Filename)

}

func GetFileHandler(c *gin.Context) {
	cid := c.Query("cid")
	if cid == "" {
		c.String(http.StatusBadRequest, "cid is required")
		return
	}
	err := catFileFromToIPFS(cid)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.File(IPFSBasePath + "/" + cid)
}

func CheckFileExistAndGet(c *gin.Context) {
	// 获取请求路径
	fp := c.Query("key")
	// 构建文件完整路径
	fullPath := filepath.Join("./ipfs", fp)

	// 检查文件是否存在
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		// 文件不存在
		err = catFileFromToIPFS(fp)
		if err != nil {
			c.String(http.StatusNotFound, "File Not Found in ipfs")
		}
		c.File(fullPath)
		return
	}

	fmt.Println("File exists:", fullPath)

	c.File(fullPath)
}
