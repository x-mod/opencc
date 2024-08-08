package opencc

import (
	"fmt"
)

type Converter struct {
	s2t  *OpenCC
	t2s  *OpenCC
	t2jp *OpenCC
	jp2t *OpenCC
}

func NewConverter() *Converter {
	return &Converter{}
}

func (c *Converter) Open() error {
	s2t, err := New("s2t")
	if err != nil {
		return fmt.Errorf("s2t open failed")
	}
	c.s2t = s2t

	t2s, err := New("t2s")
	if err != nil {
		return fmt.Errorf("t2s open failed")
	}
	c.t2s = t2s

	jp2t, err := New("jp2t")
	if err != nil {
		return fmt.Errorf("jp2t open failed")
	}
	c.jp2t = jp2t

	t2jp, err := New("t2jp")
	if err != nil {
		return fmt.Errorf("t2jp open failed")
	}
	c.t2jp = t2jp
	return nil
}

func (c *Converter) S2T(text string) (string, error) {
	return c.s2t.Convert(text)
}

func (c *Converter) T2S(text string) (string, error) {
	return c.t2s.Convert(text)
}

func (c *Converter) T2JP(text string) (string, error) {
	return c.t2jp.Convert(text)
}

func (c *Converter) JP2T(text string) (string, error) {
	return c.jp2t.Convert(text)
}

func (c *Converter) JP2S(text string) (string, error) {
	t, err := c.jp2t.Convert(text)
	if err != nil {
		return "", err
	}
	return c.t2s.Convert(t)
}

func (c *Converter) S2JP(text string) (string, error) {
	t, err := c.s2t.Convert(text)
	if err != nil {
		return "", err
	}
	return c.t2jp.Convert(t)
}
