### Install
```
go get github.com/farzadb82/orgpics
```
### Build
```
go install github.com/farzadb82/orgpics
```
### Run
```
$GOPATH/bin/orgpics <filename>
```
The above command will use any EXIF data in the file to organize the file into a sub-directory by the DateTime value within the EXIF data. If the file lacks EXIF data or lacks the DateTime field in EXIF, it will not be moved.
#### For example:
If the DateTime value in the EXIF data is `2016:04:30 21:59:58`, the file will be placed into a sub-directory named `2016-04-30`.

----------

### License (MIT)
**Copyright (c) 2016 Farzad Battiwalla**

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
