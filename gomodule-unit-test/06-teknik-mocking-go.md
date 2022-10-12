# Teknik Membuat Mocking pada Golang

Saat kita membuat fungsi atau sedang melakukan code terkadang kita suka kesusahan untuk melakukan unit test di beberapa titik yang tidak bisa kita cover dengan unit test.
Maka berikut ini ada beberapa cara teknik melakukan unit test dengan teknik mocking. Tetapi sebenarnya kita juga bisa menggunakan Third-party yang sudah tersedia dibeberapa library sehingga tinggal langsung kita pakai saja.

Nah, kekurangannya yaitu ketika kita menggunakan Third-party kita tidak mengerti secara menyeluruh proses dari unit test Third-party tersebut untuk melakukan covering unit test. Maka, kita perlu juga tahu cara bagaimana melakukan mocking agar bisa melihat alur dari proses code yg kita jalankan.

##  Higher-Order Functions
Misalkan kita memiliki fungsi untuk melakukan koneksi ke dalam database SQL seperti dibawah ini.
```go
func OpenDB(user, password, addr, db string) (*sql.DB, error) {
	conn := fmt.Sprintf("%s:%s@%s/%s", user, password, addr, db)
	sql, err :=sql.Open("mysql", conn)
    if err != nil {
        log.Error("error open connection mysql")
    }
    return sql, nil  
}
```
agar kita bisa melakukan test pada fungsi dari `sql.Open` maka kita perlu melakukan perubahan dari sisi code kita yaitu dengan cara `mock` fungsi tersebut menjadi tipe fungsi. Agar lebih mudah bisa dilihat implementasinya dibawah ini.
```go
type (
	sqlOpener func(string, string) (*sql.DB, error)
)

func OpenDB(user, password, addr, db string, open sqlOpener) (*sql.DB, error) {
	conn := fmt.Sprintf("%s:%s@%s/%s", user, password, addr, db)
	sql, err := open("mysql", conn)
    if err != nil {
        log.Error("error open connection mysql")
    }
    return sql, nil
}
```

Pada tipe `sqlOpener`, kita akan mock fungsi tersebut untuk kebutuhan unit test nantinya sehingga bisa kita membuat test cases untuk terjadi `error` dan sukses.

Ketika memanggil fungsi `OpenDB` ini kita perlu mengirimkan fungsi `sql.Open` agar bisa *provide* sesuai dengan fungsi utamanya. Agar lebih bisa terbayang bagaimana impelementasinya bisa kita lihat code dibawah ini.

```go
    OpenDB("myUser", "myPass", "localhost", "foo", sql.Open)
```

Lalu bagaimana cara kita membuat unit test-nya? Berikut ini coba simak dan lihat impelemntasinya dibawah ini.
```go
func TestOpenDB(t *testing.T) {
	type args struct {
		user     string
		password string
		addr     string
		db       string
		open     func(string, string) (*sql.DB, error)
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case 1 : success open connection database mysql",
			args: args{
				user:     "myUser",
				password: "myPass",
				addr:     "localhost",
				db:       "foo",
				open: func(s1, s2 string) (*sql.DB, error) {
					return &sql.DB{}, nil
				},
			},
			wantErr: false,
		},
		{
			name: "case 2: failed open connection because have error",
			args: args{
				user:     "myUser",
				password: "myPass",
				addr:     "localhost",
				db:       "foo",
				open: func(s1, s2 string) (*sql.DB, error) {
					return nil, errors.New("got error")
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := OpenDB(tt.args.user, tt.args.password, tt.args.addr, tt.args.db, tt.args.open)
			if (err != nil) != tt.wantErr {
				t.Errorf("OpenDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
```
Metode ini perlu kita perhatikan ketika kita membuat mock untuk fungsi aslinya, karena bisa jadi ketika kita akan melakukan upgrade dependency tersebut lalu ada beberapa parameter yang berubah atau ada penambahan maka kita perlu juga merubah semua variabel fungsi yang dibuat `mock` tersebut agar bisa menyesuaikan fungsinya.

## Monkey Patching

Teknik ini hampir sama dengan teknik `mock` **Higher-Order Functions** bahkan mirip sekali dengan teknik tersebut yaitu kita akan menjadikan fungsi utama yg akan dipanggil `sql.Open` ini menjadi variabe global.

Alih-alih meneruskan fungsi ke `OpenDB()`, kami hanya menggunakan variabel untuk panggilan yang sebenarnya. Berikut ini implementasi dalam code-nya.

```go
var (
	SQLOpen = sql.Open
)

func OpenDB(user, password, addr, db string) (*sql.DB, error) {
	conn := fmt.Sprintf("%s:%s@%s/%s", user, password, addr, db)
	sql, err := SQLOpen("mysql", conn)
	if err != nil {
		log.Print("error open connection mysql")
		return sql, err
	}

	return sql, nil
}

```

Perbedaanya itu hanya dari tipe data yang dilakukan yaitu inisialisasi variabel untuk teknik ini. Lalu, bagaimana cara mock di dalam unit test-nya? Dibawah ini kita akan jelaskan.

```go
for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        SQLOpen = tt.args.open
        _, err := OpenDB(tt.args.user, tt.args.password, tt.args.addr, tt.args.db)
        if (err != nil) != tt.wantErr {
            t.Errorf("OpenDB() error = %v, wantErr %v", err, tt.wantErr)
            return
        }
    })
}
```
Dan di bagian unit test-nya pun perbedaannya yaitu kita melakukan `assign` ke dalam variabel `SQLOpen` tersebut yang di mock dari tiap test case sehingga bisa menggambarkan case `error` atau sukses.

Terkadang teknik ini juga bukan cara terbaik untuk memperbaiki unit test coverage karena perlu di pastikan variabel tersebut publik sehingga bisa di panggil oleh fungsi utamanya.

Ingat! Teknik ini sama dengan teknik sebelumnya maka perlu hati-hati dalam menggunakannya ketika third-party akan kita upgrade maka perlu dipastikan lg mock fungsinya harus di sesuaikan lagi jika ada perubahan pada dependency aslinya.

## Interface Substitution
Teknik ini kita gunakan untuk tipe fungsi yang interface atau concrete. Di dalamn bahasa Go, kita bisa melakukan teknik ini dengan memiliki fungsi `interface` sehingga tidak perlu secara implisit menjadikan implementasi fungsi tersebut.

Terkadang kita perlu melakukan ini `interface` agar mengurangi jangkauan unit test yang akan kita uji. Misalkan kita contohkan membuat fungsi untuk mengambil data dari file seperti dibawah ini.

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("foo.txt")
	if err != nil {
		fmt.Printf("error opening file %v \n", err)
	}
	data, err := ReadContents(f, 50)
	if err != nil {
		fmt.Printf("error from ReadContents %v \n", err)
	}
	fmt.Printf("data from file: %s", string(data))
}

func ReadContents(f *os.File, numBytes int) ([]byte, error) {
	defer f.Close()
	data := make([]byte, numBytes)
	_, err := f.Read(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
```

Kita perlu meniru fungsi yang ada di dalam `os.File` yaitu kita gunakan pada fungsi `ReadContents`.Secara khusus kita menggunakan fungsi `f.Read(data)` untuk membaca data dari file dan diakhiri dengan kita menutup file dengan `defer f.Clode()`

Dengan begitu kita akan membuat `mock` dari `os.File` tersebut yang mana ini merupakan package standar library IO dari Golang. bisa dilihat dibawah ini
```go
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

// ReadCloser is the interface that groups the basic Read and Close methods.
type ReadCloser interface {
	Reader
	Closer
}

```
Karena `os.File` ini merupakan implementasinya dari `io` library, maka kita bisa ganti fungsi `ReadContents` tersebut menjadi seperti dibawah ini.

```go
func ReadContents(rc io.ReadCloser, numBytes int) ([]byte, error) {
    defer rc.Close()
    data := make([]byte, numBytes)
    _, err := rc.Read(data)
    if err != nil {
        return nil, err
    }
    return data, nil
}
```

Dalam kebanyakan kasus, Kita mungkin perlu membuat `interface` sendiri, tetapi disini kita dapat menggunakan kembali `interface` yang ditentukan dalam package `io`. Sekarang kita coba membuat unit test dengan mudah menggunakan `mock`.

```go
package main

import (
	"errors"
	"io"
	"reflect"
	"testing"
)

type (
	mockReadCloser struct {
		expectedData []byte
		expectedErr  error
	}
)

func (mrc *mockReadCloser) Read(p []byte) (n int, err error) {
	copy(p, mrc.expectedData)
	return 0, mrc.expectedErr
}

func (mrc *mockReadCloser) Close() error { return nil }

func TestReadContents(t *testing.T) {
	errorz := errors.New("got error")
	type args struct {
		rc       io.ReadCloser
		numBytes int
	}
	tests := []struct {
		name         string
		args         args
		expectedData []byte
		expectedErr  error
	}{
		{
			name: "case success getting data read",
			args: args{
				rc: &mockReadCloser{
					expectedData: []byte(`hello`),
					expectedErr:  nil,
				},
				numBytes: 5,
			},
			expectedData: []byte(`hello`),
			expectedErr:  nil,
		},
		{
			name: "case failed getting data read",
			args: args{
				rc: &mockReadCloser{
					expectedData: []byte(`hello`),
					expectedErr:  errorz,
				},
				numBytes: 5,
			},
			expectedData: nil,
			expectedErr:  errorz,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadContents(tt.args.rc, tt.args.numBytes)
			if !reflect.DeepEqual(got, tt.expectedData) {
				t.Errorf("expected (%b), got (%b)", tt.expectedData, got)
			}
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("expected error (%v), got error (%v)", tt.expectedErr, err)
			}
		})
	}
}
```
Perlu diperhatikan bahwa `struct mockReadCloser` merupakan mock dari interface-nya, dengan cara ini, setiap pengujian dapat membuat struct dan mengembalikan nilai sesuai dengan keinginan.

## Embedding Interfaces

Embedding Interface ini adalah teknik mocking menggunakan embedded fungsi interface yang kita buat seolah-olah implementasinya sesuai dengan ekspektasi kita. Disini kita menggunakan [SDK Library AWS](https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/dynamodbiface/) yang mana ini bisa kita gunakan untuk melakukan unit testing.

Berikut ini contoh kode misalkan kita menggunakan aws dynamodb untuk mengambil data batch item.
```go
package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

func main() {
	sess := session.New()
	svc := dynamodb.New(sess)

	GetBatchItem(svc, &dynamodb.BatchGetItemInput{
		RequestItems: map[string]*dynamodb.KeysAndAttributes{
			"a": &dynamodb.KeysAndAttributes{
				AttributesToGet: []*string{},
			},
		},
	})
}

func GetBatchItem(svc dynamodbiface.DynamoDBAPI, input *dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, error) {
	batch, err := svc.BatchGetItem(input)
	if err != nil {
		log.Printf("error")
		return nil, err
	}

	return batch, nil
}
```

Unit test lengkapnya seperti ini dimana kita membuat `mockDynamoDBClient` struct berisi interface `dynamodbiface.DynamoDBAPI` yang memiliki beberapa method. Yang kita mock hanya method yg kita butuhkan saja yaitu method `BatchGetItem` jadi tidak perlu semuanya kita implementasikan.

```go
package main

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type mockDynamoDBClient struct {
	dynamodbiface.DynamoDBAPI
}

func (m *mockDynamoDBClient) BatchGetItem(d *dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, error) {
	if len(d.RequestItems) == 0 {
		return nil, errors.New("got error")
	}
	return &dynamodb.BatchGetItemOutput{
		Responses: map[string][]map[string]*dynamodb.AttributeValue{},
	}, nil
}

func TestGetBatchItem(t *testing.T) {
	type args struct {
		svc   dynamodbiface.DynamoDBAPI
		input *dynamodb.BatchGetItemInput
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success get batch items",
			args: args{
				svc: &mockDynamoDBClient{},
				input: &dynamodb.BatchGetItemInput{
					RequestItems: map[string]*dynamodb.KeysAndAttributes{
						"a": {
							AttributesToGet: []*string{},
						},
					},
				},
			},
		},
		{
			name: "failed get batch items",
			args: args{
				svc:   &mockDynamoDBClient{},
				input: &dynamodb.BatchGetItemInput{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetBatchItem(tt.args.svc, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBatchItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

```

## Mocking out Downstream HTTP Calls
Membuat mock untuk eksternal call HTTP memang agak tricky jika kita ingin mengimplemenasikannya. tetapi dengan teknik ini kita bisa secara lengkap bisa meng-cover semua case yang akan kita buat.

Misalkan kita memiliki fungsi yang mana akan mengakses Rest API eksternal, lebih lengkapnya sebagai berikut.
```go
type Response struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
}

func MakeHTTPCall(url string) (*Response, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    r := &Response{}
    if err := json.Unmarshal(body, r); err != nil {
        return nil, err
    }
    return r, nil
}
```

Nah lalu bagaimana caranya agar bisa kita buat unit test-nya?
Ini biasnaya kita menggunakan `httptest` library standar-nya dari Golang yang nantinya seolah-olah bisa membuat API external dengan response yang disesuaikan.

Lebih lengkapnya yuk kita coba langkah-langkahnya sebagai berikut.
* Arahkan kursor pada fungsi `MakeHTTPCall` lalu klik kanan dan pilih `Go: Generate Unit Tests For Function`, maka akan dilakukan generate code default unit test seperti ini.
    ```go
    func TestMakeHTTPCall(t *testing.T) {
        type args struct {
            url string
        }
        tests := []struct {
            name    string
            args    args
            want    *Response
            wantErr bool
        }{
            // TODO: Add test cases.
        }
        for _, tt := range tests {
            t.Run(tt.name, func(t *testing.T) {
                got, err := MakeHTTPCall(tt.args.url)
                if (err != nil) != tt.wantErr {
                    t.Errorf("MakeHTTPCall() error = %v, wantErr %v", err, tt.wantErr)
                    return
                }
                if !reflect.DeepEqual(got, tt.want) {
                    t.Errorf("MakeHTTPCall() = %v, want %v", got, tt.want)
                }
            })
        }
    }
    ```

* Selanjutnya tambahkan di dalam struct `args` variable ini `server *httptest.Server` yang berfungsi untuk membuat mock data http external call.
* Lalu dibawah sebelum pemanggilan `MakeHTTPCall` perlu ada yg diupdate seperti ini
    ```go
    defer tt.args.server.Close()
    var url string
    if tt.args.url == "" {
        url = tt.args.server.URL
    }
    got, err := MakeHTTPCall(url)
    ```
    Info:
    * `defer tt.args.server.Close()` dimaksudkan agar tiap `NewServer` testing kita perlu close server-nya agar tidak bentrok.
    * `var url string` ini digunakan untuk kondisi pengecekan url yang benar dan salah

* Terakhir kita tambahkan test case yang perlu kita butuhkan sesuai dengan code yang kita buat.
    ```go
    {
        name: "success call http",
        args: args{
            server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                w.WriteHeader(http.StatusOK)
                w.Write([]byte(`{"id": 1, "name": "santekno", "description": "santekno jaya"}`))
            })),
        },
        want: &Response{
            ID:          1,
            Name:        "santekno",
            Description: "santekno jaya",
        },
        wantErr: false,
    },
    {
        name: "failed call http when http 400",
        args: args{
            server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                w.WriteHeader(http.StatusBadRequest)
            })),
        },
        want:    nil,
        wantErr: true,
    },
    {
        name: "failed url http call",
        args: args{
            url: "localhost",
            server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                w.WriteHeader(http.StatusBadRequest)
            })),
        },
        want:    nil,
        wantErr: true,
    },
    ```

Semua sudah terisi tinggal kita coba jalankan apakah tiap test case tersebut sudah meng-cover semua code kita atau belum.

Ingin tahu lebih lengkap code unit test-nya?
Ini nih kita informasikan lebih detail lagi yaa
```go
func TestMakeHTTPCall(t *testing.T) {
	type args struct {
		url    string
		server *httptest.Server
	}
	tests := []struct {
		name    string
		args    args
		want    *Response
		wantErr bool
	}{
		{
			name: "success call http",
			args: args{
				server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(`{"id": 1, "name": "santekno", "description": "santekno jaya"}`))
				})),
			},
			want: &Response{
				ID:          1,
				Name:        "santekno",
				Description: "santekno jaya",
			},
			wantErr: false,
		},
		{
			name: "failed call http when http 400",
			args: args{
				server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusBadRequest)
				})),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failed url http call",
			args: args{
				url: "localhost",
				server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusBadRequest)
				})),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer tt.args.server.Close()
			var url string
			if tt.args.url == "" {
				url = tt.args.server.URL
			}
			got, err := MakeHTTPCall(url)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeHTTPCall() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeHTTPCall() = %v, want %v", got, tt.want)
			}
		})
	}
}
```

## Conclusion
Jika kita tidak mempelajari secara manual, kita tidak akan tahu alur dari uni test itu bekerja maka diharapkan sebelum kita menggunakan Third-Party yang mendukung pemenuhan unit test, ada baiknya kita juga perlu tahu bagaimana mekanisme didalamnya.

Tujuan dari unit test sebenarnya kan menguji code kita apakah sesuai dengan kebutuhan bisnis,produk yang sedang kita kembangkan dan agar minim bug saat kita jalankan di production (live).
