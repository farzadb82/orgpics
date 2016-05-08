package main

import (
    "os"
    "fmt"
    "log"
    "flag"
    "strings"
    //"path/filepath"
    "math/big"
    exif "github.com/m0rcq/go-exif"
)

func main() {
    var outputDir string
    var makeCopy bool
    var overwrite bool

    flag.StringVar(&outputDir, "outputdir", ".", "Directory to re-write file to")
    flag.BoolVar(&makeCopy, "copy", true, "Make a copy of the original file rather than moving")
    flag.BoolVar(&overwrite, "overwrite", true, "Overwrite the target, if it exists")
    flag.Parse()
    fmt.Println(" Writing output to: ", outputDir)
    fmt.Println("     Making a copy: ", makeCopy)
    fmt.Println("Overwriting target: ", overwrite)
    fmt.Println("-----------------------------------------------------------------------")

    for _, inFile := range flag.Args() {
        fmt.Println("Image file to process: ", inFile)
        f, err := os.Open(inFile)
        if err != nil {
            log.Fatal(err)
        }

        var tags map[string]map[string]string = ProcessExifStream(f)
        fmt.Printf("%+v\n", tags)
        fmt.Println("-----------------------------------------------------------------------")
    }
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

    for _, v := range exifEntries {
        // v.Values is an array of interface{} typed items, that can hold items of any type
        // The command below converts the interface{} type back into an array of a specific type
        // We then use teh swtch statement to convert any particular type into a string
        lval := v.Values.([]interface{})
        var values string
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
