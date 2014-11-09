package main

import "fmt"
import "os"
import "strings"

type strtostr func(string) string

type strtobool func(string) bool

func mapto(item_s []string, f strtostr) []string {
    item_s_new := []string{}

    for _, item := range item_s {
        item_new := f(item)
        item_s_new = append(item_s_new, item_new)
    }

    return item_s_new
}

func filter(item_s []string, f strtobool) []string {
    item_s_new := []string{}

    for _, item := range item_s {
        if f(item) {
            item_s_new = append(item_s_new, item)
        }
    }

    return item_s_new
}

func any(item_s []string, f strtobool) bool {
    for _, item := range item_s {
        if f(item) {
            return true
        }
    }

    return false
}

func contain(item_s []string, item string) bool {
    for _, x := range item_s {
        if x == item {
            return true
        }
    }
    return false
}

func append_uniq(item_s []string, item string) []string {
    if contain(item_s, item) {
        return item_s
    } else {
        return append(item_s, item)
    }
}

func uniq(item_s []string) []string {
    item_s_new := []string{}

    for _, item := range item_s {
        item_s_new = append_uniq(item_s_new, item)
    }

    return item_s_new;
}

// Modified from |http://stackoverflow.com/a/12527546|.
// ---BEG
func file_exists(path string) bool {
    _, err := os.Stat(path)

    if err != nil {
        if os.IsNotExist(err) {
            return false
        }
    }

    return true
}
// ---END

func find_executable(prog string) []string {
    // 8f1kRCu
    env_var_PATHEXT := os.Getenv("PATHEXT")
    /// can be ""

    // 6qhHTHF
    // split into a list of extensions
    val_sep := string(os.PathListSeparator)

    var ext_s []string = nil

    if env_var_PATHEXT == "" {
        ext_s = []string{}
    } else {
        ext_s = strings.Split(env_var_PATHEXT, val_sep)
    }

    // 2pGJrMW
    // strip
    ext_s = mapto(ext_s, func(x string) string {
        return strings.TrimSpace(x)
    })

    // 2gqeHHl
    // remove empty
    ext_s = filter(ext_s, func(x string) bool {
        return x != ""
    })

    // 2zdGM8W
    // convert to lowercase
    ext_s = mapto(ext_s, func(x string) string {
        return strings.ToLower(x)
    })

    // 2fT8aRB
    // uniquify
    ext_s = uniq(ext_s);

    // 4ysaQVN
    env_var_PATH := os.Getenv("PATH")
    /// can be ""

    // 6mPI0lg
    var dir_path_s []string = nil

    if env_var_PATH == "" {
        dir_path_s = []string{}
    } else {
        dir_path_s = strings.Split(env_var_PATH, val_sep)
    }

    // 5rT49zI
    // insert empty dir path to the beginning
    //
    // Empty dir handles the case that |prog| is a path, either relative or
    //  absolute. See code 7rO7NIN.
    dir_path_s = append([]string{""}, dir_path_s...)

    // 2klTv20
    // uniquify
    dir_path_s = uniq(dir_path_s)

    //
    prog_lower := strings.ToLower(prog)

    prog_has_ext := any(ext_s, func(x string) bool {
        return strings.HasSuffix(prog_lower, x)
    })

    // 6bFwhbv
    path_sep := string(os.PathSeparator)

    exe_path_s := []string{}

    for _, dir_path := range dir_path_s {
        // 7rO7NIN
        // synthesize a path with the dir and prog
        path := ""

        if dir_path == "" {
            path = prog
        } else {
            path = dir_path + path_sep + prog
        }

        // 6kZa5cq
        // assume the path has extension, check if it is an executable
        if prog_has_ext && file_exists(path) {
            exe_path_s = append(exe_path_s, path)
        }

        // 2sJhhEV
        // assume the path has no extension
        for _, ext := range ext_s {
            // 6k9X6GP
            // synthesize a new path with the path and the executable extension
            path_plus_ext := path + ext

            // 6kabzQg
            // check if it is an executable
            if file_exists(path_plus_ext) {
                exe_path_s = append(exe_path_s, path_plus_ext)
            }
        }
    }

    // 8swW6Av
    // uniquify
    exe_path_s = uniq(exe_path_s);

    //
    return exe_path_s
}

func main() {
    //
    println := fmt.Println

    // 9mlJlKg
    // check if one cmd arg is given
    args := os.Args[1:]

    if (len(args) != 1) {
        // 7rOUXFo
        // print program usage
        println(`Usage: aoikwinwhich PROG`);
        println(``);
        println(`#/ PROG can be either name or path`);
        println(`aoikwinwhich notepad.exe`);
        println(`aoikwinwhich C:\Windows\notepad.exe`);
        println(``);
        println(`#/ PROG can be either absolute or relative`);
        println(`aoikwinwhich C:\Windows\notepad.exe`);
        println(`aoikwinwhich Windows\notepad.exe`);
        println(``);
        println(`#/ PROG can be either with or without extension`);
        println(`aoikwinwhich notepad.exe`);
        println(`aoikwinwhich notepad`);
        println(`aoikwinwhich C:\Windows\notepad.exe`);
        println(`aoikwinwhich C:\Windows\notepad`);

        // 3nqHnP7
        return;
    }

    // 9m5B08H
    // get name or path of a program from cmd arg
    prog := args[0]

    // 8ulvPXM
    // find executables
    path_s := find_executable(prog);

    // 5fWrcaF
    // has found none, exit
    if (len(path_s) == 0) {
        // 3uswpx0
        return;
    }

    // 9xPCWuS
    // has found some, output
    txt := strings.Join(path_s, "\n")

    println(txt)

    // 4s1yY1b
    return;
}
