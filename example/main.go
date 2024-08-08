package main

import (
	"fmt"
	"log"

	"github.com/x-mod/opencc"
)

func main() {
	s2t, err := opencc.New("s2t")
	if err != nil {
		log.Fatal(err)
	}

	in := `自然语言处理是人工智能领域中的一个重要方向。`
	out, err := s2t.Convert(in)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n%s\n", in, out)
	//自然语言处理是人工智能领域中的一个重要方向。
	//自然語言處理是人工智能領域中的一個重要方向。

	t2jp, err := opencc.New("jp2t")
	if err != nil {
		log.Fatal(err)
	}
	in1 := `保険証`
	out1, err := t2jp.Convert(in1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n%s\n", in1, out1)

	t2s, err := opencc.New("t2s")
	if err != nil {
		log.Fatal(err)
	}
	out2, err := t2s.Convert(out1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n%s\n", out1, out2)
}
