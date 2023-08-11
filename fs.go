/*
   Create: 2023/8/4
   Project: Gungnir
   Github: https://github.com/landers1037
   Copyright Renj
*/

package gungnir

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"path"
	"sort"
	"strings"
)

type fileHandler struct {
	rootPath string
	root     http.FileSystem
}

type anyDirs interface {
	len() int
	name(i int) string
	isDir(i int) bool
}

type fileInfoDirs []fs.FileInfo

func (d fileInfoDirs) len() int          { return len(d) }
func (d fileInfoDirs) isDir(i int) bool  { return d[i].IsDir() }
func (d fileInfoDirs) name(i int) string { return d[i].Name() }

type dirEntryDirs []fs.DirEntry

func (d dirEntryDirs) len() int          { return len(d) }
func (d dirEntryDirs) isDir(i int) bool  { return d[i].IsDir() }
func (d dirEntryDirs) name(i int) string { return d[i].Name() }

func (f *fileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upath := r.URL.Path
	if filterAgents(r) || filterIPs(r) {
		serveDeny(w, nil)
	}
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		r.URL.Path = upath
	}
	customHeaders(w)
	serveFile(w, r, f, path.Clean(upath), true)
}

func serveFile(w http.ResponseWriter, r *http.Request, fd *fileHandler, name string, redirect bool) {
	const indexPage = "/index.html"
	httpFs := fd.root
	rootPath := fd.rootPath

	// redirect .../index.html to .../
	// can't use Redirect() because that would make the path absolute,
	// which would be a problem running under StripPrefix
	if strings.HasSuffix(r.URL.Path, indexPage) {
		localRedirect(w, r, "./")
		return
	}

	if filterDenyDirs(name) {
		serveDeny(w, "permission denied")
		return
	}

	f, err := httpFs.Open(name)
	if err != nil {
		log.Printf("open fs file error: %s\n", err.Error())
		if strings.Contains(err.Error(), "Access") {
			serveError(w, "file access denied")
		} else {
			serveNotFound(w)
		}
		return
	}
	defer f.Close()

	d, err := f.Stat()
	if err != nil {
		log.Printf("stat fs file error: %s\n", err.Error())
		serveError(w, "file stat error")
		return
	}

	if redirect {
		// redirect to canonical path: / at end of directory url
		// r.URL.Path always begins with /
		url := r.URL.Path
		if d.IsDir() {
			if url[len(url)-1] != '/' {
				localRedirect(w, r, path.Base(url)+"/")
				return
			}
		} else {
			if url[len(url)-1] == '/' {
				localRedirect(w, r, "../"+path.Base(url))
				return
			}
		}
	}

	if d.IsDir() {
		url := r.URL.Path
		// redirect if the directory name doesn't end in a slash
		if url == "" || url[len(url)-1] != '/' {
			localRedirect(w, r, path.Base(url)+"/")
			return
		}

		// use contents of index.html for directory, if present
		index := strings.TrimSuffix(name, "/") + indexPage
		ff, err := httpFs.Open(index)
		if err == nil {
			defer ff.Close()
			dd, err := ff.Stat()
			if err == nil {
				d = dd
				f = ff
			}
		}
	}

	// Still a directory? (we didn't find an index.html file)
	if d.IsDir() {
		dirList(w, f, rootPath, name)
		return
	}

	stat, err := f.Stat()
	if err != nil {
		serveError(w, nil)
	}

	// file preview mode?
	if r.URL.Query().Get("mode") != "" {
		serveFilePage(w, stat, d.Name())
		return
	}

	if filterDenyFiles(stat.Name()) {
		serveDeny(w, "permission denied")
		return
	}

	// serveContent will check modification time
	http.ServeContent(w, r, d.Name(), d.ModTime(), f)
}

func localRedirect(w http.ResponseWriter, r *http.Request, newPath string) {
	if q := r.URL.RawQuery; q != "" {
		newPath += "?" + q
	}
	w.Header().Set("Location", newPath)
	w.WriteHeader(500)
}

func Error(w http.ResponseWriter, error string, code int) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	fmt.Fprintln(w, error)
}

func toHTTPError(err error) (msg string, httpStatus int) {
	if errors.Is(err, fs.ErrNotExist) {
		return "404 page not found", 404
	}
	if errors.Is(err, fs.ErrPermission) {
		return "403 Forbidden", 403
	}
	// Default:
	return "500 Internal Server Error", 500
}

func dirList(w http.ResponseWriter, f http.File, rootPath, dir string) {
	// Prefer to use ReadDir instead of Readdir,
	// because the former doesn't require calling
	// Stat on every entry of a directory on Unix.
	var dirs anyDirs
	var err error
	if d, ok := f.(fs.ReadDirFile); ok {
		var list dirEntryDirs
		list, err = d.ReadDir(-1)
		dirs = list
	} else {
		var list fileInfoDirs
		list, err = f.Readdir(-1)
		dirs = list
	}

	if err != nil {
		log.Printf("http: error reading directory: %v\n", err)
		serveError(w, "Error reading directory")
		return
	}
	sort.Slice(dirs, func(i, j int) bool { return dirs.name(i) < dirs.name(j) })

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var data DataModel

	for i, n := 0, dirs.len(); i < n; i++ {
		name := dirs.name(i)
		if dirs.isDir(i) {
			name += "/"
			data.add(Entry{
				Name:    name,
				Size:    "",
				ModTime: "",
				IsDir:   true,
			})
		} else {
			filePath := getFilePath(rootPath, dir, name)
			size, mod := getFileInfo(filePath)
			data.add(Entry{
				Name:    htmlReplacer.Replace(name),
				Size:    size,
				ModTime: mod,
				IsDir:   false,
			})
		}
		// name may contain '?' or '#', which must be escaped to remain
		// part of the URL path, and not indicate the start of a query
		// string or fragment.
	}
	serveIndex(w, data)
}

// 转义
var htmlReplacer = strings.NewReplacer(
	"&", "&amp;",
	"<", "&lt;",
	">", "&gt;",
	// "&#34;" is shorter than "&quot;".
	`"`, "&#34;",
	// "&#39;" is shorter than "&apos;" and apos was not in HTML until HTML5.
	"'", "&#39;",
)
