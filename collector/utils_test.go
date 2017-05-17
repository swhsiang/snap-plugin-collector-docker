// +build small

package collector

import (
	"fmt"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	. "github.com/intelsdi-x/snap-plugin-collector-docker/collector"
)

func TestHasAnyPrefix(t *testing.T) {
	Convey("Filter label with specified prefix string", t, func() {
		type args struct {
			s          string
			prefixList []string
		}
		tests := []struct {
			name string
			args args
			want bool
		}{
			// TODO: Add test cases.
			struct {
				name string
				args args
				want bool
			}{
				args: args{
					s: "annotation.io.kubernetes.container.ports",
					prefixList: []string{
						"license",
						"maintainer",
						"vendor",
						"version",
						"build-date",
						"works.weave.role",
						"org.label-schema",
						"annotation.kubernetes.io",
						"annotation.scheduler.alpha.kubernetes.io",
					},
				},
				name: "full list",
				want: true,
			},
		}

		for _, tt := range tests {
			Convey(fmt.Sprintf("Given %s want %v", tt.name, tt.want),
				func() {
					got := HasAnyPrefix(tt.args.s, tt.args.prefixList)
					So(got, convey.ShouldEqual, tt.want)
				})
		}
	})
}
