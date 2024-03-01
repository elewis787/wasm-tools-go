package bindgen

// Option represents a single configuration option for this package.
type Option interface {
	applyOption(*options) error
}

type optionFunc func(*options) error

func (f optionFunc) applyOption(opts *options) error {
	return f(opts)
}

type options struct {
	// generatedBy is the name of the program that generates code with this package.
	generatedBy string

	// packageRoot is the root Go package or module path used in generated code.
	packageRoot string

	// cmPackage is the package path to the "cm" or Component Model package with basic types.
	// Default: github.com/ydnar/wasm-tools-go/cm.
	cmPackage string

	// versioned determines if Go packages are generated with version numbers.
	versioned bool

	// idents maps WIT identifiers to Go identifiers. Examples:
	// "wasi:clocks" -> "wasi/clocks"
	// "wasi:clocks/wall-clock" -> "wasi/clocks/wall"
	// "wasi:clocks/wall-clock" -> "wasi/clocks/wall#wallclock" (for a Go short package name of wallclock)
	// "wasi:clocks/wall-clock.datetime" -> "wasi/clocks/wall#DateTime"
	// "wasi:clocks/wall-clock.now" -> "wasi/clocks/wall#Now"
	idents map[string]string
}

func (opts *options) apply(o ...Option) error {
	for _, o := range o {
		err := o.applyOption(opts)
		if err != nil {
			return err
		}
	}
	return nil
}

// GeneratedBy returns an [Option] that specifies the name of the program or package
// that will appear in the "Code generated by ..." header on generated files.
func GeneratedBy(name string) Option {
	return optionFunc(func(opts *options) error {
		opts.generatedBy = name
		return nil
	})
}

// PackageRoot returns an [Option] that specifies the root Go package path for generated Go packages.
func PackageRoot(path string) Option {
	return optionFunc(func(opts *options) error {
		opts.packageRoot = path
		return nil
	})
}

// CMPackage returns an [Option] that specifies the package path to the
// Component Model utility package (default: github.com/ydnar/wasm-tools-go/cm).
func CMPackage(path string) Option {
	return optionFunc(func(opts *options) error {
		opts.cmPackage = path
		return nil
	})
}

// Versioned returns an [Option] that that specifies that all generated Go packages
// will have versions that match WIT versions.
func Versioned(versioned bool) Option {
	return optionFunc(func(opts *options) error {
		opts.versioned = versioned
		return nil
	})
}

// MapIdent returns an [Option] that maps a [WIT] identifier to a Go identifier.
//
// Acceptable values for from include: package names like "wasi:clocks",
// qualified interface or world names like "wasi:clocks/wall-clock",
// or fully-qualified identifiers like "wasi:clocks/wall-clock.datetime".
//
// Acceptable values for to include: Go package paths like "wasi/clocks/wallclock",
// qualified package names like "wasi/clocks/wall#wallclock", or
// qualified identifier names like "wasi/clocks/wallclock#DateTime".
//
// [WIT]: https://github.com/WebAssembly/component-model/blob/main/design/mvp/WIT.md
func MapIdent(from, to string) Option {
	return optionFunc(func(opts *options) error {
		if opts.idents == nil {
			opts.idents = make(map[string]string)
		}
		opts.idents[from] = to
		return nil
	})
}

// MapIdents returns an [Option] that maps [WIT] identifiers to Go identifiers.
// See [MapIdent] for more information.
//
// [WIT]: https://github.com/WebAssembly/component-model/blob/main/design/mvp/WIT.md
func MapIdents(idents map[string]string) Option {
	return optionFunc(func(opts *options) error {
		for from, to := range idents {
			err := MapIdent(from, to).applyOption(opts)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
