package utils

import (
	"AIZO-Projekt/algo"
	"bufio"
	"io"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func ReadFile[T algo.AllowedTypes](path string) (algo.Array[T], error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	r := bufio.NewReader(f)

	var arr *algo.Array[T]

	read, err := r.ReadString('\n')
	if err != nil && err != io.EOF {
		log.Fatalln("Error reading file")
	}
	read = strings.Trim(read, "\r\n")
	atoi, err := strconv.Atoi(read)
	if err != nil {
		return nil, err
	}
	a := algo.NewArray[T](atoi)
	arr = &a

	read, err = r.ReadString('\n')
	read = strings.Trim(read, "\r\n")
	for p := 0; p < atoi; p++ {
		if err != nil && err != io.EOF {
			log.Fatalln("Error reading file")
		}

		t := reflect.ValueOf((*arr)[0])
		switch t.Kind() {
		case reflect.Int64:
			i, err := strconv.ParseInt(read, 10, 64)
			if err != nil {
				return nil, err
			}
			(*arr)[p] = T(i)
			break
		case reflect.Int32:
			i, err := strconv.ParseInt(read, 10, 64)
			if err != nil {
				return nil, err
			}
			(*arr)[p] = T(i)
			break
		case reflect.Int:
			i, err := strconv.ParseInt(read, 10, 64)
			if err != nil {
				return nil, err
			}
			(*arr)[p] = T(i)
			break
		case reflect.Float64:
			i, err := strconv.ParseFloat(read, 64)
			if err != nil {
				return nil, err
			}
			(*arr)[p] = T(i)
			break
		case reflect.Float32:
			i, err := strconv.ParseFloat(read, 64)
			if err != nil {
				return nil, err
			}
			(*arr)[p] = T(i)
			break
		default:
			log.Fatalln("Provided array of wrong type")
		}

		read, err = r.ReadString('\n')
		read = strings.Trim(read, "\r\n")
	}

	return *arr, nil
}
