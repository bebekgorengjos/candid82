// This file is generated by generate-std.joke script. Do not edit manually!

package filepath

import (
	. "github.com/candid82/joker/core"
	"path/filepath"
)

var filepathNamespace = GLOBAL_ENV.EnsureNamespace(MakeSymbol("joker.filepath"))

var list_separator_ String
var separator_ String

var __abs__P ProcFn = __abs_
var abs_ Proc = Proc{Fn: __abs__P, Name: "abs_", Package: "std/filepath"}

func __abs_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		path := ExtractString(_args, 0)
		 _res, err := filepath.Abs(path)
		PanicOnErr(err)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __isabs__P ProcFn = __isabs_
var isabs_ Proc = Proc{Fn: __isabs__P, Name: "isabs_", Package: "std/filepath"}

func __isabs_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		path := ExtractString(_args, 0)
		_res := filepath.IsAbs(path)
		return MakeBoolean(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __base__P ProcFn = __base_
var base_ Proc = Proc{Fn: __base__P, Name: "base_", Package: "std/filepath"}

func __base_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		path := ExtractString(_args, 0)
		_res := filepath.Base(path)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __clean__P ProcFn = __clean_
var clean_ Proc = Proc{Fn: __clean__P, Name: "clean_", Package: "std/filepath"}

func __clean_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		path := ExtractString(_args, 0)
		_res := filepath.Clean(path)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __dir__P ProcFn = __dir_
var dir_ Proc = Proc{Fn: __dir__P, Name: "dir_", Package: "std/filepath"}

func __dir_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		path := ExtractString(_args, 0)
		_res := filepath.Dir(path)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __eval_symlinks__P ProcFn = __eval_symlinks_
var eval_symlinks_ Proc = Proc{Fn: __eval_symlinks__P, Name: "eval_symlinks_", Package: "std/filepath"}

func __eval_symlinks_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		path := ExtractString(_args, 0)
		 _res, err := filepath.EvalSymlinks(path)
		PanicOnErr(err)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __ext__P ProcFn = __ext_
var ext_ Proc = Proc{Fn: __ext__P, Name: "ext_", Package: "std/filepath"}

func __ext_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		path := ExtractString(_args, 0)
		_res := filepath.Ext(path)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __file_seq__P ProcFn = __file_seq_
var file_seq_ Proc = Proc{Fn: __file_seq__P, Name: "file_seq_", Package: "std/filepath"}

func __file_seq_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		root := ExtractString(_args, 0)
		_res := fileSeq(root)
		return _res

	default:
		PanicArity(_c)
	}
	return NIL
}

var __from_slash__P ProcFn = __from_slash_
var from_slash_ Proc = Proc{Fn: __from_slash__P, Name: "from_slash_", Package: "std/filepath"}

func __from_slash_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		path := ExtractString(_args, 0)
		_res := filepath.FromSlash(path)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __glob__P ProcFn = __glob_
var glob_ Proc = Proc{Fn: __glob__P, Name: "glob_", Package: "std/filepath"}

func __glob_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		pattern := ExtractString(_args, 0)
		 _res, err := filepath.Glob(pattern)
		PanicOnErr(err)
		return MakeStringVector(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __join__P ProcFn = __join_
var join_ Proc = Proc{Fn: __join__P, Name: "join_", Package: "std/filepath"}

func __join_(_args []Object) Object {
	_c := len(_args)
	switch {
	case true:
		CheckArity(_args, 0, 999)
		elems := ExtractStrings(_args, 0)
		_res := filepath.Join(elems...)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __ismatches__P ProcFn = __ismatches_
var ismatches_ Proc = Proc{Fn: __ismatches__P, Name: "ismatches_", Package: "std/filepath"}

func __ismatches_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 2:
		pattern := ExtractString(_args, 0)
		name := ExtractString(_args, 1)
		 _res, err := filepath.Match(pattern, name)
		PanicOnErr(err)
		return MakeBoolean(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __rel__P ProcFn = __rel_
var rel_ Proc = Proc{Fn: __rel__P, Name: "rel_", Package: "std/filepath"}

func __rel_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 2:
		basepath := ExtractString(_args, 0)
		targpath := ExtractString(_args, 1)
		 _res, err := filepath.Rel(basepath, targpath)
		PanicOnErr(err)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __split__P ProcFn = __split_
var split_ Proc = Proc{Fn: __split__P, Name: "split_", Package: "std/filepath"}

func __split_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		path := ExtractString(_args, 0)
		 _dir, _file := filepath.Split(path)
		_res := NewVectorFrom(MakeString(_dir), MakeString(_file))
		return _res

	default:
		PanicArity(_c)
	}
	return NIL
}

var __split_list__P ProcFn = __split_list_
var split_list_ Proc = Proc{Fn: __split_list__P, Name: "split_list_", Package: "std/filepath"}

func __split_list_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		path := ExtractString(_args, 0)
		_res := filepath.SplitList(path)
		return MakeStringVector(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __to_slash__P ProcFn = __to_slash_
var to_slash_ Proc = Proc{Fn: __to_slash__P, Name: "to_slash_", Package: "std/filepath"}

func __to_slash_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		path := ExtractString(_args, 0)
		_res := filepath.ToSlash(path)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __volume_name__P ProcFn = __volume_name_
var volume_name_ Proc = Proc{Fn: __volume_name__P, Name: "volume_name_", Package: "std/filepath"}

func __volume_name_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		path := ExtractString(_args, 0)
		_res := filepath.VolumeName(path)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

func Init() {
	list_separator_ = MakeString(string(filepath.ListSeparator))
	separator_ = MakeString(string(filepath.Separator))


	filepathNamespace.ResetMeta(MakeMeta(nil, `Implements utility routines for manipulating filename paths.`, "1.0"))

	filepathNamespace.InternVar("list-separator", list_separator_,
		MakeMeta(
			nil,
			`OS-specific path list separator.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	filepathNamespace.InternVar("separator", separator_,
		MakeMeta(
			nil,
			`OS-specific path separator.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	filepathNamespace.InternVar("abs", abs_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("path"))),
			`Returns an absolute representation of path. If the path is not absolute it will be
  joined with the current working directory to turn it into an absolute path.
  The absolute path name for a given file is not guaranteed to be unique.
  Calls clean on the result.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	filepathNamespace.InternVar("abs?", isabs_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("path"))),
			`Reports whether the path is absolute.`, "1.0").Plus(MakeKeyword("tag"), String{S: "Boolean"}))

	filepathNamespace.InternVar("base", base_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("path"))),
			`Returns the last element of path. Trailing path separators are removed before
  extracting the last element. If the path is empty, returns ".". If the path consists
  entirely of separators, returns a single separator.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	filepathNamespace.InternVar("clean", clean_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("path"))),
			`Returns the shortest path name equivalent to path by purely lexical processing.
  Applies the following rules iteratively until no further processing can be done:

1. Replace multiple separator elements with a single one.
2. Eliminate each . path name element (the current directory).
3. Eliminate each inner .. path name element (the parent directory)
   along with the non-.. element that precedes it.
4. Eliminate .. elements that begin a rooted path:
   that is, replace "/.." by "/" at the beginning of a path,
   assuming separator is '/'.
The returned path ends in a slash only if it represents a root directory, such as "/" on Unix or ` + "`" + `C:\` + "`" + ` on Windows.

Finally, any occurrences of slash are replaced by separator.

If the result of this process is an empty string, returns the string ".".`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	filepathNamespace.InternVar("dir", dir_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("path"))),
			`Returns all but the last element of path, typically the path's directory.
  After dropping the final element, calls clean on the path and trailing slashes are removed.
  If the path is empty, returns ".". If the path consists entirely of separators,
  returns a single separator. The returned path does not end in a separator unless it is the root directory.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	filepathNamespace.InternVar("eval-symlinks", eval_symlinks_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("path"))),
			`Returns the path name after the evaluation of any symbolic links. If path is relative the result will be
  relative to the current directory, unless one of the components is an absolute symbolic link.
  Calls clean on the result.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	filepathNamespace.InternVar("ext", ext_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("path"))),
			`Returns the file name extension used by path. The extension is the suffix beginning at the final dot
  in the final element of path; it is empty if there is no dot.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	filepathNamespace.InternVar("file-seq", file_seq_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("root"))),
			`Returns a seq of maps with info about files or directories under root.`, "1.0"))

	filepathNamespace.InternVar("from-slash", from_slash_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("path"))),
			`Returns the result of replacing each slash ('/') character in path with a separator character.
  Multiple slashes are replaced by multiple separators.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	filepathNamespace.InternVar("glob", glob_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("pattern"))),
			`Returns the names of all files matching pattern or nil if there is no matching file.
  The syntax of patterns is the same as in Match. The pattern may describe hierarchical
  names such as /usr/*/bin/ed (assuming the separator is '/').

  Ignores file system errors such as I/O errors reading directories.
  Throws exception when pattern is malformed.`, "1.0").Plus(MakeKeyword("tag"), String{S: "[String]"}))

	filepathNamespace.InternVar("join", join_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("&"), MakeSymbol("elems"))),
			`Joins any number of path elements into a single path, adding a separator if necessary.
  Calls clean on the result; in particular, all empty strings are ignored. On Windows,
  the result is a UNC path if and only if the first path element is a UNC path.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	filepathNamespace.InternVar("matches?", ismatches_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("pattern"), MakeSymbol("name"))),
			`Reports whether name matches the shell file name pattern.
  Requires pattern to match all of name, not just a substring.
  Throws exception if pattern is malformed.
  On Windows, escaping is disabled. Instead, '\' is treated as path separator.`, "1.0").Plus(MakeKeyword("tag"), String{S: "Boolean"}))

	filepathNamespace.InternVar("rel", rel_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("basepath"), MakeSymbol("targpath"))),
			`Returns a relative path that is lexically equivalent to targpath when joined to basepath
  with an intervening separator. On success, the returned path will always be relative to basepath,
  even if basepath and targpath share no elements. An exception is thrown if targpath can't be made
  relative to basepath or if knowing the current working directory would be necessary to compute it.
  Calls clean on the result.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	filepathNamespace.InternVar("split", split_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("path"))),
			`Splits path immediately following the final separator, separating it into a directory and file name component.
  If there is no separator in path, returns an empty dir and file set to path. The returned values have
  the property that path = dir+file.`, "1.0"))

	filepathNamespace.InternVar("split-list", split_list_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("path"))),
			`Splits a list of paths joined by the OS-specific list-separator, usually found in PATH or GOPATH environment variables.
  Returns an empty slice when passed an empty string.`, "1.0").Plus(MakeKeyword("tag"), String{S: "[String]"}))

	filepathNamespace.InternVar("to-slash", to_slash_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("path"))),
			`Returns the result of replacing each separator character in path with a slash ('/') character.
  Multiple separators are replaced by multiple slashes.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	filepathNamespace.InternVar("volume-name", volume_name_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("path"))),
			`Returns leading volume name. Given "C:\foo\bar" it returns "C:" on Windows. Given "\\host\share\foo"
  returns "\\host\share". On other platforms it returns "".`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

}

func init() {
	filepathNamespace.Lazy = Init
}
