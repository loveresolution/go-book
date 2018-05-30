/**
 * FileName:   switch.go
 * Author:     Fasion Chan
 * @contact:   fasionchan@gmail.com
 * @version:   $Id$
 *
 * Description:
 *
 * Changelog:
 *
 **/

package main

import (
    "fmt"
    "runtime"
)


func main() {
    fmt.Print("Go runs on ")

    switch os := runtime.GOOS; os {
        case "darwin":
            fmt.Println("OS X.")

        case "linux":
            fmt.Println("Linux.")

        default:
            // freebsd, openbsd
            // plan9, windows...
            fmt.Printf("%s.\n", os)
    }
}
