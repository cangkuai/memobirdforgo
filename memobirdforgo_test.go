package memobirdforgo

import (
	"fmt"
	"testing"
)

func TestRestApiPrintPaper(t *testing.T) {
	memobirdtest := newmemobirdapi("你的ak", "你的设备编号", "你的设备api")
	fmt.Println(memobirdtest.printpaperbytext("--------------------------------").getprintstate(memobirdtest.ak))
	//fmt.Println(memobirdtest.getSignalBase64Pic("/9j/2wCEAAEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQECAgICAgICAgICAgMDAwMDAwMDAwMBAQEBAQEBAgEBAgICAQICAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDA//dAAQADf/uAA5BZG9iZQBkwAAAAAH/wAARCABkAGQDABEAAREBAhEB/8QAbwABAQEAAwEBAAAAAAAAAAAAAAsJBQYIBwoBAQAAAAAAAAAAAAAAAAAAAAAQAAAGAgEDAgQEBgMAAAAAAAABAgMEBQYHEQgJIRITChQiMTlCeLYVJDRBUWF0gbMRAQAAAAAAAAAAAAAAAAAAAAD/2gAMAwAAARECEQA/AJ/4AAAAAAAAAAAAAAAAD//Qn/gAAAAAAAAAAAAAAAAP/9Gf+AAAAAAAAAAAAAAAAA//0p/4AAAAAAAAAAAAAAAAD//Tn/gAAA5zGMZyHNclx7DcRpLPJcry28qcZxjHKSE/Y3N/kN9Pj1dLSVFfFQ5Jn2drZSmmI7LaVOOuuJSkjMyIBt7qv4abvWbXfr/kOifJcJrJxpN662psLUmuWKxpSDUT9hS5JncTL+CPhJtsVr75GflHBGZBprqX4LTuPZXJZe23vfpQ1HTLbSbyavItjbLyphxSk+pH8Fr9fY5jjiUI55UV35URERGR+og1F038EDo6saS71Bdd2182fW4S1wdN6nxDVrUVr0pJUdNrm1/uBc5z1EoyeOHGLgyL2vBmoPcWffCudmzp46aN35taav3DtXIdd6S2dl8DMdl74zmHaxrXFcFvLqHeOwNWydZYo9Jhy4CX/acrlQ1qT6VsqbM0GEsoAAAAAAf/1J/4AAAOzYVmOR67zLEtgYfYqp8uwbJqHMcWt0MRpK6rI8ZtYt1SWKI0xmRDkKg2cJp0kOtraWaOFJUkzIw/TVhfxgPeIxZphu8uum3ZK2koJyRmmjm4Lso0kRGp8td5VgLCVLMuT9pDZcn4IvBEHvfp1+NC61MnzzXuB7d6Uuly/LMM1xfFJl9gc7amvCgRchvoNQux+RyHMtlNOPQmZnuKL3UIcUj8hH4CjiA8Dd1e+k4x2we4vfwnfYnVfQz1Xya9/nj2bAtFZ0iA6R+DM25i0GRckZmXHgBEAAAAAAAH/9Wf+AAAAAAO/wCqLFNRtLWtstaWkVmf4dYrcUhbiW0wsirpKlqbQSluJQTXJpIjM/sQC9yAzL70Er5PtK9x14vz9Gu/4v2M/E7XV7CP7Gk/tI/6/wB/YwidgAAAAAD/1p/4AAAAAA5CosFVNtWWqGkvLrLCFYJZUo0JdVCktSUtKWRKNKXDb4M+D45AX6AGWXe9UpPaI7ixpUaTPpN26kzSZkfpXjUpC08lx9KkKMjL+5GAiqAAAAAAD//Xn/gAAAAAAAv7kZGRGRkZGRGRkfJGR/YyP7GRkAyy73/4Q/cW/Sftv9uSAEVUAAAAAAf/0J/4AAAAAAAL90H+ih/8WP8A+SAGXfe//CH7i36T9t/tyQAiqgAAAAAD/9Gf+AAAAAAAC/mw17LDLPPq9pptr1ccer20EnnjzxzwAy273/4Q/cW/Sftv9uSAEVUAAAAAAf/Sn/gAAAAADncWgs2eTY7WyGTksWF7UQX46TcJUhmXYR2HWUm0aXSN1DhpL0mSvPjyAvtgMuu9ox8z2jO4y3ws/T0jbnf+guT/AJXEJ8rk/B/QXs/V/hPICKYAAAAAAP/Tn/gAAAAAD790oYxXZv1S9NWF3C5TdTl2/tOYxaOQXW2Zrddf7ExyqmrhvOsyGmpSI0tRtqU2tKV8GaVF4MLvQDwz3O8Nqtg9uHryw27dms1d70gdRcWW9XPMsTmkt6myuU25GdfjymEOIejpP621pPjgyMgEO0AAAAAAf//Z"))
	//fmt.Println(memobirdtest.printpaperbyimgfile("D:\\logo.jpg"))
}
