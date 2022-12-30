package mtproto

import (
	"io"
	"math/rand"
	"os"
	"path/filepath"
)

func MakeInputFileByLocal(
	file string,
	cb func(fileId int64, filePart int32, bytes []byte) error) (*InputFile, error) {

	fi, err := os.Stat(file)
	if err != nil {
		return nil, err
	}

	if fi.Size() > 10240000 {
		//
	}

	//if fi.
	inputFile := MakeTLInputFile(&InputFile{
		Id:          rand.Int63(),
		Parts:       0,
		Name:        "", // file,
		Md5Checksum: "",
	}).To_InputFile()
	_, inputFile.Name = filepath.Split(file)

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data := make([]byte, 512*1024)
	i := 0
	for {
		n, err2 := f.Read(data)
		if err2 != nil {
			if err2 != io.EOF {
				return nil, err2
			} else {
				break
			}
		}

		if err = cb(inputFile.Id, int32(i), data[:n]); err != nil {
			return nil, err
		}

		i += 1
		inputFile.Parts = int32(i)
	}

	return inputFile, nil
}

func MakeInputFile(
	fileName string,
	fileData []byte,
	cb func(fileId int64, filePart int32, bytes []byte) error) (*InputFile, error) {
	file := MakeTLInputFile(&InputFile{
		Id:          rand.Int63(),
		Parts:       0,
		Name:        fileName,
		Md5Checksum: "",
	}).To_InputFile()

	// fileId := rand.Int63()
	for i := 0; i < len(fileData); i = i + 512*1024 {
		if i+512*1024 >= len(fileData) {
			if err := cb(file.Id, file.Parts, fileData[i:]); err != nil {
				return nil, err
			}
		} else {
			if err := cb(file.Id, file.Parts, fileData[i:i+512*1024]); err != nil {
				return nil, err
			}
		}
		file.Parts += 1
	}

	return file, nil
}
