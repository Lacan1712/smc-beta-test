package services

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

type ProjectConfig struct {
	GroupId     string
	ArtifactId  string
	Name        string
	Description string
	PackageName string
	Version     string
}

func DownloadSpringBootProject(config ProjectConfig) {

	baseURL := "https://start.spring.io/starter.zip"
	params := url.Values{}
	params.Add("type", "maven-project")
	params.Add("groupId", config.GroupId)
	params.Add("artifactId", config.ArtifactId)
	params.Add("name", config.Name)
	params.Add("description", config.Description)
	params.Add("packageName", config.PackageName)
	params.Add("version", config.Version)

	url := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Erro ao fazer a requisição:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Erro ao baixar o projeto, código de status: %d\n", resp.StatusCode)
		return
	}

	outFile, err := os.Create("demo.zip")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		fmt.Println("Erro ao salvar o arquivo:", err)
		return
	}

	fmt.Println("Projeto Spring Boot baixado com sucesso como demo.zip!")

	if err := Unzip("demo.zip", "./demo"); err != nil {
		fmt.Println("Erro ao descompactar o arquivo:", err)
		return
	}

	if err := os.Remove("demo.zip"); err != nil {
		fmt.Println("Erro ao excluir o arquivo:", err)
		return
	}

	fmt.Println("Arquivo demo.zip excluído com sucesso!")
}

// Unzip descompacta um arquivo ZIP em um diretório de destino
func Unzip(src string, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	if err := os.MkdirAll(dest, 0755); err != nil {
		return err
	}

	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(fpath, f.Mode()); err != nil {
				return err
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fpath), 0755); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fpath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}

		_, err = io.Copy(outFile, rc)
		if err != nil {
			return err
		}

		outFile.Close()
		rc.Close()
	}

	return nil
}
