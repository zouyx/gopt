package build


import (
	"github.com/zouyx/gopt/input"
	"fmt"
)


const (
	COVERAGE_FILE_TEMPLATE="set -e\n" +
		"\n" +
		"workdir=.cover\n" +
		"profile=\"$workdir/cover.out\"\n" +
		"mode=count\n" +
		"\n" +
		"generate_cover_data() {\n" +
		"    rm -rf \"$workdir\"\n" +
		"    mkdir \"$workdir\"\n" +
		"\n" +
		"    for pkg in \"$@\"; do\n" +
		"        f=\"$workdir/$(echo $pkg | tr / -).cover\"\n" +
		"        go test -covermode=\"$mode\" -coverprofile=\"$f\" \"$pkg\"\n" +
		"    done\n" +
		"\n" +
		"    echo \"mode: $mode\" >\"$profile\"\n" +
		"    grep -h -v \"^mode:\" \"$workdir\"/*.cover >>\"$profile\"\n" +
		"}\n" +
		"\n" +
		"show_html_report() {\n" +
		"    go tool cover -html=\"$profile\" -o=\"$workdir\"/coverage.html\n" +
		"}\n" +
		"\n" +
		"show_csv_report() {\n" +
		"    go tool cover -func=\"$profile\" -o=\"$workdir\"/coverage.csv\n" +
		"}\n" +
		"\n" +
		"push_to_coveralls() {\n" +
		"    echo \"Pushing coverage statistics to coveralls.io\"\n" +
		"    # ignore failure to push - it happens\n" +
		"    $HOME/gopath/bin/goveralls -coverprofile=\"$profile\" \\\n" +
		"                               -service=travis-ci  || true\n" +
		"}\n" +
		"\n" +
		"generate_cover_data $(go list ./...)\n" +
		"show_csv_report\n" +
		"\n" +
		"case \"$1\" in\n" +
		"\"\")\n" +
		"    ;;\n" +
		"--html)\n" +
		"    show_html_report ;;\n" +
		"--coveralls)\n" +
		"    push_to_coveralls ;;\n" +
		"*)\n" +
		"    echo >&2 \"error: invalid option: $1\"; exit 1 ;;\n" +
		"esac"
	COVERAGE_FILE="/coverage.sh"
)

type CoverageBuilder struct {

}

// coverage.sh build method
func (this *CoverageBuilder) Build(params *input.Params) {
	src := getSrc(params)
	fileName := src + COVERAGE_FILE

	write(fileName,fmt.Sprintf(COVERAGE_FILE_TEMPLATE))
}
