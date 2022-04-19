package dictionary

import "github.com/vishalkuo/bimap"

const AlphabetSize = 256

type Dictionary struct {
	dict bimap.BiMap
}

func Dictionary_CreateDictionary() *Dictionary {

	d := &Dictionary{dict: *bimap.NewBiMap()}

	for i := 0; i < AlphabetSize; i++ {
		d.dict.Insert(string([]byte{byte(i)}), i)
	}

	return d
}

func (d *Dictionary) Dictionary_IsKeyContained(bytes []byte) bool {
	return d.dict.Exists(string(bytes))
}

func (d *Dictionary) Dictionary_IsValueContained(val int) bool {
	return d.dict.ExistsInverse(val)
}

func (d *Dictionary) Dictionary_AddKey(bytes []byte) {
	d.dict.Insert(string(bytes), d.dict.Size())
}

func (d *Dictionary) Dictionary_GetVal(bytes []byte) int {
	res, _ := d.dict.Get(string(bytes))
	return res.(int)
}

func (d *Dictionary) Dictionary_GetKey(val int) []byte {
	res, _ := d.dict.GetInverse(val)
	return []byte(res.(string))
}
