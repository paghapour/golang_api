package common

import "encoding/json"


func TypeConverter[T any](data any) (*T, error){
	var result T
	dataJson, err := json.Marshal(&data)
	if err != nil{
		return nil, err
	}
	err = json.Unmarshal(dataJson, &result) // این میاد اون دیتایی که تبدیل به جیسون شده بودو میگیره میریزه تو ریزالت
	if err != nil {
		return nil, err
	}
	return &result, nil
}