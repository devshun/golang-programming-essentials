package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"regexp"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestFindEntries(t *testing.T) {

	// mockを定義
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.String())
		if r.URL.String() == "/" {
			w.Write([]byte(`
            <table summary="作家データ">
            <tr><td class="header">作家名：</td><td><font size="+2">テスト 太郎</font></td></tr>
            <tr><td class="header">作家名読み：</td><td>テスト 太郎</td></tr>
            <tr><td class="header">ローマ字表記：</td><td>Test, Taro</td></tr>
            </table>
            <ol>
            <li><a href="../cards/999999/card001.html">テスト書籍001</a></li> 
            <li><a href="../cards/999999/card002.html">テスト書籍002</a></li> 
            <li><a href="../cards/999999/card003.html">テスト書籍003</a></li> 
            </ol>
            `))
		} else {
			pat := regexp.MustCompile(`.*/cards/([0-9]+)/card([0-9]+).html$`)
			token := pat.FindStringSubmatch(r.URL.String())
			w.Write([]byte(fmt.Sprintf(`
            <table summary="作家データ">
            <tr><td class="header">作家名：</td><td><font size="+2">テスト 太郎</font></td></tr>
            <tr><td class="header">作家名読み：</td><td>テスト 太郎</td></tr>
            <tr><td class="header">ローマ字表記：</td><td>Test, Taro</td></tr>
            </table>
            <table border="1" summary="ダウンロードデータ" class="download">
            <tr>
                <td><a href="./files/%[1]s_%[2]s.zip">%[1]s_%[2]s.zip</a></td>
            </tr>
            </table>
            `, token[1], token[2])))
		}
	}))
	defer ts.Close()

	tmp := pageURLFormat
	pageURLFormat = ts.URL + "/cards/%s/card%s.html"
	defer func() {
		pageURLFormat = tmp
	}()

	got, err := findEntries(ts.URL)
	if err != nil {
		t.Error(err)
		return
	}
	want := []Entry{
		{
			AuthorID: "999999",
			Author:   "テスト 太郎",
			TitleID:  "001",
			Title:    "テスト書籍001",
			SiteURL:  ts.URL,
			ZipURL:   ts.URL + "/cards/999999/files/999999_001.zip",
		},
		{
			AuthorID: "999999",
			Author:   "テスト 太郎",
			TitleID:  "002",
			Title:    "テスト書籍002",
			SiteURL:  ts.URL,
			ZipURL:   ts.URL + "/cards/999999/files/999999_002.zip",
		},
		{
			AuthorID: "999999",
			Author:   "テスト 太郎",
			TitleID:  "003",
			Title:    "テスト書籍003",
			SiteURL:  ts.URL,
			ZipURL:   ts.URL + "/cards/999999/files/999999_003.zip",
		},
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %+v, but got %+v", want, got)
	}
}

func TestFindAuthorAndZIP(t *testing.T) {
	type args struct {
		siteURL string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := findAuthorAndZIP(tt.args.siteURL)
			if got != tt.want {
				t.Errorf("findAuthorAndZIP() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("findAuthorAndZIP() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFindEntries(t *testing.T) {
	type args struct {
		siteURL string
	}
	tests := []struct {
		name    string
		args    args
		want    []Entry
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findEntries(tt.args.siteURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("findEntries() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findEntries() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extractText(t *testing.T) {
	type args struct {
		zipURL string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractText(tt.args.zipURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("extractText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setupDB(t *testing.T) {
	type args struct {
		dsn string
	}
	tests := []struct {
		name    string
		args    args
		want    *sql.DB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := setupDB(tt.args.dsn)
			if (err != nil) != tt.wantErr {
				t.Errorf("setupDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setupDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addEntry(t *testing.T) {
	type args struct {
		db      *sql.DB
		entry   *Entry
		content string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := addEntry(tt.args.db, tt.args.entry, tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("addEntry() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
