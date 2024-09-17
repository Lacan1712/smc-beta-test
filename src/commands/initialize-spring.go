package commands

import (
	createProject "SpringManagerCLI/src/services"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func InitCommand() {
	var rootCmd = &cobra.Command{
		Use:   "smc",
		Short: "Spring Manager CLI para gerenciar aplicações Java Spring Boot",
		Long:  "Uma ferramenta de linha de comando para gerenciar projetos Spring Boot.",
	}

	// Comando init
	var initCmd = &cobra.Command{
		Use:   "init",
		Short: "Baixa um projeto Spring Boot da API pública",
		Run: func(cmd *cobra.Command, args []string) {
			custom, _ := cmd.Flags().GetBool("custom")

			if custom {
				fmt.Println("Instalação customizada!")

				// Solicitar informações do usuário
				var groupId, artifactId, name, description, packageName, version string

				fmt.Print("Digite o Group ID: ")
				fmt.Scanln(&groupId)

				fmt.Print("Digite o Artifact ID: ")
				fmt.Scanln(&artifactId)

				fmt.Print("Digite o Nome do Projeto: ")
				fmt.Scanln(&name)

				fmt.Print("Digite a Descrição do Projeto: ")
				fmt.Scanln(&description)

				fmt.Print("Digite o Package Name: ")
				fmt.Scanln(&packageName)

				fmt.Print("Digite a Versão")
				fmt.Scanln(&version)

				config := createProject.ProjectConfig{
					GroupId:     groupId,
					ArtifactId:  artifactId,
					Name:        name,
					Description: description,
					PackageName: packageName,
					Version:     version,
				}
				createProject.DownloadSpringBootProject(config)
			} else {
				fmt.Println("Baixando o projeto Spring Boot com configuração padrão...")
				fmt.Print("Baixando projeto spring...")
				config := createProject.ProjectConfig{
					GroupId:     "smc.example",
					ArtifactId:  "A project create by SMC",
					Name:        "smc-demo",
					Description: "A project create by SMC",
					PackageName: "smc.example.demo",
					Version:     "0.0.1-SNAPSHOT",
				}
				createProject.DownloadSpringBootProject(config)
			}
		},
	}

	initCmd.Flags().BoolP("custom", "c", false, "Ativa o comportamento personalizado")

	rootCmd.AddCommand(initCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
