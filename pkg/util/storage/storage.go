package storage

import (
	"io"
	"os"
	"fmt"
	"log"
	"errors"
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
	"math/rand"

	"go.beyondstorage.io/v5/pairs"
	"go.beyondstorage.io/v5/services"
	"go.beyondstorage.io/v5/types"
	"go.beyondstorage.io/v5/pkg/randbytes"
)

// store type = s3 (AWS), azblob (Azure), gcs (Google Cloud), hdfs (Hadoop Storage)
func Init() (types.Storager, error) {
	connStr := fmt.Sprintf(
		"%s://%s%s?credential=%s&endpoint=%s&location=%s&enbale_virtual_dir",
		os.Getenv("STORAGE_TYPE"),
		os.Getenv("STORAGE_NAME"),
		os.Getenv("STORAGE_WORKDIR"),
		os.Getenv("STORAGE_CREDENTIAL"),
		os.Getenv("STORAGE_ENDPOINT"),
		os.Getenv("STORAGE_LOCATION"),
	)
	return services.NewStoragerFromString(connStr)
}

func InitByString()  (types.Storager, error){
	// exampe: s3://bucket_name/path/to/workdir
	store, err := services.NewStoragerFromString(os.Getenv("STORAGE_STRING"))
    return store, err
}

func List(store types.Storager, path string) (*types.ObjectIterator, error) {
	return store.List(path)
}

func ListAll(store types.Storager) {
	it, err := store.List("")
	if err != nil {
		log.Fatalf("list: %v", err)
	}

	for {
		o, err := it.Next()
		if err != nil && !errors.Is(err, types.IterateDone) {
			log.Fatalf("Next: %v", err)
		}

		if err != nil {
			log.Printf("list completed")
			break
		}

		log.Printf("object path: %v", o.Path)
	}
}

func ListDir(store types.Storager, path string) {
	it, err := store.List(path, pairs.WithListMode(types.ListModeDir))
	if err != nil {
		log.Fatalf("list %v: %v", path, err)
	}

	for {
		o, err := it.Next()
		if err != nil && !errors.Is(err, types.IterateDone) {
			log.Fatalf("Next %v: %v", path, err)
		}

		if err != nil {
			log.Printf("list directory completed: %v", path)
			break
		}

		log.Printf("object path: %v", o.Path)
	}
}

func ListPrefix(store types.Storager, path string) {
	it, err := store.List(path, pairs.WithListMode(types.ListModePrefix))
	if err != nil {
		log.Fatalf("list %v: %v", path, err)
	}

	for {
		o, err := it.Next()
		if err != nil && !errors.Is(err, types.IterateDone) {
			log.Fatalf("Next %v: %v", path, err)
		}

		if err != nil {
			log.Printf("list with prefix completed: %v", path)
			break
		}

		log.Printf("object path: %v", o.Path)
	}
}

func ListPart(store types.Storager, path string) {
	it, err := store.List(path, pairs.WithListMode(types.ListModePart))
	if err != nil {
		log.Fatalf("list %v: %v", path, err)
	}

	for {
		o, err := it.Next()
		if err != nil && !errors.Is(err, types.IterateDone) {
			log.Fatalf("Next %v: %v", path, err)
		}

		if err != nil {
			log.Printf("list multipart uploads completed: %v", path)
			break
		}

		log.Printf("object path: %v", o.Path)
		log.Printf("object multipartID: %v", o.MustGetMultipartID())
	}
}

func Read(store types.Storager, file string, w io.Writer) (int64, error) {
	return store.Read(file, w)
}

func ReadWhole(store types.Storager, path string) {
	var buf bytes.Buffer

	n, err := store.Read(path, &buf)
	if err != nil {
		log.Fatalf("read %v: %v", path, err)
	}

	log.Printf("read size: %d", n)
	log.Printf("read content: %s", buf.Bytes())
}

func ReadRange(store types.Storager, path string, offset, size int64) {
	var buf bytes.Buffer

	n, err := store.Read(path, &buf,
		pairs.WithOffset(offset),
		pairs.WithSize(size),
	)
	if err != nil {
		log.Fatalf("read %v: %v", path, err)
	}

	log.Printf("read size: %d", n)
	log.Printf("read content: %s", buf.Bytes())
}

func ReadWithCallback(store types.Storager, path string) {
	var buf bytes.Buffer

	cur := int64(0)
	fn := func(bs []byte) {
		cur += int64(len(bs))
		log.Printf("read %d bytes already", cur)
	}

	n, err := store.Read(path, &buf, pairs.WithIoCallback(fn))
	if err != nil {
		log.Fatalf("read %v: %v", path, err)
	}

	log.Printf("read size: %d", n)
	log.Printf("read content: %s", buf.Bytes())
}

func ReadWithSignedURL(store types.Storager, path string, expire time.Duration) {
	signer, ok := store.(types.StorageHTTPSigner)
	if !ok {
		log.Fatalf("StorageHTTPSigner unimplemented")
	}

	req, err := signer.QuerySignHTTPRead(path, expire)
	if err != nil {
		log.Fatalf("read %v: %v", path, err)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("send HTTP request for reading %v: %v", path, err)
	}

	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Fatalf("close HTTP response body for reading %v: %v", path, err)
		}
	}()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("read from HTTP response body for reading %v: %v", path, err)
	}

	log.Printf("read size: %d", resp.ContentLength)
	log.Printf("read content: %s", buf)
}

func Write(store types.Storager, file string, r io.Reader, length int64) (int64, error) {
	return store.Write("hello.txt", r, length)
}

func WriteData(store types.Storager, path string) {
	size := rand.Int63n(4 * 1024 * 1024)
	r := io.LimitReader(randbytes.NewRand(), size)

	n, err := store.Write(path, r, size)
	if err != nil {
		log.Fatalf("write %v: %v", path, err)
	}

	log.Printf("write size: %d", n)
}

func WriteWithCallback(store types.Storager, path string) {
	size := rand.Int63n(4 * 1024 * 1024)
	r := io.LimitReader(randbytes.NewRand(), size)

	cur := int64(0)
	fn := func(bs []byte) {
		cur += int64(len(bs))
		log.Printf("write %d bytes already", cur)
	}

	n, err := store.Write(path, r, size, pairs.WithIoCallback(fn))
	if err != nil {
		log.Fatalf("write %v: %v", path, err)
	}

	log.Printf("write size: %d", n)
}

func WriteWithSignedURL(store types.Storager, path string, expire time.Duration) {
	signer, ok := store.(types.StorageHTTPSigner)
	if !ok {
		log.Fatalf("StorageHTTPSigner unimplemented")
	}

	size := rand.Int63n(4 * 1024 * 1024)
	r := io.LimitReader(randbytes.NewRand(), size)

	req, err := signer.QuerySignHTTPWrite(path, size, expire)
	if err != nil {
		log.Fatalf("write %v: %v", path, err)
	}

	req.Body = ioutil.NopCloser(r)

	client := http.Client{}
	_, err = client.Do(req)
	if err != nil {
		log.Fatalf("send HTTP request for writing %v: %v", path, err)
	}

	log.Printf("write size: %d", size)
}

func AppendToNewFile(appender types.Appender, path string) {
	size := rand.Int63n(4 * 1024 * 1024)
	content, _ := ioutil.ReadAll(io.LimitReader(randbytes.NewRand(), size))
	r := bytes.NewReader(content)

	o, err := appender.CreateAppend(path)
	if err != nil {
		log.Fatalf("CreateAppend %v: %v", path, err)
	}

	n, err := appender.WriteAppend(o, r, size)
	if err != nil {
		log.Fatalf("WriteAppend %v: %v", path, err)
	}

	err = appender.CommitAppend(o)
	if err != nil {
		log.Fatalf("CommitAppend %v: %v", path, err)
	}

	log.Printf("append size: %d", n)
}

func AppendToExistingFile(store types.Storager, path string) {
	size := rand.Int63n(4 * 1024 * 1024)
	content, _ := ioutil.ReadAll(io.LimitReader(randbytes.NewRand(), size))
	r := bytes.NewReader(content)

	appender, ok := store.(types.Appender)
	if !ok {
		log.Fatalf("Appender unimplemented")
	}

	o, err := store.Stat(path)
	if err != nil {
		log.Fatalf("Stat %v: %v", path, err)
	}

	n, err := appender.WriteAppend(o, r, size)
	if err != nil {
		log.Fatalf("WriteAppend %v: %v", path, err)
	}

	err = appender.CommitAppend(o)
	if err != nil {
		log.Fatalf("CommitAppend %v: %v", path, err)
	}

	log.Printf("append size: %d", n)
}


func Multipart(store types.Storager, path string) {
	size := rand.Int63n(4 * 1024 * 1024)
	content, _ := ioutil.ReadAll(io.LimitReader(randbytes.NewRand(), size))
	r := bytes.NewReader(content)

	multiparter, ok := store.(types.Multiparter)
	if !ok {
		log.Fatalf("Multiparter unimplemented")
	}

	o, err := multiparter.CreateMultipart(path)
	if err != nil {
		log.Fatalf("CreateMultipart %v: %v", path, err)
	}

	n, part, err := multiparter.WriteMultipart(o, r, size, 0)
	if err != nil {
		log.Fatalf("WriteMultipart %v: %v", path, err)
	}

	err = multiparter.CompleteMultipart(o, []*types.Part{part})
	if err != nil {
		log.Fatalf("CompleteMultipart %v: %v", path, err)
	}

	log.Printf("multipart upload size: %d", n)
}

func ResumeMultipart(store types.Storager, path string) {
	size := rand.Int63n(4 * 1024 * 1024)
	content, _ := ioutil.ReadAll(io.LimitReader(randbytes.NewRand(), size))
	r := bytes.NewReader(content)

	multiparter, ok := store.(types.Multiparter)
	if !ok {
		log.Fatalf("Multiparter unimplemented")
	}

	o, err := multiparter.CreateMultipart(path)
	if err != nil {
		log.Fatalf("CreateMultipart %v: %v", path, err)
	}

	multipartId := o.MustGetMultipartID()
	mo := store.Create(path, pairs.WithMultipartID(multipartId))
	var partNumber = -1
	var totalSize int64 = 0

	it, err := multiparter.ListMultipart(mo)
	if err != nil {
		log.Fatalf("ListMultipart %v: %v", path, err)
	}

	var parts []*types.Part
	for {
		p, err := it.Next()
		if err != nil && !errors.Is(err, types.IterateDone) {
			log.Fatalf("Next %v: %v", path, err)
		}

		if err != nil {
			log.Printf("ListMultipart completed: %v", path)
			break
		}

		partNumber = p.Index
		totalSize += p.Size
		parts = append(parts, p)
	}

	n, part, err := multiparter.WriteMultipart(mo, r, size, partNumber+1)
	if err != nil {
		log.Fatalf("WriteMultipart %v: %v", path, err)
	}

	totalSize += n
	parts = append(parts, part)

	err = multiparter.CompleteMultipart(mo, parts)
	if err != nil {
		log.Fatalf("CompleteMultipart %v: %v", path, err)
	}

	log.Printf("total upload size: %d", totalSize)
}

func CancelMultipart(store types.Storager, path string) {
	multiparter, ok := store.(types.Multiparter)
	if !ok {
		log.Fatalf("Multiparter unimplemented")
	}

	o, err := multiparter.CreateMultipart(path)
	if err != nil {
		log.Fatalf("CreateMultipart %v: %v", path, err)
	}

	err = store.Delete(path, pairs.WithMultipartID(o.MustGetMultipartID()))
	if err != nil {
		log.Fatalf("Delete with multipartId %v: %v", path, err)
	}

	log.Printf("cancel multipart: %v", path)
}

func Delete(store types.Storager) error {
	return store.Delete("hello.txt")
}