package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	fmt.Printf("*****INICIO: log_commits.go *******\n")
	// --------ini---- eea
	logDirxx := "" //  filepath.Join("..", "log")
	logFilexx := filepath.Join(logDirxx, "commits_xx.txt")
	contingutxx := "--CONTENT--"
	resultadoxx := os.WriteFile(logFilexx, []byte(contingutxx), 0644)
	fmt.Printf("****Resultado de WriteFile(): %v - %f\n", resultadoxx, logFilexx)
	fmt.Printf("****fin mi test, que grabo un arch.\n")
	// --------eea---- fin
	
	cmd := exec.Command("git", "log", "-n", "3", "--pretty=format:%h - %an, %ar : %s")
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error ejecutando git log: %v\n", err)
		os.Exit(1)
	}

  	// Creacio de la carpeta log
	logDir := filepath.Join("..", "log")
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		//Definimos permisos de escritura
		err = os.Mkdir(logDir, 0755)
		if err != nil {
			fmt.Printf("Error creando directorio %s: %v\n", logDir, err)
			os.Exit(1)
		}
	}

	//Generar nom del archivo
	currentTime := time.Now().Format("2006-01-02_15-04-05") //Mascara de YYYY-mm-dd HH-ii-ss
	logFile := filepath.Join(logDir, fmt.Sprintf("commits_%s.txt",currentTime))

	//Escriure l'arxiu
	contingut := fmt.Sprintf("Se han de ecribir los ultimos 3 commits del reposuitorio:\n%s",string(out))
	err = os.WriteFile(logFile, []byte(contingut), 0644)
	if err != nil {
		fmt.Printf("Se ha producido un error creando en %s: %v\n",logFile, err)
		os.Exit(1)
	}                         
	
	fmt.Printf("Se ha creado el archivo de log en: %s\n", logFile)     
	fmt.Printf("*****FIN: log_commits.go *******\n")
}
