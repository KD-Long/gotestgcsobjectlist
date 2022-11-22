package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
//	"log"
	"time"
	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

// listFilesWithPrefix lists objects using prefix and delimeter.
func listFilesWithPrefix(w io.Writer, bucket, prefix, delim string) error {
	//  bucket := "kdlong-bucket"
	//  prefix := "/folder"
	//  delim := ""
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
			return fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	// Prefixes and delimiters can be used to emulate directory listings.
	// Prefixes can be used to filter objects starting with prefix.
	// The delimiter argument can be used to restrict the results to only the
	// objects in the given "directory". Without the delimiter, the entire tree
	// under the prefix is returned.
	//
	// For example, given these blobs:
	//   /a/1.txt
	//   /a/b/2.txt
	//
	// If you just specify prefix="a/", you'll get back:
	//   /a/1.txt
	//   /a/b/2.txt
	//
	// However, if you specify prefix="a/" and delim="/", you'll get back:
	//   /a/1.txt
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	it := client.Bucket(bucket).Objects(ctx, &storage.Query{
			Prefix:    prefix,
			Delimiter: delim,
	})
	fmt.Println("enter func bkt"+bucket+prefix)
	for {	
			attrs, err := it.Next()
			if err == iterator.Done {
				//log.Fatalf("Failed to create bucket: %v", err)
					break
			}

			if err != nil {
					return fmt.Errorf("Bucket(%q).Objects(): %v", bucket, err)
			}
			fmt.Fprintln(w, attrs.Name)
	}
	return nil
}

func main() {

	var buf bytes.Buffer
	bucket := "kdlong-bucket"
	prefix := "folder10/folder11/"
	delim := ""
	
    fmt.Println("### Config")
	fmt.Println("Bucket= "+ bucket)
	fmt.Println("prefix= "+ prefix)
	fmt.Println("delim= "+ delim)
	fmt.Println("######")

	fmt.Println("### Start")

	listFilesWithPrefix(&buf, bucket, prefix, delim)
	fmt.Println(buf.String())
	fmt.Println("### End")
	
}

// [END storage_list_buckets]``