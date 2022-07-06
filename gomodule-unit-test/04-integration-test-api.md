## Integration Test

Melakukan integration test untuk API setidaknya kita harus bisa menjalankan aplikasi-nya terlebih dahulu agar bisa dilakukan pengetesan secara terintegrasi.
Hal ini perlu kita siapkan beberapa case, test case yang mencakup kebutuhan dari integration Test tersebut. Misalkan dari `API Endpoint` kita yang sudah dikerjakan itu memiliki resource `database`, `cache` ataupun eksternal lain yang berhubungan dengan keberlangsungannya suatu API Endpoint tersebut.

Maka, kita akan simulasikan bagaimana cara melakukan integration test API. Tahapan yang pertama yaitu kita buat API Service yang mana API tersebut digunakan untuk operasi penjumlahan angka.

## Membuat Program API
Berikut program API Endpoint yang akan diuji atau dibuat integration test-nya.
```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/add", HandleAddInts)
	address := ":9000"
	fmt.Printf("Memulai server at %s\n", address)
	fmt.Println()
	fmt.Println("contoh query: /add?a=2&b=2&authtoken=abcdef123")
	fmt.Println("tokennya adalah 'abcdef123'")
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Printf("error ketika start server: %v", err)
	}
}

type TambahResponse struct {
	Result int `json:"result"`
}

func HandleAddInts(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	paramA := params.Get("a")
	paramB := params.Get("b")
	token := params.Get("authtoken")

	if paramA == "" || paramB == "" || token == "" {
		http.Error(w, "tidak ada parameters", http.StatusBadRequest)
		return
	}

	if token != "abcdef123" {
		http.Error(w, "token tidak valid", http.StatusUnauthorized)
		return
	}

	intA, err := strconv.Atoi(paramA)
	if err != nil {
		http.Error(w, "parameter 'a' harus integer", http.StatusBadRequest)
		return
	}

	intB, err := strconv.Atoi(paramB)
	if err != nil {
		http.Error(w, "parameter 'b' harus integer", http.StatusBadRequest)
		return
	}

	response := TambahResponse{
		Result: intA + intB,
	}

	json, err := json.MarshalIndent(&response, "", " ")
	if err != nil {
		http.Error(w, "error while marshalling", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(json))
}
```
Di dalam program ini ada beberapa validasi jika saat program dijalankan, yaitu * Parameter yg diinput wajib `integer` 
* Tidak mengisi parameter `a` dan `b` maka akan menjadi 
* Mengujian jika tidak mengirimkan `authToken`

Setelah program diatas selesai maka kita jalankan dengan perintah
```bash
➜  integration-test git:(main) ✗ go run app/main.go
Memulai server at :9000

contoh query: /add?a=2&b=2&authtoken=abcdef123
tokennya adalah 'abcdef123'
```
Jika program sudah sukses berjalan maka akan terlihat atau bisa diakses menggunakan postman dan `curl` dengan endpoint dan port `:9000`.
```bash
➜  materi curl 'http://localhost:9000/add?a=2&b=2&authtoken=abcdef123'
{
 "result": 4
}%   
```

Jika program berhasil seperti diatas maka, kita sudah menjalankan program API untuk menambah angka menggunakan protokol `http`.

# Membuat Test Integration
Selanjutnya, kita kan membuat test integration API yg tadi sudah kita buat itu dengan beberapa case yang sudah disebutkan diatas.

Siapkan program atau file `main.go` untuk membuat API Testing.
```go
package integrationtest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type MathClient struct {
	Token string
	Host  string
}

type AddResult struct {
	Result int `json:"result"`
}

func (c *MathClient) APISum(i, j int) (int, error) {
	query := fmt.Sprintf("http://%s/add?a=%v&b=%v&authtoken=%v", c.Host, i, j, c.Token)

	response, err := http.Get(query)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}

	a := AddResult{}
	err = json.Unmarshal(data, &a)
	if err != nil {
		return 0, err
	}

	return a.Result, nil
}
```

Method `APISum` ini akan kita gunakan untuk mengakses `http` API yang sudah kita buat sebagai main server yang akan kita test menggunakan integration test ini.

## Membuat Struct Client
Buat struct untuk struktur API dan responsenya seperti ini.
```go
type MathClient struct {
	Token string
	Host  string
}

type AddResult struct {
	Result int `json:"result"`
}
```
## Membuat Method Akses API Sum
lalu kita akan buat method untuk ambil data penjumlahan API.
```go
func (c *MathClient) APISum(i, j string) (int, int, error) {
	query := fmt.Sprintf("http://%s/add?a=%s&b=%s&authtoken=%v", c.Host, i, j, c.Token)

	var statusCode int = http.StatusBadRequest
	response, err := http.Get(query)
	if err != nil {
		return 0, statusCode, err
	}
	defer response.Body.Close()
	statusCode = response.StatusCode
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, statusCode, err
	}

	a := AddResult{}
	err = json.Unmarshal(data, &a)
	if err != nil {
		return 0, statusCode, err
	}

	return a.Result, statusCode, nil
}
```
## Membuat Unit Test yang dijadikan Integration Test
Selanjutnya, kita akan membuat API Integration testing ini di method yang ini. Bisa kita buat program seperti ini atau `generate` saja pada method `APISum`.
```go
func TestMathClient_APISum(t *testing.T) {
	type fields struct {
		Token string
		Host  string
	}
	type args struct {
		i string
		j string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     int
		wantErr  bool
		wantCode int
	}{
		// testing test case
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &MathClient{
				Token: tt.fields.Token,
				Host:  tt.fields.Host,
			}
			got, gotCode, err := c.APISum(tt.args.i, tt.args.j)
			if (err != nil) != tt.wantErr {
				t.Errorf("MathClient.APISum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MathClient.APISum() = %v, want %v", gotCode, tt.wantCode)
			}
			if gotCode != tt.wantCode {
				t.Errorf("MathClient.APISum() = %v, want %v", gotCode, tt.wantCode)
			}
		})
	}
}

```
## Tambahkan Test Case
Lalu agar kita bisa melakukan pengujian `case by case` kita perlu isi berapa tesitng case pada kurung kurawal diatas dengan isi seperti ini.
```go
{
  name: "case sukses jumlah data",
  fields: fields{
    Token: "abcdef123",
    Host:  "localhost:9000",
  },
  args: args{
    i: "2",
    j: "2",
  },
  want:     4,
  wantErr:  false,
  wantCode: http.StatusOK,
},
{
  name: "token tidak diset",
  fields: fields{
    Token: "abc",
    Host:  "localhost:9000",
  },
  args: args{
    i: "2",
    j: "2",
  },
  want:     0,
  wantErr:  true,
  wantCode: http.StatusUnauthorized,
},
{
  name: "host tidak sesuai",
  fields: fields{
    Token: "abcdef123",
    Host:  "localhost:500",
  },
  args: args{
    i: "2",
    j: "2",
  },
  want:     0,
  wantErr:  true,
  wantCode: http.StatusBadRequest,
},
{
  name: "case parameter kosong",
  fields: fields{
    Token: "abcdef123",
    Host:  "localhost:9000",
  },
  args: args{
    i: "",
    j: "",
  },
  want:     0,
  wantErr:  true,
  wantCode: http.StatusBadRequest,
},
{
  name: "case parameter bukan integer",
  fields: fields{
    Token: "abcdef123",
    Host:  "localhost:9000",
  },
  args: args{
    i: "a",
    j: "a",
  },
  want:     0,
  wantErr:  true,
  wantCode: http.StatusBadRequest,
},
```
## Jalankan Integration Test
Semua sudah selesai untuk pengujian dan pembuatan testing jangan lupa karena ini adalah code `integration test` maka kita perlu ada penambahan agar saat melakukan `go test -v` bisa kita tambahkan code diatas seperti ini.
```go
//go:build integration
```
Kita lanjut dengan menjalankan go test dengan parameter `integration`
```bash
$ go test -v -tags=integration
```
Pastikan juga program API Endpoint berjalan dengan
```bash
➜  integration-test git:(main) ✗ go run app/main.go
Memulai server at :9000

contoh query: /add?a=2&b=2&authtoken=abcdef123
tokennya adalah 'abcdef123'
```