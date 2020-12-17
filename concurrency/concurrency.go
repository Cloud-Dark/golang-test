package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

type queueMuseum struct {
	location string
	museums  []jsonInfoMuseum
}

type jsonDataMuseum struct {
	DataMuseum []jsonInfoMuseum `json:"data"`
}

type jsonInfoMuseum struct {
	ID        string `json:"museum_id"`
	Kode      string `json:"kode_pengelolaan"`
	Nama      string `json:"nama"`
	Sdm       string `json:"sdm"`
	Alamat    string `json:"alamat_jalan"`
	Kelurahan string `json:"desa_kelurahan"`
	Kecamatan string `json:"kecamatan"`
	Kabupaten string `json:"kabupaten_kota"`
	Propinsi  string `json:"propinsi"`
	Lintang   string `json:"lintang"`
	Bujur     string `json:"bujur"`
	Koleksi   string `json:"koleksi"`
	Sumber    string `json:"sumber_dana"`
	Pengelola string `json:"pengelola"`
	Tipe      string `json:"tipe"`
	Standar   string `json:"standar"`
	Tahun     string `json:"tahun_berdiri"`
	Bangunan  string `json:"bangunan"`
	Luas      string `json:"luas_tanah"`
	Pemilik   string `json:"status_kepemilikan"`
}

func main() {

	argBuff := flag.Int("concurrent_limit", 2, "an Int")
	argLoc := flag.String("output", "./museum", "a String directory")
	flag.Parse()

	fmt.Println("BEGIN")
	res, err := http.Get("http://jendela.data.kemdikbud.go.id/api/index.php/CcariMuseum/searchGET?nama=museum")
	if err != nil {
		fmt.Println("http.Get : ", err)
		log.Fatal(err)
	}
	data, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	data = bytes.TrimPrefix(data, []byte("\xef\xbb\xbf"))

	var jsonData jsonDataMuseum

	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		fmt.Println("unmarshal : ", err)
	}

	queues := filterData(jsonData)

	if _, err := os.Stat(*argLoc); os.IsNotExist(err) {
		os.Mkdir(*argLoc, 0700)
	}

	c := make(chan queueMuseum, *argBuff)
	var wg sync.WaitGroup

	go queueChannel(queues, c)

	for queue := range c {
		csvFile, err := os.Create(string(*argLoc + "/" + queue.location + ".csv"))
		if err != nil {
			fmt.Println("os.Create : ", err)
		}
		defer csvFile.Close()

		writer := csv.NewWriter(csvFile)

		titleCSV(writer)

		wg.Add(1)
		go writing(queue, writer, &wg)
		wg.Wait()
	}
	fmt.Println("DONE")
}

func filterData(data jsonDataMuseum) (queues []queueMuseum) {
	make := true
	for _, info := range data.DataMuseum {
		for i, queue := range queues {
			if info.Kabupaten == queue.location {
				queues[i].museums = append(queue.museums, info)
				make = false
				break
			} else {
				make = true
				continue
			}
		}
		if make == true {
			var temp queueMuseum
			temp.location = info.Kabupaten
			temp.museums = append(temp.museums, info)

			queues = append(queues, temp)
		}
	}
	return queues
}

func queueChannel(queues []queueMuseum, c chan queueMuseum) {
	for _, queue := range queues {
		c <- queue
	}
	close(c)
}

func writing(queue queueMuseum, writer *csv.Writer, wg *sync.WaitGroup) {
	defer wg.Done()

	for _, info := range queue.museums {
		// fmt.Println(info.Kabupaten, i)
		buildCSV(info, writer)
	}
	writer.Flush()
}

func buildCSV(data jsonInfoMuseum, file *csv.Writer) {
	var row []string
	row = append(row, data.ID)
	row = append(row, data.Kode)
	row = append(row, data.Nama)
	row = append(row, data.Sdm)
	row = append(row, data.Alamat)
	row = append(row, data.Kelurahan)
	row = append(row, data.Kecamatan)
	row = append(row, data.Kabupaten)
	row = append(row, data.Propinsi)
	row = append(row, data.Lintang)
	row = append(row, data.Bujur)
	row = append(row, data.Koleksi)
	row = append(row, data.Sumber)
	row = append(row, data.Pengelola)
	row = append(row, data.Tipe)
	row = append(row, data.Standar)
	row = append(row, data.Tahun)
	row = append(row, data.Bangunan)
	row = append(row, data.Luas)
	row = append(row, data.Pemilik)
	file.Write(row)
}

func titleCSV(file *csv.Writer) {
	var row []string
	row = append(row, "ID")
	row = append(row, "Kode")
	row = append(row, "Nama")
	row = append(row, "Sdm")
	row = append(row, "Alamat")
	row = append(row, "Kelurahan")
	row = append(row, "Kecamatan")
	row = append(row, "Kabupaten")
	row = append(row, "Propinsi")
	row = append(row, "Lintang")
	row = append(row, "Bujur")
	row = append(row, "Koleksi")
	row = append(row, "Sumber")
	row = append(row, "Pengelola")
	row = append(row, "Tipe")
	row = append(row, "Standar")
	row = append(row, "Tahun")
	row = append(row, "Bangunan")
	row = append(row, "Luas")
	row = append(row, "Pemilik")
	file.Write(row)
}
