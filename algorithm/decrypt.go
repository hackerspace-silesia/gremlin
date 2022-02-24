package algorithm

import (
	"fmt"
	"os"
	"sync"

	"github.com/unidoc/unipdf/v3/model"
)

func DecryptPdf(filePath string, passChannel chan string, wg *sync.WaitGroup) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer f.Close()

	pdfReader, err := model.NewPdfReader(f)
	if err != nil {
		return err
	}

	isEncrypted, err := pdfReader.IsEncrypted()
	if err != nil {
		return err
	}

	if isEncrypted {
		for password := range passChannel {
			auth, err := pdfReader.Decrypt([]byte(password))
			if err != nil {
				return err
			}
			if auth {
				fmt.Println(password)
				wg.Done()
			} else {
				fmt.Errorf("Wrong password")
			}
		}
	}
	return nil
}
