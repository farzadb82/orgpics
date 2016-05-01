package main

import (
    "os"
    "fmt"
    "log"
    "math/big"
    "strings"
    //"path/filepath"
    exif "github.com/m0rcq/go-exif"
)

func main() {
    var imageFile string = "test_data/test_exif.jpg"

    fmt.Println("Image file to process: ", imageFile)
    f, err := os.Open(imageFile)
    if err != nil {
        log.Fatal(err)
    }

    var tags map[string]map[string]string = ProcessExifStream(f)
    fmt.Printf("%+v\n", tags)
    //filepath.Walk(".", walkFunc)
}

func ProcessExifStream(exifFd *os.File) map[string]map[string]string {
    var exifData = make(map[string]map[string]string)

    exif := &exif.ExifData{}
    exif.ProcessExifStream(exifFd)

    for k, v := range exif.IfdData {
        exifData[k] = decodeExifData(v)
    }

    return exifData
}

func decodeExifData(exifEntries []exif.IfdEntries) map[string]string {
    var exifTags = make(map[string]string)

    fmt.Printf("exifEntries: %+v\n\n", exifEntries)
    for i, v := range exifEntries {
        fmt.Printf("v[%d]: %+v\n", i, v)
        lval, ok := v.Values.([]interface{})
        fmt.Printf("  lval: %+v\n  ok: %+v\n", lval, ok)
        var values string
        if ok {
            switch val := lval[0].(type) {
            case string:
                values = fmt.Sprintf("'%s'", val)
            case byte:
                values = fmt.Sprintf("%#x", val)
            case []uint8:
                var lstr []string
                for _, v := range lval {
                    lstr = append(lstr, fmt.Sprintf("%#x", v))
                }
                values = strings.Join(lstr, ", ")
            case int16:
                values = fmt.Sprintf("%d", val)
            case int32:
                values = fmt.Sprintf("%d", val)
            case int64:
                values = fmt.Sprintf("%d", val)
            case uint16:
                values = fmt.Sprintf("%d", val)
            case uint32:
                values = fmt.Sprintf("%d", val)
            case uint64:
                values = fmt.Sprintf("%d", val)
            case *big.Rat:
                values = fmt.Sprintf("%s", val.RatString())
            default:
                values = fmt.Sprintf("%v", lval)
            }
        }
        /*
            TagSection: exif.IfdSeqMap[v.IfdSeq]
            TagId: v.Tag
            TagName: v.TagDesc
            TagFormatType: v.Format
            TagFormatTypeName: exif.FormatType[int(v.Format)]
            TagValue: values
        */
        exifTags[v.TagDesc] = values
    }

    return exifTags
}

/*
func walkFunc(path string, info os.FileInfo, err error) error {
    if err != nil {
        log.Fatal(err)
    }

    isDir := ""
    if info.IsDir() {
        isDir = "[ DIR ]"
    }

    fmt.Printf("%s\t\t\t%s\t%d\n", path, isDir, info.Size())

    return nil
}
*/
