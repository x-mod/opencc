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

	c := opencc.NewConverter()
	if err := c.Open(); err != nil {
		log.Fatal(err)
	}

	if s, err := c.JP2S("無事だ"); err == nil {
		fmt.Println("s => ", s)
	}

	if t, err := c.JP2T("無事だ"); err == nil {
		fmt.Println("t => ", t)
	}

}
